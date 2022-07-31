package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sabilimaulana/sebelbucks-backoffice-gateway/pkg/auth"
	"github.com/sabilimaulana/sebelbucks-backoffice-gateway/pkg/config"
	"github.com/sabilimaulana/sebelbucks-backoffice-gateway/pkg/product"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, authSvc)

	r.Run(c.Port)
}
