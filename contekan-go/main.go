package main

import (
	"errors"
	"fmt"
)

// Link mengelompokkan data; tag json dibaca oleh library encoding/json.
type Link struct {
	Code   string `json:"code"`
	URL    string `json:"url"`
	Clicks int    `json:"clicks"`
}

// Receiver pointer membuat perubahan berlaku pada struct asli.
func (l *Link) tambahKlik() {
	l.Clicks++
}

func bagi(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("tidak bisa dibagi nol")
	}
	return a / b, nil
}

func main() {
	fmt.Println("1. Variabel dan kondisi")
	var nama string = "Budi"
	umur := 20
	const batasDewasa = 17
	fmt.Printf("Nama: %s, umur: %d\n", nama, umur)
	if umur >= batasDewasa {
		fmt.Println("Dewasa")
	} else {
		fmt.Println("Belum dewasa")
	}

	fmt.Println("\n2. Function dan error")
	hasil, err := bagi(10, 2)
	if err != nil {
		fmt.Println("Gagal:", err)
		return
	}
	fmt.Println("10 / 2 =", hasil)
	if _, err := bagi(10, 0); err != nil {
		fmt.Println("Error yang ditangani:", err)
	}

	fmt.Println("\n3. Slice dan perulangan")
	peserta := []string{"Ani", "Budi"}
	peserta = append(peserta, "Citra")
	for i, nama := range peserta {
		fmt.Printf("%d. %s\n", i+1, nama)
	}
	fmt.Println("Jumlah peserta:", len(peserta))

	fmt.Println("\n4. Map dan pengecekan kunci")
	judul := map[string]string{"workshop": "Goravel Link"}
	if nilai, ada := judul["workshop"]; ada {
		fmt.Println("Judul:", nilai)
	}
	if _, ada := judul["lokasi"]; !ada {
		fmt.Println("Lokasi belum diisi")
	}

	fmt.Println("\n5. Struct, salinan, dan pointer")
	link := Link{Code: "abc123", URL: "https://example.com"}
	salinan := link
	salinan.tambahKlik()
	fmt.Printf("Ubah salinan: asli=%d, salinan=%d\n", link.Clicks, salinan.Clicks)
	alamat := &link
	alamat.tambahKlik()
	fmt.Printf("Ubah lewat pointer: asli=%d\n", link.Clicks)
	fmt.Printf("Link: /%s menuju %s\n", link.Code, link.URL)

	fmt.Println("\n6. Slice struct: ubah elemen asli")
	links := []Link{link, {Code: "xyz789", URL: "https://go.dev"}}
	for i := range links {
		links[i].tambahKlik()
	}
	for _, item := range links {
		fmt.Printf("/%s: %d klik\n", item.Code, item.Clicks)
	}
}
