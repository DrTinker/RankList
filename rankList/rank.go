package ranklist

import (
	"RankList/models"
	"RankList/utils"

	"github.com/go-redis/redis"
)

type rankList struct {
	key string
	rdb *redis.Client
}

func InitRankList(addr, pwd, key string) *rankList {
	return &rankList{
		key: key,
		rdb: utils.RedisInit(addr, pwd),
	}
}

func (r *rankList) AddElement(element *models.RankElement) error {
	z := redis.Z{
		Member: element.Member,
		Score:  element.Score,
	}
	_, err := r.rdb.ZAdd(r.key, z).Result()
	return err
}

func (r *rankList) AddElementBatch(elements []*models.RankElement) (int64, error) {
	zList := make([]redis.Z, len(elements))
	for i := 0; i < len(elements); i++ {
		zList[i] = redis.Z{
			Member: elements[i].Member,
			Score:  elements[i].Score,
		}
	}
	return r.rdb.ZAdd(r.key, zList...).Result()
}

func (r *rankList) IncrScore(element *models.RankElement) (float64, error) {
	member := utils.Strval(element.Member)
	return r.rdb.ZIncrBy(r.key, element.Score, member).Result()
}

func (r *rankList) GetTopRank(top int64) ([]string, error) {
	return r.rdb.ZRevRange(r.key, 0, int64(top-1)).Result()
}

func (r *rankList) GetRankBetween(start, end int64) ([]string, error) {
	return r.rdb.ZRevRange(r.key, int64(start), int64(end)).Result()
}

func (r *rankList) RemoveElement(element *models.RankElement) error {
	_, err := r.rdb.ZRem(r.key, element.Member).Result()
	return err
}

func (r *rankList) RemoveElementBatch(elements []*models.RankElement) (int64, error) {
	members := make([]interface{}, len(elements))
	for i := 0; i < len(elements); i++ {
		members[i] = elements[i].Member
	}
	return r.rdb.ZRem(r.key, members...).Result()
}
