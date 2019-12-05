package eaglesong

import (
	"bytes"
	"hash"
)

// Eaglesong func is just a shortcut for EaglesongSponge, it's silly, no time saving on reusing
func Eaglesong(input []byte) (output []byte) {
	output := make([]byte, 32)
	EaglesongSponge(output, 32, input, uint(len(input)))
	return
}

// types
const Size = 32

type digest struct {
	state  []uint32
	length uint
	msg    []byte
}

func (d *digest) Write(p []byte) (n int, err error) {
	if p != nil {
		d.length += uint(len(p))
		d.msg = bytes.Join([][]byte{d.msg, p}, nil)
		EaglesongUpdate(d.state, d.msg)
		remLen := d.length % (RATE / 8)
		remMsg := d.msg[d.length-remLen:]
		d.length = remLen
		d.msg = remMsg
		return len(p), nil
	} else {
		return 0, nil
	}
}

func (d *digest) Sum(b []byte) []byte {
	output := make([]byte, 32)
	EaglesongFinalize(d.state, d.msg, output, 32)
	return output
}

func (d *digest) Reset() {
	d.state = nil
	d.msg = nil
	d.length = 0
}

func (d *digest) Size() int {
	return 32
}

func (d *digest) BlockSize() int {
	return 4 * 16
}

// New returns a new hash.Hash computing the Eaglesong.
func New() hash.Hash {
	return &digest{
		state:  make([]uint32, 16),
		length: 0,
		msg:    nil,
	}
}
