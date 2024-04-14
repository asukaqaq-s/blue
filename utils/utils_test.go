package utils

import (
	"fmt"
	"testing"
)

func TestMD5(t *testing.T) {
	md := MD5([]byte("1234"))
	fmt.Println(md)
}

func TestGetFilePrefix(t *testing.T) {
	fmt.Println(GetFilePrefix("name.png"))
	fmt.Println(GetFilePrefix("1.name.png"))
	fmt.Println(GetFilePrefix("1.aaa.name.png"))
}

func TestDeduplicationList(t *testing.T) {
	fmt.Println(DeduplicationList([]uint{1, 1, 2, 6, 3, 4, 5, 2, 4, 5}))
}
