package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var mysqlClient *sql.DB

func InitMySQL() (*sql.DB, error) { // 假设 GetConfig 返回的类型是 ConfigType，需根据实际情况替换
	cfg, err := GetConfig()
	if err != nil {
		return nil, err
	}
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.MySQL.Username, cfg.MySQL.Password, cfg.MySQL.Host, cfg.MySQL.Port, cfg.MySQL.Database)

	mysqlClient, err = sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	err = mysqlClient.Ping()
	if err != nil {
		return nil, err
	}

	return mysqlClient, nil
}

func GetMySQLClient() *sql.DB {
	return mysqlClient
}
