package test

import (
	"flag"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"testing"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

// use default options, update http to websocket
/*  gorilla/websocket cross-origin issue ref:
	https://pkg.go.dev/github.com/gorilla/websocket#section-readme
	https://github.com/gorilla/websocket/issues/367
*/
var upgrader = websocket.Upgrader{
	// same to line 31, https://blog.csdn.net/taoshihan/article/details/109214077
	CheckOrigin:func(r *http.Request) bool { return true },
}    // Cross-Origin setting

// 07 add var
var ws = make(map[*websocket.Conn]struct{})

func echo(w http.ResponseWriter, r *http.Request) {

	//upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// c 为 websocket定义的connection
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	ws[c] = struct{}{} // ws add connection
	for {
		mt, message, err := c.ReadMessage() // 接收到一个消息后要给所有人发
		if err != nil {
			log.Println("read: ", err)
			break
		}
		log.Printf("[D] recv: %s", message)

		// loop
		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read: ", err)
				break
			}
			log.Printf("[D] recv: %s", message)

			for conn := range ws {
				err = conn.WriteMessage(mt, message)

				if err != nil {
					log.Println("[D] ws write: ", err)
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

// 注意会有跨域问题，https://blog.csdn.net/taoshihan/article/details/109214077
// 或者用官方插件全局解决跨域问题
// 看了下其实是 github.com/gorilla/websocket 的问题，见上方注释, https://blog.csdn.net/taoshihan/article/details/109214077
func TestGinWebsocketServer(t *testing.T) {
	r := gin.Default()
	//r.Use(cors.Default())  // 全局跨域

	fmt.Println("[*] current mode: ", gin.Mode())
	if gin.Mode() == "debug" {
		r.Use(cors.Default()) // 在 debug 模式下, 允许跨域访问
	}

	// 路由处理
	r.GET("/echo", func(ctx *gin.Context) {
		echo(ctx.Writer, ctx.Request)
	})
	r.Run(":8080")

}
