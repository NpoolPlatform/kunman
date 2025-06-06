package db

import (
	"context"
	"database/sql"
	"sync"

	"github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	ent "github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated"
	servicename "github.com/NpoolPlatform/kunman/middleware/agi/servicename"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/framework/mysql"

	// ent policy runtime
	_ "github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated/runtime"
)

var db *mysql.DB
var mutex sync.Mutex

func client(f func(cli *ent.Client) error) error {
	mutex.Lock()
	defer mutex.Unlock()

	if db == nil {
		var err error
		db, err = mysql.Initialize(servicename.ServiceDomain)
		if err != nil {
			return err
		}

		if err := db.SafeRun(func(db *sql.DB) error {
			drv := entsql.OpenDB(dialect.MySQL, db)
			cli := ent.NewClient(ent.Driver(drv))

			cli.Use()
			return cli.Schema.Create(context.Background())
		}); err != nil {
			return err
		}
	}

	return db.SafeRun(func(db *sql.DB) error {
		drv := entsql.OpenDB(dialect.MySQL, db)
		cli := ent.NewClient(ent.Driver(drv))
		return f(cli)
	})
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
	return client(func(cli *ent.Client) error {
		tx, err := cli.Tx(ctx)
		if err != nil {
			return wlog.Errorf("fail get client transaction: %v", err)
		}
		return txRun(ctx, tx, fn)
	})
}

func WithDebugTx(ctx context.Context, fn func(ctx context.Context, tx *ent.Tx) error) error {
	return client(func(cli *ent.Client) error {
		tx, err := cli.Debug().Tx(ctx)
		if err != nil {
			return wlog.Errorf("fail get client transaction: %v", err)
		}
		return txRun(ctx, tx, fn)
	})
}

func WithClient(ctx context.Context, fn func(ctx context.Context, cli *ent.Client) error) error {
	return client(func(cli *ent.Client) error {
		if err := fn(ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
}
