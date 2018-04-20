package core

import (
	"context"

	uuid "github.com/satori/go.uuid"
)

//RunWorker starts a background job and notifies all session from the user when completed
func (c *Context) RunWorker(user string, entity string, job func(ctx context.Context) (interface{}, error)) (string, error) {
	id := uuid.Must(uuid.NewV4()).String()
	c.NotifyUser(user, &Message{
		Action: "JOB_STARTED",
		Entity: id,
		Data:   entity,
	})
	_, exists := c.Workers[user]
	if !exists {
		c.Workers[user] = make(map[string]*Worker)
	}
	c.Workers[user][id] = &Worker{
		ID:     id,
		Entity: entity,
	}
	go func() {
		data, err := job(c)
		if err != nil {
			c.NotifyUser(user, &Message{
				Action: "JOB_FAILED",
				Entity: id,
				Data:   err.Error(),
			})
		} else {
			c.NotifyUser(user, &Message{
				Action: "JOB_COMPLETED",
				Entity: id,
				Data:   data,
			})
		}
		delete(c.Workers[user], id)
	}()
	return id, nil
}
