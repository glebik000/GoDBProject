package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"GoDBProject/internal/application/config"
	"GoDBProject/internal/application/models"
	"GoDBProject/internal/application/storage/postgres"
)

func GetAllData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pgPoolConf := config.GetConfig()
	pools, err := postgres.NewStorage(pgPoolConf)
	if err != nil {
		fmt.Printf("error IN HANDLE %v", err)
		return
	}
	res, err := pools.GetAll(context.Background())
	if err != nil {
		fmt.Printf("error IN REQUEST %v", err)
		return
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		fmt.Printf("ошибка при ENCODE JSON %v", err)
	}
}

func GetAllDataBy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pgPoolConf := config.GetConfig()
	pools, err := postgres.NewStorage(pgPoolConf)
	if err != nil {
		fmt.Printf("error IN HANDLE %v", err)
		return
	}

	params := mux.Vars(r)
	fmt.Println(params["id"])
	fmt.Println(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Printf("Не получилось преобразовать строку в int %v", err)
	}

	res, err := pools.GetMaterialByIdService(context.Background(), id)
	if err != nil {
		fmt.Printf("error IN REQUEST %v", err)
		return
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		fmt.Printf("ошибка при ENCODE JSON %v", err)
	}
}

func InsertProductToService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pgPoolConf := config.GetConfig()
	pools, err := postgres.NewStorage(pgPoolConf)
	if err != nil {
		fmt.Printf("error IN HANDLE %v", err)
		return
	}
	var (
		a, b, c int
		// buf     []byte
		// testMap map[string]interface{}
		testStruct models.Attacher
	)
	err = json.NewDecoder(r.Body).Decode(&testStruct)
	if err != nil {
		fmt.Printf("error IN DECODE %v", err)
		return
	}
	fmt.Println(testStruct)

	err = pools.InsertPTS(context.TODO(), a, b, c)
	if err != nil {
		fmt.Printf("error IN REQUEST %v", err)
		return
	}
	err = json.NewEncoder(w).Encode(err)
	if err != nil {
		fmt.Printf("ошибка при ENCODE JSON %v", err)
	}
}
