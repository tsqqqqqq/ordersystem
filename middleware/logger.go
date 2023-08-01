package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"time"
)

func Logger() *logrus.Logger {
	// 获取日志文件
	f, _ := os.OpenFile("./log/order.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// 实例化
	logger := logrus.New()
	// 设置颜色
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})
	// 设置输出
	out := io.MultiWriter(f, os.Stdout)
	logger.SetOutput(out)
	// 设置日志级别
	logger.SetLevel(logrus.InfoLevel)
	return logger
}

func LoggerToFile() gin.HandlerFunc {
	logger := Logger()
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		// 根据状态码设置日志格式
		switch statusCode {
		case http.StatusNotFound:
			logger.WithFields(logrus.Fields{
				"client_ip":    clientIP,
				"latency_time": latencyTime,
				"req_method":   reqMethod,
				"req_uri":      reqUri}).Warning("404")
		case http.StatusInternalServerError:
			logger.WithFields(logrus.Fields{
				"client_ip":    clientIP,
				"latency_time": latencyTime,
				"req_method":   reqMethod,
				"req_uri":      reqUri}).Error("500")
		}
	}
}
