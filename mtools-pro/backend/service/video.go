package service

import (
	"context"
	"sync"

	"mtools-pro/backend/api"
)

type VideoService struct {
	ctx context.Context
}

var videoOnce sync.Once
var videoService *VideoService

func NewVideoService() *VideoService {
	if videoService == nil {
		videoOnce.Do(func() {
			videoService = &VideoService{}
		})
	}
	return videoService
}

func (s *VideoService) Startup(ctx context.Context) {
	s.ctx = ctx
}

func (s *VideoService) DownloadVideo(req api.VideoDownloadReq) api.Resp {
	return api.Success(nil)
}
