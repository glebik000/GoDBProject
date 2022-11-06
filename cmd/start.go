package cmd

import (
	"fmt"

	"GoDBProject/internal/application/rest"
)

func Start() {
	fmt.Println("start REST serv")
	rest.RunServer()
	fmt.Println("end REST serv")
}
