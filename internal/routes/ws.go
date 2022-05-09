package routes

import (
	"election-service/pkg/ws"

	"github.com/antoniodipinto/ikisocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func WebSocket(app *fiber.App, hub *ws.Hub) {
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}

		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/election/:id", ikisocket.New(func(kws *ikisocket.Websocket) {
		userId := kws.Params("id")
		hub.Election[userId] = kws.UUID

		kws.SetAttribute("user_id", userId)
	}))

	app.Get("/ws/vote/:id", ikisocket.New(func(kws *ikisocket.Websocket) {
		userId := kws.Params("id")
		hub.Vote[userId] = kws.UUID

		kws.SetAttribute("user_id", userId)
	}))

	InitEvent(hub)
}

func InitEvent(hub *ws.Hub) {
	ikisocket.On(ikisocket.EventDisconnect, func(ep *ikisocket.EventPayload) {
		uid := ep.Kws.GetStringAttribute("user_id")

		if hub.Election[uid] != "" {
			delete(hub.Election, uid)
		}
		if hub.Vote[uid] != "" {
			delete(hub.Vote, uid)
		}
	})

	ikisocket.On(ikisocket.EventClose, func(ep *ikisocket.EventPayload) {
		uid := ep.Kws.GetStringAttribute("user_id")

		if hub.Election[uid] != "" {
			delete(hub.Election, uid)
		}
		if hub.Vote[uid] != "" {
			delete(hub.Vote, uid)
		}
	})

	ikisocket.On(ws.ElectionEvent, func(ep *ikisocket.EventPayload) {
		var list []string
		for _, uuid := range hub.Election {
			list = append(list, uuid)
		}
		ep.Kws.EmitToList(list, ep.Data)
	})

	ikisocket.On(ws.VoteEvent, func(ep *ikisocket.EventPayload) {
		var list []string
		for _, uuid := range hub.Vote {
			list = append(list, uuid)
		}
		ep.Kws.EmitToList(list, ep.Data)
	})

}
