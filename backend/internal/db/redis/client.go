package redis

import (
	"context"
	"lemon-traefik-limiter-backend/internal/config"
	"lemon-traefik-limiter-backend/internal/util"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	clientInstance *redis.Client
	ipCountScriptInstance *redis.Script
	once           sync.Once
)

const ipCountScript = `
local current = redis.call('INCR', KEYS[1])
if current == 1 then
    redis.call('expire', KEYS[1], ARGV[1])
end
return current
`

func GetRedisClient() *redis.Client {
	once.Do(func() {
		clientInstance = redis.NewClient(&redis.Options{
			Addr:     config.GlobalConfig.Redis.Addr,
			Password: config.GlobalConfig.Redis.Password,
			DB:       config.GlobalConfig.Redis.DB,
		})

		ctx := context.Background()
		ipCountScriptInstance = redis.NewScript(ipCountScript)
		ipCountScriptInstance.Load(ctx, clientInstance)
	})
	return clientInstance
}

type RedisService struct {
	client *redis.Client
	script *redis.Script
}

func GetRedisService(client *redis.Client) *RedisService {
	return &RedisService{
		client: client,
		script: ipCountScriptInstance,
	}
}

func (r *RedisService) GetIpCallCountPerSecond(ip string) (int64, error) {
	ctx := context.Background()
	key := "ip_count:"+ util.GetCurrentTimeString(util.Seconds) + ":" + ip
	expireSeconds := 1
	
	result, err := r.script.Run(ctx, r.client, []string{key}, expireSeconds).Int64()	
	if err != nil {
		return 0, err
	}
	return result, nil
}

