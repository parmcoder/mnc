package mnc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Base MNC base services
type Base interface {
	Broadcast(message *CreateBroadcast) (CreateBroadcastResponse, error)
	Monitor(message *GetTransaction) (GetTransactionResponse, error)
}

// base MNC service struct for default implementation
type base struct {
}

// CreateAPIConnector create MNC service instance
func CreateAPIConnector() Base {
	return base{}
}

// Broadcast will send transaction message to the `Node` url and receive transaciton hash.
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

	rawResponse, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		return
	}

	defer rawResponse.Body.Close()

	body, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}

	return response, nil
}

// Monitor will get transaction status from the `Node` url given transaction hash.
func (b base) Monitor(message *GetTransaction) (response GetTransactionResponse, err error) {
	if message == nil {
		err = ErrNoInput
		return
	}

	url := fmt.Sprintf("%s/check/%s", Node, message.TxHash)

	rawResponse, err := http.Get(url)
	if err != nil {
		return
	}

	defer rawResponse.Body.Close()

	body, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}

	return
}
