package endpoints

import (
	"errors"
	"math/rand"
	"strings"

	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/config"
)

var (
	ErrEndpointExhausted = errors.New("all endpoints is peeked")
	ErrEndpointsEmpty    = errors.New("endpoints empty")
)

const (
	AddrSplitter = ","
	AddrMinLen   = 3
)

type Manager struct {
	len         int
	localAddrs  []string
	publicAddrs []string
}

func NewManager() (*Manager, error) {
	localWalletAddrs := config.GetENV().LocalWalletAddr
	publicWalletAddrs := config.GetENV().PublicWalletAddr

	localWalletAddrs = strings.Trim(localWalletAddrs, " ")
	publicWalletAddrs = strings.Trim(publicWalletAddrs, " ")

	_localAddrs := strings.Split(localWalletAddrs, AddrSplitter)
	localAddrs := []string{}
	_publicAddrs := strings.Split(publicWalletAddrs, AddrSplitter)
	publicAddrs := []string{}

	for i := range _localAddrs {
		if len(_localAddrs[i]) > 0 {
			localAddrs = append(localAddrs, _localAddrs[i])
		}
	}

	for i := range _publicAddrs {
		if len(_publicAddrs[i]) > 0 {
			publicAddrs = append(publicAddrs, _publicAddrs[i])
		}
	}

	if len(localAddrs) == 0 &&
		len(publicAddrs) == 0 {
		return nil, ErrEndpointsEmpty
	}

	// TODO:probability is not equal, should use Fisher-Yates algorithm
	if len(localAddrs) > 1 {
		rand.Shuffle(len(localAddrs), func(i, j int) {
			localAddrs[i], localAddrs[j] = localAddrs[j], localAddrs[i]
		})
	}
	if len(publicAddrs) > 1 {
		rand.Shuffle(len(publicAddrs), func(i, j int) {
			publicAddrs[i], publicAddrs[j] = publicAddrs[j], publicAddrs[i]
		})
	}

	// random start
	return &Manager{
		len:         len(localAddrs) + len(publicAddrs),
		localAddrs:  localAddrs,
		publicAddrs: publicAddrs,
	}, nil
}

func (m *Manager) Peek() (addr string, err error) {
	ll := len(m.localAddrs)
	pl := len(m.publicAddrs)
	if ll > 0 {
		addr = m.localAddrs[ll-1]
		m.localAddrs = m.localAddrs[0 : ll-1]
		return addr, nil
	}

	if pl > 0 {
		addr = m.publicAddrs[pl-1]
		m.publicAddrs = m.publicAddrs[0 : pl-1]
		return addr, nil
	}

	return "", ErrEndpointExhausted
}

func (m *Manager) Len() int {
	return m.len
}
