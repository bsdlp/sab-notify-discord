package main

import (
	"errors"
	"os"
)

type parameters struct {
	NotificationType    string
	NotificationTitle   string
	NotificationMessage string
	WebhookURL          string
}

func parseParameters() (parameters, error) {
	if len(os.Args) != 5 {
		return parameters{}, errors.New("must provide all required arguments")
	}

	return parameters{
		NotificationType:    os.Args[1],
		NotificationTitle:   os.Args[2],
		NotificationMessage: os.Args[3],
		WebhookURL:          os.Args[4],
	}, nil
}
