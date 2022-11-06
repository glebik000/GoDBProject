package cmd

import (
	"fmt"

	"GoDBProject/internal/application/config"
	"GoDBProject/internal/application/rest"
	"GoDBProject/internal/application/storage/postgres"
)

func Start() {
	fmt.Println("start REST serv")
	rest.RunServer()
	fmt.Println("end REST serv")
}
