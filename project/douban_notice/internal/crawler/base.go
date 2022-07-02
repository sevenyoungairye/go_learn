// @author echo lovely
// @since 2022/06/29
// https://movie.douban.com/j/search_subjects?type=movie&tag=热门&sort=recommend&page_limit=20&page_start=0

package crawler

import (
	"fmt"
	"top.lel.dn/main/pkg/httpclient"
	"top.lel.dn/main/pkg/serial"
	"top.lel.dn/main/pkg/yaml"
)

var douBanConfig *yaml.D

func init() {
	douBanConfig = yaml.GetDouBan()
}

func GetMovieTag() *Tag {
	retVal := Tag{}
	data := httpclient.HttpWithGet(douBanConfig.DouBan.Tag.Movie, nil)
	serial.Json2Instant(data, &retVal)
	return &retVal
}

func GetTvTag() *Tag {
	retVal := Tag{}
	data := httpclient.HttpWithGet(douBanConfig.DouBan.Tag.Tv, nil)
	serial.Json2Instant(data, &retVal)
	return &retVal
}

type Tag struct {
	Tags []string `json:"tags"`
}

type DouBanRequest struct {
	Tag       string
	Sort      string
	PageLimit string
	PageStart string
}

func (param *DouBanRequest) getMovieUrl() string {
	if param.PageStart == "" {
		param.PageStart = "0"
	}
	if param.PageLimit == "" {
		param.PageLimit = "20"
	}
	return fmt.Sprintf(douBanConfig.DouBan.Movie.Url, param.Tag, param.Sort, param.PageLimit, param.PageStart)
}

func getMovieData(url string) BaseSubject {
	json := httpclient.HttpWithGet(url, nil)
	subs := BaseSubject{}
	serial.Json2Instant(json, &subs)
	return subs
}

// GetLastMovie 获取最新
func (param *DouBanRequest) GetLastMovie() BaseSubject {
	param.Tag = "最新"
	param.Sort = "time"

	return getMovieData(param.getMovieUrl())
}

// GetRmdMovie 获取热门
func (param *DouBanRequest) GetRmdMovie() BaseSubject {
	param.Tag = "热门"
	param.Sort = "recommend"
	return getMovieData(param.getMovieUrl())
}

// GetRankMovie 获取口碑 搞评价电影
func (param *DouBanRequest) GetRankMovie() BaseSubject {
	param.Tag = "热门"
	param.Sort = "rank"
	return getMovieData(param.getMovieUrl())
}

// GetMovie 根据标签获取电影
func (param *DouBanRequest) GetMovie(tagType string) BaseSubject {
	if "" == tagType {
		return BaseSubject{}
	}
	param.Tag = tagType
	return getMovieData(param.getMovieUrl())
}

// GetCustomMovie 自定义参数...
func (param *DouBanRequest) GetCustomMovie() BaseSubject {
	return getMovieData(param.getMovieUrl())
}
