// 代码生成时间: 2025-10-04 19:02:48
package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql" // 导入MySQL驱动
    "log"
)

// DatabaseConfig 包含数据库连接的配置信息
type DatabaseConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    DBName   string
}

// DBPool 管理数据库连接池
type DBPool struct {
    *sql.DB
}

// NewDBPool 创建一个新的数据库连接池
func NewDBPool(cfg *DatabaseConfig) (*DBPool, error) {
    // 构建DSN（数据源名称）
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        cfg.User,
        cfg.Password,
        cfg.Host,
        cfg.Port,
        cfg.DBName)
    
    // 打开数据库连接
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    
    // 设置数据库连接池参数
    db.SetMaxOpenConns(100) // 设置最大打开连接数
    db.SetMaxIdleConns(10)  // 设置最大空闲连接数
    db.SetConnMaxLifetime(5 * 60 * 60) // 设置连接最大存活时间
    
    // 测试连接
    if err := db.Ping(); err != nil {
        return nil, err
    }
    
    return &DBPool{db}, nil
}

// Close 关闭数据库连接池
func (p *DBPool) Close() error {
    return p.DB.Close()
}

func main() {
    // 数据库配置
    config := &DatabaseConfig{
        Host:     "localhost",
        Port:     3306,
        User:     "root",
        Password: "password",
        DBName:   "testdb",
    }
    
    // 创建数据库连接池
    dbPool, err := NewDBPool(config)
    if err != nil {
        log.Fatal(err)
    }
    defer dbPool.Close()
    
    // 使用数据库连接池
    fmt.Println("Database connection pool is ready.")
}
