package db

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/framework/mysql"

	// ent policy runtime
	_ "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/runtime"
)

func client() (*ent.Client, error) {
	conn, err := mysql.GetConn()
	if err != nil {
		return nil, err
	}
	logger.Sugar().Infow("client", "DBStats", conn.Stats())
	drv := entsql.OpenDB(dialect.MySQL, conn)
	return ent.NewClient(ent.Driver(drv)), nil
}

func Init(hooks ...ent.Hook) error {
	cli, err := client()
	if err != nil {
		return wlog.WrapError(err)
	}
	cli.Use(hooks...)
	return cli.Schema.Create(context.Background())
}

func Client() (*ent.Client, error) {
	return client()
}

func txRun(ctx context.Context, tx *ent.Tx, fn func(ctx context.Context, tx *ent.Tx) error) error {
	succ := false
	defer func() {
		if !succ {
			err := tx.Rollback()
			if err != nil {
				logger.Sugar().Errorf("fail rollback: %v", err)
				return
			}
		}
	}()

	if err := fn(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}

	if err := tx.Commit(); err != nil {
		return wlog.Errorf("committing transaction: %v", err)
	}

	succ = true
	return nil
}

func WithTx(ctx context.Context, fn func(ctx context.Context, tx *ent.Tx) error) error {
	cli, err := Client()
	if err != nil {
		return wlog.WrapError(err)
	}
	tx, err := cli.Tx(ctx)
	if err != nil {
		return wlog.Errorf("fail get client transaction: %v", err)
	}
	return txRun(ctx, tx, fn)
}

func WithDebugTx(ctx context.Context, fn func(ctx context.Context, tx *ent.Tx) error) error {
	cli, err := Client()
	if err != nil {
		return wlog.WrapError(err)
	}
	tx, err := cli.Debug().Tx(ctx)
	if err != nil {
		return wlog.Errorf("fail get client transaction: %v", err)
	}
	return txRun(ctx, tx, fn)
}

func WithClient(ctx context.Context, fn func(ctx context.Context, cli *ent.Client) error) error {
	cli, err := Client()
	if err != nil {
		return wlog.Errorf("fail get db client: %v", err)
	}

	if err := fn(ctx, cli); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}
