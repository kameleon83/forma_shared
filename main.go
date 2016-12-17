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

	controller.CreateDatabase()

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
