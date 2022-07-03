package movie

import (
	mongodbRepo "top.lel.dn/main/internal/repository/mongodb"
	"top.lel.dn/main/internal/services/episode/tag"
	"top.lel.dn/main/pkg/db/mongodb"
	"top.lel.dn/main/pkg/db/redisdb"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	RestSave(param InfoCreateVo)
	SaveLastMovie()
	SaveRmdMovie()
	SaveRankMovie()
}

type service struct {
	mongoCtx mongodb.MongoCtx
	cache    redisdb.RedisCacheRepo
}

func New(ctx mongodb.MongoCtx, cache redisdb.RedisCacheRepo) Service {
	return &service{
		mongoCtx: ctx,
		cache:    cache,
	}
}

func (s *service) i() {
}

// 获取所有tag
func (s *service) getAllTag() []mongodbRepo.TagInfo {
	tagService := tag.New(s.mongoCtx, s.cache)
	return append(tagService.MovieList(), tagService.TvList()...)
}

// 筛选tag.
func (s *service) getTag(tagList []mongodbRepo.TagInfo, tagName string, tagType string) mongodbRepo.TagInfo {
	for _, item := range tagList {
		if tagName == item.TagName && tagType == item.TagType {
			return item
		}
	}
	return mongodbRepo.TagInfo{
		Id:      "",
		TagType: "",
		TagName: "",
	}
}
