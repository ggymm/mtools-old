package base64

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
)

func DecodeImage(fp string) (str string, err error) {
	file, _ := os.Open(fp)
	defer func() {
		_ = file.Close()
	}()

	var (
		fb           []byte
		contentType  string
		base64String string
	)

	// 图片转 base64
	fb, err = io.ReadAll(file)
	if err != nil {
		return
	}
	contentType = http.DetectContentType(fb)
	base64String = base64.StdEncoding.EncodeToString(fb)

	str = fmt.Sprintf("data:%s;base64,%s", contentType, base64String)
	return
}
