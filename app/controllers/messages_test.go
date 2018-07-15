package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"os"
)

func TestMain(m *testing.M) {
	result := m.Run()
	os.Exit(result)
}

func TestMessages(t *testing.T) {
	body := "{\"update_id\":394352518,\"message\":{\"message_id\":5,\"from\":{\"id\":160589750,\"is_bot\":false,\"first_name\":\"Dmitry\",\"last_name\":\"Sadovnikov\",\"username\":\"DmitrySadovnikov\",\"language_code\":\"en-RU\"},\"chat\":{\"id\":160589750,\"first_name\":\"Dmitry\",\"last_name\":\"Sadovnikov\",\"username\":\"DmitrySadovnikov\",\"type\":\"private\"},\"date\":1527367962,\"text\":\"привет\"}}"
	r, err := http.NewRequest("POST", "messages", bytes.NewBufferString(body))
	w := httptest.NewRecorder()

	Router().ServeHTTP(w, r)

	if err != nil {
		t.Fatal(err)
	}

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	assert.Equal(t, 200, w.Code)
}
