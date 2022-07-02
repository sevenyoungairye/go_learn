package tag

import (
	mongoRepo "top.lel.dn/main/internal/repository/mongodb"
	"top.lel.dn/main/pkg/db/mongodb"
	"top.lel.dn/main/pkg/db/redisdb"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	SaveMovieTag()
	SaveTvTag()
	DeleteById(id string) error

	MovieList() []mongoRepo.TagInfo
	TvList() []mongoRepo.TagInfo
}

type service struct {
	mongoCtx mongodb.MongoCtx
	cache    redisdb.RedisCacheRepo
}

func New(mongoCtx mongodb.MongoCtx, cache redisdb.RedisCacheRepo) Service {

	return &service{
		mongoCtx: mongoCtx,
		cache:    cache,
	}
}

// 适配器
func (s *service) i() {
}
