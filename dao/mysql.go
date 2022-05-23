package dao

import (
	"database/sql"
	"fmt"
	"lantu/setting"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)

func InitMySQL(cfg *setting.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	print(DB)
	// 最大链接时长
	DB.SetConnMaxLifetime(time.Minute * 3)
	// 最大连接数
	DB.SetMaxOpenConns(10)
	// 空闲连接数
	DB.SetMaxIdleConns(10)

	err = DB.Ping()

	if err != nil {
		return err
	}

	return nil
}

func Close() {
	DB.Close()
}
