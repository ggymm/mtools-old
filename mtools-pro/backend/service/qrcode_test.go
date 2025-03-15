package service

import (
	"bytes"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"mtools-pro/backend/api"
	"os"
	"testing"
)

func TestQrcodeService_toBase64(t *testing.T) {
	fp := "C:\\Users\\19679\\Pictures\\qrcode\\mtools.png"
	f, _ := os.Open(fp)
	defer func() {
		_ = f.Close()
	}()
	buf, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}

	s := NewQrcodeService()
	t.Log(s.toBase64(buf))
}

func TestQrcodeService_parseContent(t *testing.T) {
	fp := "C:\\Users\\19679\\Pictures\\qrcode\\mtools.png"
	f, _ := os.Open(fp)
	defer func() {
		_ = f.Close()
	}()
	buf, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}

	s := NewQrcodeService()
	str, err := s.parseContent(bytes.NewBuffer(buf))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(str)
}

func TestQrcodeService_LoadLocalDecode(t *testing.T) {
	t.Log("not implement")
}

func TestQrcodeService_LoadClipboardDecode(t *testing.T) {
	req := api.DecodeQrcodeReq{}

	s := NewQrcodeService()
	resp := s.LoadClipboardDecode(req)
	if resp.Success {
		t.Log(resp.Data)
	} else {
		t.Error(resp.Msg)
	}
}

func TestQrcodeService_GenerateContentEncode(t *testing.T) {
	req := api.EncodeQrcodeReq{
		Content: "mtools",
	}

	s := NewQrcodeService()
	resp := s.GenerateContentEncode(req)
	if resp.Success {
		t.Log(resp.Data)
	} else {
		t.Error(resp.Msg)
	}
}
