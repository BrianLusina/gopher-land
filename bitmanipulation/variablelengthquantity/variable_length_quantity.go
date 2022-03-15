package variablelengthquantity

import (
	"errors"
)

func encodeInt(x uint32) []byte {
	if x>>7 == 0 {
		return []byte{
			byte(x),
		}
	}

	if x>>14 == 0 {
		return []byte{
			byte(0x80 | x>>7),
			byte(127 & x),
		}
	}

	if x>>21 == 0 {
		return []byte{
			byte(0x80 | x>>14),
			byte(0x80 | x>>7),
			byte(127 & x),
		}
	}

	if x>>28 == 0 {
		return []byte{
			byte(0x80 | x>>21),
			byte(0x80 | x>>14),
			byte(0x80 | x>>6),
			byte(127 & x),
		}
	}

	return []byte{
		byte(0x80 | x>>28),
		byte(0x80 | x>>21),
		byte(0x80 | x>>14),
		byte(0x80 | x>>7),
		byte(127 & x),
	}
}

// decodeInt reads a varint-encoded integer from the slice.
// It returns the integer and the number of bytes consumed, or
// zero if there is not enough.
func decodeInt(buf []byte) (x uint32, n int, err error) {
	if len(buf) < 1 {
		return 0, 0, nil
	}

	if buf[0] < 0x80 {
		return uint32(buf[0]), 1, nil
	}

	var b byte
	for n, b = range buf {
		x <<= 7
		x |= uint32(b) & 0x7f
		if (b & 0x80) == 0 {
			return x, n + 1, nil
		}
	}

	return x, 0, errors.New("unterminated sequence")
}

func EncodeVarint(input []uint32) []byte {
	result := make([]byte, 0)
	for _, x := range input {
		result = append(result, encodeInt(x)...)
	}
	return result
}

func DecodeVarint(input []byte) ([]uint32, error) {
	if len(input) == 0 {
		return []uint32{0}, nil
	}
	usedBytes := 0
	result := make([]uint32, 0)
	for usedBytes < len(input) {
		r, nUsed, err := decodeInt(input[usedBytes:])
		if err != nil {
			return nil, err
		}
		result = append(result, r)
		usedBytes += nUsed
	}
	return result, nil
}
