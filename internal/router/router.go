package router

import (
	"net/http"
	"os"

	"github.com/kurokobo/dify-unsandboxed-sandbox/internal/executor"
	"github.com/kurokobo/dify-unsandboxed-sandbox/internal/models"

	"github.com/gin-gonic/gin"
)

func Default() *gin.Engine {
	r := gin.Default()
	r.Use(RequireAPIKey)
	r.POST("/v1/sandbox/run", SandboxRun)
	r.GET("/v1/sandbox/dependencies", Dependencies)
	return r
}

func RequireAPIKey(c *gin.Context) {
	allowedApiKey, exists := os.LookupEnv("API_KEY")
	if !exists {
		allowedApiKey = "dify-sandbox"
	}
	apiKey := c.GetHeader("X-API-Key")
	if apiKey != allowedApiKey {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}
	c.Next()
}

func SandboxRun(c *gin.Context) {
	var req models.Request
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Language == "python3" {
		python3 := executor.Python3Executor{}
		res, err := python3.Run(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported language"})
		return
	}
}

func Dependencies(c *gin.Context) {
	language := c.Query("language")
	if language == "python3" {
		python3 := executor.Python3Executor{}
		c.JSON(http.StatusOK, python3.Dependencies())
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported language"})
		return
	}
}
