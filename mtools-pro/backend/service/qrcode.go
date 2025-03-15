package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/makiuchi-d/gozxing"
	gozxingqrcode "github.com/makiuchi-d/gozxing/qrcode"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.design/x/clipboard"

	"mtools-pro/backend/api"
)

type QrcodeService struct {
	ctx context.Context
}

var qrcodeOnce sync.Once
var qrcodeService *QrcodeService

func NewQrcodeService() *QrcodeService {
	if qrcodeService == nil {
		qrcodeOnce.Do(func() {
			qrcodeService = &QrcodeService{}
		})
	}
	return qrcodeService
}

func (s *QrcodeService) Startup(ctx context.Context) {
	s.ctx = ctx
}

func (s *QrcodeService) toBase64(buf []byte) (str string) {
	str = base64.StdEncoding.EncodeToString(buf)
	str = fmt.Sprintf("data:%s;base64,%s", http.DetectContentType(buf), str)
	return
}

func (s *QrcodeService) parseContent(r io.Reader) (str string, err error) {
	var (
		img    image.Image
		bmp    *gozxing.BinaryBitmap
		result *gozxing.Result
	)

	// 解析二维码
	img, _, err = image.Decode(r)
	if err != nil {
		return
	}
	bmp, err = gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return
	}
	result, err = gozxingqrcode.NewQRCodeReader().Decode(bmp, nil)
	if err != nil {
		return
	}

	str = result.GetText()
	return
}

func (s *QrcodeService) LoadLocalDecode(_ api.DecodeQrcodeReq) api.Resp {
	fp, err := runtime.OpenFileDialog(s.ctx, runtime.OpenDialogOptions{
		Title:           "选择图片",
		ShowHiddenFiles: true,
	})
	if err != nil {
		return api.Error(err.Error())
	}

	f, _ := os.Open(fp)
	defer func() {
		_ = f.Close()
	}()
	buf, err := io.ReadAll(f)
	if err != nil {
		return api.Error(err.Error())
	}

	var (
		b64     string
		content string
	)

	// 图片转 base64
	b64 = s.toBase64(buf)

	// 解析二维码
	content, err = s.parseContent(bytes.NewBuffer(buf))
	if err != nil {
		return api.Error(err.Error())
	}

	return api.Success(api.DecodeQrcodeResp{
		Base64:  b64,
		Content: content,
	})
}

func (s *QrcodeService) LoadClipboardDecode(_ api.DecodeQrcodeReq) api.Resp {
	buf := clipboard.Read(clipboard.FmtImage)

	// 图片转 base64
	b64 := s.toBase64(buf)

	// 解析二维码
	content, err := s.parseContent(bytes.NewBuffer(buf))
	if err != nil {
		return api.Error(err.Error())
	}

	return api.Success(api.DecodeQrcodeResp{
		Base64:  b64,
		Content: content,
	})
}

func (s *QrcodeService) GenerateContentEncode(req api.EncodeQrcodeReq) api.Resp {
	hints := map[gozxing.EncodeHintType]interface{}{
		gozxing.EncodeHintType_MARGIN:        1,
		gozxing.EncodeHintType_CHARACTER_SET: "UTF-8",
	}
	img, err := gozxingqrcode.NewQRCodeWriter().Encode(
		req.Content,
		gozxing.BarcodeFormat_QR_CODE,
		req.Size,
		req.Size,
		hints,
	)
	if err != nil {
		return api.Error(err.Error())
	}

	buf := bytes.NewBuffer(nil)
	err = png.Encode(buf, img)
	if err != nil {
		return api.Error(err.Error())
	}

	// 图片转 base64
	b64 := s.toBase64(buf.Bytes())
	return api.Success(api.EncodeQrcodeResp{
		Base64: b64,
	})
}
