package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/YJinHai/MyIm/internal/app/agent"
	"github.com/YJinHai/MyIm/internal/app/chat"
	"github.com/YJinHai/MyIm/internal/app/router/middleware"
	"github.com/YJinHai/MyIm/internal/app/user"
)

// swagger:model
type InfoRequest struct {
	// 用户ID
	// required: true
	Uid int `json:"uid"`
}

// Load loads the middlewares, routes, handlers.
func Load() *gin.Engine {
	g := gin.New()

	//pprof.Register(g)
	//
	//g.Use(gin.Recovery())
	//g.Use(middleware.Logging())
	//g.Use(model.MysqlMid())
	g.Use(middleware.Cors())


	//authMiddleware := middleware.NewMid()

	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))



	//gAuth := g.Group("/auth")
	//{
	//	gAuth.GET("/refresh_token", authMiddleware.RefreshHandler)
	//	gAuth.POST("/update", authMiddleware.LoginHandler)
	//}


	v1 := g.Group("/v1")
	//  v1.Use(authMiddleware.MiddlewareFunc())
	v1.Use()
	{
		gUser := v1.Group("/user")
		{
			//gUser.POST("/update", user.Update)
			gUser.POST("/info", user.Info)
			gUser.POST("/login",user.Login)
			gUser.GET("/test", func (c *gin.Context){
				c.JSON(http.StatusOK,InfoRequest{
					Uid:11111,
				})
			})
		}

		gChannel := v1.Group("/channels")
		{
			gChannel.GET("/:name", chat.ListMembers)
			gChannel.POST("/register",chat.Register)
		}

		gChannelAd := v1.Group("/admin/channels")
		{
			gChannelAd.GET("", chat.ListChannels)
			gChannelAd.POST("",chat.CreateChannel)
			gChannelAd.GET("/:chanName/user/:uid",chat.UnreadCount)
		}

		gWS := v1.Group("/connect")
		{
			gWS.GET("",agent.WSConnect)
		}
	}

	return g
}
