package tools

import (
	"time"
	"crypto/md5"
	"encoding/hex"
)

func PreventExit() {
	timer := time.NewTimer(time.Hour)
	for range timer.C {

	}
}

func MD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
	//hasher := md5.New()
	//hasher.Write([]byte(text))
	//return hex.EncodeToString(hasher.Sum(nil))
}
