package app

import (
	"controller"
	"fmt"
	"net/http"
)

func StartApp() {

	fmt.Println("Starting the app....")

	mux := http.NewServeMux()

	mux.HandleFunc("/user", controller.GetUserController)

	http.ListenAndServe(":8080", mux)

}
