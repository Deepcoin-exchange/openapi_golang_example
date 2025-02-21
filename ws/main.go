package main

import (
	"log"

	"github.com/gorilla/websocket"
)

type WSClient struct {
	conn *websocket.Conn
}

func NewWSClient() (*WSClient, error) {
	conn, _, err := websocket.DefaultDialer.Dial("wss://test-wss.goodtest.cc/v1/private?listenKey=3f8021a44a262e69344d5c522b613006", nil)
	if err != nil {
		return nil, err
	}

	client := &WSClient{
		conn: conn,
	}

	go client.handleMessages()

	return client, nil
}

func (c *WSClient) handleMessages() {
	log.Println("start...")
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Printf("read message error: %v", err)
			return
		}
		log.Printf("recv message: %s", message)
	}
}

func main() {
	_, err := NewWSClient()
	if err != nil {
		log.Fatal(err)
	}
	select {}
}
