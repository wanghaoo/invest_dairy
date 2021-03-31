package common

import (
  "gopkg.in/redis.v5"
)

var RedisCache *redis.Client

const (
  ADDR = "great-redis-out.redis.ap-south-1.rds.aliyuncs.com:6379"
  PWD  = "20RG8Yh72E"
)

type (
  ArticleRecVO struct {
    ID                 int
    ArticleID          int
    RecArticleID       int
    Ctime              int
    RecArticleTitle    string
    RecArticleType     int
    RecArticleColumnID int
  }
)

func OpenRedis() *redis.Client {
  RedisCache = redis.NewClient(&redis.Options{Addr: ADDR, Password: PWD, DB: 36})
  Mlog.Debug("Redis conection set up successfully")
  return RedisCache
}

func CloseRedis() {
  RedisCache.Close()
}
