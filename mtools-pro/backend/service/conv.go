package service

import (
	"strconv"
	"sync"

	"mtools-pro/backend/api"
)

type ConvService struct {
}

var convOnce sync.Once
var convService *ConvService

func NewConvService() *ConvService {
	if convService == nil {
		convOnce.Do(func() {
			convService = &ConvService{}
		})
	}
	return convService
}

func (s *ConvService) NumberBase(req api.NumberBaseReq) api.Resp {
	i, err := strconv.ParseInt(req.Input, req.InputBase, 64)
	if err != nil {
		return api.Error(err.Error())
	}

	return api.Success(api.NumberBaseResp{
		Binary:      strconv.FormatInt(i, 2),
		Octal:       strconv.FormatInt(i, 8),
		Decimal:     strconv.FormatInt(i, 10),
		Hexadecimal: strconv.FormatInt(i, 16),
	})
}
