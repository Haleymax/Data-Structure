package main

import (
	"fmt"
	"github.com/OneOfOne/xxhash"
)

// 将一个键进行Hash
func XXHash(key []byte) uint64 {
	h := xxhash.New64()
	h.Write(key)

	//返回哈希值，将哈希值与表长取余便可以得到位置
	return h.Sum64()
}

func xxHash() {
	keys := []string{"hi", "my", "friend", "I", "love", "you", "my", "apple"}
	for _, key := range keys {
		fmt.Printf("xxhash('%s')=%d\n", key, XXHash([]byte(key)))
	}
}
