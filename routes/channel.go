package routes

import (
	"encoding/json"
	"log"
	"net/http"
)

type Channel struct {
	ChannelID   int    `json:"channel_id"`
	ChannelName string `json:"channel_name:"`
	AdminID     string `json:"admin_id"`
}

func (router *Router) SetChannelRoutes() {
	router.ChiMux.Get("/api/channels", router.getAllChannels)
}

func (router *Router) getAllChannels(w http.ResponseWriter, r *http.Request) {
	rows, err := router.Database.Query("SELECT channel_id, channel_name, admin_id FROM channel")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var channels []Channel
	for rows.Next() {
		var channel Channel

		err = rows.Scan(&channel.ChannelID, &channel.ChannelName, &channel.AdminID)
		if err != nil {
			log.Fatal(err)
		}

		channels = append(channels, channel)
	}

	jsonData, err := json.Marshal(channels)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(jsonData)
}
