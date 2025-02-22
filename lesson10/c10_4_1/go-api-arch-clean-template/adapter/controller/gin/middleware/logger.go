package middleware

import (
	"go-api-arch-clean-template/pkg/logger"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

func GinZap() gin.HandlerFunc {
	return ginzap.Ginzap(logger.ZapLogger, time.RFC3339, true)
}

func RecoveryWithZap() gin.HandlerFunc {
	return ginzap.RecoveryWithZap(logger.ZapLogger, true)
}
