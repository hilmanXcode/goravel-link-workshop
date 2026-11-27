package main

import "testing"

func TestBagi(t *testing.T) {
	for _, kasus := range []struct {
		nama    string
		a, b    int
		ingin   int
		errorYa bool
	}{
		{nama: "habis dibagi", a: 10, b: 2, ingin: 5},
		{nama: "pecahan dibuang", a: 7, b: 2, ingin: 3},
		{nama: "pembilang nol", a: 0, b: 2, ingin: 0},
		{nama: "pembagi nol", a: 10, b: 0, errorYa: true},
	} {
		t.Run(kasus.nama, func(t *testing.T) {
			hasil, err := bagi(kasus.a, kasus.b)
			if (err != nil) != kasus.errorYa {
				t.Fatalf("bagi(%d, %d): error = %v", kasus.a, kasus.b, err)
			}
			if !kasus.errorYa && hasil != kasus.ingin {
				t.Fatalf("hasil = %d, ingin %d", hasil, kasus.ingin)
			}
		})
	}
}

func TestTambahKlik(t *testing.T) {
	link := Link{Code: "abc123", URL: "https://example.com", Clicks: 2}
	salinan := link
	salinan.tambahKlik()
	if link.Clicks != 2 || salinan.Clicks != 3 {
		t.Fatalf("salinan mengubah sumber: asli=%d, salinan=%d", link.Clicks, salinan.Clicks)
	}

	alamat := &link
	alamat.tambahKlik()
	if link.Clicks != 3 {
		t.Fatalf("pointer tidak mengubah sumber: klik=%d", link.Clicks)
	}
	if link.Code != "abc123" || link.URL != "https://example.com" {
		t.Fatalf("field selain Clicks berubah: %+v", link)
	}
}
