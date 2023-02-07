package test

import (
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"testing"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options, update http to websocket  todo

// 07 add var
var ws = make(map[*websocket.Conn]struct{})

func echo(w http.ResponseWriter, r *http.Request) {
	// c 为 websocket定义的connection
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	ws[c] = struct{}{}   // ws add connection
	for {
		mt, message, err := c.ReadMessage()   // 接收到一个消息后要给所有人发
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)

		// loop
		for {
			mt, message, err := c.ReadMessage()
			if err != nil{
				log.Println("read: ",err)
				break
			}
			log.Printf("[D recv: %s", message)

			for conn := range ws{
				err = conn.WriteMessage(mt, message)

				if err != nil{
					log.Println("[D] write: ", err)
					break
				}
			}
		}

		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func TestWebsocketServer(t *testing.T) {
	// websocket的服务端， websocket的客户端用postman
	// 在test中调用方法
	//flag.Parse()
	//log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	//http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(*addr, nil))

}

func TestGinWebsocketServer(t *testing.T) {
	
	
}
