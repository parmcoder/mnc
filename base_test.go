package mnc

import (
	"reflect"
	"testing"
)

/*
 Disclaimer:
 This test is run against the real server. So, we can use
*/
func TestBaseImpl_Broadcast(t *testing.T) {
	type args struct {
		message *CreateBroadcast
	}

	message1 := CreateBroadcast{
		Symbol:    "BTC",
		Price:     1000,
		Timestamp: 100123,
	}

	message2 := CreateBroadcast{}

	tests := []struct {
		name         string
		args         args
		wantResponse bool
		wantErr      bool
	}{
		{
			name:         "Success",
			args:         args{message: &message1},
			wantResponse: true,
			wantErr:      false,
		},
		{
			name:         "Success - empty",
			args:         args{message: &message2},
			wantResponse: true,
			wantErr:      false,
		},
		{
			name:         "Failed - no input",
			args:         args{message: nil},
			wantResponse: false,
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := base{}
			gotResponse, err := b.Broadcast(tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("BaseImpl.Broadcast() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantResponse && len(gotResponse.TxHash) <= 0 {
				t.Errorf("BaseImpl.Broadcast() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestBaseImpl_Monitor(t *testing.T) {
	type args struct {
		message *GetTransaction
	}

	message1 := GetTransaction{
		TxHash: "abc",
	}

	message2 := GetTransaction{
		TxHash: "",
	}

	tests := []struct {
		name         string
		b            base
		args         args
		wantResponse GetTransactionResponse
		wantErr      bool
	}{
		{
			name: "Success",
			args: args{message: &message1},
			wantResponse: GetTransactionResponse{
				TxStatus: DoNotExist,
			},
			wantErr: false,
		},
		{
			name:         "Failed - cannot process input",
			args:         args{message: &message2},
			wantResponse: GetTransactionResponse{},
			wantErr:      true,
		},
		{
			name:         "Failed - no input",
			args:         args{message: nil},
			wantResponse: GetTransactionResponse{},
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := base{}
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
