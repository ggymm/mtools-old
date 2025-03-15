package base64

import "testing"

func Test_DecodeImage(t *testing.T) {
	fp := "C:\\Users\\19679\\Pictures\\qrcode\\mtools.png"
	str, err := DecodeImage(fp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(str)
}
