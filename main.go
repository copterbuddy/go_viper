package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	initConfig()

	r := gin.Default()
	r.GET("/", Greeting)

	r.Run(fmt.Sprintf(":%v", viper.GetString("app.port")))
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func Greeting(c *gin.Context) {
	c.JSON(http.StatusOK, fmt.Sprintf("service run on environment %v", viper.GetString("app.env")))
}
