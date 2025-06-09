package types

import (
	"context"

	basetype "github.com/NpoolPlatform/kunman/message/basetypes/v1"
)

type AssetsBalance struct {
	Balance              float64
	Paid                 float64
	TotalIncome          float64
	YesterdayIncome      float64
	EstimatedTodayIncome float64
}

type PoolManager interface {
	CheckAuth(ctx context.Context) error
	// mining user
	AddMiningUser(ctx context.Context) (userName string, readPageLink string, err error)
	ExistMiningUser(ctx context.Context, userName string) (bool, error)
	DeleteMiningUser(ctx context.Context, userName string) error

	// read page link
	AddReadPageLink(ctx context.Context, userName string) (string, error)
	GetReadPageLink(ctx context.Context, userName string) (string, error)
	DeleteReadPageLink(ctx context.Context, userName string) error

	// revenue
	SetRevenueProportion(ctx context.Context, distributor string, recipient string, proportion string) error
	GetRevenueProportion(ctx context.Context, distributor string, recipient string) (*float64, error)
	SetRevenueAddress(ctx context.Context, userName string, address string) error
	GetRevenueAddress(ctx context.Context, userName string) (string, error)

	// auto pay
	PausePayment(ctx context.Context, userName string) (bool, error)
	ResumePayment(ctx context.Context, userName string) (bool, error)

	// withdraw FractionWithdrawal
	WithdrawFractionWithdrawal(ctx context.Context, userName string) (int64, error)

	// hash rate
	GetHashRate(ctx context.Context, name string, cointypes []basetype.CoinType) (float64, error)

	GetAssetsBalance(ctx context.Context, name string) (*AssetsBalance, error)
}
