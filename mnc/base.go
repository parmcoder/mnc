package mnc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Base interface {
	Broadcast(message *CreateBroadcast) (CreateBroadcastResponse, error)
	Monitor(message *GetTransaction) (GetTransactionResponse, error)
}

type base struct {
}

func CreateAPIConnector() Base {
	return base{}
}

func (b base) Broadcast(message *CreateBroadcast) (response CreateBroadcastResponse, err error) {
	if message == nil {
		err = ErrNoInput

		return response, err
	}

	url := fmt.Sprintf("%s/broadcast", Node)

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

func (b base) Monitor(message *GetTransaction) (response GetTransactionResponse, err error) {
	if message == nil {
		err = ErrNoInput
		return
	}

	url := fmt.Sprintf("%s/check/%s", Node, message.TxHash)

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
