package route

import (
	"go/internal/pkg/api/app/handler/mst_user"
	"go/internal/pkg/api/app/handler/session"
	"go/internal/pkg/config"
	"go/internal/pkg/middleware"
	"path"

	"github.com/gin-gonic/gin"
)

//ip:port/api/v/1/group/endrouteoptional
func Invoke(conf *config.AppConfig, engine *gin.Engine,
	mst_user mst_user.UserHandler,
	session session.SessionHandler) {
	rootPath := path.Join("api",
		conf.GetString("app.version.tag"),
		conf.GetString("app.version.value"))

	route := engine.Group(rootPath)

	route.Use(middleware.InitCors())
	route.Use(gin.Logger())
	route.Use(gin.Recovery())
	route.Use(gin.ErrorLogger())

	route.OPTIONS("/*path", middleware.InitCors())

	// route master user
	{
		userRoute := route.Group("user")
		userRoute.POST("", mst_user.CreateHandler)
		userRoute.POST("login", mst_user.LoginHandler)
	}

	{
		sessionRoute := route.Group("session")
		sessionRoute.POST("", middleware.ValidationToken(conf.GetString("app.signature")), session.CreateHandler)
		sessionRoute.PUT("update/:id", middleware.ValidationToken(conf.GetString("app.signature")), session.UpdateHandler)
		sessionRoute.DELETE("delete", middleware.ValidationToken(conf.GetString("app.signature")), session.DeleteHandler)
		sessionRoute.GET("find/detail/:id", session.FindDetailHandler)
		sessionRoute.GET("find/all", session.FindAllHandler)
	}
}
