package demo

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleHttp(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("Client connected:", conn.LocalAddr())
	reader(conn)
}

func reader(conn *websocket.Conn) {

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		fmt.Println("Server:" + string(msg))

		conn.WriteMessage(msgType, msg)
	}
}

func Client() {
	time.Sleep(10 * time.Second)
	fmt.Println("Client is connecting....")
	socketUrl := "ws://localhost:8080" + "/"
	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		log.Fatal("Error connecting to Websocket Server:", err)
	}
	//defer conn.Close()
	for {
		var st string
		fmt.Scanln(&st)
		if st == "quit" {
			conn.Close()
		}
		conn.WriteMessage(1, []byte(st))

		_, msg, _ := conn.ReadMessage()
		fmt.Println("Client:", string(msg))
	}

}

func WebsocketDemo() {
	go Client()
	http.HandleFunc("/", handleHttp)
	http.ListenAndServe(":8080", nil)

}
