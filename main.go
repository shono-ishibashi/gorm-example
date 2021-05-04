package main

import (
	"fmt"
	"gorm_example/controllers"
	"gorm_example/models"
)

func main() {
	fmt.Println(models.Db)
	controllers.StartMainServer()
}
