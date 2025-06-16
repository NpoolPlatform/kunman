package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/NpoolPlatform/kunman/framework/config"
	"github.com/NpoolPlatform/kunman/framework/logger"
	constant "github.com/NpoolPlatform/kunman/framework/mysql/const"
	_ "github.com/go-sql-driver/mysql" // nolint
	"github.com/hashicorp/consul/api"
)

type DB struct {
	db    *sql.DB
	dsn   string
	mutex sync.Mutex
}

const (
	keyUsername = "username"
	keyPassword = "password"
	keyDBName   = "database_name"

	checkDuration = time.Second * 10
	pingTimeout   = time.Second * 5
)

func Initialize(domain string) (*DB, error) {
	dsn, err := mysqlConfig(domain)
	if err != nil {
		return nil, err
	}

	_db, err := newConnection(dsn)
	if err != nil {
		return nil, err
	}

	db := &DB{
		db:    _db,
		dsn:   dsn,
		mutex: sync.Mutex{},
	}

	go db.heartbeat()

	return db, nil
}

func (db *DB) SafeRun(f func(db *sql.DB) error) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	return f(db.db)
}

func (db *DB) heartbeat() {
	for range time.After(checkDuration) {
		ctx, cancel := context.WithTimeout(context.Background(), pingTimeout)
		err := db.db.PingContext(ctx)
		cancel()

		if err == nil {
			continue
		}
		logger.Sugar().Warnw("Failed ping database", "Error", err)
		if err != nil && strings.Contains(err.Error(), "Too many connections") {
			continue
		}

		db.mutex.Lock()

		func() {
			defer db.mutex.Unlock()

			db.db, err = newConnection(db.dsn)
			if err != nil {
				logger.Sugar().Warnf("call open error: %v", err)
			}
		}()
	}
}

func newConnection(dsn string) (conn *sql.DB, err error) {
	conn, err = open("mysql", dsn)
	if err != nil {
		logger.Sugar().Warnf("call open error: %v", err)
		return nil, err
	}

	return
}

func apolloConfig() (*api.AgentService, error) {
	return config.PeekService(constant.MysqlServiceName)
}

func mysqlConfig(domain string) (string, error) {
	username := config.GetStringValueWithNameSpace(constant.MysqlServiceName, keyUsername)
	password := config.GetStringValueWithNameSpace(constant.MysqlServiceName, keyPassword)

	dbname := config.GetStringValueWithNameSpace(domain, keyDBName)
	if dbname == "" {
		logger.Sugar().Warnw("Invalid database", "Domain", domain)
		return "", fmt.Errorf("Invalid database")
	}

	svc, err := apolloConfig()
	if err != nil {
		logger.Sugar().Warnw("Failed get apollo config", "Error", err)
		return "", err
	}

	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&interpolateParams=true&multiStatements=true",
		username,
		password,
		svc.Address,
		svc.Port,
		dbname,
	), nil
}

func open(driverName, dataSourceName string) (conn *sql.DB, err error) {
	logger.Sugar().Infof("[Re]open database %v: %v", driverName, dataSourceName)
	conn, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		logger.Sugar().Warnw("Failed open database", "Error", err)
		return nil, err
	}

	// https://github.com/go-sql-driver/mysql
	// See "Important settings" section.
	conn.SetConnMaxLifetime(time.Minute * 3)
	conn.SetMaxOpenConns(2)
	conn.SetMaxIdleConns(1)

	return conn, nil
}
