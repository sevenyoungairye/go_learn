package tag

import (
	"fmt"
	"time"
	"top.lel.dn/main/internal/crawler"
	"top.lel.dn/main/internal/repository/mongodb"
	"top.lel.dn/main/pkg/db"
	"top.lel.dn/main/pkg/logger"
)

var paramMap map[string]string

const (
	movie = "movie"
	tv    = "tv"
)

func init() {
	paramMap = make(map[string]string)
	paramMap[tv] = "电视剧"
	paramMap[movie] = "电影"
}

func (*service) SaveMovieTag() {
	save(movie)
}

func (*service) SaveTvTag() {
	save(tv)
}

func save(t string) {
	var data *crawler.Tag
	if t == movie {
		data = crawler.GetMovieTag()
	} else if t == tv {
		data = crawler.GetTvTag()
	} else {
		return
	}

	infos := make([]mongodb.TagInfo, 0)
	for _, item := range data.Tags {
		// logger.Info(fmt.Sprintf("%d, %s", i, item))
		infos = append(infos, mongodb.TagInfo{
			Id:         db.GenID(),
			TagType:    t,
			TagName:    item,
			CreateTime: time.Now(),
			Creator:    "system",
			Updater:    "system",
			UpdateTime: time.Now(),
		})
	}

	for _, tagInfo := range infos {
		// mongodb入库
		tagInfo.SaveOrUpdTag()
	}

	logger.Info(fmt.Sprintf("保存%s标签个数成功, 共计抓取标签个数: %d", paramMap[t], len(infos)))

}
