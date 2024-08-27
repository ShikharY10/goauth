// This file contains routes for admin
package routes_v1

import (
	"github.com/ShikharY10/goauth/cmd/controllers/controller_v1"
	"github.com/ShikharY10/goauth/cmd/middleware"
	"github.com/gin-gonic/gin"
)

// contitution of all the admin route
func Admin(router *gin.RouterGroup, controller *controller_v1.AdminController, jwt *middleware.JWT) {

	adminRoutes := router.Group("/admin")
	adminRoutes.Use(jwt.APIV1_Authorization())
	adminRoutes.Use(middleware.RoleBasedAccess("admin"))

	adminRoutes.GET("/user/:id", controller.GetOneUserData)
	adminRoutes.GET("/users", controller.GetAllUserData)
	adminRoutes.POST("/user", controller.CreateNewUser)
	adminRoutes.DELETE("/user/:id", controller.DeleteOneUser)

	adminRoutes.PUT("/new/:id", controller.CreateNewAdmin)
}
