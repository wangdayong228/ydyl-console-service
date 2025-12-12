package servers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	pkgmiddlewares "github.com/nft-rainbow/rainbow-goutils/middlewares"
	"github.com/sirupsen/logrus"
	"github.com/wangdayong228/ydyl-console-service/internal/configs"
	"github.com/wangdayong228/ydyl-console-service/internal/routers"
)

func Init() {
}

func Start() []*WrapHttpServer {
	cfg := configs.Get()
	publicServer := NewServer(mustInitHandler(), cfg.Server.Port)
	publicServer.Start()
	logrus.WithField("port", cfg.Server.Port).Info("public server started")

	return []*WrapHttpServer{publicServer}
}

func mustInitHandler() *gin.Engine {
	engine := gin.New()

	engine.Use(gin.Logger())
	engine.Use(pkgmiddlewares.ApiLogMiddleware(GetPublicBodyIgnoredPaths()))
	engine.Use(pkgmiddlewares.Recovery())

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	engine.Use(cors.New(corsConfig))

	routers.MustNewPublicRouter().Setup(engine)

	return engine
}

func GetPublicBodyIgnoredPaths() []string {
	return []string{
		// "/v1/user/sign/nonce",
	}
}

func Stop(servers []*WrapHttpServer) {
	for _, v := range servers {
		v.Shutdown()
	}
}
