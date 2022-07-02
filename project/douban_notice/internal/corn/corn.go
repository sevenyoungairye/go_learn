// 定时拉取豆瓣资源.
// @since 2022/07/02

package corn

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
	mongodbRepo "top.lel.dn/main/internal/repository/mongodb"
	"top.lel.dn/main/internal/services/episode/movie"
	"top.lel.dn/main/pkg/db/mongodb"
	"top.lel.dn/main/pkg/db/redisdb"
	"top.lel.dn/main/pkg/logger"
)

// Demo https://pkg.go.dev/github.com/robfig/cron/v3@v3.0.0#hdr-Usage
func Demo() {

}

func init() {
	c := cron.New()
	id, err := c.AddFunc("50 21 * * *", func() {
		ctx := mongodb.GetCtxCollection(mongodbRepo.MovieInfoCollect)
		service := movie.New(*ctx, *redisdb.New())
		start := time.Now().UnixMilli()
		logger.Info("抓取电影定时任务开始, ")
		service.SaveLastMovie()
		//service.SaveRmdMovie()
		//service.SaveRankMovie()
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
