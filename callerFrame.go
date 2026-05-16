/*
 * @Author: nijineko
 * @Date: 2026-05-17 02:34:34
 * @LastEditTime: 2026-05-17 02:53:54
 * @LastEditors: nijineko
 * @Description: caller frame utils
 * @FilePath: \noa-gorm\callerFrame.go
 */
package noagorm

import (
	"runtime"
	"strconv"
	"strings"
)

/**
 * @description: Get the caller frame of the SQL log
 * @return {string} caller frame
 */
func getSQLCallerFrame() string {
	PCS := [13]uintptr{}
	// the third caller usually from gorm internal
	Len := runtime.Callers(3, PCS[:])
	Frames := runtime.CallersFrames(PCS[:Len])
	var SQLFrame runtime.Frame
	for range Len {
		Frame, _ := Frames.Next()
		if (!strings.Contains(Frame.File, "noa-gorm") && !strings.Contains(Frame.File, "gorm.io/gorm") ||
			strings.HasSuffix(Frame.File, "_test.go")) && !strings.HasSuffix(Frame.File, ".gen.go") {
			SQLFrame = Frame
			break
		}
	}

	if SQLFrame.PC != 0 {
		return string(strconv.AppendInt(append([]byte(SQLFrame.File), ':'), int64(SQLFrame.Line), 10))
	}

	return ""
}
