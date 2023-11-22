package graceful

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// ShutdownGin 停止server
func ShutdownGin(instance *http.Server, timeout time.Duration) {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("WAF服务关闭中 ...")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := instance.Shutdown(ctx); err != nil {
		log.Fatal("WAF 关闭：", err)
	}
	select {
	case <-ctx.Done():
		log.Println("--------------")
	}
	log.Println("WAF已关闭")
}
