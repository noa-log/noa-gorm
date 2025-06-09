# Noa Gorm
Noa Log 的 Gorm 集成模块，可以快速地将 Noa 集成到 Gorm 中，以通过统一的方式打印 SQL 日志等。

## 安装
```bash
go get -u github.com/noa-log/noa-gorm
```

## 快速开始
```go
package main

import (
    "github.com/noa-log/noa"
    noagorm "github.com/noa-log/noa-gorm"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func main() {
    // 创建一个新的日志实例
    logger := noa.NewLog()

    db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
        Logger: noagorm.New(logger), // 使用 Noa Gorm 日志实例替换默认日志实例
    })
    if err != nil {
        panic("failed to connect database")
    }

    // ... 进行数据库操作
}
```

## 配置
Noa Gorm 共享 Noa 实例的配置，部分针对配置项可以在 Gorm 日志实例修改
```go
package main

import (
    "time"
    "github.com/noa-log/noa"
    noagorm "github.com/noa-log/noa-gorm"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func main() {
    // 创建一个新的日志实例
    logger := noa.NewLog()

    // 使用 Noa 实例创建一个 Gorm 日志实例
    gormLogger := noagorm.New(logger)

    // 忽略未找到记录错误
    gormLogger.IgnoreRecordNotFoundError = true
    // 设置慢SQL阈值
    gormLogger.SlowThreshold = 200 * time.Millisecond

    db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
        Logger: gormLogger, // 使用 Noa Gorm 日志实例替换默认日志实例
    })
    if err != nil {
        panic("failed to connect database")
    }

    // ... 进行数据库操作
}
```

## 许可
本项目基于[Apache License 2.0](https://www.apache.org/licenses/LICENSE-2.0)协议开源。使用时请遵守协议的条款。