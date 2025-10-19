package main

import (
	"log"
	"sync"

	"urlShortner/config"
	"urlShortner/handler"
	"urlShortner/pkg/snowflake"
	"urlShortner/repository/cache"
	"urlShortner/repository/database"
	"urlShortner/service"
	"urlShortner/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	configuration := config.NewConfig()
	conf, err := configuration.Load()
	if err != nil {
		log.Fatalln(err.Error())
	}

	postgresDB := storage.NewPostgres(conf.Postgres)
	err = postgresDB.Connect()
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = postgresDB.MigrateUp()
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = snowflake.Init(36)
	if err != nil {
		log.Fatalln(err.Error())
	}

	redisClient := storage.NewRedis(conf.Redis)
	redisCache := cache.NewRedisCache(redisClient)

	mapCache, err := cache.NewMapCache(map[string]string{}, &sync.Mutex{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	operator, err := initCacheOperator(redisCache, mapCache)
	if err != nil {
		log.Fatalln(err.Error())
	}

	newCache, err := operator.NewCache(cache.Operator(conf.CacheOperator))
	if err != nil {
		log.Fatalln(err.Error())
	}

	newPostgresDB := database.NewPostgresDB(postgresDB)
	newService := service.NewService(conf, newPostgresDB, newCache)
	newHandler := handler.NewHandler(newService)

	server := gin.New()
	server.Use(gin.Logger())

	v1 := server.Group("/v1")
	v1.GET("/health", newHandler.Health)
	v1.GET("/:short", newHandler.Redirect)
	v1.POST("/create", newHandler.Create)

	err = server.Run("localhost:8080")
	if err != nil {
		log.Println(err.Error())
	}
}

func initCacheOperator(redisCache *cache.RedisCache, mapCache *cache.MapCache) (*cache.Operators, error) {
	return cache.InitialCacheOperators(map[cache.Operator]cache.Cache{
		cache.REDISOPERATOR: redisCache,
		cache.MAPOPERATOR:   mapCache,
	}), nil
}

// if www.hello/1 --> https://www.alleycat.org/wp-content/uploads/2019/03/FELV-cat.jpg

// 0 health endpoint ---> get

// 1 user click on a short url ---> get
// short url = inserted url in storage having an id encoded with base62
// if is cache? return
// if is in db? return
// if none error

// 2 user adds a new url for having a short one ---> post
// if long url in cache? return
// if long url in db? return
// if none? insert and return the short one
