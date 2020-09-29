package readerFile

import (
	"io"
	"io/ioutil"
)

func WriteUint16(w io.Writer, num uint16, buf []byte) (int, error) {
	if len(buf) == 0 {
		buf = make([]byte, 2)
	}
	buf[0] = byte(num)
	buf[1] = byte(num >> 8)
	return w.Write(buf)
}

func ReadUint16(r io.Reader, buf []byte) (res uint16, n int, err error) {
	if len(buf) == 0 {
		buf = make([]byte, 2)
	}
	n, err = r.Read(buf)
	if err != nil {
		return
	}
	res = uint16(buf[0]) + uint16(buf[1])<<8
	return
}

func ReadAllUint16(r io.Reader) (res []uint16, n int, err error) {
	bts, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	a := 0
	b := 1
	for {
		if b >= len(bts) {
			break
		}
		num := uint16(bts[a]) + uint16(bts[b])<<8
		res = append(res, num)
		a++
		b++
	}
	return
}
