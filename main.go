package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/gorilla/mux"

	"forma_shared/controller"
	"forma_shared/lib"
	"forma_shared/model"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// linux : go build -ldflags "-X main.buildstamp=`date '+%d-%m-%Y_%H:%M:%S'` -X main.githash=`git rev-parse HEAD` -X main.version=1.0.0-BETA"
// windows : go build -ldflags "-X main.buildstamp=$((get-date).tostring('d-MM-yyyy_H:mm:ss')) -X main.githash=$(git rev-parse HEAD) -X main.version=1.0.0-BETA "
// go build => Par défaut

var buildstamp, githash, version string

func versioning() {
	vers := flag.Bool("version", false, "prints current version")
	flag.BoolVar(vers, "v", false, "prints current version")
	flag.Parse()

	if *vers {
		fmt.Printf("Date last build : %s\n", buildstamp)
		fmt.Printf("Git Hash : %s\n", githash)
		fmt.Printf("Version : %s\n", version)
		os.Exit(0)
	}
}

func main() {
	versioning()

	f, err := os.OpenFile("forma_shared.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Println(err)
	}
	// don't forget to close it
	defer f.Close()

	// assign it to the standard logger
	log.SetOutput(f)

	model.ConnDB()

	// addrs, err := net.InterfaceAddrs()
	// if err != nil {
	// 	panic(err)
	// }
	// for i, addr := range addrs {
	// 	fmt.Printf("%d %v\n", i, addr)
	// }

	http.Get("/refresh")

	port := ":9000"
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controller.Index)

	router.HandleFunc("/download/{folder}/{file}", controller.Download)
	router.HandleFunc("/upload", controller.Upload)
	router.HandleFunc("/annuaire/{nameCol}/{sort:(?:asc|desc)}", controller.Annuaire)
	// J'ai remplacé : lib.RefreshListFilesAndFolder() par une route refresh
	router.HandleFunc("/refresh", controller.RefreshList)
	router.HandleFunc("/delete/{fileID}", lib.DeleteFile).Methods("POST")

	router.HandleFunc("/config", controller.Config).Methods("GET")
	router.HandleFunc("/config", controller.ConfigPost).Methods("POST")

	router.HandleFunc("/register", controller.Register)
	router.HandleFunc("/valid", controller.Valid)
	router.HandleFunc("/login", controller.Login)
	router.HandleFunc("/logout", controller.Logout)
	router.HandleFunc("/passwordForgot", controller.PasswordForgot)
	router.HandleFunc("/newPassword", controller.NewPassword)
	// router.HandleFunc("/autorized", lib.ClientAutorize)
	router.HandleFunc("/follow/{email}/{niveau}", controller.ChangeLevelByEmail).Methods("POST")
	router.HandleFunc("/checkpoint", controller.Checkpoint)
	router.HandleFunc("/checkpoint/{email}", controller.CheckpointUser).Methods("POST")
	router.HandleFunc("/checkpoint_reset", controller.CheckpointReset)
	router.HandleFunc("/random-user", controller.RandomUser)
	// router.HandleFunc("/notify", controller.InjectJavaScript)
	router.HandleFunc("/countfiles", controller.NewFileCheck)

	router.HandleFunc("/question&a", controller.QuestionAndA)
	router.HandleFunc("/q&answer", controller.QAndAnswer)
	router.HandleFunc("/q&anotif", lib.CheckNewQuestionAndAnswer)
	router.HandleFunc("/q&alike/{postID}/{like}", controller.QAndALike)

	router.HandleFunc("/stagiaire", controller.Stagiaire).Methods("GET")
	router.HandleFunc("/stagiaire", controller.StagiairePost).Methods("POST")
	router.HandleFunc("/pdf", controller.Pdf).Methods("GET")

	config := new(model.Config)
	router.PathPrefix("/files").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir(config.SendDirectory()))))
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server start : ", time.Now(), " to port "+port)

	fmt.Println(config.Count())
	// if config.Count() > 0 {
	// 	go lib.AutoStartBrowser("http://localhost" + port)
	// } else {
	// 	go lib.AutoStartBrowser("http://localhost" + port + "/config")
	// }

	// go http.ListenAndServeTLS(":9001", "cert.pem", "key.pem", router)
	// http.ListenAndServe(port, http.HandlerFunc(redirectToHttps))
	http.ListenAndServe(port, router)

}

func redirectToHttps(w http.ResponseWriter, r *http.Request) {
	// Redirect the incoming HTTP request. Note that "127.0.0.1:8081" will only work if you are accessing the server from your local machine.
	http.Redirect(w, r, "http://localhost:9001"+r.RequestURI, http.StatusMovedPermanently)
}
