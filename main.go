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
	r.GET("/myconf", getConf)
	r.GET("/", helloWorld)

	fmt.Printf("Servier started on port %v", viper.GetString("app.port"))
	fmt.Println()
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

func getConf(c *gin.Context) {
	c.JSON(http.StatusOK, viper.GetString("app.color"))
}

func helloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello World")
}
