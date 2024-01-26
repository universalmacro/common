package server

import (
	"net/http"
	"strconv"

	"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"
)

func MetricsMiddleware(router *gin.Engine) {
	p := ginprom.New(
		ginprom.Engine(router),
		ginprom.Subsystem("gin"),
		ginprom.Path("/metrics"),
	)
	router.Use(p.Instrument())
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
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
