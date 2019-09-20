package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"text/template"
)

func notify(webhookURL string, params parameters) error {
	msg, err := formatMessage(params)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(map[string]string{"content": msg})
	if err != nil {
		return err
	}

	_, err = http.DefaultClient.Post(webhookURL, "application/json", &buf)
	return err
}

func formatMessage(params parameters) (string, error) {
	const msgFormat = `__{{.notificationType}}__
{{.notificationTitle}}
{{.notificationMessage}}`

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
