package main

import (
	"fmt"
	"github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

func main() {
	server := socketio.NewServer(nil)



	//sockets
	server.OnConnect("/", func(so socketio.Conn) error {
		so.SetContext("")
		so.Join("chat_room")
		fmt.Println("New User Connected", so.ID())
		return nil
	})

	server.OnEvent("/", "chat-message", func(so socketio.Conn, msg string){
		fmt.Println(msg)
		// serve.emit("chat-message")
		server.BroadcastToRoom("/","chat_room", "chat-message", msg)

	})

	go server.Serve()
	defer server.Close()

	//Modulo Http
	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Println("Server on Port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}