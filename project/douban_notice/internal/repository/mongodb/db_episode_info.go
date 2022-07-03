package mongodb

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"top.lel.dn/main/pkg/db/mongodb"
	"top.lel.dn/main/pkg/logger"
)

const (
	MovieInfoCollect = "lel_db_episode_info"
)

type EpisodeInfo struct {
	Id           string    `bson:"_id" json:"id"`
	TagType      string    `bson:"tag_type"`
	TagList      []TagInfo `bson:"tag_list"` // one to many.
	DbId         string    `bson:"db_id"`
	Title        string    `bson:"title"`
	Url          string    `bson:"url"`
	Cover        string    `bson:"cover"`
	Rate         string    `bson:"rate"`
	IsNew        bool      `bson:"is_new"`
	Playable     bool      `bson:"playable"`
	CoverX       int       `bson:"cover_x"`
	CoverY       int       `bson:"cover_y"`
	EpisodesInfo string    `bson:"episodes_info"`
	// 需要解析页面
	PublicDate *time.Time `bson:"public_date"`
	Torrent    string     `bson:"torrent"`
	DriverUrl  string     `bson:"driver_url"`
	// now start is subject abstract info
	Actors           []string `bson:"actors" fmt:"actors"`
	Blacklisted      string   `bson:"blacklisted" fmt:"blacklisted"`
	CollectionStatus string   `bson:"collection_status" fmt:"collection_status"`
	Directors        []string `bson:"directors" fmt:"directors"`
	Duration         string   `bson:"duration" fmt:"duration"`
	EpisodesCount    string   `bson:"episodes_count" fmt:"episodes_count"`
	IsTv             bool     `bson:"is_tv" fmt:"is_tv"`
	Region           string   `bson:"region" fmt:"region"`
	ReleaseYear      string   `bson:"release_year" fmt:"release_year"`
	Star             int      `bson:"star" fmt:"star"`
	Subtype          string   `bson:"subtype" fmt:"subtype"`
	Types            []string `bson:"types" fmt:"types"`
	ShortComment     struct {
		Author  string `bson:"author" fmt:"author"`
		Content string `bson:"content" fmt:"content"`
	} `bson:"short_comment"`
	Created    time.Time `bson:"created"`
	Creator    string    `bson:"creator"`
	Updater    string    `bson:"updater"`
	UpdateTime time.Time `bson:"update_time"`
}

func (m *EpisodeInfo) SaveOrUpd(ctx *mongodb.MongoCtx) {
	var info *EpisodeInfo
	_ = ctx.Collection.FindOne(ctx.TODO, bson.D{{"db_id", m.DbId}}).Decode(&info)
	if info != nil {
		// upd
		filter := bson.D{{"db_id", m.DbId}}
		tagList := append(info.TagList, m.TagList...)
		update := bson.D{{
			"$set", bson.D{
				{"tag_type", m.TagType},
				{"tag_list", distinct(tagList)},
				{"title", m.Title},
				{"url", m.Url},
				{"cover", m.Cover},
				{"rate", m.Rate},
				{"is_new", m.IsNew},
				{"playable", m.Playable},
				{"cover_x", m.CoverX},
				{"cover_y", m.CoverY},
				{"episodes_info", m.EpisodesInfo},
				{"torrent", m.Torrent},
				{"driver_url", m.DriverUrl},
				// subject
				{"actors", m.Actors},
				{"blacklisted", m.Blacklisted},
				{"collection_status", m.CollectionStatus},
				{"directors", m.Directors},
				{"duration", m.Duration},
				{"episodes_count", m.EpisodesCount},
				{"is_tv", m.IsTv},
				{"region", m.Region},
				{"release_year", m.ReleaseYear},
				{"star", m.Star},
				{"subtype", m.Subtype},
				{"types", m.Types},
				{"short_comment", m.ShortComment},
				{"updater", m.Updater},
				{"update_time", m.UpdateTime},
			},
		}}
		_, err := ctx.Collection.UpdateOne(ctx.TODO, filter, update)
		if err != nil {
			logger.Warn("upd movie info err, e: " + fmt.Sprint(err))
		}
	} else {
		// insert
		_, _ = ctx.Collection.InsertOne(ctx.TODO, m)
	}
}

// 根据标签id去重, 如果切片包含空串将会被剔除
func distinct(tagList []TagInfo) []TagInfo {
	// 设置1位缓冲用于重复比较
	data := make([]TagInfo, 1)
	for _, item := range tagList {
		if isDistinct(data, item) {
			continue
		}
		data = append(data, item)
	}
	return data[1:]
}

func isDistinct(data []TagInfo, item TagInfo) bool {
	for _, raw := range data {
		if raw.Id == item.Id {
			return true
		}
	}
	return false
}
