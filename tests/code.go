package main

import (
	"crypto/md5"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	hmd5 := md5.Sum([]byte("string"))
	fmt.Println(hmd5)
	generated := uuid.New()
	fmt.Println(generated.String())
}
