package server

import (
	"net/http"
	"strconv"

	"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/universalmacro/common/snowflake"
)

func RequestIDMiddleware() gin.HandlerFunc {
	requestIdGenetator := snowflake.NewIdGenertor(0)
	return func(ctx *gin.Context) {
		ctx.Set("requestID", requestIdGenetator.Uint())
		ctx.Next()
	}
}

func GetRequestID(ctx *gin.Context) uint {
	id, _ := ctx.Get("requestID")
	return id.(uint)
}

func MetricsMiddleware(router *gin.Engine) {
	p := ginprom.New(
		ginprom.Engine(router),
		ginprom.Subsystem("gin"),
		ginprom.Path("/metrics"),
	)
	router.Use(p.Instrument())
}

func CorsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}
		ctx.Next()
	}
}

func IndexAndLimit(ctx *gin.Context) (index, limit int64) {
	index = 0
	limit = 10
	if marks, err := strconv.Atoi(ctx.Query("index")); err != nil {
		index = int64(marks)
	}
	if marks, err := strconv.Atoi(ctx.Query("limit")); err != nil {
		limit = int64(marks)
	}
	if limit == 0 {
		limit = 1
	}
	return
}

func Int64ID(ctx *gin.Context, path string) int64 {
	return int64(IntID(ctx, path))
}

func IntID(ctx *gin.Context, path string) int {
	id, _ := strconv.Atoi(ctx.Param(path))
	return id
}

func UintID(ctx *gin.Context, path string) uint {
	return uint(IntID(ctx, path))
}
