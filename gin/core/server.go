package core

import (
	"context"
	"fmt"
	"gin-template/global"
	"gin-template/initialize"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func initServer(router *gin.Engine, addr string) {
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

func RunServer() {

	Router := initialize.Routers()

	address := fmt.Sprintf(":%s", global.GVA_CONFIG.System.Addr)

	logger := global.GVA_LOG

	logger.Info("start server", zap.String("addr", address))

	Router.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	Router.Use(ginzap.RecoveryWithZap(logger, true))

	initServer(Router, address)

}
