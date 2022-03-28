package cipher

import "strings"

type shift struct {
	shiftKey int
}

type vigenere struct {
	Key string
}

func normalize(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + 'a' - 'A'
	}
	if r >= 'a' && r <= 'z' {
		return r
	}
	return -1
}

func encode(r rune, shiftKey int) rune {
	switch {
	case r >= 'a' && r <= 'z':
		return (r-'a'+rune(shiftKey))%26 + 'a'
	case r >= 'A' && r <= 'Z':
		return (r-'A'+rune(shiftKey))%26 + 'a'
	}
	return -1
}

func decode(r rune, shiftKey int) rune {
	switch {
	case r >= 'a' && r <= 'z':
		return (r-'a'+rune(26-shiftKey))%26 + 'a'
	case r >= 'A' && r <= 'Z':
		return (r-'A'+rune(26-shiftKey))%26 + 'a'
	}
	return -1
}

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance > 1 && distance < 26 {
		return shift{distance}
	} else if distance > -26 && distance < 0 {
		return shift{distance + 26}
	}
	return nil
}

func (c shift) Encode(input string) string {
	var output []rune
	plain := strings.Map(normalize, input)

	for _, r := range plain {
		base := 0

		switch {
		case r >= 'A' && r <= 'Z':
			base = int('A')
		case r >= 'a' && r <= 'z':
			base = int('a')
		}
		if base != 0 {
			r = rune(((int(r) - base + c.shiftKey) % 26) + base)
		}
		output = append(output, r)
	}
	return string(output)
}

func (c shift) Decode(input string) string {
	var output []rune

	for _, r := range input {
		if r >= 'a' && r <= 'z' {
			r = (r-'a'+rune(26-c.shiftKey))%26 + 'a'
		}
		output = append(output, r)
	}
	return string(output)
}

func NewVigenere(key string) Cipher {
	ok := false
	for _, r := range key {
		if r < 'a' || r > 'z' {
			return nil
		} else if r > 'a' {
			ok = true
		}
	}
	if ok {
		return vigenere{key}
	}
	return nil
}

func (v vigenere) Encode(input string) string {
	x := 0
	return strings.Map(func(r rune) rune {
		if r = encode(r, int(v.Key[x]-'a')); r >= 0 {
			x = (x + 1) % len(v.Key)
		}
		return r
	}, input)
}

func (v vigenere) Decode(input string) string {
	x := 0
	return strings.Map(func(r rune) rune {
		if r = decode(r, int(v.Key[x]-'a')); r >= 0 {
			x = (x + 1) % len(v.Key)
		}
		return r
	}, input)
}
