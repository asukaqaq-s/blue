package redis_service

import (
	"context"
	"encoding/json"
	"fim_server/common/models/ctype"
	"fim_server/fim_user/user_rpc/types/user_rpc"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// GetUserBaseInfo 获取用户信息
func GetUserBaseInfo(client *redis.Client, userRpc user_rpc.UsersClient, userID uint) (userInfo ctype.UserInfo, err error) {
	key := fmt.Sprintf("fim_server_user_%d", userID)
	str, err := client.Get(key).Result()
	if err != nil {
		// 没找到
		userBaseResponse, err1 := userRpc.UserBaseInfo(context.Background(), &user_rpc.UserBaseInfoRequest{
			UserId: uint32(userID),
		})
		if err1 != nil {
			err = err1
			return
		}
		err = nil
		userInfo.ID = userID
		userInfo.Avatar = userBaseResponse.Avatar
		userInfo.NickName = userBaseResponse.NickName

		byteData, _ := json.Marshal(userInfo)

		// 设置进缓存
		client.Set(key, string(byteData), time.Hour) // 1个小时过期
		return
	}
	err = json.Unmarshal([]byte(str), &userInfo)
	if err != nil {
		return
	}
	return
}
