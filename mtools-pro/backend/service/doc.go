package service

import (
	"sync"
)

type DocService struct {
}

var docOnce sync.Once
var docService *DocService

func NewDocService() *DocService {
	if docService == nil {
		docOnce.Do(func() {
			docService = &DocService{}
		})
	}
	return docService
}
