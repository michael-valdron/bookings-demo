package web

import (
	"log"
	"os"

	"bookings.server/pkg/db"
	"bookings.server/pkg/web/api"
	middleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	swagger, err := api.GetSwagger()
	if err != nil {
		log.Fatalf("Error loading swagger spec: %v", err)
	}

	swagger.Servers = nil

	conn, err := db.ConnectDB()
	if err != nil {
		os.Exit(1)
	}
	defer conn.Close()

	server := &api.Server{
		Dao: db.DAO{
			Connection: conn,
		},
	}

	router := gin.Default()

	router.Use(middleware.OapiRequestValidator(swagger))

	router = api.RegisterHandlers(router, server)

	router.Run(":8000")
}
