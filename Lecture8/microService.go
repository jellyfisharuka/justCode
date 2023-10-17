package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	authorized := r.Group("/", AuthMiddleware("tokenXcxzcasdKLD SAdxc"))
	authorized.POST("/endpoint", CreateEntity)
	authorized.GET("/endpoint", GetEntities)
	r.Run(":3001")
}

func AuthMiddleware(expectedT string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tok := c.Request.Header.Get("Authorization")
		if tok != expectedT {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

type Entity struct {
	Id int64 `json:"id"`
}

var entities []Entity

func CreateEntity(c *gin.Context) {
	var newEntity Entity
	if err := c.ShouldBindJSON(&newEntity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entities = append(entities, newEntity)
	c.JSON(http.StatusCreated, newEntity)
}

func GetEntities(c *gin.Context) {
	c.JSON(http.StatusOK, entities)
}
