package mongodb

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"top.lel.dn/main/pkg/db/mongodb"
	"top.lel.dn/main/pkg/logger"
)

const (
	MovieInfoCollect = "lel_db_movie_info"
)

type MovieInfo struct {
	Id           string `bson:"_id" json:"id"`
	MvTagId      string `bson:"mv_tag_id"`
	MvTagInfo    string `bson:"mv_tag_info"`
	DbId         string `bson:"db_id"`
	Title        string `bson:"title"`
	Url          string `bson:"url"`
	Cover        string `bson:"cover"`
	Rate         string `bson:"rate"`
	IsNew        bool   `bson:"is_new"`
	Playable     bool   `bson:"playable"`
	CoverX       int    `bson:"cover_x"`
	CoverY       int    `bson:"cover_y"`
	EpisodesInfo string `bson:"episodes_info"`
	// 需要解析页面
	PublicDate *time.Time `bson:"public_date"`
	Torrent    string     `bson:"torrent"`
	DriverUrl  string     `bson:"driver_url"`
	Created    time.Time  `bson:"created"`
	Creator    string     `bson:"creator"`
	Updater    string     `bson:"updater"`
	UpdateTime time.Time  `bson:"update_time"`
}

func (m *MovieInfo) SaveOrUpd(ctx *mongodb.MongoCtx) {
	var info *MovieInfo
	_ = ctx.Collection.FindOne(ctx.TODO, bson.D{{"db_id", m.DbId}}).Decode(&info)
	if info != nil {
		// upd
		filter := bson.D{{"db_id", m.DbId}}
		update := bson.D{{
			"$set", bson.D{
				{"mv_tag_id", m.MvTagId},
				{"mv_tag_info", m.MvTagInfo},
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
