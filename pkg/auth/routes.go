package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/sabilimaulana/sebelbucks-backoffice-gateway/pkg/auth/routes"
	"github.com/sabilimaulana/sebelbucks-backoffice-gateway/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/auth/admin")
	routes.POST("/register", svc.RegisterAdmin)
	routes.POST("/login", svc.LoginAdmin)

	return svc
}

func (svc *ServiceClient) RegisterAdmin(ctx *gin.Context) {
	routes.RegisterAdmin(ctx, svc.Client)
}

func (svc *ServiceClient) LoginAdmin(ctx *gin.Context) {
	routes.LoginAdmin(ctx, svc.Client)
}
