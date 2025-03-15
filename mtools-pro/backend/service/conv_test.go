package service

import (
	"testing"

	"mtools-pro/backend/api"
)

func TestConvService_NumberBase(t *testing.T) {
	req := api.NumberBaseReq{
		Input:     "123456",
		InputBase: 10,
	}

	s := NewConvService()
	resp := s.NumberBase(req)
	if resp.Success {
		t.Log(resp.Data)
	} else {
		t.Error(resp.Msg)
	}
}
