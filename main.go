package main

import (
	"ant-frame/bootstrap/global"
	"ant-frame/bootstrap/initialize"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 初始化配置文件
	initialize.InitConfig()
	// 初始化日志
	initialize.InitLogger()
	// 初始化数据库

	//初始化Casbin

	// 初始化路由

}

// 优雅的重启服务器
func startHttp(r *gin.Engine) {
	host := global.Conf.Server.Host
	port := global.Conf.Server.Port
	srv := http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Error("listen error:", err)
		}
	}()
	global.Logger.Info(fmt.Sprintf("server is rinning at %s:%d", host, port))
	//创建通道
	quit := make(chan os.Signal)
	//signal.Notify 监听和捕获信号量
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.Logger.Info("shutting down server...")
	//控制goRouting的生命周期
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.Logger.Error("Server forced to shutdown:", err)
	}
	global.Logger.Info("Server exiting")
}
