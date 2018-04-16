package kubefront

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/rctl/kubefront/src/backend/kubefront/core"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (s *Server) upstreamMessageHandler(c *core.Upstream, m core.Message) *core.Message {
	//Find upstream
	if m.Action == "SUBSCRIBE" {
		c.Subscriptions[m.Entity] = true
		c.Connection.WriteJSON(&core.Message{
			Action: "SUBSCRIPTIONS_UPDATED",
			Data:   c.Subscriptions,
		})
	}
	if m.Action == "UNSUBSCRIBE" {
		c.Subscriptions[m.Entity] = false
		c.Connection.WriteJSON(&core.Message{
			Action: "SUBSCRIPTIOS_UPDATED",
			Data:   c.Subscriptions,
		})
	}
	return nil
}

func (s *Server) upstreamHandler(c *gin.Context) {
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: ", err.Error())
		return
	}

	u := c.MustGet("username").(string)
	id := c.MustGet("session").(string)

	if s.Upstreams == nil {
		s.Upstreams = make(map[string]map[string]*core.Upstream)
	}

	//Add connection to connection tracker
	_, exists := s.Upstreams[u]
	if !exists {
		s.Upstreams[u] = make(map[string]*core.Upstream)
	}
	ustream := &core.Upstream{
		User:          u,
		Session:       id,
		Connection:    conn,
		Subscriptions: make(map[string]bool),
	}
	s.Upstreams[u][id] = ustream

	conn.SetCloseHandler(func(code int, text string) error {
		//Remove connection tracker
		delete(s.Upstreams[u], id)
		return nil
	})

	//Setup handler for incoming messages
	for {
		var m core.Message
		err := conn.ReadJSON(&m)
		if err != nil {
			fmt.Printf("Failed to decode websocket message: %s\n", err.Error())
			break
		}
		r := s.upstreamMessageHandler(ustream, m)
		if r != nil {
			fmt.Printf("Failed to encode websocket message: %s\n", err.Error())
			conn.WriteJSON(r)
		}
	}
}
