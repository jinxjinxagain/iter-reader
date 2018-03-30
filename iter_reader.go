package ireader

import "errors"

// ReadOneFunc supply the buffer to reader
// For example, when read streaming data from db
// ReadOneFunc return one row's bytes from mysql or one ducument's bytes from mongodb
type ReadOneFunc func() ([]byte, error)

// Reader implements an io.Reader
type Reader struct {
	ReadOne func() ([]byte, error)
	Index   int
	Buf     []byte
	Err     error
}

// NewReader generates a reader with an iterator.
func NewReader(iter ReadOneFunc) *Reader {
	return &Reader{
		ReadOne: iter,
		Buf:     make([]byte, 0),
	}
}

// Read implements a basic read method
// refs: https://github.com/golang/go/blob/master/src/io/io.go#L77
func (r *Reader) Read(p []byte) (n int, err error) {
	if cap(p) == 0 {
		return 0, errors.New("empty buffer")
	}

	if len(r.Buf) == 0 {
		r.Buf, err = r.ReadOne()
		if err != nil {
			r.Err = err
		}
	}

	var cur = 0
	for {
		if cur+len(r.Buf)-r.Index < cap(p) {
			var remain = len(r.Buf) - r.Index
			copy(p[cur:cur+remain], r.Buf[r.Index:])
			if r.Err != nil {
				return cur + remain, r.Err
			}
			r.Buf, err = r.ReadOne()
			if err != nil {
				r.Err = err
			}
			r.Index = 0
			cur += remain
		} else {
			var remain = cap(p) - cur
			copy(p[cur:], r.Buf[r.Index:r.Index+remain])
			r.Index += remain
			cur = cap(p)
			break
		}
	}
	return cur, nil
}
