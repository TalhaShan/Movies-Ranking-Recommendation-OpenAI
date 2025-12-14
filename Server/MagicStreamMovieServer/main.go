package main

import (
	"fmt"

	"github.com/GavinLonDigital/MagicStream/Server/MagicStreamServer/database"
	"github.com/TalhaShan/Movies-Ranking-Recommendation-OpenAI/Server/MagicStreamMovieServer/controllers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func main() {
	fmt.Print("asdas")
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello movie fun")

	})

	var client *mongo.Client = database.Connect()

	router.GET("/movies", controllers.GetMovies(client))
	router.GET("/movie/:imdb_id", controllers.GetMovie(client))
	router.POST("/addmovie", controllers.AddMovie(client))
	router.POST("/register", controllers.RegisterUser(client))
	router.POST("/login", controllers.LoginUser(client))

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server", err)
	}
}
