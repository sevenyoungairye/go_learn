// @author echo lovely
// @since 2022/06/29
// https://movie.douban.com/j/search_subjects?type=movie&tag=热门&sort=recommend&page_limit=20&page_start=0

package crawler

import (
	"fmt"
	"strings"
	"top.lel.dn/main/pkg/httpclient"
	"top.lel.dn/main/pkg/serial"
	"top.lel.dn/main/pkg/yaml"
)

var douBanConfig *yaml.D

func init() {
	douBanConfig = yaml.GetDouBan()
}

// GetSubject 获取更多剧集明细信息
func (param *DouBanRequest) GetSubject(id string) *Subject {
	return getSub(id)
}

func GetSubject(id string) *Subject {
	return getSub(id)
}
func getSub(id string) *Subject {
	json := httpclient.HttpWithGet(fmt.Sprintf(douBanConfig.DouBan.SubjectAbstractUrl, id), nil)
	sa := SubjectAbstract{}
	serial.Json2Instant(json, &sa)
	return &sa.Subject
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

// GetRankTv 好评热门
func (param *DouBanRequest) GetRankTv() BaseSubject {
	param.Tag = "热门"
	param.Sort = "rank"
	return param.GetTv()
}

// GetRmdTv 推荐热门
func (param *DouBanRequest) GetRmdTv() BaseSubject {
	param.Tag = "热门"
	param.Sort = "recommend"
	return param.GetTv()
}

// GetTv 根据tag获取电视剧
func (param *DouBanRequest) GetTv() BaseSubject {
	if param.Tag == "" {
		return BaseSubject{}
	}
	return getEpisodeData(param.getTvUrl())
}

func (param *DouBanRequest) getTvUrl() string {
	if param.PageStart == "" {
		param.PageStart = "0"
	}
	if param.PageLimit == "" {
		param.PageLimit = "20"
	}
	if param.Sort == "" {
		// 替换sort.
		return fmt.Sprintf(strings.ReplaceAll(douBanConfig.DouBan.Tv.Url, "&sort=%s", ""), param.Tag, param.PageLimit, param.PageStart)
	}
	return fmt.Sprintf(douBanConfig.DouBan.Tv.Url, param.Tag, param.Sort, param.PageLimit, param.PageStart)
}

func (param *DouBanRequest) getMovieUrl() string {
	if param.PageStart == "" {
		param.PageStart = "0"
	}
	if param.PageLimit == "" {
		param.PageLimit = "20"
	}
	if param.Sort == "" {
		return fmt.Sprintf(strings.ReplaceAll(douBanConfig.DouBan.Movie.Url, "&sort=%s", ""), param.Tag, param.PageLimit, param.PageStart)
	}
	return fmt.Sprintf(douBanConfig.DouBan.Movie.Url, param.Tag, param.Sort, param.PageLimit, param.PageStart)
}

func getEpisodeData(url string) BaseSubject {
	json := httpclient.HttpWithGet(url, nil)
	subs := BaseSubject{}
	serial.Json2Instant(json, &subs)
	return subs
}

// GetLastMovie 获取最新
func (param *DouBanRequest) GetLastMovie() BaseSubject {
	param.Tag = "最新"
	param.Sort = "time"

	return getEpisodeData(param.getMovieUrl())
}

// GetRmdMovie 获取热门
func (param *DouBanRequest) GetRmdMovie() BaseSubject {
	param.Tag = "热门"
	param.Sort = "recommend"
	return getEpisodeData(param.getMovieUrl())
}

// GetRankMovie 获取口碑 搞评价电影
func (param *DouBanRequest) GetRankMovie() BaseSubject {
	param.Tag = "热门"
	param.Sort = "rank"
	return getEpisodeData(param.getMovieUrl())
}

// GetMovie 根据标签获取电影
func (param *DouBanRequest) GetMovie() BaseSubject {
	if "" == param.Tag {
		return BaseSubject{}
	}
	return getEpisodeData(param.getMovieUrl())
}

// GetCustomMovie 自定义参数...
func (param *DouBanRequest) GetCustomMovie() BaseSubject {
	return getEpisodeData(param.getMovieUrl())
}
