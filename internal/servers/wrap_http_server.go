package servers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type WrapHttpServer struct {
	Handler    *gin.Engine
	Port       int
	httpServer *http.Server
}

func NewServer(handler *gin.Engine, port int) *WrapHttpServer {
	httpServer := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
		Handler: handler,
	}

	server := &WrapHttpServer{
		Handler:    handler,
		Port:       port,
		httpServer: httpServer,
	}

	return server
}

func (s *WrapHttpServer) Start() {
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
}

func (s *WrapHttpServer) Shutdown() {
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
}
