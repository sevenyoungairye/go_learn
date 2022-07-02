package mongodb

// mongodb context
// get collection and connect context.
// https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#pkg-examples1

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"top.lel.dn/main/pkg/logger"
	"top.lel.dn/main/pkg/yaml"
)

var collectName = "numbers"

func Insert() {
	mongoCtx := GetCtxCollection(collectName)
	ctx, collection := mongoCtx.Context, mongoCtx.Collection

	//res, err := collection.InsertOne(ctx, bson.D{{"name", "pi"}, {"value", 3.14159}})
	//res, err := collection.InsertOne(ctx, bson.D{{"name", "jack"}, {"value", "rose"}})
	user := User{}
	user.setUser("hh", "dd")
	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		logger.Warn(fmt.Sprint(err))
	}
	if res != nil {
		id := res.InsertedID
		logger.Debug(fmt.Sprintf("intert success! id: %s", id))
	}

	mongoCtx.Release()
}

func SelectAll() {
	mongoCtx := GetCtxCollection(collectName)
	ctx, collection := mongoCtx.Context, mongoCtx.Collection
	cursor, err := collection.Find(ctx, bson.D{{"name", "jack"}})
	if err != nil {
		logger.Warn(fmt.Sprint(err))
		return
	}
	for cursor.TryNext(ctx) {
		s := cursor.Current.String()
		logger.Debug(fmt.Sprintf("get data from mongodb: %s", s))
	}
	mongoCtx.Release()
}

// GetCtxCollection get mongodb collection, for crud.
func GetCtxCollection(collectName string) *MongoCtx {
	mongodbCtx := GetMongodbContext()
	ctx, cancel, client := mongodbCtx.Context, mongodbCtx.CancelFunc, mongodbCtx.Client
	collection := client.Database(yaml.GetMongodb().Mongodb.Database).Collection(collectName)
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	if ctx.Err() != nil {
		defer cancel()
	}
	mongodbCtx.Collection = collection
	mongodbCtx.Context = ctx
	mongodbCtx.CancelFunc = cancel
	return mongodbCtx
}

// Release cancel the source
func (mongoCtx *MongoCtx) Release() {
	logger.Debug("touch cancel func, release.")
	defer mongoCtx.CancelFunc()
	defer func() {
		// if some error, system will exit.
		if err := mongoCtx.Client.Disconnect(mongoCtx.Context); err != nil {
			logger.Warn(fmt.Sprint(err))
			panic(err)
		}
	}()
}

// GetMongodbContext get mongodb context.
func GetMongodbContext() *MongoCtx {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(yaml.GetMongodb().Mongodb.Uri))
	if err != nil {
		logger.Warn(fmt.Sprint(err))
		panic(err)
	}

	return &MongoCtx{Context: ctx, CancelFunc: cancel, Client: client, TODO: context.TODO()}
}

// MongoCtx mongodb context object.
type MongoCtx struct {
	Collection *mongo.Collection

	TODO       context.Context // 用于操作增删改
	Context    context.Context // 用于关闭连接
	CancelFunc context.CancelFunc
	Client     *mongo.Client
}

type User struct {
	Name  string `bson:"name"`
	Value string `bson:"value"`
}

func (u *User) setUser(k, v string) {
	u.Name = k
	u.Value = v
}
