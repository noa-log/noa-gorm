/*
 * @Author: nijineko
 * @Date: 2025-06-09 16:20:09
 * @LastEditTime: 2025-06-09 17:16:58
 * @LastEditors: nijineko
 * @Description: noa gorm log
 * @FilePath: \noa-gorm\log_test.go
 */
package noagorm

import (
	"fmt"
	"testing"

	"github.com/noa-log/noa"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGormLog(t *testing.T) {
	Log := noa.NewLog()

	// Create a new gorm logger instance
	GormLogger := New(Log)

	Dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"password",
		"localhost",
		3306,
		"mysql",
	)

	DB, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{
		Logger: GormLogger,
	})
	if err != nil {
		t.Fatal(err)
	}

	// info log test
	var UserCount int64
	if err := DB.Table("user").Count(&UserCount).Error; err != nil {
		t.Fatal(err)
	}
	t.Log(fmt.Sprintf("User count: %d", UserCount))

	// slow log test
	DB.Exec("SELECT SLEEP(3)")

	// error log test
	DB.Table("users").Count(&UserCount)
}
