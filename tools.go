package tools

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

func PreventExit() {
	timer := time.NewTimer(time.Hour)
	for range timer.C {

	}
}

func MD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
