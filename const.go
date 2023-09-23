package main

import (
	"sync"
	"time"
)

type ContentType int32
const (
	MESSAGE ContentType = 0
	FILE    ContentType = 1
)

var MessageKey = "message"

var FileKey = "file"

// kv map to store message
var MessageKV = new(sync.Map)

// kv map to store file name
var FileKV = new(sync.Map)

// Default 10 minutes expiration time
var Expire = time.Minute * 30

// Max length of kv map
var MaxLength = 1000

var FilesDir = "./files"
