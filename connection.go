package gouse

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/go-redis/redis/v8"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

/* Connect to Redis */

func ConnectRedis(addr, pass string, dbNo ...int) *redis.Client {
	nums := 0

	if len(dbNo) > 0 {
		nums = dbNo[0]
	}

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       nums,
	})

	pong, err := client.Ping(client.Context()).Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Redis: ", pong)

	defer client.Close()

	return client
}

func ConnectRedisUri(uri string) *redis.Client {
	opt, err := redis.ParseURL(uri)
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)

	pong, err := client.Ping(client.Context()).Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Redis: ", pong)

	defer client.Close()

	return client
}

/* Connect to Postgres */

type PgDB struct {
	Client *sql.DB
}

func ConnectPostgres(dsn string) (*PgDB, error) {
	const maxOpenDbConn = 10
	const maxIdleDbConn = 5
	const maxDbLifeTime = 5 * time.Minute

	dbConn := &PgDB{}

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetMaxIdleConns(maxIdleDbConn)
	db.SetConnMaxLifetime(maxDbLifeTime)

	err = db.Ping()
	if err != nil {
		fmt.Println("*** Pinged database successfully ***")
		log.Fatal(err)
	}

	dbConn.Client = db

	defer func() {
		db.Close()
	}()

	return dbConn, nil
}

/* Connect to MinIO */

type MinioConf struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	UseSSL    bool
	Bucket    string
	Location  string
}

func ConnectMinio(ctx context.Context, conf MinioConf) (*minio.Client, error) {
	minioClient, errInit := minio.New(conf.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.AccessKey, conf.SecretKey, ""),
		Secure: conf.UseSSL,
	})
	if errInit != nil {
		log.Fatalln(errInit)
	}

	exists, errBucketExists := minioClient.BucketExists(ctx, conf.Bucket)
	if errBucketExists == nil && exists {
		log.Printf("We already own %s\n", conf.Bucket)
		return minioClient, errInit
	} else {
		err := minioClient.MakeBucket(ctx, conf.Bucket, minio.MakeBucketOptions{
			Region: conf.Location,
		})
		if err != nil {
			return minioClient, errInit
		} else {
			log.Printf("Successfully created %s\n", conf.Bucket)
			return minioClient, errInit
		}
	}
}

/* Connect to MongoDB */

func ConnectMongo(ctx context.Context, uri string) (*mongo.Client, error) {
	if uri == "" {
		return nil, fmt.Errorf("uri connection is not set")
	}

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	defer func() {
		client.Disconnect(ctx)
	}()

	return client, nil
}
