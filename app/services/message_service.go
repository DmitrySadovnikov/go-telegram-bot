package services

import (
	"strings"
)

type MessageService struct{}

type Result struct {
	Message string `json:"message"`
}

func (_ MessageService) Call(message string) Result {
	message = strings.TrimPrefix(message, "/")
	return Result{Message: message}
}
