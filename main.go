package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"forma_shared/controller"
	"forma_shared/controller/view"
)

// DIRFILE constante link download and upload file

func main() {

	controller.CreateDatabase()

	controller.ReadDir()

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for i, addr := range addrs {
		fmt.Printf("%d %v\n", i, addr)
	}

	http.Get("/refresh")

	port := ":9000"
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controllerView.Index)
	router.HandleFunc("/download/{folder}/{file}", controllerView.Download)
	router.HandleFunc("/upload", controllerView.Upload)
	router.HandleFunc("/annuaire/{nameCol}/{sort:(?:asc|desc)}", controllerView.Annuaire)
	// J'ai remplacé : controller.RefreshListFilesAndFolder() par une route refresh
	router.HandleFunc("/refresh", controllerView.RefreshList)
	router.HandleFunc("/register", controllerView.Register)
	router.HandleFunc("/valid", controllerView.Valid)
	router.HandleFunc("/login", controllerView.Login)
	router.HandleFunc("/logout", controllerView.Logout)
	router.HandleFunc("/passwordForgot", controllerView.PasswordForgot)
	router.HandleFunc("/newPassword", controllerView.NewPassword)
	// router.HandleFunc("/autorized", controller.ClientAutorize)
	router.HandleFunc("/follow/{user}/{niveau}", controllerView.ChangeLevelByName)
	router.HandleFunc("/followed_reset", controllerView.ResetFollow)
	router.HandleFunc("/checkpoint", controllerView.ProblemFollowed)
	// router.HandleFunc("/notify", controllerView.InjectJavaScript)

	router.PathPrefix("/files").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir(controller.DIRFILE))))
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server start : ", time.Now(), " to port 9000")

	// http.ListenAndServeTLS(port, "cert.pem", "key.pem", router)
	http.ListenAndServe(port, router)

}
