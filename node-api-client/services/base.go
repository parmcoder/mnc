package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/parmcoder/mnc/configs"
	"github.com/parmcoder/mnc/models"
)

type Base interface {
	Broadcast(message *models.CreateBroadcast) (models.CreateBroadcastResponse, error)
	Monitor(message *models.GetTransaction) (models.GetTransactionResponse, error)
}

type BaseImpl struct {
}

func CreateAPIConnector() Base {
	return BaseImpl{}
}

func (b BaseImpl) Broadcast(message *models.CreateBroadcast) (response models.CreateBroadcastResponse, err error) {
	if message == nil {
		err = configs.ErrNoInput

		return response, err
	}

	url := fmt.Sprintf("%s/broadcast", configs.Node)

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

func (b BaseImpl) Monitor(message *models.GetTransaction) (response models.GetTransactionResponse, err error) {
	if message == nil {
		err = configs.ErrNoInput
		return
	}

	url := fmt.Sprintf("%s/check/%s", configs.Node, message.TxHash)

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
