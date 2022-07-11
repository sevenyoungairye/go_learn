package crawler

import (
	"time"
	"top.lel.dn/main/pkg/httpclient"
)

// SubjectAbstract 豆瓣剧集信息异步接口
// https://movie.douban.com/j/subject_abstract?subject_id=35235813
type SubjectAbstract struct {
	R       int     `fmt:"r"`
	Subject Subject `fmt:"subject"`
}

type Subject struct {
	Actors           []string     `json:"actors" fmt:"actors"`
	Blacklisted      string       `json:"blacklisted" fmt:"blacklisted"`
	CollectionStatus string       `json:"collection_status" fmt:"collection_status"`
	Directors        []string     `json:"directors" fmt:"directors"`
	Duration         string       `json:"duration" fmt:"duration"`
	EpisodesCount    string       `json:"episodes_count" fmt:"episodes_count"`
	ID               string       `json:"id" fmt:"id"`
	IsTv             bool         `json:"is_tv" fmt:"is_tv"`
	Playable         bool         `json:"playable" fmt:"playable"`
	Rate             string       `json:"rate" fmt:"rate"`
	Region           string       `json:"region" fmt:"region"`
	ReleaseYear      string       `json:"release_year" fmt:"release_year"`
	ShortComment     ShortComment `json:"short_comment" fmt:"short_comment"`
	Star             int          `json:"star" fmt:"star"`
	Subtype          string       `json:"subtype" fmt:"subtype"`
	Title            string       `json:"title" fmt:"title"`
	Types            []string     `json:"types" fmt:"types"`
	URL              string       `json:"url" fmt:"url"`
}

type ShortComment struct {
	Author  string `json:"author" fmt:"author"`
	Content string `json:"content" fmt:"content"`
}

type Episode struct {
	Cover        string `json:"cover" fmt:"cover"`
	CoverX       int64  `json:"cover_x" fmt:"cover_x"`
	CoverY       int64  `json:"cover_y" fmt:"cover_y"`
	EpisodesInfo string `json:"episodes_info" fmt:"episodes_info"`
	ID           string `json:"id" fmt:"id"`
	IsNew        bool   `json:"is_new" fmt:"is_new"`
	Playable     bool   `json:"playable" fmt:"playable"`
	Rate         string `json:"rate" fmt:"rate"`
	Title        string `json:"title" fmt:"title"`
	URL          string `json:"url" fmt:"url"`
}

type BaseSubject struct {
	Subjects []Episode `json:"subjects" fmt:"subjects"`
}

// GetPublicDate 获取剧集上映日期
func (e *Episode) GetPublicDate() *time.Time {
	// https://movie.douban.com/subject/35558660/
	dateStr := httpclient.GetDateByAttrSelector(e.URL, "span[property=\"v:initialReleaseDate\"]")
	l := len(dateStr)
	if l > 8 {
		dateValue, _ := time.Parse("2006-01-02", dateStr) // convert 'String' to 'Time' data type
		return &dateValue
	}
	if l > 4 {
		dateValue, _ := time.Parse("2006-01", dateStr)
		return &dateValue
	}
	if l > 0 {
		dateValue, _ := time.Parse("2006", dateStr)
		return &dateValue
	}
	return nil
}

/*
func (b *BaseSubject) convert() *[]mongodb.MovieInfo {
	data := make([]mongodb.MovieInfo, 0)
	for _, item := range b.Subjects {
		data = append(data, *item.covert())
	}
	return &data
}

func (e *Episode) covert() *mongodb.MovieInfo {
	return &mongodb.MovieInfo{
		Id:           "",
		MvTagId:      "",
		MvTagInfo:    "",
		DbId:         "",
		Title:        "",
		Url:          "",
		Cover:        "",
		Rate:         "",
		IsNew:        false,
		Playable:     false,
		CoverX:       0,
		CoverY:       0,
		EpisodesInfo: "",
		PublicDate:   time.Time{},
		Torrent:      "",
		DriverUrl:    "",
		Created:      time.Time{},
		Creator:      "",
		Updater:      "",
		UpdateTime:   time.Time{},
	}
}
*/
