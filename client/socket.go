package client

import (
	"encoding/json"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/srepio/sdk/types"
)

type ws struct {
	conn *websocket.Conn
	mu   *sync.Mutex
}

func newWs(conn *websocket.Conn) *ws {
	return &ws{
		conn: conn,
		mu:   &sync.Mutex{},
	}
}

func (ws *ws) Read() (*types.TerminalMessage, error) {
	_, raw, err := ws.conn.ReadMessage()
	if err != nil {
		return nil, err
	}
	msg := &types.TerminalMessage{}
	if err := json.Unmarshal(raw, msg); err != nil {
		return nil, err
	}
	return msg, nil
}

func (ws *ws) Write(msg *types.TerminalMessage) error {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	return ws.conn.WriteJSON(msg)
}
