package config

import (
	"testing"
)

func Test_Init(t *testing.T) {
	Init()
	t.Log(AppName)
	t.Log(BasePath)
}
