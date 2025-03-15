package service

import (
	"mtools-pro/backend/api"
	"testing"
)

func TestCryptoService_BcryptEncode(t *testing.T) {
	req := api.BcryptEncodeReq{
		Input:     "123456",
		SaltCount: 10,
	}

	s := NewCryptoService()
	resp := s.BcryptEncode(req)
	if resp.Success {
		t.Log(resp.Data)
	} else {
		t.Error(resp.Msg)
	}
}

func TestCryptoService_BcryptCompare(t *testing.T) {
	req := api.BcryptCompareReq{
		Origin: "123456",
		Encode: "$2a$10$4oixu3xDMJBjTpo5/jP8XOu9huxYaPzoP20DdO2x1TQaRTmiSMUii",
	}

	s := NewCryptoService()
	resp := s.BcryptCompare(req)
	if resp.Success {
		t.Log(resp.Data)
	} else {
		t.Error(resp.Msg)
	}
}
