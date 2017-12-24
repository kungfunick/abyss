package engine

import (
    "time"

    "github.com/teris-io/shortid"
)

type (
    // Event contains a platform event details
    Event struct {
        ID        string      `json:"id"`
        Name      string      `json:"name"`
        Data      interface{} `json:"data"`
        Inbound   bool        `json:"inbound"`
        ClientId  string      `json:"clientId"`
        CreatedAt time.Time   `json:"createdAt"`
    }

    // EventDefinition stores the definition of an event
    // used for documentation
    EventDefinition struct {
        Name        string `json:"name"`
        Description string `json:"description"`
        Broadcast   bool   `json:"broadcast"`
        Websocket   bool   `json:"websocket"`
        Internal    bool   `json:"internal"`
    }

    // EventHandlers are used to process events
    EventHandler func(e Event) bool
)

func NewEvent(n string, d interface{}) Event {
    id, _ := shortid.Generate()

    return Event{
        ID:        id,
        Name:      n,
        Data:      d,
        Inbound:   false,
        CreatedAt: time.Now(),
    }
}
