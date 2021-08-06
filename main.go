// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("hello Mili !")

// }

package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This is Mili doormachine, HTTP server",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
