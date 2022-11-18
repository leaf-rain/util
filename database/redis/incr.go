package redis

import (
	"context"
	redis2 "github.com/go-redis/redis/v8"
)

var IncrScript = redis2.NewScript(`if (redis.call('exists', KEYS[1]) == 1) then
				local stock = tonumber(redis.call('get', KEYS[1]));
				local num = tonumber(ARGV[1]);
				if (stock == -1) then
					return -1;
				end;
				if (stock >= num) then
					return redis.call('incrby', KEYS[1], 0-num);
				end;
				return -2;
			end;
			return -3;`)

func (c *Client) IncrUnMinus(ctx context.Context, key string, num int64) (int64, error) {
	return IncrScript.Run(ctx, c, []string{"testKey"}, 1).Int64()
}
