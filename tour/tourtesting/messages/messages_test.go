package messages

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"testing"
)

func TestGreeting(t *testing.T) {
	got := Greet("Lu")
	expect := "Hello, Lu!\n"

	if got != expect {
		t.Errorf("Expected %v, Got: %v", expect, got)
	}
}

func TestGreetTableDriven(t *testing.T) {
	scenarios := []struct {
		input  string
		expect string
	}{
		{input: "Gopher", expect: "Hello, Gopher!\n"},
		{input: "", expect: "Hello, !\n"},
	}

	for _, s := range scenarios {
		got := Greet(s.input)
		if got != s.expect {
			t.Errorf("Expected %v, Got: %v", s.expect, got)
		}
	}
}

func TestDepart(t *testing.T) {
	got := depart("Lu")
	expect := "Goodbye, Lu\n"

	if got != expect {
		t.Errorf("Expected %v, Got: %v", expect, got)
	}
}

func BenchmarkSHA1(b *testing.B) {
	data := []byte("Mayr had a little lamb")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sha1.Sum(data)
	}
}

func BenchmarkSHA256(b *testing.B) {
	data := []byte("Mayr had a little lamb")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sha256.Sum256(data)
	}
}

func BenchmarkSHA512(b *testing.B) {
	data := []byte("Mayr had a little lamb")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sha512.Sum512(data)
	}
}
