//nolint:dupl
package client

import (
	"context"

	"github.com/NpoolPlatform/kunman/middleware/miningpool/pools/f2pool/types"
)

func (cli *Client) BlocksPaging(ctx context.Context, req *types.BlocksPagingReq) (*types.BlocksPagingResp, error) {
	resp := &types.BlocksPagingResp{}
	err := cli.post(types.BlocksPaging, req, resp)
	return resp, err
}

func (cli *Client) BlocksDateRange(ctx context.Context, req *types.BlocksDateRangeReq) (*types.BlocksDateRangeResp, error) {
	resp := &types.BlocksDateRangeResp{}
	err := cli.post(types.BlocksDateRange, req, resp)
	return resp, err
}
