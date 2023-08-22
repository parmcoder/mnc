package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/parmcoder/mnc"
)

type Base interface {
	Broadcast(message *mnc.CreateBroadcast) (mnc.CreateBroadcastResponse, error)
	Monitor(message *mnc.GetTransaction) (mnc.GetTransactionResponse, error)
}

type BaseImpl struct {
}

func CreateAPIConnector() Base {
	return BaseImpl{}
}

func (b BaseImpl) Broadcast(message *mnc.CreateBroadcast) (response mnc.CreateBroadcastResponse, err error) {
	if message == nil {
		err = mnc.ErrNoInput

		return response, err
	}

	url := fmt.Sprintf("%s/broadcast", mnc.Node)

	postBody, err := json.Marshal(*message)
	if err != nil {
		return
	}

	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}

	return response, nil
}

func (b BaseImpl) Monitor(message *mnc.GetTransaction) (response mnc.GetTransactionResponse, err error) {
	if message == nil {
		err = mnc.ErrNoInput
		return
	}

	url := fmt.Sprintf("%s/check/%s", mnc.Node, message.TxHash)

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}

	return
}
