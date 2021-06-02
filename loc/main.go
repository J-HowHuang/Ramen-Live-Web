package main

import (
	// "context"
	"fmt"
	"log"
	"net/http"
	// "os"
	// "time"

	"github.com/J-HowHuang/Ramen-Live/loc/pkg/websocket"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	fmt.Println("start")
	log.Println("start")
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Println(err)
	}
	// go websocket.Writer(ws)
	websocket.Reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/ws", serveWs)
}

func main() {
	// db_url := os.Getenv("DB_URL")
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI(db_url))

	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }()
	// if err == nil {
	// 	err := client.Ping(ctx, nil)
	// 	if err == nil {
	// 		log.Println("DB connected at " + db_url)
	// 	} else {
	// 		log.Println("Failed to connect DB at " + db_url)
	// 	}
	// }

	// db.InitDB(client)
	setupRoutes()
	http.ListenAndServe(":18763", nil)

}