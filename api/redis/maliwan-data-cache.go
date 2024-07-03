package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"github.com/redis/go-redis/v9"
)

func GetCacheFilter(ctx context.Context, key string) ([]string, error) {
	client := database.Redis_cache()

	val, err := client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err // Some other error
	}

	var result []string
	err = json.Unmarshal([]byte(val), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func SetCacheFilter(ctx context.Context, key string, value []string) error {
	client := database.Redis_cache()

	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = client.Set(ctx, key, data, 24*time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

func Setcachemaliwan(key string, value []*model.Maliwan_data) error {
	ctx := context.Background()
	client := database.Redis_cache()

	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	expiration := 24 * time.Hour

	err = client.Set(ctx, key, jsonValue, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func SetCountMaliwanDataByBranchCache(key string, value []*model.Maliwan_data) error {
	ctx := context.Background()
	client := database.Redis_cache()

	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	expiration := 24 * time.Hour

	err = client.Set(ctx, key, jsonValue, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func Getcachemaliwan(key string) ([]*model.Maliwan_data, error) {
	ctx := context.Background()
	client := database.Redis_cache()

	jsonValue, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var value []*model.Maliwan_data
	err = json.Unmarshal([]byte(jsonValue), &value)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func GetCountMaliwanDataByBranchCache(key string) ([]*model.Maliwan_data, error) {
	ctx := context.Background()
	client := database.Redis_cache()

	jsonValue, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var value []*model.Maliwan_data
	err = json.Unmarshal([]byte(jsonValue), &value)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func Lpushmaliwan(key string, valuearray []*model.Maliwan_data) (int64, error) {
	ctx := context.Background()
	client := database.Redis_cache()
	var interfaceSlice []interface{}
	for _, item := range valuearray {
		interfaceSlice = append(interfaceSlice, *item)
	}
	value, err := client.LPush(ctx, key, interfaceSlice...).Result()
	if err != nil {
		return 0, err
	}
	return value, nil
}

// func Lpushmaliwan(key string, valuearray []*model.Maliwan_data) error {
// 	ctx := context.Background()
// 	client := database.Redis_cache()

// 	var interfaceSlice []interface{}
// 	for _, item := range valuearray {
// 		interfaceSlice = append(interfaceSlice, item)
// 	}

// 	err := client.LPush(ctx, key, interfaceSlice...).Err()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
