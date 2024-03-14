package apis

import (
	"log"

	"wgm/cli/apis/data"
	"wgm/cli/apis/peers"
	"wgm/cli/apis/routes"
	"wgm/cli/apis/users"
	"wgm/common/configs"
	"wgm/common/services/common/logs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Run 执行服务
func Run(listenAddr string, handle gin.HandlerFunc) error {
	log.Println("Listen: " + listenAddr)

	logger, err := logs.GinLogger("access.log")
	if err != nil {
		return err
	}

	if configs.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	if handle != nil {
		router.Use(handle)
		log.Println("Mode: embed")
	} else {
		log.Println("Mode: api")
	}
	router.SetTrustedProxies(nil)

	// 跨域
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTION"}
	router.Use(cors.New(config))

	router.Use(gin.Recovery())
	router.Use(logger)

	v1 := router.Group("/api/v1")
	{

		v1.GET("/user/list", users.ListUsers)
		v1.POST("/user/add", users.AddUser)
		v1.POST("/user/update/:userID", users.UpdateUser)
		v1.POST("/user/delete/:userID", users.DeleteUser)

		v1.GET("/route/list", routes.ListRoutes)
		v1.POST("/route/add", routes.AddRoute)
		v1.POST("/route/update/:routeID", routes.UpdateRoute)
		v1.POST("/route/delete/:routeID", routes.DeleteRoute)

		v1.GET("/peer/list", peers.ListPeers)
		v1.POST("/peer/add", peers.AddPeer)
		v1.POST("/peer/update/:peerID", peers.UpdatePeer)
		v1.POST("/peer/delete/:peerID", peers.DeletePeer)

		v1.GET("/data", data.ListPeerData)
	}

	return router.Run(listenAddr)
}
