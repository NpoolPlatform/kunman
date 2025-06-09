//nolint:dupl
package client

import (
	"context"

	"github.com/NpoolPlatform/kunman/middleware/miningpool/pools/f2pool/types"
)

func (cli *Client) MiningUserGet(ctx context.Context, req *types.MiningUserGetReq) (*types.MiningUserGetResp, error) {
	resp := &types.MiningUserGetResp{}
	err := cli.post(types.MiningUserGet, req, resp)
	return resp, err
}

func (cli *Client) MiningUserAdd(ctx context.Context, req *types.MiningUserAddReq) (*types.MiningUserAddResp, error) {
	resp := &types.MiningUserAddResp{}
	err := cli.post(types.MiningUserAdd, req, resp)
	return resp, err
}

func (cli *Client) MiningUserList(ctx context.Context, req *types.MiningUserListReq) (*types.MiningUserListResp, error) {
	resp := &types.MiningUserListResp{}
	err := cli.post(types.MiningUserList, req, resp)
	return resp, err
}

func (cli *Client) MiningUserWalletUpdate(ctx context.Context, req *types.MiningUserWalletUpdateReq) (*types.MiningUserWalletUpdateResp, error) {
	resp := &types.MiningUserWalletUpdateResp{}
	err := cli.post(types.MiningUserWalletUpdate, req, resp)
	return resp, err
}

func (cli *Client) MiningUserReadOnlyPageAdd(ctx context.Context, req *types.MiningUserReadOnlyPageAddReq) (*types.MiningUserReadOnlyPageAddResp, error) {
	resp := &types.MiningUserReadOnlyPageAddResp{}
	err := cli.post(types.MiningUserReadOnlyPageAdd, req, resp)
	return resp, err
}

func (cli *Client) MiningUserReadOnlyPageDelete(ctx context.Context, req *types.MiningUserReadOnlyPageDeleteReq) (*types.MiningUserReadOnlyPageDeleteResp, error) {
	resp := &types.MiningUserReadOnlyPageDeleteResp{}
	err := cli.post(types.MiningUserReadOnlyPageDelete, req, resp)
	return resp, err
}

func (cli *Client) MiningUserBalanceWithdraw(ctx context.Context, req *types.MiningUserBalanceWithdrawReq) (*types.MiningUserBalanceWithdrawResp, error) {
	resp := &types.MiningUserBalanceWithdrawResp{}
	err := cli.post(types.MiningUserBalanceWithdraw, req, resp)
	return resp, err
}

func (cli *Client) MiningUserPaymentPause(ctx context.Context, req *types.MiningUserPaymentPauseReq) (*types.MiningUserPaymentPauseResp, error) {
	resp := &types.MiningUserPaymentPauseResp{}
	err := cli.post(types.MiningUserPaymentPause, req, resp)
	return resp, err
}

func (cli *Client) MiningUserPaymentResume(ctx context.Context, req *types.MiningUserPaymentResumeReq) (*types.MiningUserPaymentResumeResp, error) {
	resp := &types.MiningUserPaymentResumeResp{}
	err := cli.post(types.MiningUserPaymentResume, req, resp)
	return resp, err
}

func (cli *Client) MiningUserLeastTransferAmountUpdate(ctx context.Context, req *types.MiningUserLeastTransferAmountUpdateReq) (*types.MiningUserLeastTransferAmountUpdateResp, error) {
	resp := &types.MiningUserLeastTransferAmountUpdateResp{}
	err := cli.post(types.MiningUserLeastTransferAmountUpdate, req, resp)
	return resp, err
}
