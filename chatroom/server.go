package chatroom

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"github.com/kataras/iris"
)
var clients = make(map[*websocket.Conn]string)
var broadcast = make(chan Message)

var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
    Message string `json:"message"`
    Uid string `json:"uid"`
}

//注册成为 websocket
func HandleConnections(request iris.Context) {
	w := request.ResponseWriter()
	r := request.Request()
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
			log.Fatal(err)
	}
	defer ws.Close()
	uid, _ := uuid.NewV4()
	id := uid.String()
	clients[ws] = id

	//不断的从页面上获取数据 然后广播发送出去
	for {
			// 将从页面上接收数据改为不接收 直接发送
			var msg Message
			err := ws.ReadJSON(&msg)
			msg.Uid = id
			if err != nil {
					log.Printf("error: %v", err)
					delete(clients, ws)
					break
			}
			broadcast <- msg
	}
}

//广播发送至页面
func HandleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
		if clients[client] != msg.Uid {
			err := client.WriteJSON(msg)
				if err != nil {
					log.Printf("client.WriteJSON error: %v", err)
					client.Close()
					delete(clients, client)
				}
			}
		}
	}
}