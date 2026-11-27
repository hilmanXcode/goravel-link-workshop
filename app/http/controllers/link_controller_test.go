package controllers

import "testing"

func TestNewCode(t *testing.T) {
	seen := map[string]bool{}
	for range 1000 {
		code := newCode()
		if len(code) != codeLength {
			t.Fatalf("panjang kode %d, harusnya %d", len(code), codeLength)
		}
		for _, c := range code {
			if !isBase62(c) {
				t.Fatalf("karakter tidak valid %q di kode %q", c, code)
			}
		}
		seen[code] = true
	}
	if len(seen) < 990 {
		t.Fatalf("terlalu banyak duplikat: %d unik dari 1000", len(seen))
	}
}

func isBase62(c rune) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9'
}
