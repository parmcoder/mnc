package services

import (
	"reflect"
	"testing"

	"github.com/parmcoder/mnc/configs"
	"github.com/parmcoder/mnc/models"
)

func TestBaseImpl_Broadcast(t *testing.T) {
	type args struct {
		message *models.CreateBroadcast
	}

	message1 := models.CreateBroadcast{
		Symbol:    "BTC",
		Price:     1000,
		Timestamp: 100123,
	}

	message2 := models.CreateBroadcast{}

	tests := []struct {
		name         string
		args         args
		wantResponse models.CreateBroadcastResponse
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			name:         "Success",
			args:         args{message: &message1},
			wantResponse: models.CreateBroadcastResponse{},
			wantErr:      false,
		},
		{
			name:         "Success - empty",
			args:         args{message: &message2},
			wantResponse: models.CreateBroadcastResponse{},
			wantErr:      false,
		},
		{
			name:         "Failed - no input",
			args:         args{message: nil},
			wantResponse: models.CreateBroadcastResponse{},
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BaseImpl{}
			gotResponse, err := b.Broadcast(tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("BaseImpl.Broadcast() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("BaseImpl.Broadcast() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestBaseImpl_Monitor(t *testing.T) {
	type args struct {
		message *models.GetTransaction
	}

	message1 := models.GetTransaction{
		TxHash: "abc",
	}

	message2 := models.GetTransaction{
		TxHash: "",
	}

	tests := []struct {
		name         string
		b            BaseImpl
		args         args
		wantResponse models.GetTransactionResponse
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			name: "Success",
			args: args{message: &message1},
			wantResponse: models.GetTransactionResponse{
				TxStatus: configs.DoNotExist,
			},
			wantErr: false,
		},
		{
			name:         "Failed - cannot process input",
			args:         args{message: &message2},
			wantResponse: models.GetTransactionResponse{},
			wantErr:      true,
		},
		{
			name:         "Failed - no input",
			args:         args{message: nil},
			wantResponse: models.GetTransactionResponse{},
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BaseImpl{}
			gotResponse, err := b.Monitor(tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("BaseImpl.Monitor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("BaseImpl.Monitor() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
