// This file contains routes for users
package routes_v1

import (
	"github.com/ShikharY10/goauth/cmd/controllers/controller_v1"
	"github.com/ShikharY10/goauth/cmd/middleware"
	"github.com/gin-gonic/gin"
)

// Constitution of all the user routes.
func User(router *gin.RouterGroup, controller *controller_v1.UserController, jwt *middleware.JWT) {
	router.POST("/signup", controller.SignUp)
	router.POST("/login", controller.Login)
	router.PUT("/refresh/:id", controller.RefreshAccessToken)

	authorizedRoutes := router.Group("/")
	authorizedRoutes.Use(jwt.APIV1_Authorization())

	authorizedRoutes.DELETE("/logout", controller.Logout)
	authorizedRoutes.GET("/user/:username", controller.GetOneUser)
	authorizedRoutes.GET("/users", controller.GetMultipleUsers)

}
