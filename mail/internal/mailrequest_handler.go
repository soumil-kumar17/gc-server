package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/p-society/gc-server/mail/internal/models"
)

func MailRequestHandler(w http.ResponseWriter, r *http.Request) {
	var paramInstance models.MailingParams

	senderName := os.Getenv("EMAIL_SENDER_NAME")
	senderAddress := os.Getenv("EMAIL_SENDER_ADDRESS")
	senderPassword := os.Getenv("EMAIL_SENDER_PASSWORD")

	sender := NewGmailSender(senderName, senderAddress, senderPassword)

	err := json.NewDecoder(r.Body).Decode(&paramInstance)

	if err != nil {
		jsonPayload := `{"err": "` + err.Error() + `"}`
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(jsonPayload)
		return
	}

	fmt.Println(paramInstance)

	err = sender.SendEmail(paramInstance.Subject, paramInstance.Content, paramInstance.To, nil, nil, nil)

	if err != nil {
		jsonPayload := `{"err": "` + err.Error() + `"}`
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(jsonPayload)
		return
	}

	fmt.Println("Sucess")
	w.Write([]byte("E-mail sent Successfully."))
}