package services

import (
	"reflect"
	"testing"

	"github.com/parmcoder/transaction-tracker/models"
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
			name:         "Failed",
			args:         args{message: &message1},
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
