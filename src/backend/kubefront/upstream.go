package kubefront

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/rctl/kubefront/src/backend/kubefront/core"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (s *Server) upstreamMessageHandler(c *gin.Context, m core.Message) *core.Message {
	return nil
}

func (s *Server) upstreamHandler(c *gin.Context) {
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %s", err.Error())
		return
	}

	u := c.MustGet("username").(string)
	id := c.MustGet("session").(string)

	//Add connection to connection tracker
	_, exists := s.Upstreams[u]
	if !exists {
		s.Upstreams[u] = make(map[string]*core.Upstream)
	}
	s.Upstreams[u][id] = &core.Upstream{
		User:       u,
		Session:    id,
		Connection: conn,
	}

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
			fmt.Println("Failed to decode websocket message: %s", err.Error())
			break
		}
		r := s.upstreamMessageHandler(c, m)
		if r != nil {
			fmt.Println("Failed to encode websocket message: %s", err.Error())
			conn.WriteJSON(r)
		}
	}
}
