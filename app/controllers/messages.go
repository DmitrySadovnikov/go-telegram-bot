package controllers

import (
	"encoding/json"
	"go-telegram-bot/app/services"
	"net/http"
)

//	{
//	 "update_id": 394352518,
//	 "message": {
//	   "message_id": 5,
//	   "from": {
//	     "id": 123456789,
//	     "is_bot": false,
//	     "first_name": "Dmitry",
//	     "last_name": "Sadovnikov",
//	     "username": "DmitrySadovnikov",
//	     "language_code": "en-RU"
//	   },
//	   "chat": {
//	     "id": 123456789,
//	     "first_name": "Dmitry",
//	     "last_name": "Sadovnikov",
//	     "username": "DmitrySadovnikov",
//	     "type": "private"
//	   },
//	   "date": 1527367962,
//	   "text": "hi"
//	 }
//	}
type TelegramRequest struct {
	UpdateId int `json:"update_id"`
	Message  struct {
		MessageId int    `json:"message_id"`
		Date      int    `json:"date"`
		Text      string `json:"text"`
		From      struct {
			Id           int    `json:"id"`
			IsBot        bool   `json:"is_bot"`
			FirstName    string `json:"first_name"`
			LastName     string `json:"last_name"`
			Username     string `json:"username"`
			LanguageCode string `json:"language_code"`
		} `json:"from"`
		Chat struct {
			Id        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
			Type      string `json:"type"`
		} `json:"chat"`
	} `json:"message"`
}

func Messages(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var tr TelegramRequest
	err := decoder.Decode(&tr)

	if err != nil {
		panic(err)
	}

	message := tr.Message
	service := services.MessageService{}
	result := service.Call(message.Text)
	js, err := json.Marshal(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(js)
}
