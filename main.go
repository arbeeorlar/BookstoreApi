package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "gorm.io/driver/sqlite"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json: "title"`
	Artist string  `json: "artist"`
	Price  float64 `json:"price`
}

var albums = []album{
	{ID: "1", Title: "Emi lokan", Artist: "Tinubu", Price: 56.90},
	{ID: "2", Title: "Iwo lokan", Artist: "Sowore", Price: 100.90},
	{ID: "3", Title: "Obidient", Artist: "Obi", Price: 51.90},
	{ID: "4", Title: "Atikulate", Artist: "Atiku", Price: 70.90},
}

func getAllAbum(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
	var id = c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

}

func main() {
	router := gin.Default()
	router.GET("/albums", getAllAbum)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}
