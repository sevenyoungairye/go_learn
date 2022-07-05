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
	_ = ctx.Collection.FindOne(ctx.TODO, bson.D{{Key: "db_id", Value: m.DbId}}).Decode(&info)
	if info != nil {
		// upd
		filter := bson.D{{Key: "db_id", Value: info.DbId}}
		tagList := append(info.TagList, m.TagList...)
		update := bson.D{{
			Key: "$set", Value: bson.D{
				{Key: "tag_type", Value: m.TagType},
				{Key: "tag_list", Value: distinct(tagList)},
				{Key: "title", Value: m.Title},
				{Key: "url", Value: m.Url},
				{Key: "cover", Value: m.Cover},
				{Key: "rate", Value: m.Rate},
				{Key: "is_new", Value: m.IsNew},
				{Key: "playable", Value: m.Playable},
				{Key: "cover_x", Value: m.CoverX},
				{Key: "cover_y", Value: m.CoverY},
				{Key: "episodes_info", Value: m.EpisodesInfo},
				{Key: "torrent", Value: m.Torrent},
				{Key: "driver_url", Value: m.DriverUrl},
				{Key: "public_date", Value: m.PublicDate},
				// subject
				{Key: "actors", Value: m.Actors},
				{Key: "blacklisted", Value: m.Blacklisted},
				{Key: "collection_status", Value: m.CollectionStatus},
				{Key: "directors", Value: m.Directors},
				{Key: "duration", Value: m.Duration},
				{Key: "episodes_count", Value: m.EpisodesCount},
				{Key: "is_tv", Value: m.IsTv},
				{Key: "region", Value: m.Region},
				{Key: "release_year", Value: m.ReleaseYear},
				{Key: "star", Value: m.Star},
				{Key: "subtype", Value: m.Subtype},
				{Key: "types", Value: m.Types},
				{Key: "short_comment", Value: m.ShortComment},
				{Key: "updater", Value: m.Updater},
				{Key: "update_time", Value: m.UpdateTime},
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
