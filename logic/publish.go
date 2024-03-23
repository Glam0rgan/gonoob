package logic

import (
	"bytes"
	"encoding/json"
)

var RedisClient *redis.Client
var RedisSessClient *redis.Client

func (logic *Logic) RedisPublishChannel(serverId string, toUserId int, msg []byte) (err error) {
	redisMsg := proto.RedisMsg{
		Op:       config.OpSingleSend,
		serverId: serverId,
		UserId:   toUserId,
		Msg:      msg,
	}

	redisMsgStr, err := json.Marshal(redisMsg)

	if err != nil {
		logrus.Errorf("logic,RedisPublishChannel Marshal err:%s", err.Error())
		return err
	}

	redisChannel := config.QueueName
	if err := RedisClient.LPush(redisChannel, redisMsgStr).Err(); err != nil {
		logrus.Errorf("logic,lpush err:%s", err.Error())
		return err
	}
	return
}

func (logic *Logic) RedisPublishRoomInfo(roomId int, count int, RoomUserInfo map[string]string, msg []byte) (err error) {
	var redisMsg = &proto.RedisMsg{
		Op:           config.OpRoomSend,
		RoomId:       roomId,
		Count:        count,
		Msg:          msg,
		RoomUserInfo: RoomUserInfo,
	}
	redisMsgByte, err := json.Marshal(redisMsg)
	if err != nil {
		logrus.Errorf("logic,RedisPublishRoomInfo redisMsg error : %s", err.Error())
		return
	}
	err = RedisClient.LPush(config.QueueName, redisMsgByte).Err()
	if err != nil {
		logrus.Errorf("logic,RedisPublishRoomInfo redisMsg error : %s", err.Error())
		return
	}
	return
}

func (logic *Logic) RedisPushRoomCount(roomId int, count int) (err error) {
	var redisMsg = &proto.RedisMsg{
		Op:     config.OpRoomCountSend,
		RoomId: roomId,
		Count:  count,
	}
	redisMsgByte, err := json.Marshal(redisMsg)
	if err != nil {
		logrus.Errorf("logic,RedisPushRoomCount redisMsg error : %s", err.Error())
		return
	}
	err = RedisClient.LPush(config.QueueName, redisMsgByte).Err()
	if err != nil {
		logrus.Errorf("logic,RedisPushRoomCount redisMsg error : %s", err.Error())
		return
	}
	return
}

func (logic *Logic) RedisPushRoomInfo(roomId int, count int, roomUserInfo map[string]string) (err error) {
	var redisMsg = &proto.RedisMsg{
		Op:           config.OpRoomInfoSend,
		RoomId:       roomId,
		Count:        count,
		RoomUserInfo: roomUserInfo,
	}
	redisMsgByte, err := json.Marshal(redisMsg)
	if err != nil {
		logrus.Errorf("logic,RedisPushRoomInfo redisMsg error : %s", err.Error())
		return
	}
	err = RedisClient.LPush(config.QueueName, redisMsgByte).Err()
	if err != nil {
		logrus.Errorf("logic,RedisPushRoomInfo redisMsg error : %s", err.Error())
		return
	}
	return
}

func (logic *Logic) getRoomUserKey(authKey string) string {
	var returnKey bytes.Buffer
	returnKey.WriteString(config.RedisRoomPrefix)
	returnKey.WriteString(authKey)
	return returnKey.String()
}

func (logic *Logic) getRoomOnlineCountKey(authKey string) string {
	var returnKey bytes.Buffer
	returnKey.WriteString(config.RedisRoomOnlinePrefix)
	returnKey.WriteString(authKey)
	return returnKey.String()
}

func (logic *Logic) getUserKey(authKey string) string {
	var returnKey bytes.Buffer
	returnKey.WriteString(config.RedisPrefix)
	returnKey.WriteString(authKey)
	return returnKey.String()
}