package main

import (
	"os"

	_ "mygram-byferdiansyah/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MyGram By Ferdiansya
// @version 1.0
// @description This API was made as a primary purpose for one of the requirements (final project) in the Hack8tiv and FGA Kominfo courses. MyGram is a website similar to Instagram. On this website, users can register by login (if they are over eight years old), post images, and comments.
// @termOfService http://swagger.io/terms/
// @contact.name ferdi
// @contact.email ferdicompany@gmail.com
// @license.name MIT License
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey  Bearer
// @in                          header
// @name                        Authorization
// @description					        Description for what is this security definition being used
func main() {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("Error loading .env file: ", err)
	// }

	// db := database.StartDB()

	routers := gin.Default()

	routers.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	})

	routers.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("Error loading .env file: ", err)
	// }

	port := os.Getenv("PORT")

	if len(os.Args) > 1 {
		reqPort := os.Args[1]

		if reqPort != "" {
			port = reqPort
		}
	}

	if port == "" {
		port = "8080"
	}

	routers.Run(":" + port)
}
