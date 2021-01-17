package cachepool

import (
	"github.com/iikira/iikira-go-utils/utils/converter"
	"runtime"
	"sync"
)

var (
	syncPoolSize     = int(64 * converter.KB)
	syncPoolFirstNew = false
	SyncPool         = sync.Pool{
		New: func() interface{} {
			syncPoolFirstNew = true
			return RawMallocByteSlice(syncPoolSize)
		},
	}
)

func SetSyncPoolSize(size int) {
	if syncPoolFirstNew && size != syncPoolSize {
		runtime.GC()
	}
	syncPoolSize = size
}
