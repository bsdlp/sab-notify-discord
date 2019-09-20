package main

import "flag"

type parameters struct {
	notificationType       string
	notificationTitle      string
	notificationMessage    string
	notificationParameters string
}

func parseParameters() parameters {
	return parameters{
		notificationType:       flag.Arg(1),
		notificationTitle:      flag.Arg(2),
		notificationMessage:    flag.Arg(3),
		notificationParameters: flag.Arg(4),
	}
}
