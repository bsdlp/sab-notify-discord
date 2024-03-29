package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"text/template"
)

func notify(params parameters) error {
	if params.WebhookURL == "" {
		return errors.New("webhook url is a required parameter")
	}

	msg, err := formatMessage(params)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(map[string]string{"content": msg})
	if err != nil {
		return err
	}

	_, err = http.DefaultClient.Post(params.WebhookURL, "application/json", &buf)
	return err
}

func formatMessage(params parameters) (string, error) {
	const msgFormat = `__{{.NotificationType}}__
{{.NotificationTitle}}
{{.NotificationMessage}}`

	tmpl, err := template.New("msg").Parse(msgFormat)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, params)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
