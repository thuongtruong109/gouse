package samples

import (
	// "github.com/minio/minio-go/v7"
	"github.com/thuongtruong109/gouse"
)

/*
Description: Connect to Redis(using parameters)
Input params: (address string, password string, dbNo int)
*/
func ConnectRedis() {
	redis := gouse.Redis("localhost:6379", "password", 0)

	// ... do something with redis
	redis.Set(redis.Context(), "key", "value", 0)
}

/*
Description: Connect to Redis cloud (using uri)
Input params: (uri string)
*/
func ConnectRedisUri() {
	redis := gouse.RedisCloud("redis://localhost:6379/0")

	// ... do something with redis
	redis.Set(redis.Context(), "key", "value", 0)

}

/*
Description: Connect to Postgres
Input params: (uri string)
*/
func ConnectPostgres() {
	pg, err := gouse.Postgres("localhost:5432")
	if err != nil {
		panic(err)
	}

	// ... do something with pg
	pg.Client.QueryRow("SELECT * FROM users")
}

/*
Description: Connect to MongoDB
Input params: (context.Context, uri string)
*/
func ConnectMongo() {
	ctx := gouse.CtxBg

	mongoClient, err := gouse.Mongo(ctx, "mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	// ... do something with mongo
	mongoClient.Database("test").Collection("users").FindOne(ctx, nil)
}

/*
Description: Connect to Minio
Input params: (context.Context, config gouse.MinioConfig)
*/
func ConnectMinio() {
	// ctx := gouse.CtxBg

	// minioClient, err := gouse.ConnectMinio(ctx, gouse.MinioConf{
	// 	Endpoint:  "localhost:9000",
	// 	AccessKey: "minio",
	// 	SecretKey: "minio123",
	// 	UseSSL:    false,
	// 	Bucket:    "bucket",
	// 	Location:  "us-east-1",
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// // ... do something with minio
	// minioClient.FPutObject(ctx, "bucket", "object", "file", minio.PutObjectOptions{
	// 	ContentType: "application/octet-stream",
	// })
}
