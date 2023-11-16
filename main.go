package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/", helloworld)

	r.GET("/aboutme", aboutMe)

	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

}

func helloworld(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")

}

func aboutMe(c *gin.Context) {
	c.HTML(http.StatusOK, "aboutme.html", nil)
}
