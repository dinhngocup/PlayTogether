package redis

import (
	"PlayTogether/model"
	_redisValueGenerator "PlayTogether/utils/redis"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis"
)

type UserRepositoryHandler struct {
	client *redis.Client
}

func NewRedisUserRepository(redisClient *redis.Client) model.UserRepository {
	return &UserRepositoryHandler{
		client: redisClient,
	}
}

func (u *UserRepositoryHandler) GetByID(id string) (model.User, error) {
	userKey := _redisValueGenerator.GenPrefixKey("user", id, "")
	mapUser, _ := u.client.HGetAll(userKey).Result()

	jsonStr, _ := json.Marshal(mapUser)
	userInfo := model.User{}
	json.Unmarshal(jsonStr, &userInfo)

	if userInfo.Id == "" {
		return model.User{}, errors.New("this user id not exists")
	}

	return userInfo, nil
}

func (u *UserRepositoryHandler) CreateUser(user model.User) error {
	_, err := u.GetByID(user.Id)

	if err == nil {
		return errors.New("this user exists")
	}
	userKey := _redisValueGenerator.GenPrefixKey("user", user.Id, "")

	// convert struct to map
	var userMap map[string]interface{}
	inrecUser, _ := json.Marshal(user)
	json.Unmarshal(inrecUser, &userMap)

	createUserResult := u.client.HMSet(userKey, userMap).Err()

	if createUserResult != nil {
		println("create user failed: " + createUserResult.Error())
		return errors.New("create user failed")
	}
	return createUserResult
}
