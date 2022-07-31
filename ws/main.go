package main

import (
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"
)

var u = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Barrage() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// 建立连接
		conn, err := u.Upgrade(writer, request, nil)
		if err != nil {
			zap.S().Error("websocket can not upgrade", err)
			return
		}
		// 关闭连接
		defer func(conn *websocket.Conn) {
			_ = conn.Close()
		}(conn)

		// 防止退出
		ch := make(chan int)

		if err := conn.WriteMessage(websocket.TextMessage, []byte("连接成功")); err != nil {
			zap.S().Error("发送消息失败", err)
			ch <- 1
		}

		// 读取消息
		go func() {
			for {
				_, data, err := conn.ReadMessage()
				if err != nil {
					zap.S().Error("读取消息失败", err)
					ch <- 1
				}
				if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
					zap.S().Error("发送消息失败", err)
					ch <- 1
				}
			}
		}()
		<-ch
	}
}

func main() {
	http.HandleFunc("/test", Barrage())
	_ = http.ListenAndServe("0.0.0.0:9501", nil)
}
