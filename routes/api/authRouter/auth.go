package authRouter

import (
	"github.com/BioforestChain/dweb-browser-matrix-service-search/internal/app/http/controller/api/authController"
	config2 "github.com/BioforestChain/dweb-browser-matrix-service-search/internal/helper/config"
	"github.com/gin-gonic/gin"
)

func RegisterAuth(router *gin.Engine) {
	domainMark := config2.GetDomainMark("domain")

	route := router.Group(domainMark + "/auth")
	{
		route.GET("/:token/*action", func(ctx *gin.Context) {

			authController.NewController(ctx).GetAuthProtocol()
		})

		route.GET("/test-auth", func(ctx *gin.Context) {
			authController.NewController(ctx).TestAuth()
		})
	}
}
