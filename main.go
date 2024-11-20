package main

import (
	"be-mklinik/database"
	"be-mklinik/handlers"
	"be-mklinik/middlewares"
	"be-mklinik/websockets"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	// Memuat variabel lingkungan dari file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {

	db, err := database.ConnectPostgreSQL()
	if err != nil {
		log.Fatal("Connection ", err.Error())
	}

	fmt.Println(db)
	mux := mux.NewRouter()
	v1 := mux.PathPrefix("/v1").Subrouter()

	v1.HandleFunc("/check-version", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("check-version " + os.Getenv("VERSION")))
	}).Methods("GET")

	pageRequestHandler := handlers.NewPageTokenHandler()
	v1.HandleFunc("/request-page-token", pageRequestHandler.RequestPageTokenHandler).Methods("GET")

	pageMonitorAntrianHandler := handlers.NewPageMonitorAntrianHandler(db)
	v1.HandleFunc("/monitor-antrian", pageMonitorAntrianHandler.MonitorAntrianHandler).Methods("GET")

	wsPanggilAntrianHandler := websockets.NewPanggilAntrianHandler(db)
	v1.HandleFunc("/panggil-antrian", wsPanggilAntrianHandler.PanggilAntrian).Methods("POST")

	loketHandler := handlers.NewLoketHandler(db)
	v1.HandleFunc("/loket", loketHandler.LoketHandler).Methods("GET")

	v1.HandleFunc("/ws-antrian", websockets.WsAntrian)

	var handler http.Handler = mux
	// handler = middleware.MiddlewareAuth(handler)
	handler = middlewares.CheckLoginToken(handler)
	handler = middlewares.CheckPageToken(handler)
	handler = middlewares.EnableCORS(handler)

	server := new(http.Server)
	server.Addr = ":3001"
	server.Handler = handler
	// server.Handler = mux
	fmt.Println("Server is running on port 3001")
	server.ListenAndServe()

	//sample websocket

	// sample := mux.PathPrefix("/sample").Subrouter()
	// sample.HandleFunc("/ws-sample", websockets.WsSample)
	// sample.HandleFunc("/ws-chat", websockets.WsChat)
	// sample.HandleFunc("/ws-listen", websockets.ListenCall)
	// sample.HandleFunc("/ws-call", websockets.CallAntrian).Methods("POST")

	// go websockets.HandleChatMessages()

}
