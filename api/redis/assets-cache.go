package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
)

func SetcacheAssets(key string, value []*model.Assets) error {
	ctx := context.Background()
	client := database.Redis_cache()

	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	expiration := 1 * time.Hour

	err = client.Set(ctx, key, jsonValue, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func GetcacheAssets(key string) ([]*model.Assets, error) {
	ctx := context.Background()
	client := database.Redis_cache()

	jsonValue, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var value []*model.Assets
	err = json.Unmarshal([]byte(jsonValue), &value)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func LpushmaAssets(key string, valuearray []*model.Assets) (int64, error) {
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
