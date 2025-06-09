/*
 * @Author: nijineko
 * @Date: 2025-06-09 16:20:09
 * @LastEditTime: 2025-06-09 17:03:18
 * @LastEditors: nijineko
 * @Description: noa gorm log
 * @FilePath: \noa-gorm\log.go
 */
package noagorm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/noa-log/noa"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

// Implementation Gorm logger interface
type gormLogger struct {
	IgnoreRecordNotFoundError bool          // ignore record not found error
	SlowThreshold             time.Duration // slow SQL threshold

	log *noa.LogConfig // noa log instance
}

/**
 * @description: Create a new gorm logger instance
 * @param {noa.LogConfig} Log noa log configuration
 * @return {*gormLogger} gorm logger instance
 */
func New(Log *noa.LogConfig) *gormLogger {
	return &gormLogger{
		IgnoreRecordNotFoundError: false,
		SlowThreshold:             200 * time.Millisecond,
		log:                       Log,
	}
}

// set log mode
func (gl *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	// Ignore the log level parameter
	return gl
}

// print info log
func (gl *gormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	gl.log.Info(DEFAULT_LOG_SOURCE, fmt.Sprintf("%s %s", append([]any{utils.FileWithLineNum()}, data...)...))
}

// print warn log
func (gl *gormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	gl.log.Warning(DEFAULT_LOG_SOURCE, fmt.Sprintf("%s %s", append([]any{utils.FileWithLineNum()}, data...)...))
}

// print error log
func (gl *gormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	gl.log.Error(DEFAULT_LOG_SOURCE, fmt.Sprintf("%s %s", append([]any{utils.FileWithLineNum()}, data...)...))
}

const (
	TraceFormat     = "%s | %.3fms | rows:%v | %s"
	TraceWarnFormat = "%s | %.3fms | rows:%v | %s\n%s"
	TraceErrFormat  = "%s | %.3fms | rows:%v | %s\n"
)

// print trace log
func (gl *gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	Elapsed := time.Since(begin)
	switch {
	// error log
	case err != nil && (!errors.Is(err, gorm.ErrRecordNotFound) || !gl.IgnoreRecordNotFoundError):
		SQL, Rows := fc()
		if Rows == -1 {
			gl.log.Error(DEFAULT_LOG_SOURCE, fmt.Sprintf(TraceErrFormat, utils.FileWithLineNum(), float64(Elapsed.Nanoseconds())/1e6, "-", SQL), err)
		} else {
			gl.log.Error(DEFAULT_LOG_SOURCE, fmt.Sprintf(TraceErrFormat, utils.FileWithLineNum(), float64(Elapsed.Nanoseconds())/1e6, Rows, SQL), err)
		}
	// slow log
	case Elapsed > gl.SlowThreshold && gl.SlowThreshold != 0:
		SQL, Rows := fc()
		SlowLog := fmt.Sprintf("SLOW SQL >= %v", gl.SlowThreshold)
		if Rows == -1 {
			gl.log.Warning(DEFAULT_LOG_SOURCE, fmt.Sprintf(TraceWarnFormat, utils.FileWithLineNum(), float64(Elapsed.Nanoseconds())/1e6, "-", SQL, SlowLog))
		} else {
			gl.log.Warning(DEFAULT_LOG_SOURCE, fmt.Sprintf(TraceWarnFormat, utils.FileWithLineNum(), float64(Elapsed.Nanoseconds())/1e6, Rows, SQL, SlowLog))
		}
	// normal log
	default:
		SQL, Rows := fc()
		if Rows == -1 {
			gl.log.Info(DEFAULT_LOG_SOURCE, fmt.Sprintf(TraceFormat, utils.FileWithLineNum(), float64(Elapsed.Nanoseconds())/1e6, "-", SQL))
		} else {
			gl.log.Info(DEFAULT_LOG_SOURCE, fmt.Sprintf(TraceFormat, utils.FileWithLineNum(), float64(Elapsed.Nanoseconds())/1e6, Rows, SQL))
		}
	}
}
