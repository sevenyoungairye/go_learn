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

// todo: 定时任务设置tag.
func (s *service) getAllTag() []mongodbRepo.TagInfo {
	return tag.New(s.mongoCtx, s.cache).MovieList()
}

func (s *service) getMovieTag(tagList []mongodbRepo.TagInfo, tagName string) mongodbRepo.TagInfo {
	for _, item := range tagList {
		if tagName == item.TagName {
			return item
		}
	}
	return mongodbRepo.TagInfo{
		Id:      "",
		TagType: "",
		TagName: "",
	}
}
