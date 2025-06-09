# Noa Gorm
The Gorm integration module for Noa Log, enabling quick integration of Noa into Gorm to print SQL logs and more in a unified manner.

## Installation
```bash
go get -u github.com/noa-log/noa-gorm
```

## Quick Start
```go
package main

import (
    "github.com/noa-log/noa"
    noagorm "github.com/noa-log/noa-gorm"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func main() {
    // Create a new Noa logger instance
    logger := noa.NewLog()

    db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
        Logger: noagorm.New(logger), // Replace the default logger with Noa Gorm logger
    })
    if err != nil {
        panic("failed to connect database")
    }

    // ... perform database operations
}
```

## Configuration
Noa Gorm shares configuration with the Noa instance. Some sql log options can be customized within the Gorm logger instance.
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
    // Create a new Noa logger instance
    logger := noa.NewLog()

    // Create a Gorm logger using the Noa instance
    gormLogger := noagorm.New(logger)

    // Ignore 'record not found' errors
    gormLogger.IgnoreRecordNotFoundError = true
    // Set the threshold for slow SQL logs
    gormLogger.SlowThreshold = 200 * time.Millisecond

    db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
        Logger: gormLogger, // Replace the default logger with Noa Gorm logger
    })
    if err != nil {
        panic("failed to connect database")
    }

    // ... perform database operations
}
```

## License
This project is open-sourced under the [Apache License 2.0](https://www.apache.org/licenses/LICENSE-2.0). Please comply with the terms when using it.