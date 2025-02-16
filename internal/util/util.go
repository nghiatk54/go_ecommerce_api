package util

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

// get user key for redis
func GetUserKey(hashKey string) string {
	return fmt.Sprintf("u:%s:otp", hashKey)
}

// generate cli token uuid
func GenerateCliTokenUuid(userId int) string {
	newUuid := uuid.New()
	uuidString := strings.ReplaceAll(newUuid.String(), "-", "")
	return strconv.Itoa(userId) + "clitoken" + uuidString

}
