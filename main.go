package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/gorilla/mux"

	"forma_shared/controller"
	"forma_shared/controllerView"
	"forma_shared/model"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {

	f, err := os.OpenFile("forma_shared.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Println(err)
	}
	// don't forget to close it
	defer f.Close()

	// assign it to the standard logger
	log.SetOutput(f)

	var wg sync.WaitGroup
	model.ConnDB()

	wg.Add(1)
	go controller.Config(&wg)
	wg.Wait()

	// addrs, err := net.InterfaceAddrs()
	// if err != nil {
	// 	panic(err)
	// }
	// for i, addr := range addrs {
	// 	fmt.Printf("%d %v\n", i, addr)
	// }

	http.Get("/refresh")

	port := ":9001"

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
	router.HandleFunc("/checkpoint", controllerView.Checkpoint)
	router.HandleFunc("/checkpoint_reset", controllerView.CheckpointReset)
	// router.HandleFunc("/notify", controllerView.InjectJavaScript)
	router.HandleFunc("/countfiles", controllerView.NewFileCheck)

	router.HandleFunc("/question&a", controllerView.QuestionAndA)
	router.HandleFunc("/q&answer", controllerView.QAndAnswer)

	router.PathPrefix("/files").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir(controller.DIRFILE))))
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server start : ", time.Now(), " to port "+port)

	// http.ListenAndServeTLS(port, "server.crt", "server.key", router)
	// go http.ListenAndServeTLS(":9001", "cert.pem", "key.pem", router)
	http.ListenAndServe(port, router)

}
