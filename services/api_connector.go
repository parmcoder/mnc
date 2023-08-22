package services

import "github.com/parmcoder/transaction-tracker/domains"

func Broadcast(*message domains.CreateBroadcast) {
	// TODO: send POST request
}

func Monitor(*hash string) {
	// TODO: send GET request
}