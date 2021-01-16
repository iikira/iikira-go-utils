package uploader

import (
	"fmt"
	rio2 "github.com/iikira/iikira-go-utils/requester/rio"
	"github.com/iikira/iikira-go-utils/utils/converter"
)

// DoUpload 执行上传
func DoUpload(uploadURL string, readerlen64 rio2.ReaderLen64, checkFunc CheckFunc) {
	u := NewUploader(uploadURL, readerlen64)
	u.SetCheckFunc(checkFunc)

	exitChan := make(chan struct{})

	u.OnExecute(func() {
		statusChan := u.GetStatusChan()
		for {
			select {
			case <-exitChan:
				return
			case v, ok := <-statusChan:
				if !ok {
					return
				}

				fmt.Printf("\r ↑ %s/%s %s/s in %s ............",
					converter.ConvertFileSize(v.Uploaded(), 2),
					converter.ConvertFileSize(v.TotalSize(), 2),
					converter.ConvertFileSize(v.SpeedsPerSecond(), 2),
					v.TimeElapsed(),
				)
			}
		}
	})

	u.Execute()
	close(exitChan)

	return
}
