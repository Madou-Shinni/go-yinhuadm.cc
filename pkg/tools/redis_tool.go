package tools

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"reflect"
	"time"
)

// GetRedisStrResult 获取string类型的redis值
func GetRedisStrResult[T any](rdb redis.Cmdable, key string) (T, error) {
	var result T

	value, err := rdb.Get(key).Result()
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal([]byte(value), &result); err != nil {
		return result, err
	}

	return result, nil
}

// SetRedisStrResult 设置string类型的redis值
func SetRedisStrResult[T any](rdb redis.Cmdable, key string, data T, expiration time.Duration) (T, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return data, err
	}

	if err := rdb.Set(key, jsonData, expiration).Err(); err != nil {
		return data, err
	}

	return data, nil
}

// DelRedisStrResult 删除string类型的redis值
func DelRedisStrResult(rdb redis.Cmdable, key ...string) error {
	if err := rdb.Del(key...).Err(); err != nil {
		return err
	}

	return nil
}

// MGetRedisStrResult 获取string类型的redis值
func MGetRedisStrResult[T any](rdb redis.Cmdable, key ...string) ([]T, error) {
	var result []T
	var zero T
	slice, err := rdb.MGet(key...).Result()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}

	for _, v := range slice {
		var item T
		if v == nil {
			result = append(result, zero)
			continue
		}
		s, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("类型转化错误 %t to string", reflect.TypeOf(v))
		}
		if err := json.Unmarshal([]byte(s), &item); err != nil {
			return result, err
		}

		result = append(result, item)
	}

	return result, nil
}
