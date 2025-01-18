package cmd

import (
	"e-wallet-notification/helpers"
	"github.com/gin-gonic/gin"
	"log"
)

func ServeHTTP() {

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	err := r.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal("Failed to set trusted proxies", err)
	}

	err = r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal("Failed to start server", err)
	}
	log.Println("Server started")
}
