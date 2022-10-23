package cmd

import (
	"fmt"

	"GoDBProject/internal/application/config"
	"GoDBProject/internal/application/rest"
	"GoDBProject/internal/application/storage/postgres"
)

func Start() {
	pgPoolConf := config.GetConfig()
	pools, err := postgres.NewStorage(pgPoolConf)
	defer pools.Close()
	if err != nil {
		fmt.Printf("test :\n%v", err)
		return
	}
	// fmt.Println(pools, "1")
	// all, err := pools.GetAll(context.TODO())
	// if err != nil {
	// 	return
	// }
	// fmt.Println("2")
	// fmt.Println(all)
	fmt.Println("start REST serv")
	rest.RunServer()
	fmt.Println("end REST serv")
}
