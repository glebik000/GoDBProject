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
	r.HandleFunc("/service/{id}/material-details", handlers.GetAllDataBy).Methods("GET")
	r.HandleFunc("/insert-product-to-service", handlers.InsertProductToService).Methods("POST")
	r.HandleFunc("/update-product-price", handlers.UpdateProductPrice).Methods("PUT")
	r.HandleFunc("/update-service-price", handlers.UpdateServicePrice).Methods("PUT")
	r.HandleFunc("/update-product-hidden", handlers.UpdateProductHidden).Methods("PUT")
	r.HandleFunc("/update-service-hidden", handlers.UpdateServiceHidden).Methods("PUT")
	r.HandleFunc("/update-service-group-hidden", handlers.UpdateServiceGroupHidden).Methods("PUT")
	r.HandleFunc("/service", handlers.DeleteService).Methods("DELETE")
	return nil
}
