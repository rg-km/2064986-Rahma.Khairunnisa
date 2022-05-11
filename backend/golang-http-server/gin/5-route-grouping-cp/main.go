package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	Title string
}

var movies = map[int]Movie{
	1: {
		"Spiderman",
	},
	2: {
		"Joker",
	},
	3: {
		"Escape Plan",
	},
	4: {
		"Lord of the Rings",
	},
}

var MovieListHandler = func(c *gin.Context) {
	var movieList []Movie
	for _, movie := range movies {
		movieList = append(movieList, movie)
	}
	c.JSON(http.StatusOK, gin.H{"data": movieList}) // TODO: replace this
}

var MovieGetHandler = func(c *gin.Context) {
	paramID := c.Param("id")
	movieID, _ := strconv.Atoi(paramID)
	if _, ok := movies[movieID]; !ok {
		c.JSON(http.StatusNotFound, gin.H{"message": "data not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": movies[movieID]}) // TODO: replace this
}

func GetRouter() *gin.Engine {
	router := gin.Default()
	// TODO: answer here
	movieRoute := router.Group("/movie")
	{
		movieRoute.GET("/List", MovieListHandler)
		movieRoute.GET("/get/:id", MovieGetHandler)
	}
	return router
}

func main() {
	router := GetRouter()

	router.Run(":8080")
}
