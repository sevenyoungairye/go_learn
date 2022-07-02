package movie

import (
	"fmt"
	"time"
	"top.lel.dn/main/internal/crawler"
	"top.lel.dn/main/internal/repository/mongodb"
	"top.lel.dn/main/internal/services/episode"
	"top.lel.dn/main/pkg/db"
)

type InfoCreateVo struct {
	PageStart int    `json:"pageStart"`
	PageLimit int    `json:"pageLimit"`
	Tag       string `json:"tag"`
	Sort      string `json:"sort"`
}

// RestSave 后台rest 接口保存数据用
func (s *service) RestSave(param InfoCreateVo) {
	request := crawler.DouBanRequest{
		Tag:       param.Tag,
		Sort:      param.Sort,
		PageStart: fmt.Sprint(param.PageStart),
		PageLimit: fmt.Sprint(param.PageLimit),
	}
	data := getMovieData(request.GetCustomMovie().Subjects)
	for _, item := range data {
		// todo 定时任务 根据标签名和标签类型获取 标签id...
		item.MvTagInfo = request.Tag
		item.SaveOrUpd(&s.mongoCtx)
	}
}

func (s *service) SaveRankMovie() {
	s.saveMovie(episode.Rank)
}

func (s *service) SaveRmdMovie() {
	s.saveMovie(episode.Rmd)
}

func (s *service) SaveLastMovie() {
	s.saveMovie(episode.Last)
}

func (s *service) saveMovie(flg int) {
	for pageNo := 1; pageNo <= episode.TotalPage; pageNo++ {
		pageStart, pageLimit := episode.ComputePageData(pageNo)
		request := crawler.DouBanRequest{
			PageStart: fmt.Sprint(pageStart),
			PageLimit: fmt.Sprint(pageLimit),
		}

		// 保存最新电影...
		var movie = crawler.BaseSubject{}
		if episode.Last == flg {
			movie = request.GetLastMovie()
		}
		if episode.Rank == flg {
			movie = request.GetRankMovie()
		}
		if episode.Rmd == flg {
			movie = request.GetRmdMovie()
		}
		for _, item := range getMovieData(movie.Subjects) {
			item.SaveOrUpd(&s.mongoCtx)
		}
	}
}

func getMovieData(episodeList []crawler.Episode) []mongodb.MovieInfo {
	data := make([]mongodb.MovieInfo, 0)
	for _, item := range episodeList {
		data = append(data, convert(item))
	}
	return data
}

func convert(item crawler.Episode) mongodb.MovieInfo {
	return mongodb.MovieInfo{
		Id:           db.GenID(),
		MvTagId:      "",
		MvTagInfo:    "",
		DbId:         item.ID,
		Title:        item.Title,
		Url:          item.URL,
		Cover:        item.Cover,
		Rate:         item.Rate,
		IsNew:        item.IsNew,
		Playable:     item.Playable,
		CoverX:       int(item.CoverX),
		CoverY:       int(item.CoverY),
		EpisodesInfo: item.EpisodesInfo,
		PublicDate:   nil,
		Torrent:      "",
		DriverUrl:    "",
		Created:      time.Now(),
		Creator:      "system",
		Updater:      "system",
		UpdateTime:   time.Now(),
	}
}
