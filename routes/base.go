package router

import (
	"github.com/BioforestChain/dweb-browser-matrix-service-search/routes/api/authRouter"
	"github.com/BioforestChain/dweb-browser-matrix-service-search/routes/api/helloRouter"
	"github.com/BioforestChain/dweb-browser-matrix-service-search/routes/api/rsaGenRouter"
	"github.com/BioforestChain/dweb-browser-matrix-service-search/routes/api/userRouter"
	"github.com/gin-gonic/gin"
	"strings"
)

// RegisterConfig 路由配置
func RegisterConfig() []func(r *gin.Engine) {
	// 配置路由列表
	return []func(r *gin.Engine){
		// 其他
		RegisterOther,

		// 功能模块路由
		helloRouter.RegisterHello,
		userRouter.RegisterUser,
		authRouter.RegisterAuth,
		rsaGenRouter.RegisterRsaGen,
	}
}

// RegisterRouter 注册路由
func RegisterRouter(router *gin.Engine) {
	if router == nil {
		return
	}
	// 注册所有路由
	for _, register := range RegisterConfig() {
		register(router)
	}
}

// AddRouter 兼容路由斜杆，路由( test/和test )
// 解决问题：HTTP/1.1 301 Moved Permanently redirecting request 301: /test/ --> /test
func AddRouter(method func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes, relativePath string, handlers ...gin.HandlerFunc) {
	strings.TrimRight(relativePath, "/")
	method(relativePath, handlers...)
	method(relativePath+"/", handlers...)
}
