//nolint:dupl
package client

import (
	"context"

	"github.com/NpoolPlatform/kunman/middleware/miningpool/pools/f2pool/types"
)

func (cli *Client) HashRateInfo(ctx context.Context, req *types.HashRateInfoReq) (*types.HashRateInfoResp, error) {
	resp := &types.HashRateInfoResp{}
	err := cli.post(types.HashRateInfo, req, resp)
	return resp, err
}

func (cli *Client) HashRateInfoList(ctx context.Context, req *types.HashRateInfoListReq) (*types.HashRateInfoListResp, error) {
	resp := &types.HashRateInfoListResp{}
	err := cli.post(types.HashRateInfoList, req, resp)
	return resp, err
}

func (cli *Client) HashRateHistory(ctx context.Context, req *types.HashRateHistoryReq) (*types.HashRateHistoryResp, error) {
	resp := &types.HashRateHistoryResp{}
	err := cli.post(types.HashRateHistory, req, resp)
	return resp, err
}

func (cli *Client) HashRateWorkerList(ctx context.Context, req *types.HashRateWorkerListReq) (*types.HashRateWorkerListResp, error) {
	resp := &types.HashRateWorkerListResp{}
	err := cli.post(types.HashRateWorkerList, req, resp)
	return resp, err
}

func (cli *Client) HashRateWorkerHistory(ctx context.Context, req *types.HashRateWorkerHistoryReq) (*types.HashRateWorkerHistoryResp, error) {
	resp := &types.HashRateWorkerHistoryResp{}
	err := cli.post(types.HashRateWorkerHistory, req, resp)
	return resp, err
}

func (cli *Client) HashRateDistributionInfo(ctx context.Context, req *types.HashRateDistributionInfoReq) (*types.HashRateDistributionInfoResp, error) {
	resp := &types.HashRateDistributionInfoResp{}
	err := cli.post(types.HashRateDistributionInfo, req, resp)
	return resp, err
}

func (cli *Client) HashRateDistributionOrders(ctx context.Context, req *types.HashRateDistributionOrdersReq) (*types.HashRateDistributionOrdersResp, error) {
	resp := &types.HashRateDistributionOrdersResp{}
	err := cli.post(types.HashRateDistributionOrders, req, resp)
	return resp, err
}

func (cli *Client) HashRateDistributionSettlements(ctx context.Context, req *types.HashRateDistributionSettlementsReq) (*types.HashRateDistributionSettlementsResp, error) {
	resp := &types.HashRateDistributionSettlementsResp{}
	err := cli.post(types.HashRateDistributionSettlements, req, resp)
	return resp, err
}

func (cli *Client) HashRateDistribute(ctx context.Context, req *types.HashRateDistributeReq) (*types.HashRateDistributeResp, error) {
	resp := &types.HashRateDistributeResp{}
	err := cli.post(types.HashRateDistribute, req, resp)
	return resp, err
}

func (cli *Client) HashRateDistributionTerminate(ctx context.Context, req *types.HashRateDistributionTerminateReq) (*types.HashRateDistributionTerminateResp, error) {
	resp := &types.HashRateDistributionTerminateResp{}
	err := cli.post(types.HashRateDistributionTerminate, req, resp)
	return resp, err
}
