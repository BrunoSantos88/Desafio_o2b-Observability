package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

func main() {
	// Configuração do Jaeger
	cfg, err := config.FromEnv()
	if err != nil {
		fmt.Println("Erro ao configurar o Jaeger:", err)
		return
	}
	cfg.ServiceName = "meu-app-go"

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		fmt.Println("Erro ao criar o Tracer:", err)
		return
	}
	defer closer.Close()

	opentracing.SetGlobalTracer(tracer)

	// Configuração do Gin
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		span := opentracing.StartSpan("hello-handler")
		defer span.Finish()

		c.JSON(http.StatusOK, gin.H{"message": "Bem-vindo ao meu aplicativo Go!"})
	})

	// Inicia o servidor Gin
	port := getPort()
	router.Run(":" + port)
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	return port
}
