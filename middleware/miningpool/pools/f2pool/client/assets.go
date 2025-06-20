package client

import (
	"context"

	"github.com/NpoolPlatform/kunman/middleware/miningpool/pools/f2pool/types"
)

func (cli *Client) AssetsBalance(ctx context.Context, req *types.AssetsBalanceReq) (*types.AssetsBalanceResp, error) {
	resp := &types.AssetsBalanceResp{}
	err := cli.post(types.AssetsBalance, req, resp)
	return resp, err
}

func (cli *Client) AssetsTransactionsList(ctx context.Context, req *types.AssetsTransactionsListReq) (*types.AssetsTransactionsListResp, error) {
	resp := &types.AssetsTransactionsListResp{}
	err := cli.post(types.AssetsTransactionsList, req, resp)
	return resp, err
}

func (cli *Client) AssetsSettleModeSwitch(ctx context.Context, req *types.AssetsSettleModeSwitchReq) (*types.AssetsSettleModeSwitchResp, error) {
	resp := &types.AssetsSettleModeSwitchResp{}
	err := cli.post(types.AssetsSettleModeSwitch, req, resp)
	return resp, err
}
