package service

import (
	"sync"

	"golang.org/x/crypto/bcrypt"

	"mtools-pro/backend/api"
)

type CryptoService struct {
}

var cryptoOnce sync.Once
var cryptoService *CryptoService

func NewCryptoService() *CryptoService {
	if cryptoService == nil {
		cryptoOnce.Do(func() {
			cryptoService = &CryptoService{}
		})
	}
	return cryptoService
}

func (s *CryptoService) BcryptEncode(req api.BcryptEncodeReq) api.Resp {
	result, err := bcrypt.GenerateFromPassword([]byte(req.Input), req.SaltCount)
	if err != nil {
		return api.Error(err.Error())
	}
	return api.Success(string(result))
}

func (s *CryptoService) BcryptCompare(req api.BcryptCompareReq) api.Resp {
	err := bcrypt.CompareHashAndPassword([]byte(req.Encode), []byte(req.Origin))
	if err != nil {
		return api.Success("false")
	}
	return api.Success("true")
}
