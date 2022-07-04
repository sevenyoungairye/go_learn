package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"top.lel.dn/main/pkg/db/mongodb"
	"top.lel.dn/main/pkg/serial"
)

const (
	lelDbSubscribe = "lel_db_subscribe"
)

type Subscribe struct {
	Id            string    `bson:"_id"`
	SubscribeName string    `bson:"subscribe_name"`
	SubTagId      string    `bson:"sub_tag_id"`
	SubTagType    string    `bson:"sub_tag_type"`
	SubTagName    string    `bson:"sub_tag_name"`
	SubSort       string    `bson:"sub_sort"`
	Created       time.Time `bson:"created"`
	Creator       string    `bson:"creator"`
	Updater       string    `bson:"updater"`
	UpdateTime    time.Time `bson:"update_time"`
}

func (s *Subscribe) GetList(tagType string) []Subscribe {
	mongoCtx := mongodb.GetCtxCollection(lelDbSubscribe)
	data, err := mongoCtx.Collection.Find(mongoCtx.Context, bson.D{{Key: "sub_tag_type", Value: tagType}})
	if err != nil {
		return make([]Subscribe, 0)
	}
	retVal := make([]Subscribe, 0)
	for data.TryNext(mongoCtx.Context) {
		item := data.Current.String()
		s := Subscribe{}
		serial.Json2Instant(item, &s)
		retVal = append(retVal, s)
	}
	mongoCtx.Release()

	return retVal
}
