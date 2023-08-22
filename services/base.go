package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/parmcoder/transaction-tracker/configs"
	"github.com/parmcoder/transaction-tracker/models"
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

	return models.CreateBroadcastResponse{}, nil
}

func (b BaseImpl) Monitor(message *models.GetTransaction) (response models.GetTransactionResponse, err error) {
	url := fmt.Sprintf("%s/check/%s", configs.Node, *message)

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
