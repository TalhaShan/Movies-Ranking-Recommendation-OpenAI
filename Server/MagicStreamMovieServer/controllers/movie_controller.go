package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/TalhaShan/Movies-Ranking-Recommendation-OpenAI/Server/MagicStreamMovieServer/database"
	"github.com/TalhaShan/Movies-Ranking-Recommendation-OpenAI/Server/MagicStreamMovieServer/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func GetMovies(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 100*time.Second)
		defer cancel()
		// c.JSON(200, gin.H{"message": "List of movies"}) //quick check API
		var movieCollection *mongo.Collection = database.OpenCollection("movies", client)

		cursor, err := movieCollection.Find(ctx, bson.D{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies."})
		}
		defer cursor.Close(ctx)

		var movies []models.Movie

		if err = cursor.All(ctx, &movies); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode movies."})
			return
		}

		c.JSON(http.StatusOK, movies)
	}
}
func GetMovie(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 100*time.Second)
		defer cancel()

		var movieCollection *mongo.Collection = database.OpenCollection("movies", client)

		movieID := c.Param("imbd_id")
		if movieID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to fetch movies."})
			return
		}

		var movie models.Movie

		err := movieCollection.FindOne(ctx, bson.M{"imbd_id": movieID}).Decode(&movie)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Movuew not found"})
			return
		}

		c.JSON(http.StatusOK, movie)
	}
}
