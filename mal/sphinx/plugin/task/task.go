package task

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/client"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/log"
	proxypb "github.com/NpoolPlatform/kunman/message/sphinx/proxy"

	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/getter"
	coins_register "github.com/NpoolPlatform/kunman/mal/sphinx/plugin/coins/register"

	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/config"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/env"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/rpc"
	"google.golang.org/grpc"
)

var (
	chanBuff             = 1000
	maxRetries           = 3
	retryDuration        = time.Second * 2
	registerCoinDuration = time.Second * 5
	logInterval          = 10
)

type pluginClient struct {
	closeBadConn chan struct{}
	exitChan     chan struct{}
	sendChannel  chan *proxypb.ProxyPluginResponse

	once        sync.Once
	conn        *grpc.ClientConn
	proxyClient proxypb.SphinxProxy_ProxyPluginClient
}

func Plugin(exitSig chan os.Signal, cleanChan chan struct{}) {
	for i := 1; i <= maxRetries; i++ {
		if err := newClient(exitSig, cleanChan); err != nil {
			log.Errorf("failed to connect proxy, err %v, will retry %v times (max retries: %v)", err, i, maxRetries)
			time.Sleep(retryDuration)
			continue
		}
		return
	}
	log.Errorf("failed to connect proxy, retries exhausted , exit!")
	close(cleanChan)
}

func newClient(exitSig chan os.Signal, cleanChan chan struct{}) error {
	proxyClient := &pluginClient{
		closeBadConn: make(chan struct{}),
		exitChan:     make(chan struct{}),
		sendChannel:  make(chan *proxypb.ProxyPluginResponse, chanBuff),
	}

	conn, pc, err := proxyClient.newProxyClient()
	if err != nil {
		log.Errorf("create new proxy client error: %v", err)
		return fmt.Errorf("create new proxy client error: %v", err)
	}

	proxyClient.conn, proxyClient.proxyClient = conn, pc

	go proxyClient.watch(exitSig, cleanChan)
	go proxyClient.register()
	go proxyClient.send()
	go proxyClient.recv()
	return nil
}

func (c *pluginClient) closeProxyClient() {
	c.once.Do(func() {
		log.Info("close proxy conn and client")
		if c != nil {
			close(c.exitChan)
			if c.proxyClient != nil {
				if err := c.proxyClient.CloseSend(); err != nil {
					log.Warnf("close proxy conn and client error: %v", err)
				}
			}
			if c.conn != nil {
				if err := c.conn.Close(); err != nil {
					log.Warnf("close conn error: %v", err)
				}
			}
		}
	})
}

func (c *pluginClient) newProxyClient() (*grpc.ClientConn, proxypb.SphinxProxy_ProxyPluginClient, error) {
	log.Info("start new proxy client")
	conn, err := client.GetGRPCConn(config.GetENV().Proxy)
	if err != nil {
		log.Errorf("call GetGRPCConn error: %v", err)
		return nil, nil, err
	}

	pClient := proxypb.NewSphinxProxyClient(conn)
	proxyClient, err := pClient.ProxyPlugin(context.Background())
	if err != nil {
		log.Errorf("call Transaction error: %v", err)
		return nil, nil, err
	}

	log.Info("start new proxy client ok")
	return conn, proxyClient, nil
}

func (c *pluginClient) watch(exitSig chan os.Signal, cleanChan chan struct{}) {
	for {
		select {
		case <-c.closeBadConn:
			logger.Sugar().Info("start watch proxy client")
			<-c.closeBadConn
			c.closeProxyClient()
			logger.Sugar().Info("start watch proxy client exit")

			time.Sleep(retryDuration)
			if err := newClient(exitSig, cleanChan); err != nil {
				log.Errorf("failed to connect proxy, err %v, will exit", err)
				close(cleanChan)
			}
		case <-exitSig:
			c.closeProxyClient()
			close(cleanChan)
			return
		}
	}
}

func (c *pluginClient) register() {
	logCount := 0

	for {
		select {
		case <-c.exitChan:
			log.Info("register new coin exit")
			return
		case <-time.After(registerCoinDuration):
			coinInfo, err := env.GetCoinInfo()
			if err != nil {
				log.Errorf("register new coin error: %v", err)
				continue
			}

			coinType := coins.CoinStr2CoinType(coinInfo.NetworkType, coinInfo.CoinType)
			tokenInfos := getter.GetTokenInfos(coinType)

			tokensLen := 0
			// TODO: send a msg,which contain all tokentype bellow this plugin
			for _, tokenInfo := range tokenInfos {
				if tokenInfo.DisableRegiste {
					continue
				}
				env.CheckAndSetChainInfo(tokenInfo)

				resp := &proxypb.ProxyPluginResponse{
					CoinType:            tokenInfo.CoinType,
					ChainType:           tokenInfo.ChainType,
					ChainNativeUnit:     tokenInfo.ChainNativeUnit,
					ChainAtomicUnit:     tokenInfo.ChainAtomicUnit,
					ChainUnitExp:        tokenInfo.ChainUnitExp,
					ChainID:             tokenInfo.ChainID,
					ChainNickname:       tokenInfo.ChainNickname,
					ChainNativeCoinName: tokenInfo.ChainNativeCoinName,
					GasType:             tokenInfo.GasType,
					Name:                tokenInfo.Name,
					TransactionType:     proxypb.TransactionType_RegisterCoin,
					ENV:                 tokenInfo.Net,
					Unit:                tokenInfo.Unit,
					PluginWanIP:         config.GetENV().WanIP,
					PluginPosition:      config.GetENV().Position,
				}
				tokensLen++
				c.sendChannel <- resp
			}
			if logCount%logInterval == 0 {
				log.Infof("register new coin: %v for %s network,has %v tokens,registered %v", coinType, coinInfo.NetworkType, len(tokenInfos), tokensLen)
				logCount = 0
			}
			logCount++
		}
	}
}

func (c *pluginClient) recv() {
	log.Info("plugin client start recv")
	for {
		select {
		case <-c.exitChan:
			log.Info("plugin client start recv exit")
			return
		default:
			req, err := c.proxyClient.Recv()
			if err != nil {
				log.Errorf("receiver info error: %v", err)
				if rpc.CheckCode(err) {
					c.closeBadConn <- struct{}{}
					break
				}
			}

			go func() {
				coinType := req.GetCoinType()
				transactionType := req.GetTransactionType()
				transactionID := req.GetTransactionID()

				log.Infof(
					"sphinx plugin recv info TransactionID: %v CoinType: %v TransactionType: %v",
					transactionID,
					coinType,
					transactionType,
				)

				now := time.Now()
				defer func() {
					log.Infof(
						"plugin handle coinType: %v transaction type: %v id: %v use: %v",
						coinType,
						transactionType,
						transactionID,
						time.Since(now).String(),
					)
				}()

				var resp *proxypb.ProxyPluginResponse
				var err error
				var handler coins_register.HandlerDef
				// handler, err := coins.GetCoinBalancePlugin(coinType, transactionType)
				tokenInfo := getter.GetTokenInfo(req.Name)
				if tokenInfo == nil {
					log.Errorf("GetCoinPlugin get handler error: %v", err)
					resp = &proxypb.ProxyPluginResponse{
						TransactionType: req.GetTransactionType(),
						CoinType:        req.GetCoinType(),
						TransactionID:   req.GetTransactionID(),
						RPCExitMessage:  err.Error(),
					}
					goto send
				}

				switch transactionType {
				case proxypb.TransactionType_Balance:
					handler, err = getter.GetTokenHandler(tokenInfo.TokenType, coins_register.OpGetBalance)
				case proxypb.TransactionType_EstimateGas:
					handler, err = getter.GetTokenHandler(tokenInfo.TokenType, coins_register.OpEstimateGas)
				}

				if err != nil {
					log.Errorf("GetCoinPlugin get handler error: %v", err)
					resp = &proxypb.ProxyPluginResponse{
						TransactionType: req.GetTransactionType(),
						CoinType:        req.GetCoinType(),
						TransactionID:   req.GetTransactionID(),
						RPCExitMessage:  err.Error(),
					}
					goto send
				}
				{
					respPayload, err := handler(context.Background(), req.GetPayload(), tokenInfo)
					if err != nil {
						log.Errorf("GetCoinPlugin handle deal transaction error: %v", err)
						resp = &proxypb.ProxyPluginResponse{
							TransactionType: req.GetTransactionType(),
							CoinType:        req.GetCoinType(),
							TransactionID:   req.GetTransactionID(),
							RPCExitMessage:  err.Error(),
						}
						goto send
					}

					resp = &proxypb.ProxyPluginResponse{
						TransactionType: req.GetTransactionType(),
						CoinType:        req.GetCoinType(),
						TransactionID:   req.GetTransactionID(),
						Payload:         respPayload,
					}
				}

			send:
				c.sendChannel <- resp
			}()
		}
	}
}

func (c *pluginClient) send() {
	log.Info("plugin client start send")
	for {
		select {
		case <-c.exitChan:
			log.Info("plugin client start send exit")
			return
		case resp := <-c.sendChannel:
			err := c.proxyClient.Send(resp)
			if err != nil {
				log.Errorf("send info error: %v", err)
				if rpc.CheckCode(err) {
					c.closeBadConn <- struct{}{}
				}
			}
		}
	}
}
