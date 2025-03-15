package service

import (
	"context"
	"sync"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"mtools-pro/backend/api"
)

type RuntimeService struct {
	ctx context.Context
}

var runtimeOnce sync.Once
var runtimeService *RuntimeService

func NewRuntimeService() *RuntimeService {
	if runtimeService == nil {
		runtimeOnce.Do(func() {
			runtimeService = &RuntimeService{}
		})
	}
	return runtimeService
}

func (s *RuntimeService) Startup(ctx context.Context) {
	s.ctx = ctx
}

func (s *RuntimeService) ChooseFile(req api.ChooseFileReq) api.Resp {
	filter := make([]runtime.FileFilter, 0)
	for _, f := range req.Filter {
		filter = append(filter, runtime.FileFilter{
			Pattern:     f.Pattern,
			DisplayName: f.DisplayName,
		})
	}
	fp, err := runtime.OpenFileDialog(s.ctx, runtime.OpenDialogOptions{
		Title:           req.Title,
		Filters:         filter,
		ShowHiddenFiles: true,
	})
	if err != nil {
		return api.Success(nil)
	}
	return api.Success(fp)
}
