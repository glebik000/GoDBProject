package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"GoDBProject/internal/application/rest/v1/handlers"
)

func RunServer() {
	router := mux.NewRouter()
	err := setHandler(router)
	if err != nil {
		fmt.Printf("ошибка при добавлении хендлеров %v", err)
		return
	}
	err = http.ListenAndServe(":80", router)
	if err != nil {
		fmt.Printf("ошибка при прослушивании запросов %v", err)
		return
	}

}

func setHandler(r *mux.Router) error {
	r.HandleFunc("/all-price", handlers.GetAllData).Methods("GET")
	// r.HandleFunc("/service/{id}/material-details", handlers.GetAllData).Methods("GET")
	r.HandleFunc("/insert-product-to-service", handlers.InsertProductToService).Methods("GET")
	return nil
}
