package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"forma_shared/controller"
	"forma_shared/controller/view"
	"forma_shared/model"
)

// DIRFILE constante link download and upload file

func main() {

	controller.ReadDir()

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for i, addr := range addrs {
		fmt.Printf("%d %v\n", i, addr)
	}

	http.Get("/refresh")

	model.ReadUserJSON(false, "lastname")

	port := ":9000"
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controllerView.Index)
	router.HandleFunc("/download/{folder}/{file}", controllerView.Download)
	router.HandleFunc("/upload", controllerView.Upload)
	router.HandleFunc("/not_access", controllerView.NotAccess)
	router.HandleFunc("/annuaire/{reverse}/{col}", controllerView.Annuaire)
	// J'ai remplacé : controller.RefreshListFilesAndFolder() par une route refresh
	router.HandleFunc("/refresh", controllerView.RefreshList)
	// router.HandleFunc("/autorized", controller.ClientAutorize)

	router.PathPrefix("/files").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir(controller.DIRFILE))))
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server start : ", time.Now(), " to port 9000")

	http.ListenAndServe(port, router)

}
