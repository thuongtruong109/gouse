package gouse

import (
	"github.com/go-redis/redis/v8"
)

func ConnectRedis(addr, pass string, dbNo ...int) *redis.Client {
	nums := 0

	if len(dbNo) > 0 {
		nums = dbNo[0]
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       nums,
	})
	return rdb
}

// import (
// 	"github.com/go-redis/redis"
// 	"github.com/spf13/viper"
// )

// var RedisClient *redis.Client

// func InitRedisClient() {
// 	RedisClient = redis.NewClient(&redis.Options{
// 		Addr:     viper.GetString("redis.host"),
// 		Password: viper.GetString("redis.password"),
// 		DB:       viper.GetInt("redis.db"),
// 	})
// }

// func CloseRedisClient() {
// 	_ = RedisClient.Close()
// }

// import (
// 	"context"
// 	"log"

// 	"github.com/minio/minio-go/v7"
// 	"github.com/minio/minio-go/v7/pkg/credentials"
// )

// type MinioClient = *minio.Client

// type MinioConf struct {
// 	Endpoint  string
// 	AccessKey string
// 	SecretKey string
// 	UseSSL    bool
// 	Bucket    string
// 	Location  string
// }

// func Minio(conf MinioConf) (MinioClient, error) {
// 	ctx := context.Background()

// 	minioClient, errInit := minio.New(conf.Endpoint, &minio.Options{
// 		Creds:  credentials.NewStaticV4(conf.AccessKey, conf.SecretKey, ""),
// 		Secure: conf.UseSSL,
// 	})
// 	if errInit != nil {
// 		log.Fatalln(errInit)
// 	}

// 	exists, errBucketExists := minioClient.BucketExists(ctx, conf.Bucket)
// 	if errBucketExists == nil && exists {
// 		log.Printf("We already own %s\n", conf.Bucket)
// 		return minioClient, errInit
// 	} else {
// 		err := minioClient.MakeBucket(ctx, conf.Bucket, minio.MakeBucketOptions{Region: conf.Location})
// 		if err != nil {
// 			return minioClient, errInit
// 		} else {
// 			log.Printf("Successfully created %s\n", conf.Bucket)
// 			return minioClient, errInit
// 		}
// 	}
// }

// import (
// 	"context"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// 	"go.mongodb.org/mongo-driver/mongo/readpref"
// )

// type MongoConnection struct {
// 	client *mongo.Client
// 	ctx    context.Context
// }

// func NewMongo(context context.Context, uri string) (*MongoConnection, error) {
// 	client, err := Connect(uri)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &MongoConnection{
// 		client: client,
// 		ctx:    context,
// 	}, nil
// }

// func (connection *MongoConnection) Connect(uri string) (*mongo.Client, error) {
// 	client, err := mongo.Connect(connection.context.TODO(), options.Client().ApplyURI(uri))
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := connection.Ping(); err != nil {
// 		return nil, err
// 	}

// 	return client, nil
// }

// func (connection *MongoConnection) Ping() error {
// 	return connection.Client.Ping(connection.context.Background(), readpref.Primary())
// }

// func (connection *MongoConnection) Close() error {
// 	return connection.Client.Disconnect(connection.context.Background())
// }

// func (connection *MongoConnection) CreateDatabase(dbName string) *mongo.Database {
// 	return connection.Client.Database(dbName)
// }

// func (connection *MongoClient) CreateCollection(dbName string, collectionName string) (*mongo.Collection, error) {
// 	return connection.Client.Database(dbName).Collection(collectionName), nil
// }

// func (connection *MongoConnection) CountDocuments(collection *mongo.Collection, filter interface{}) (int64, error) {
// 	return collection.CountDocuments(connection.context.Background(), filter)
// }

// func (connection *MongoConnection) InsertOne(collection *mongo.Collection, document interface{}) (*mongo.InsertOneResult, error) {
// 	return collection.InsertOne(connection.context.Background(), document)
// }

// func (connection *MongoConnection) InsertMany(collection *mongo.Collection, documents []interface{}) (*mongo.InsertManyResult, error) {
// 	return collection.InsertMany(connection.context.Background(), documents)
// }

// func (connection *MongoConnection) FindOne(collection *mongo.Collection, filter interface{}) *mongo.SingleResult {
// 	return collection.FindOne(connection.context.Background(), filter)
// }

// func (connection *MongoConnection) FindMany(collection *mongo.Collection, filter interface{}) (*mongo.Cursor, error) {
// 	return collection.Find(connection.context.Background(), filter)
// }

// func (connection *MongoConnection) UpdateOne(collection *mongo.Collection, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
// 	return collection.UpdateOne(connection.context.Background(), filter, update)
// }

// func (connection *MongoConnection) UpdateMany(collection *mongo.Collection, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
// 	return collection.UpdateMany(connection.context.Background(), filter, update)
// }

// func (connection *MongoConnection) DeleteOne(collection *mongo.Collection, filter interface{}) (*mongo.DeleteResult, error) {
// 	return collection.DeleteOne(connection.context.Background(), filter)
// }

// func (connection *MongoConnection) DeleteMany(collection *mongo.Collection, filter interface{}) (*mongo.DeleteResult, error) {
// 	return collection.DeleteMany(connection.context.Background(), filter)
// }

// func (connection *MongoConnection) Aggregate(collection *mongo.Collection, pipeline interface{}) (*mongo.Cursor, error) {
// 	return collection.Aggregate(connection.context.Background(), pipeline)
// }

// import (
// 	"database/sql"
// 	"fmt"
// 	"time"
// )

// type MySQLConnection struct {
// 	db *sql.DB
// }

// func NewMySQLConnection(uri string) (*MySQLConnection, error) {
// 	db, err := Connect(uri)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &MySQLConnection{
// 		db: db,
// 	}, nil
// }

// func Connect(uri string) (*sql.DB, error) {
// 	db, err := sql.Open("mysql", uri)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := db.Ping(); err != nil {
// 		return nil, err
// 	}

// 	return db, nil
// }

// func (connection *MySQLConnection) Close() error {
// 	return connection.db.Close()
// }

// func (connection *MySQLConnection) CreateDatabase(dbName string) (*sql.DB, error) {
// 	_, err := connection.db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
// 	if err != nil {
// 		return nil, err
// 	}

// 	return connection.db, nil
// }

// func (connection *MySQLConnection) CreateTable(dbName string, tableName string, columns string) error {
// 	_, err := connection.db.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s.%s (%s)", dbName, tableName, columns))
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (connection *MySQLConnection) Insert(dbName string, tableName string, columns string, values string) error {
// 	_, err := connection.db.Exec(fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (%s)", dbName, tableName, columns, values))
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (connection *MySQLConnection) Query(query string) (*sql.Rows, error) {
// 	return connection.db.Query(query)
// }

// func (connection *MySQLConnection) QueryRow(query string) *sql.Row {
// 	return connection.db.QueryRow(query)
// }

// func (connection *MySQLConnection) Exec(query string) (sql.Result, error) {
// 	return connection.db.Exec(query)
// }

// func (connection *MySQLConnection) Prepare(query string) (*sql.Stmt, error) {
// 	return connection.db.Prepare(query)
// }

// func (connection *MySQLConnection) Begin() (*sql.Tx, error) {
// 	return connection.db.Begin()
// }

// func (connection *MySQLConnection) SetMaxOpenConns(n int) {
// 	connection.db.SetMaxOpenConns(n)
// }

// func (connection *MySQLConnection) SetMaxIdleConns(n int) {
// 	connection.db.SetMaxIdleConns(n)
// }

// func (connection *MySQLConnection) SetConnMaxLifetime(d time.Duration) {
// 	connection.db.SetConnMaxLifetime(d)
// }

// func (connection *MySQLConnection) SetConnMaxIdleTime(d time.Duration) {
// 	connection.db.SetConnMaxIdleTime(d)
// }

// import (
// 	"encoding/base64"
// 	"event-tracking/config"
// 	"event-tracking/pkg/logger"
// 	"fmt"
// 	"log"
// 	"time"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// type Postgres struct {
// 	Username    string `yaml:"username" mapstructure:"username"`
// 	Password    string `yaml:"password" mapstructure:"password"`
// 	Database    string `yaml:"database" mapstructure:"database"`
// 	Host        string `yaml:"host" mapstructure:"host"`
// 	Port        int    `yaml:"port" mapstructure:"port"`
// 	Schema      string `yaml:"schema" mapstructure:"schema"`
// 	MaxIdleConn int    `yaml:"max_idle_conn" mapstructure:"max_idle_conn"`
// 	MaxOpenConn int    `yaml:"max_open_conn" mapstructure:"max_open_conn"`
// }

// func InitPostgres(config config.Postgres) (*gorm.DB, error) {
// 	var (
// 		db  *gorm.DB
// 		err error
// 	)
// 	password, _ := base64.StdEncoding.DecodeString(config.Password)
// 	config.Password = string(password)
// 	logger.GLogger.Infof("connecting postgres database, user %s, dbname %s, host: %s", config.Username, config.Database, config.Host)

// 	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d ", config.Host, config.Username, config.Password, config.Database, config.Port)
// 	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Default().Println("connect postgres err:", err)
// 		return db, err
// 	}
// 	if config.IsDebug {
// 		db.Debug()
// 	}

// 	dbc, err := db.DB()
// 	if err == nil {
// 		if config.MaxIdleConn > 0 {
// 			dbc.SetMaxIdleConns(config.MaxIdleConn)
// 		}
// 		if config.MaxOpenConn > 0 {
// 			dbc.SetMaxOpenConns(config.MaxOpenConn)
// 		}
// 		if config.MaxIdleTime > 0 {
// 			dbc.SetConnMaxIdleTime(time.Duration(config.MaxIdleTime * int(time.Second)))
// 		}
// 		if config.MaxLifeTime > 0 {
// 			dbc.SetConnMaxLifetime(time.Duration(int(time.Second) * config.MaxLifeTime))
// 		}
// 	}

// 	logger.GLogger.Info("connect postgres successfully")
// 	return db, err
// }

// import (
// 	"database/sql"
// 	"fmt"
// 	"github.com/thuongtruong109/gouse/types"
// 	"log"
// 	"reflect"
// 	"strings"
// )

// type Postgres struct {
// 	Host     string
// 	Port     int
// 	Username string
// 	Password string
// 	Database string
// }

// func (p Postgres) Connect() *sql.DB {
// 	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", p.Host, p.Port, p.Username, p.Password, p.Database))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return db
// }

// func (p Postgres) Insert(data interface{}) {
// 	db := p.Connect()
// 	defer db.Close()

// 	v := reflect.ValueOf(data)
// 	t := v.Type()

// 	var fields []string
// 	var values []string

// 	for i := 0; i < v.NumField(); i++ {
// 		fieldName := t.Field(i).Name
// 		fieldValue := v.Field(i).Interface()
// 		fields = append(fields, fieldName)
// 		values = append(values, types.ToString(fieldValue))
// 	}

// 	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES ('%s')", t.Name(), strings.Join(fields, ", "), strings.Join(values, "', '"))
// 	_, err := db.Exec(query)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func (p Postgres) Update(data interface{}) {
// 	db := p.Connect()
// 	defer db.Close()

// 	v := reflect.ValueOf(data)
// 	t := v.Type()

// 	var fields []string
// 	var values []string

// 	for i := 0; i < v.NumField(); i++ {
// 		fieldName := t.Field(i).Name
// 		fieldValue := v.Field(i).Interface()
// 		fields = append(fields, fieldName)
// 		values = append(values, types.ToString(fieldValue))
// 	}

// 	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = %s", t.Name(), strings.Join(fields, ", "), strings.Join(values, "', '"))
// 	_, err := db.Exec(query)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func (p Postgres) Delete(data interface{}) {
// 	db := p.Connect()
// 	defer db.Close()

// 	v := reflect.ValueOf(data)
// 	t := v.Type()

// 	var values []string

// 	for i := 0; i < v.NumField(); i++ {
// 		fieldValue := v.Field(i).Interface()
// 		values = append(values, types.ToString(fieldValue))
// 	}

// 	query := fmt.Sprintf("DELETE FROM %s WHERE id = %s", t.Name(), strings.Join(values, "', '"))
// 	_, err := db.Exec(query)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func (p Postgres) Select(data interface{}) {
// 	db := p.Connect()
// 	defer db.Close()

// 	v := reflect.ValueOf(data)
// 	t := v.Type()

// 	var values []string

// 	for i := 0; i < v.NumField(); i++ {
// 		fieldValue := v.Field(i).Interface()
// 		values = append(values, types.ToString(fieldValue))
// 	}

// 	query := fmt.Sprintf("SELECT * FROM %s WHERE id = %s", t.Name(), strings.Join(values, "', '"))
// 	rows, err := db.Query(query)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for rows.Next() {
// 		err := rows.Scan(&data)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }

// func (p Postgres) SelectAll(data interface{}) {
// 	db := p.Connect()
// 	defer db.Close()

// 	v := reflect.ValueOf(data)
// 	t := v.Type()

// 	query := fmt.Sprintf("SELECT * FROM %s", t.Name())
// 	rows, err := db.Query(query)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for rows.Next() {
// 		err := rows.Scan(&data)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }

// func (p Postgres) SelectWhere(data interface{}, where string) {
// 	db := p.Connect()
// 	defer db.Close()

// 	v := reflect.ValueOf(data)
// 	t := v.Type()

// 	query := fmt.Sprintf("SELECT * FROM %s WHERE %s", t.Name(), where)
// 	rows, err := db.Query(query)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for rows.Next() {
// 		err := rows.Scan(&data)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }
