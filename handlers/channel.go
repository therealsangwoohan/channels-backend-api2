package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/therealsangwoohan/channels-backend-api2/database"
	"github.com/therealsangwoohan/channels-backend-api2/models"
)

func GetAllChannels(w http.ResponseWriter, r *http.Request) {
	rows, err := database.Database.Query("SELECT channel_id, channel_name, admin_id FROM channel")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var channels []models.Channel
	for rows.Next() {
		var channel models.Channel

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

func GetChannel(w http.ResponseWriter, r *http.Request) {
	channelID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid channel ID", http.StatusBadRequest)
		return
	}

	row := database.Database.QueryRow("SELECT channel_id, channel_name, admin_id FROM channel WHERE channel_id=$1", channelID)
	var channel models.Channel
	err = row.Scan(&channel.ChannelID, &channel.ChannelName, &channel.ChannelName)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}

	jsonData, err := json.Marshal(channel)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonData)
}

func CreateChannel(w http.ResponseWriter, r *http.Request) {
	var channel models.Channel
	err := json.NewDecoder(r.Body).Decode(&channel)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	queryString := "INSERT INTO channel (channel_name, admin_id) VALUES ($1, $2) RETURNING channel_id"
	row := database.Database.QueryRow(queryString, channel.ChannelName, channel.AdminID)
	var channelID string
	err = row.Scan(&channelID)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
	fmt.Println(channelID)
}
