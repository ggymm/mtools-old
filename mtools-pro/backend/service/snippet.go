package service

import (
	"context"
	"sync"

	"github.com/ggymm/gopkg/utils"

	"mtools-pro/backend/api"
	"mtools-pro/backend/pkg/consts"
	"mtools-pro/backend/pkg/database"
	"mtools-pro/backend/pkg/database/model"
)

type SnippetService struct {
}

var snippetOnce sync.Once
var snippetService *SnippetService

func NewSnippetService() *SnippetService {
	if snippetService == nil {
		snippetOnce.Do(func() {
			snippetService = &SnippetService{}
		})
	}
	return snippetService
}

func (s *SnippetService) Startup(_ context.Context) {
}

func (s *SnippetService) GetCatalogTree() api.Resp {
	var (
		err     error
		list    []*model.SnippetCatalog
		tree    = make([]*api.SnippetCatalogTree, 0)
		treeMap = make(map[int64]*api.SnippetCatalogTree)
	)

	err = database.DB.Find(&list).Error
	if err != nil {
		return api.Error(err.Error())
	}

	for _, dept := range list {
		var node api.SnippetCatalogTree
		node.Key = dept.Id
		node.Label = dept.Name
		node.Extras = dept.Extras
		node.Children = make([]*api.SnippetCatalogTree, 0)

		if dept.ParentId == consts.RootId {
			tree = append(tree, &node)
		} else {
			if treeMap[dept.ParentId] != nil {
				treeMap[dept.ParentId].Children = append(treeMap[dept.ParentId].Children, &node)
			}
		}
		treeMap[dept.Id] = &node
	}
	return api.Success(tree)
}

func (s *SnippetService) CreateCatalogItem(req api.SnippetCatalogCreateReq) api.Resp {
	snippetCatalog := model.SnippetCatalog{
		Name:     req.Name,
		ParentId: req.ParentId,
	}
	err := database.DB.Create(&snippetCatalog).Error
	if err != nil {
		return api.Error(err.Error())
	}
	return api.Success(snippetCatalog.Id)
}

func (s *SnippetService) GetPage(req api.SnippetPageReq) api.Resp {
	var (
		err       error
		where     model.Snippet
		modelList []*model.SnippetExt

		list        = make([]*api.SnippetPageResp, 0)
		total int64 = 0
	)

	// 分页参数
	size := req.GetSize()
	offset := req.GetOffset()

	// 查询条件
	where.CatalogId = req.CatalogId

	// 查询数据
	err = database.DB.Joins("SnippetCatalog").Where(&where).Limit(size).Offset(offset).Find(&modelList).Error
	if err != nil {
		return api.Error(err.Error())
	}
	err = database.DB.Model(&model.Snippet{}).Where(&where).Count(&total).Error
	if err != nil {
		return api.Error(err.Error())
	}

	// 映射为实体
	for _, item := range modelList {
		list = append(list, &api.SnippetPageResp{
			Id:          item.Id,
			Title:       item.Title,
			UpdateAt:    item.GetUpdateAt(),
			CatalogId:   item.CatalogId,
			CatalogName: item.SnippetCatalog.Name,
		})
	}
	return api.Success(api.PageResp{List: list, Total: total})
}

func (s *SnippetService) GetItem(req api.SnippetReq) api.Resp {
	var (
		err       error
		item      model.Snippet
		fragments []*model.SnippetFragment

		resp api.SnippetResp
	)

	// 查询 snippet
	err = database.DB.First(&item, req.Id).Error
	if err != nil {
		return api.Error(err.Error())
	}

	// 查询 fragments
	err = database.DB.Where("snippet_id = ?", req.Id).Find(&fragments).Error
	if err != nil {
		return api.Error(err.Error())
	}

	// 映射为实体
	resp.Id = item.Id
	resp.Title = item.Title
	resp.Fragments = make([]api.SnippetFragment, len(fragments))
	for i, fragment := range fragments {
		resp.Fragments[i].Title = fragment.Title
		resp.Fragments[i].Content = fragment.Content
		resp.Fragments[i].Language = fragment.Language
	}
	return api.Success(resp)
}

func (s *SnippetService) CreateItem(req api.SnippetCreateReq) api.Resp {
	// 保存 snippet
	item := model.Snippet{
		Title:     req.Title,
		CatalogId: utils.If(req.CatalogId == 0, consts.DefaultId, req.CatalogId),
	}
	err := database.DB.Create(&item).Error
	if err != nil {
		return api.Error(err.Error())
	}

	// 保存 fragment
	for _, fragment := range req.Fragments {
		f := model.SnippetFragment{
			SnippetId: item.Id,
			Title:     fragment.Title,
			Content:   fragment.Content,
			Language:  fragment.Language,
		}
		err = database.DB.Create(&f).Error
		if err != nil {
			return api.Error(err.Error())
		}
	}
	return api.Success(item.Id)
}

func (s *SnippetService) UpdateItem(req api.SnippetUpdateReq) api.Resp {
	// 更新 snippet
	err := database.DB.Model(&model.Snippet{Id: req.Id}).Updates(model.Snippet{
		Title: req.Title,
	}).Error
	if err != nil {
		return api.Error(err.Error())
	}

	// 删除 fragment
	err = database.DB.Delete(&model.SnippetFragment{}, "snippet_id = ?", req.Id).Error
	if err != nil {
		return api.Error(err.Error())
	}

	// 保存 fragment
	for _, fragment := range req.Fragments {
		f := model.SnippetFragment{
			SnippetId: req.Id,
			Title:     fragment.Title,
			Content:   fragment.Content,
			Language:  fragment.Language,
		}
		err = database.DB.Create(&f).Error
		if err != nil {
			return api.Error(err.Error())
		}
	}
	return api.Success(nil)
}

func (s *SnippetService) DeleteItem(req api.SnippetReq) api.Resp {
	// 删除 snippet
	err := database.DB.Delete(&model.Snippet{}, req.Id).Error
	if err != nil {
		return api.Error(err.Error())
	}

	// 删除 fragment
	err = database.DB.Delete(&model.SnippetFragment{}, "snippet_id = ?", req.Id).Error
	if err != nil {
		return api.Error(err.Error())
	}
	return api.Success(nil)
}
