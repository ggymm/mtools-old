package service

import (
	"testing"

	"mtools-pro/backend/pkg/config"
	"mtools-pro/backend/pkg/database"
	"mtools-pro/backend/pkg/log"
)

func Test_GetCatalogTree(t *testing.T) {
	config.Init()
	log.Init()
	err := database.Init()
	if err != nil {
		t.Fatal(err)
		return
	}

	s := NewSnippetService()
	resp := s.GetCatalogTree()
	t.Logf("%+v", resp)
}
