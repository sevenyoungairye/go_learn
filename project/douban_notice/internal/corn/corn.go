// 定时拉取豆瓣资源.
// @since 2022/07/02

package corn

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
	mongodbRepo "top.lel.dn/main/internal/repository/mongodb"
	"top.lel.dn/main/internal/services/episode"
	"top.lel.dn/main/internal/services/episode/movie"
	"top.lel.dn/main/internal/services/episode/tag"
	"top.lel.dn/main/pkg/db/mongodb"
	"top.lel.dn/main/pkg/db/redisdb"
	"top.lel.dn/main/pkg/logger"
)

// Demo https://pkg.go.dev/github.com/robfig/cron/v3@v3.0.0#hdr-Usage
func Demo() {

}

const (
	TvPage    = 2000
	MoviePage = 2000
)

// 周天10:30 -> 30 10 * * 0
// 周六10:30 -> 30 10 * * 6
const (
	// CrawlerMovieCorn 每月15号晚6点
	CrawlerMovieCorn = "00 18 15 * ?"

	// CrawlerTvCorn 每月15号早上8点
	CrawlerTvCorn = "00 08 15 * ?"

	// HomeMovieCorn 每周天上午10点半
	HomeMovieCorn = "30 10 * * 0"

	// HomeTvCorn 每周天上午10点
	HomeTvCorn = "00 10 * * 0"
)

func init() {

	c := cron.New()
	go SaveHomeMovie(c)
	go SaveHomeTv(c)

	go CrawlerTv(c)
	go CrawlerMovie(c)
}

// CrawlerMovie 抓取指定标签movie数据, 每个标签2000页
func CrawlerMovie(c *cron.Cron) {
	if c == nil {
		logger.Warn("the corn is nil!")
		return
	}
	id, err := c.AddFunc(CrawlerMovieCorn, func() {
		ctx := mongodb.GetCtxCollection(mongodbRepo.MovieInfoCollect)
		service := movie.New(*ctx, *redisdb.New())
		start := time.Now().UnixMilli()
		logger.Info("抓取电影定时任务开始, ")

		for _, item := range tag.New(*ctx, *redisdb.New()).MovieList() {
			if "热门" == item.TagName || "最新" == item.TagName {
				continue
			}
			count := 0
			for pageNo := 1; pageNo < MoviePage; pageNo++ {
				start, limit := episode.ComputePageData(pageNo)
				service.RestSaveMovie(movie.InfoCreateVo{
					PageStart: start,
					PageLimit: limit,
					Tag:       item.TagName,
					Sort:      "",
				})
				count += episode.PageLimit
				time.Sleep(time.Second * 2)
			}
			logger.Info(fmt.Sprintf("抓取movie完成, tagName: %s, 共计%d个", item.TagName, count))
			count = 0
		}

		end := time.Now().UnixMilli()
		logger.Info(fmt.Sprintf("抓取电影定时任务结束, 用时: %d秒", (end-start)/1000))

		defer ctx.Release()
	})
	if err != nil {
		logger.Warn(fmt.Sprint(err))
		return
	} else {
		logger.Info(fmt.Sprintf("corn job start %d", id))
	}
	c.Start()
}

func CrawlerTv(c *cron.Cron) {
	if c == nil {
		logger.Warn("the corn is nil!")
		return
	}

	id, err := c.AddFunc(CrawlerTvCorn, func() {
		ctx := mongodb.GetCtxCollection(mongodbRepo.MovieInfoCollect)
		service := movie.New(*ctx, *redisdb.New())
		start := time.Now().UnixMilli()
		logger.Info("抓取tv定时任务开始, ")
		for _, item := range tag.New(*ctx, *redisdb.New()).TvList() {
			if "热门" == item.TagName {
				continue
			}
			count := 0
			for pageNo := 1; pageNo < TvPage; pageNo++ {
				start, limit := episode.ComputePageData(pageNo)
				service.RestSaveTv(movie.InfoCreateVo{
					PageStart: start,
					PageLimit: limit,
					Tag:       item.TagName,
					Sort:      "",
				})
				count += episode.PageLimit
				time.Sleep(time.Second * 2)
			}
			logger.Info(fmt.Sprintf("抓取tv完成, tagName: %s, 共计%d个", item.TagName, count))
			count = 0
		}
		end := time.Now().UnixMilli()
		logger.Info(fmt.Sprintf("抓取tv定时任务结束, 用时: %d秒", (end-start)/1000))

		defer ctx.Release()
	})
	if err != nil {
		logger.Warn(fmt.Sprint(err))
		return
	} else {
		logger.Info(fmt.Sprintf("corn job start %d", id))
	}
	c.Start()
}

func SaveHomeMovie(c *cron.Cron) {
	if c == nil {
		logger.Warn("the corn is nil!")
		return
	}
	id, err := c.AddFunc(HomeMovieCorn, func() {
		ctx := mongodb.GetCtxCollection(mongodbRepo.MovieInfoCollect)
		service := movie.New(*ctx, *redisdb.New())
		start := time.Now().UnixMilli()
		logger.Info("抓取首页电影定时任务开始, ")
		service.SaveLastMovie()
		service.SaveRmdMovie()
		service.SaveRankMovie()
		end := time.Now().UnixMilli()
		logger.Info(fmt.Sprintf("抓取首页电影定时任务结束, 用时: %d秒", (end-start)/1000))

		defer ctx.Release()
	})
	if err != nil {
		logger.Warn(fmt.Sprint(err))
		return
	} else {
		logger.Info(fmt.Sprintf("corn job start %d", id))
	}
	c.Start()
}

func SaveHomeTv(c *cron.Cron) {
	if c == nil {
		logger.Warn("the corn is nil!")
		return
	}
	id, err := c.AddFunc(HomeTvCorn, func() {
		ctx := mongodb.GetCtxCollection(mongodbRepo.MovieInfoCollect)
		service := movie.New(*ctx, *redisdb.New())
		start := time.Now().UnixMilli()
		logger.Info("抓取home tv定时任务开始, ")
		service.SaveRmdTv()
		service.SaveRankMovie()
		end := time.Now().UnixMilli()
		logger.Info(fmt.Sprintf("抓取home tv定时任务结束, 用时: %d秒", (end-start)/1000))

		defer ctx.Release()
	})
	if err != nil {
		logger.Warn(fmt.Sprint(err))
		return
	} else {
		logger.Info(fmt.Sprintf("corn job start %d", id))
	}
	c.Start()
}
