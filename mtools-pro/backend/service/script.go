package service

import (
	"sync"
)

type ScriptService struct {
}

var scriptOnce sync.Once
var scriptService *ScriptService

func NewScriptService() *ScriptService {
	if scriptService == nil {
		scriptOnce.Do(func() {
			scriptService = &ScriptService{}
		})
	}
	return scriptService
}
