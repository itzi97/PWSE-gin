package main

import "github.com/gin-gonic/gin"

// Structure to store form data
type userForm struct {
	Name    string `form:"name"`
	Surname string `form:"surname"`
}

func main() {
	r := gin.Default()

	// Load templates
	r.LoadHTMLGlob("templates/*")

	// Set routes
	r.GET("/", index)
	r.POST("/", submit)

	// Run in 8080
	r.Run(":8080")
}

// Print template from index.html
func index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

// Post submitted data in JSON
func submit(c *gin.Context) {
	var form userForm
	c.Bind(&form)
	c.JSON(200, gin.H{
		"name":    form.Name,
		"surname": form.Surname,
	})
}
