package ireader

import (
	"io"
	"io/ioutil"
	"strconv"
	"testing"
)

func TestIterReader(t *testing.T) {
	var reader = NewReader(readone)
	var result, _ = ioutil.ReadAll(reader)
	var sres = string(result)
	var exp = "12345"
	expected(t, sres == exp, exp, sres)
}

var cnt = 0

func readone() ([]byte, error) {
	cnt++
	var scnt = strconv.Itoa(cnt)
	var err error
	if cnt == 5 {
		err = io.EOF
	}
	return []byte(scnt), err
}

func expected(t testing.TB, cond bool, expect, got interface{}) {
	if cond {
		return
	}
	t.Fatalf("expected '%v', but got '%v'", expect, got)
}
