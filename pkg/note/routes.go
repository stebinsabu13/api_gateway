package note

import (
	"github.com/gin-gonic/gin"
	"github.com/stebinsabu13/note_taking_microservice/api_gateway/pkg/auth"
	"github.com/stebinsabu13/note_taking_microservice/api_gateway/pkg/config"
	"github.com/stebinsabu13/note_taking_microservice/api_gateway/pkg/note/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authsrv *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authsrv)
	svc := &ServiceClient{
		Client: InitServCli(c),
	}

	r.Use(a.AuthRequired)
	r.POST("/notes", svc.CreateNote)
	r.GET("/notes", svc.ListAllNote)
	r.DELETE("/notes", svc.DeleteNote)
}

func (s *ServiceClient) CreateNote(c *gin.Context) {
	routes.CreateNote(c, s.Client)
}

func (s *ServiceClient) ListAllNote(c *gin.Context) {
	routes.ListAllNote(c, s.Client)
}

func (s *ServiceClient) DeleteNote(c *gin.Context) {
	routes.DeleteNote(c, s.Client)
}
