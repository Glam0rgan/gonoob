package tools

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"time"
)

const SessionPrefix = "sess_"

// Get id by snowflake.
func GetSnowflakeId() string {
	//default node id eq 1,this can modify to different serverId node
	node, _ := snowflake.NewNode(1)
	// Generate a snowflake ID.
	id := node.Generate().String()
	return id
}

func GetRandomToken(length int) string {
	r := make([]byte, length)
	io.ReadFull(rand.Reader, r)
	return base64.URLEncoding.EncodeToString(r)
}

func CreateSessionId(sessionId string) string {
	return SessionPrefix + sessionId
}


// ?????
func GetSessionIdByUserId(userId int) string {
	return fmt.Sprintf("sess_map_%d", userId)
}


func GetSessionName(sessionId string) string {
	return SessionPrefix + sessionId
}

// hash
func Sha1(s string) (str string) {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func GetNowDateTime() string {
	return time.Unix(time.Now().Unix(), 0).Format("2000-01-02 15:04:05")
}