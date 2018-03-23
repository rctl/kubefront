package core

import (
	"fmt"

	"github.com/gorilla/websocket"
)

//Upstream is an upstream to a user session
type Upstream struct {
	User       string
	Session    string
	Connection *websocket.Conn
}

//Message is used in websocket communication
type Message struct {
	Action string `json:"action"`
	Data   string `json:"data"`
}

//NotifyAll sends a message to all online users
func (c *Context) NotifyAll(m *Message) error {
	for _, u := range c.Upstreams {
		for _, s := range u {
			err := s.Connection.WriteJSON(m)
			if err != nil {
				fmt.Printf("Failed to send to websocket: %s\n", err.Error())
				return err
			}
		}
	}
	return nil
}

//NotifyUser sends a message to a specific online user
func (c *Context) NotifyUser(username string, m *Message) error {
	u, e := c.Upstreams[username]
	if e {
		for _, s := range u {
			err := s.Connection.WriteJSON(m)
			if err != nil {
				fmt.Printf("Failed to send to websocket: %s\n", err.Error())
				return err
			}
		}
	}
	return nil
}

//NotifySession sends a message to a specific active session
func (c *Context) NotifySession(username, session string, m *Message) error {
	u, e := c.Upstreams[username]
	if e {
		s, e := u[session]
		if e {
			err := s.Connection.WriteJSON(m)
			if err != nil {
				fmt.Printf("Failed to send to websocket: %s\n", err.Error())
				return err
			}
		}
	}
	return nil
}
