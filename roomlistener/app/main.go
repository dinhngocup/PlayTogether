package main

import (
	_client "PlayTogether/roomlistener"
	_room_listener "PlayTogether/roomlistener"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	flag.Parse()
	roomListener := _room_listener.NewListener()

	go roomListener.Run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hi ne")
		_client.ServeWs(w, r, roomListener)
	})

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
