package main

import (
	"coffeeshop/coffee"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

var (
	Logger *zap.Logger
	sugar  *zap.SugaredLogger
)

func init() {
	path := "logs/logs.txt"
	createFile(path)
	logger, err := NewLogger(path)
	if err != nil {
		log.Fatal(err)
	} else {
		Logger = logger
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func main() {
	sugar = Logger.Sugar()
	portNumber := getEnv("APP_PORT", "8085")

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/ping", func(c *gin.Context) {
		sugar.Infof("Hey ping %s", "pong")
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Coffeeshop!",
		})
	})
	r.GET("/home", coffeeList)

	sugar.Infof("Starting the app on port %s", portNumber)
	r.Run(fmt.Sprintf(":%s", portNumber))
}

func createFile(path string) {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		err = os.MkdirAll("logs", os.ModePerm)
		if err != nil {
			return
		}
		file, err := os.Create(path)
		if err != nil {
			return
		}
		defer file.Close()
	}
}

func NewLogger(path string) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()

	cfg.OutputPaths = []string{
		path,
	}

	return cfg.Build()
}

func coffeeList(c *gin.Context) {
	coffee, err := coffee.GetCoffees()
	if err != nil {
		sugar.Errorf("Error while getting the coffee list %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error!",
		})
		return
	}

	// Call the HTML method of the Context to render a template
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"list": coffee.List,
		},
	)
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the Coffeeshop!",
	})
}

// func setLogOutput() {
// 	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	log.SetOutput(file)
// }
