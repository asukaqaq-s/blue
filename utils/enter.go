package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"regexp"
	"strings"
)

func InList(list []string, key string) (ok bool) {
	for _, s := range list {
		if s == key {
			return true
		}
	}
	return false
}

func InListByRegex(list []string, key string) (ok bool) {
	for _, s := range list {
		regex, err := regexp.Compile(s)
		if err != nil {
			logx.Error(err)
			return
		}
		if regex.MatchString(key) {
			return true
		}
	}
	return false
}

func MD5(data []byte) string {
	h := md5.New()
	h.Write(data)
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func GetFilePrefix(fileName string) (prefix string) {
	nameList := strings.Split(fileName, ".")
	for i := 0; i < len(nameList)-1; i++ {
		if i == len(nameList)-2 {
			prefix += nameList[i]
			continue
		} else {
			prefix += nameList[i] + "."
		}
	}
	return
}
func InDir(dir []os.DirEntry, file string) bool {
	for _, entry := range dir {
		if entry.Name() == file {
			return true
		}
	}
	return false
}

// DeduplicationList 去重
func DeduplicationList[T string | int | uint | uint32](req []T) (response []T) {
	i32Map := make(map[T]bool)
	for _, i32 := range req {
		if !i32Map[i32] {
			i32Map[i32] = true
		}
	}
	for key, _ := range i32Map {
		response = append(response, key)
	}
	return
}
