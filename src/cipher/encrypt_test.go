package cipher

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestCipher_Encrypt(t *testing.T) {
	tests := []struct {
		name    string
		inData  string
		initKey uint32
		wantOut string
		wantErr bool
	}{
		{name: "test", inData: "960100010203040506070809", initKey: 0x3e05, wantOut: "df3c57c4e5aa21986104167b43f3f1ef053e0000", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cipher{
				Diverged: false,
				Key:      tt.initKey,
			}
			inBytes, err := hex.DecodeString(tt.inData)
			if err != nil {
				t.Errorf("invalid test data, %v", err)
				return
			}
			gotOut, err := c.Encrypt(inBytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(hex.EncodeToString(gotOut), tt.wantOut) {
				t.Errorf("Encrypt() gotOut = %v, want %v", hex.EncodeToString(gotOut), tt.wantOut)
			}
		})
	}
}
