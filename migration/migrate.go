package main

import (
	"github.com/majdsassi/go-url/initialziers"
	"github.com/majdsassi/go-url/models"
)
func init() {
	initialziers.LoadEnvVars()
	initialziers.Connect2DB()
}
func main() {
	initialziers.DB.AutoMigrate(&models.Url{})
}