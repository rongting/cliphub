package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func Set(key string, value string, t ContentType) (bool, string) {
	kv := getKV(t)
	if getSize(kv) >= MaxLength {
		return false, "Error: Reach max length"
	}
	kv.Store(key, value)
	time.AfterFunc(Expire, func() {
		kv.Delete(key)
		if t == FILE {
			err := os.RemoveAll(FilesDir + key)
			if err != nil {
				fmt.Println("Delete expired file failed", err)
			}
		}
	})
	return true, ""
}

func Get(key string, t ContentType) string {
	kv := getKV(t)
	val, ok := kv.Load(key)
	if ok {
		return val.(string)
	} else {
		return ""
	}
}

func Contains(key string, t ContentType) bool {
	kv := getKV(t)
	_, ok := kv.Load(key)
	return ok
}

func getSize(kv *sync.Map) int {
	len := 0
	kv.Range(func(k, v interface{}) bool {
		len++
		return true
	})
	return len
}

func getKV(t ContentType) *sync.Map {
	if t == MESSAGE {
		return MessageKV
	} else {
		return FileKV
	}
}


