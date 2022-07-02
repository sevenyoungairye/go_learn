// 针对tag做了缓存优化.

package tag

import (
	"fmt"
	"time"
	"top.lel.dn/main/internal/repository/mongodb"
	"top.lel.dn/main/internal/repository/redis"
	"top.lel.dn/main/pkg/logger"
	"top.lel.dn/main/pkg/serial"
)

func (s *service) MovieList() []mongodb.TagInfo {
	info := mongodb.TagInfo{}
	return s.cacheTag(info.GetAll(movie), redis.MovieTagRedisKey)
}

func (s *service) TvList() []mongodb.TagInfo {
	info := mongodb.TagInfo{}
	return s.cacheTag(info.GetAll(tv), redis.TvTagRedisKey)
}

func (s *service) cacheTag(data []mongodb.TagInfo, key string) []mongodb.TagInfo {
	cache := s.cache
	val := cache.GetVal(key)
	if val != "" {
		infos := make([]mongodb.TagInfo, 0)
		serial.Json2Instant(val, &infos)
		return infos
	}
	err := cache.Add(key, serial.Object2Json(data), time.Hour*24*7)
	if err != nil {
		logger.Warn("save tag to cache err: %s" + fmt.Sprint(err))
		return make([]mongodb.TagInfo, 0)
	}
	return data
}
