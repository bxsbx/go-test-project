package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	}, //CheckOrigin防止跨站点的请求伪造
}

type Conn struct {
	ConnMap map[string]*websocket.Conn
	sync.Mutex
}

var WsConn Conn

type UserParams struct {
	UserId string `form:"user_id" binding:"required"`
}

func WebSocket(c *gin.Context) {
	var params UserParams
	if err := c.ShouldBind(&params); err != nil {
		c.Error(err)
		return
	}
	conn, err := Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.Error(err)
		return
	}
	WsConn.Lock()
	WsConn.ConnMap[params.UserId] = conn
	WsConn.Unlock()
	defer func() {
		conn.Close() //返回前关闭
		WsConn.Lock()
		delete(WsConn.ConnMap, params.UserId)
		WsConn.Unlock()
	}()
	for {
		//读取数据
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			break
		}
		if len(message) <= 0 {
			continue
		}
		//写入数据
		err = conn.WriteMessage(messageType, []byte("233929"+string(message)))
		if err != nil {
			break
		}
	}
}

//// 不可有多个读
//func Read1(c *gin.Context) {
//	var params UserParams
//	if err := c.ShouldBind(&params); err != nil {
//		c.Error(err)
//		return
//	}
//	if conn, ok := WsConn.ConnMap[params.UserId]; ok {
//		//for {
//		messageType, message, err := conn.ReadMessage()
//		if err != nil {
//			c.Error(err)
//		}
//		err = conn.WriteMessage(messageType, []byte("read1"+string(message)))
//		if err != nil {
//			c.Error(err)
//		}
//		//}
//
//		//c.JSON(http.StatusOK, strconv.Itoa(messageType)+string(message))
//	}
//}
//
//func Read2(c *gin.Context) {
//	var params UserParams
//	if err := c.ShouldBind(&params); err != nil {
//		c.Error(err)
//		return
//	}
//	if conn, ok := WsConn.ConnMap[params.UserId]; ok {
//		//for {
//		messageType, message, err := conn.ReadMessage()
//		if err != nil {
//			c.Error(err)
//		}
//		err = conn.WriteMessage(messageType, []byte("read2"+string(message)))
//		if err != nil {
//			c.Error(err)
//		}
//		//}
//		//c.JSON(http.StatusOK, strconv.Itoa(messageType)+string(message))
//	}
//}

func Send1(c *gin.Context) {
	var params UserParams
	if err := c.ShouldBind(&params); err != nil {
		c.Error(err)
		return
	}
	if conn, ok := WsConn.ConnMap[params.UserId]; ok {
		err := conn.WriteMessage(websocket.TextMessage, []byte("Send1"))
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, "发送成功Send1")
	}
}

func Send2(c *gin.Context) {
	var params UserParams
	if err := c.ShouldBind(&params); err != nil {
		c.Error(err)
		return
	}
	if conn, ok := WsConn.ConnMap[params.UserId]; ok {
		err := conn.WriteMessage(websocket.TextMessage, []byte("Send2"))
		if err != nil {
			c.Error(err)
			return
		}
		fmt.Println("Send2")
		c.JSON(http.StatusOK, "发送成功Send2")
	}
}

// 不能同时有多个读等待，否则会因为一个读会影响另一个读的位置导致通信标志出错（读位置是共享的）（建议只允许一个读），但可以有多个写
func main() {
	WsConn.ConnMap = make(map[string]*websocket.Conn)
	app := gin.Default()

	app.GET("/websocket", WebSocket)

	//app.GET("/read1", Read1)
	//app.GET("/read2", Read2)
	app.GET("/send1", Send1)
	app.GET("/send2", Send2)

	app.Run(":8888")
}
