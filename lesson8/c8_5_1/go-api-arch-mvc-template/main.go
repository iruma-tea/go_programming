package main

import (
	"go-api-arch-mvc-template/app/models"
	"go-api-arch-mvc-template/pkg/logger"
)

func main() {
	if err := models.SetDatabase(models.InstanceMySQL); err != nil {
		logger.Fatal(err.Error())
	}
}
