package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/stebinsabu13/note_taking_microservice/api_gateway/pkg/auth"
	"github.com/stebinsabu13/note_taking_microservice/api_gateway/pkg/config"
	"github.com/stebinsabu13/note_taking_microservice/api_gateway/pkg/note"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	r := gin.Default()
	authsrv := auth.RegisterRoutes(r, &c)
	note.RegisterRoutes(r, &c, authsrv)
	r.Run(c.Port)
}
