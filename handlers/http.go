package handlers

import (
	"encoding/json"
	"fmt"

	"prometheus-http-sd/models"

	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var targets models.TargetsList

func HttpHandler(logger *zap.Logger) {

	logger.Info("Building Routes")
	router := gin.Default()
	gin.SetMode("release")
	router.GET("/targets", GetJson)

	MetricHandler(router)
	HealthHandlers(router)

	router.Run(":9990")
}

func GetJson(c *gin.Context) {

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	c.Header("Content-type", "application/json")
	c.Header("Access-Control-Allow-Methods", "GET")

	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &targets)

	c.IndentedJSON(200, &targets)
}
