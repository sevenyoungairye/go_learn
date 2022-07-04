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

func (s *service) RestSaveTv(param InfoCreateVo) {
	allTag := s.getAllTag()
	request := crawler.DouBanRequest{
		Tag:       param.Tag,
		Sort:      param.Sort,
		PageStart: fmt.Sprint(param.PageStart),
		PageLimit: fmt.Sprint(param.PageLimit),
	}

	// 保存tv
	var tv = request.GetTv()
	for _, item := range getEpisodeData(tv.Subjects) {
		item.TagType = episode.Tv
		tag := s.getTag(allTag, request.Tag, item.TagType)
		item.TagList = append(make([]mongodb.TagInfo, 0), tag)
		item.SaveOrUpd(&s.mongoCtx)
	}
}

func (s *service) SaveRmdTv() {
	s.saveTv(episode.Rmd)
}

func (s *service) SaveRankTv() {
	s.saveTv(episode.Rank)
}

func (s *service) saveTv(flg int) {
	allTag := s.getAllTag()
	for pageNo := 1; pageNo <= episode.TotalPage; pageNo++ {
		pageStart, pageLimit := episode.ComputePageData(pageNo)
		request := crawler.DouBanRequest{
			PageStart: fmt.Sprint(pageStart),
			PageLimit: fmt.Sprint(pageLimit),
		}

		// 保存tv
		var tv = crawler.BaseSubject{}
		if flg == episode.Rmd {
			tv = request.GetRmdTv()
		}
		if flg == episode.Rank {
			tv = request.GetRankTv()
		}
		for _, item := range getEpisodeData(tv.Subjects) {
			item.TagType = episode.Tv
			tag := s.getTag(allTag, request.Tag, item.TagType)
			item.TagList = append(make([]mongodb.TagInfo, 0), tag)
			item.SaveOrUpd(&s.mongoCtx)
		}
	}
}

// RestSaveMovie 后台rest 接口保存数据用
func (s *service) RestSaveMovie(param InfoCreateVo) {
	allTag := s.getAllTag()
	request := crawler.DouBanRequest{
		Tag:       param.Tag,
		Sort:      param.Sort,
		PageStart: fmt.Sprint(param.PageStart),
		PageLimit: fmt.Sprint(param.PageLimit),
	}
	data := getEpisodeData(request.GetCustomMovie().Subjects)
	for _, item := range data {
		item.TagType = episode.Movie
		tag := s.getTag(allTag, request.Tag, item.TagType)
		item.TagList = append(make([]mongodb.TagInfo, 0), tag)
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
	allTag := s.getAllTag()
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
		for _, item := range getEpisodeData(movie.Subjects) {
			item.TagType = episode.Movie
			tag := s.getTag(allTag, request.Tag, item.TagType)
			item.TagList = append(make([]mongodb.TagInfo, 0), tag)
			item.SaveOrUpd(&s.mongoCtx)
		}
	}
}

func getEpisodeData(episodeList []crawler.Episode) []mongodb.EpisodeInfo {
	data := make([]mongodb.EpisodeInfo, 0)
	for _, item := range episodeList {
		data = append(data, convert(item))
	}
	return data
}

func convert(item crawler.Episode) mongodb.EpisodeInfo {
	sub := crawler.GetSubject(item.ID)
	return mongodb.EpisodeInfo{
		Id:               db.GenID(),
		DbId:             item.ID,
		Title:            item.Title,
		Url:              item.URL,
		Cover:            item.Cover,
		Rate:             item.Rate,
		IsNew:            item.IsNew,
		Playable:         item.Playable,
		CoverX:           int(item.CoverX),
		CoverY:           int(item.CoverY),
		EpisodesInfo:     item.EpisodesInfo,
		PublicDate:       nil,
		Torrent:          "",
		DriverUrl:        "",
		Created:          time.Now(),
		Creator:          "system",
		Updater:          "system",
		UpdateTime:       time.Now(),
		Actors:           sub.Actors,
		Blacklisted:      sub.Blacklisted,
		CollectionStatus: sub.CollectionStatus,
		Directors:        sub.Directors,
		Duration:         sub.Duration,
		EpisodesCount:    sub.EpisodesCount,
		IsTv:             sub.IsTv,
		Region:           sub.Region,
		ReleaseYear:      sub.ReleaseYear,
		Star:             sub.Star,
		Subtype:          sub.Subtype,
		Types:            sub.Types,
		ShortComment: struct {
			Author  string `bson:"author" fmt:"author"`
			Content string `bson:"content" fmt:"content"`
		}(sub.ShortComment),
	}
}
