// 用于补全数据, 业余数据更新

package episode

import (
	"fmt"
	"top.lel.dn/main/internal/crawler"
	mongodb2 "top.lel.dn/main/internal/repository/mongodb"
	"top.lel.dn/main/pkg/component/pageable"
	"top.lel.dn/main/pkg/db/mongodb"
	"top.lel.dn/main/pkg/logger"
)

func Run() {
	ctx := mongodb.GetCtxCollection(mongodb2.MovieInfoCollect)
	//now, _ := time.Parse(time.RFC3339, "01-01-01 00:00:00.000")
	info := mongodb2.EpisodeInfo{PublicDate: nil}
	pageVo := pageable.PageVo{PageNo: 1, PageLimit: 200}
	data, total := info.PageList(ctx, pageVo)
	logger.Debug(fmt.Sprintln(total))

	var totalPage = total / int64(pageVo.PageLimit)
	if total%int64(pageVo.PageLimit) > 0 {
		totalPage += 1
	}

	for _, item := range data {
		episode := crawler.Episode{URL: item.Url}
		item.PublicDate = episode.GetPublicDate()
		logger.Debug(fmt.Sprintf("id: %v,  public date: %v", item.DbId, item.PublicDate))
		item.SaveOrUpd(ctx)
	}
	logger.Info(fmt.Sprintf("total count: %v, len: %v", total, len(data)))

	for pageNo := 2; pageNo < int(totalPage); pageNo++ {
		pageVo.PageNo = 1
		data, total = info.PageList(ctx, pageVo)
		for _, item := range data {
			episode := crawler.Episode{URL: item.Url}
			item.PublicDate = episode.GetPublicDate()
			logger.Debug(fmt.Sprintf("id: %v,  public date: %v", item.DbId, item.PublicDate))
			item.SaveOrUpd(ctx)
		}
		logger.Info(fmt.Sprintf("total count: %v, len: %v", total, len(data)))
	}

	defer ctx.Release()
}
