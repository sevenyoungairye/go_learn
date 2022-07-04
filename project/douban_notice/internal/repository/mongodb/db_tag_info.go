package mongodb

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"top.lel.dn/main/pkg/db/mongodb"
	"top.lel.dn/main/pkg/logger"
	"top.lel.dn/main/pkg/serial"
)

const (
	lelDbTagInfo = "lel_db_tag_info"
)

// TagInfo 应该把标签作为参数, 作为mongodbCtx的函数来操作
type TagInfo struct {
	Id         string    `bson:"_id" json:"_id"`
	TagType    string    `bson:"tag_type" json:"tag_type"`
	TagName    string    `bson:"tag_name" json:"tag_name"`
	CreateTime time.Time `bson:"create_time" json:"create_time"`
	Creator    string    `bson:"creator" json:"creator"`
	Updater    string    `bson:"updater" json:"updater"`
	UpdateTime time.Time `bson:"update_time" json:"update_time"`
}

func (tagInfo *TagInfo) DeleteById(ctx mongodb.MongoCtx) error {
	_, err := ctx.Collection.DeleteOne(ctx.TODO, bson.D{{Key: "_id", Value: tagInfo.Id}})
	return err
}

func (tagInfo *TagInfo) UpdById() {
	ctx := mongodb.GetCtxCollection(lelDbTagInfo)
	// bson.D{{"_id", tagInfo.Id}}
	result, err := ctx.Collection.UpdateByID(ctx.TODO, tagInfo.Id, buildUpdData(tagInfo))
	if err != nil {
		logger.Warn(fmt.Sprint(err))
		return
	}
	logger.Info(fmt.Sprint(result))

}

// SaveOrUpdTag pls see this document.
// https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Collection.UpdateOne
func (tagInfo *TagInfo) SaveOrUpdTag() {
	mongoCtx := mongodb.GetCtxCollection(lelDbTagInfo)
	collection := mongoCtx.Collection
	var tmp *TagInfo

	filter := bson.D{{Key: "tag_type", Value: tagInfo.TagType}, {Key: "tag_name", Value: tagInfo.TagName}}
	_ = collection.FindOne(mongoCtx.TODO, filter).Decode(&tmp)

	logger.Info(fmt.Sprintf("get one: %s", serial.Object2Json(tmp)))
	if tmp != nil {
		update := buildUpdData(tagInfo)
		_, err := collection.UpdateOne(mongoCtx.Context, filter, update)
		if err != nil {
			logger.Warn(fmt.Sprint(err))
		}
	} else {
		res, err := collection.InsertOne(mongoCtx.TODO, tagInfo)
		if err != nil {
			logger.Warn(fmt.Sprint(err))
			return
		}
		val := res.InsertedID
		tagInfo.Id = fmt.Sprint(val)
	}

	//d := bson.D{{"$set", bson.D{{"tag_name", "dd"}}}}
	// 根据filter 修改
	// result, err := collection.UpdateOne(mongoCtx.Context, bson.D{{"_id", tagInfo.Id}}, tagInfo, options.Update().SetUpsert(true))
	mongoCtx.Release()
}

func buildUpdData(tagInfo *TagInfo) bson.D {
	return bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "tag_type", Value: tagInfo.TagType},
				{Key: "tag_name", Value: tagInfo.TagName},
				{Key: "updater", Value: tagInfo.Updater},
				{Key: "update_time", Value: tagInfo.UpdateTime},
			},
		},
	}
}

func (tagInfo *TagInfo) GetAll(tagType string) []TagInfo {
	mongoCtx := mongodb.GetCtxCollection(lelDbTagInfo)
	filter := bson.D{}
	if tagType != "" {
		filter = bson.D{{Key: "tag_type", Value: tagType}}
	}
	data, err := mongoCtx.Collection.Find(mongoCtx.TODO, filter)
	if err != nil {
		return make([]TagInfo, 0)
	}
	retVal := make([]TagInfo, 0)
	for data.TryNext(mongoCtx.Context) {
		s := TagInfo{}
		//item := data.Current.String()
		//serial.Json2Instant(item, &s)
		_ = data.Decode(&s)
		retVal = append(retVal, s)
	}
	mongoCtx.Release()
	return retVal
}
