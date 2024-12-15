
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Connection' />


```go
import (
	"github.com/minio/minio-go/v7"
	"github.com/thuongtruong109/gouse"
)
```

## 1. Connect redis

Description: Connect to Redis(using parameters)<br>Input params: (address string, password string, dbNo int)<br>

```go
func ConnectRedis() {
	redis := gouse.ConnectRedis("localhost:6379", "password", 0)

	// ... do something with redis
	redis.Set(redis.Context(), "key", "value", 0)
}
```

## 2. Connect redis uri

Description: Connect to Redis(using uri)<br>Input params: (uri string)<br>

```go
func ConnectRedisUri() {
	redis := gouse.ConnectRedisUri("redis://localhost:6379/0")

	// ... do something with redis
	redis.Set(redis.Context(), "key", "value", 0)

}
```

## 3. Connect postgres

Description: Connect to Postgres<br>Input params: (uri string)<br>

```go
func ConnectPostgres() {
	pg, err := gouse.ConnectPostgres("localhost:5432")
	if err != nil {
		panic(err)
	}

	// ... do something with pg
	pg.Client.QueryRow("SELECT * FROM users")
}
```

## 4. Connect mongo

Description: Connect to MongoDB<br>Input params: (context.Context, uri string)<br>

```go
func ConnectMongo() {
	ctx := gouse.CtxBg

	mongoClient, err := gouse.ConnectMongo(ctx, "mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	// ... do something with mongo
	mongoClient.Database("test").Collection("users").FindOne(ctx, nil)
}
```

## 5. Connect minio

Description: Connect to Minio<br>Input params: (context.Context, config gouse.MinioConfig)<br>

```go
func ConnectMinio() {
	ctx := gouse.CtxBg

	minioClient, err := gouse.ConnectMinio(ctx, gouse.MinioConf{
		Endpoint:  "localhost:9000",
		AccessKey: "minio",
		SecretKey: "minio123",
		UseSSL:    false,
		Bucket:    "bucket",
		Location:  "us-east-1",
	})
	if err != nil {
		panic(err)
	}

	// ... do something with minio
	minioClient.FPutObject(ctx, "bucket", "object", "file", minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
}
```
