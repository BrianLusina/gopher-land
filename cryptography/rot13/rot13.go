package rot13

import "bytes"

func rotateCharacter(b byte) byte {
	if 'A' <= b && b <= 'Z' {
		return ((b-'A')+13)%26 + 'A'
	}
	if 'a' <= b && b <= 'z' {
		return ((b-'a')+13)%26 + 'a'
	}
	return b
}

func Rot13(message string) string {
	var b bytes.Buffer
	for _, character := range message {
		b.WriteByte(rotateCharacter(byte(character)))
	}
	return b.String()
}
