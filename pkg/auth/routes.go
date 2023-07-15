package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/stebinsabu13/note_taking_microservice/api_gateway/pkg/auth/routes"
	"github.com/stebinsabu13/note_taking_microservice/api_gateway/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}
	r.POST("/signup", svc.Register)
	r.POST("/login", svc.Login)

	return svc
}

func (s *ServiceClient) Register(c *gin.Context) {
	routes.Register(c, s.Client)
}

func (s *ServiceClient) Login(c *gin.Context) {
	routes.Login(c, s.Client)
}
