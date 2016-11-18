package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"forma_shared/controller/view"
)

// DIRFILE constante link download and upload file

func main() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for i, addr := range addrs {
		fmt.Printf("%d %v\n", i, addr)
	}

	port := ":9000"
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controllerView.Index)
	router.HandleFunc("/download/{folder}/{file}", controllerView.Download)
	router.HandleFunc("/upload", controllerView.Upload)
	router.HandleFunc("/not_access", controllerView.NotAccess)
	router.HandleFunc("/annuaire", controllerView.Annuaire)

	router.PathPrefix("/files").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir("files"))))
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server start : ", time.Now(), " to port 9000")

	http.ListenAndServe(port, router)
}
