package rio

import "io"

type (
	PartReaderLen64 struct {
		Part io.Reader
		Size int64
	}
)

func (p2 *PartReaderLen64) Read(p []byte) (n int, err error) {
	return p2.Part.Read(p)
}

func (p2 *PartReaderLen64) Len() int64 {
	return p2.Size
}
