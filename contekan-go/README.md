# Contekan Dasar Go untuk Pemateri

Pendamping sesi dasar Go sebelum masuk ke Goravel. Semua contoh lengkap ada di [`main.go`](main.go), menggunakan standard library tanpa server atau database.

## Jalankan

Dari root project:

```bash
go run ./contekan-go
go test ./contekan-go
```

Atau dari folder ini: `go run .` dan `go test .`. Perintah ini tidak memerlukan `.env`.

## Urutan demo singkat

1. Jalankan contoh sekali, lalu buka `main.go`.
2. Jelaskan variabel, tipe, dan kondisi. Ubah `umur` menjadi `16`, lalu jalankan lagi.
3. Jelaskan function, dua nilai kembalian, dan pengecekan `error`.
4. Jelaskan slice, `range`, serta map.
5. Bandingkan perubahan pada salinan struct dengan perubahan melalui pointer.
6. Hubungkan konsep tersebut dengan controller, request, dan model di Goravel.

## 1. Package, import, dan main

- `package main`: package untuk program yang dapat dijalankan.
- `func main()`: titik masuk program.
- `import "fmt"`: memakai standard library untuk format dan cetak teks.
- File dalam satu folder biasanya menggunakan package yang sama; file test eksternal dapat menggunakan akhiran `_test`.
- Nama yang diawali huruf besar diekspor, sehingga dapat dipakai dari package lain.
- Import yang tidak dipakai dan variabel lokal yang tidak dipakai menyebabkan compile error.

## 2. Variabel, konstanta, dan tipe

```go
var nama string = "Budi"
umur := 20
const batasDewasa = 17
```

- `var`: deklarasi variabel; tipe dapat ditulis atau disimpulkan.
- `:=`: deklarasi singkat di dalam function; harus ada minimal satu variabel baru pada scope yang sama.
- `=`: mengisi ulang variabel yang sudah ada.
- `const`: nilai konstan, tidak dapat diubah.
- Tipe umum: `string`, `int`, `float64`, `bool`.
- `any`: alias `interface{}`, dapat menyimpan nilai bertipe apa pun; bukan berarti pemeriksaan tipe hilang.
- Zero value: angka `0`, string `""`, boolean `false`, pointer/slice/map `nil`.
- Konversi tipe harus eksplisit, misalnya `float64(umur)`.

## 3. Kondisi dan perulangan

```go
if umur >= batasDewasa {
    fmt.Println("Dewasa")
} else {
    fmt.Println("Belum dewasa")
}

for i := 0; i < 3; i++ {
    fmt.Println(i)
}
```

Go menggunakan `for` untuk perulangan, termasuk bentuk seperti `while`: `for kondisi { ... }`.
`break` keluar dari loop; `continue` melanjutkan ke iterasi berikutnya. `switch` memilih kasus tanpa memerlukan `break` pada setiap kasus.

## 4. Function dan error

```go
func bagi(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("tidak bisa dibagi nol")
    }
    return a / b, nil
}
```

`a, b int` berarti kedua parameter bertipe `int`. `(int, error)` berarti dua nilai kembalian.
Pembagian bilangan bulat membuang bagian pecahan menuju nol: `7 / 2` menghasilkan `3`.

```go
hasil, err := bagi(10, 2)
if err != nil {
    fmt.Println("Gagal:", err)
    return
}
fmt.Println(hasil)
```

`nil` pada `error` berarti tidak ada error. Jangan gunakan `hasil` sebelum menangani error.
Go tidak memiliki `try/catch`; error biasa dikembalikan sebagai nilai. `panic` bukan pengganti validasi input.
`defer` menunda pemanggilan function sampai function pembungkus selesai, misalnya `defer file.Close()` setelah file berhasil dibuka.

## 5. Array, slice, dan map

- `[3]string`: array dengan panjang tetap; panjang bagian dari tipe.
- `[]string`: slice, tampilan atas underlying array; panjangnya dapat berubah dengan `append`.
- Simpan hasil `append`: `nama = append(nama, "Citra")`.
- `len(nama)` memberi jumlah elemen. Index mulai dari `0`; index di luar batas menyebabkan panic.
- `for i, nilai := range nama`: membaca index dan nilai.
- `_` mengabaikan nilai, misalnya `for _, nilai := range nama`.

```go
peserta := []string{"Ani", "Budi"}
judul := map[string]string{"workshop": "Goravel Link"}
nilai, ada := judul["workshop"]
```

Map menyimpan pasangan kunci dan nilai. Gunakan bentuk `nilai, ada` untuk membedakan kunci tidak ada dari zero value.
Urutan iterasi map tidak dijamin. Membaca nil map aman, tetapi menulisnya panic; inisialisasi dengan literal atau `make`.

## 6. Struct, method, dan pointer

```go
type Link struct {
    Code   string `json:"code"`
    URL    string `json:"url"`
    Clicks int    `json:"clicks"`
}

func (l *Link) tambahKlik() {
    l.Clicks++
}
```

- Struct mengelompokkan data; Go tidak memakai class inheritance.
- Method adalah function dengan receiver, di sini `(l *Link)`.
- `*Link` adalah tipe pointer ke `Link`; `&link` mengambil alamat variabel.
- `*pointer` mengakses nilai yang ditunjuk; akses field dapat langsung memakai `pointer.Clicks`.
- Menyalin struct menyalin nilai field. Field berupa slice, map, atau pointer masih dapat berbagi data yang dirujuk.
- Receiver pointer dapat mengubah struct asli. Jangan dereference pointer `nil`.
- `tambahKlik` tidak diekspor karena diawali huruf kecil.
- Tag seperti `json:"code"` adalah metadata untuk library; tag sendiri tidak menjalankan kode.

## 7. Hubungkan ke Goravel

| Kode workshop | Konsep Go |
|---|---|
| `func (r *LinkController) Store(...)` | Method dengan receiver pointer |
| `var link models.Link` | Variabel struct dengan zero value |
| `First(&link)` | Mengirim alamat agar ORM dapat mengisi struct |
| `[]models.Link` | Slice berisi model |
| `map[string]any{"title": "Goravel Link"}` | Data template dengan tipe nilai beragam |
| `if err := query.Create(&link); err != nil` | Deklarasi pendek di scope kondisi dan penanganan error |
| `form:"url"` | Metadata yang dibaca framework untuk input form |

## Latihan 10 menit

1. Tambahkan field `Aktif bool` ke `Link` dan cetak nilainya.
2. Buat slice berisi tiga link, lalu cetak dengan `range`.
3. Tambahkan klik ke elemen pertama melalui `links[0].tambahKlik()`.
4. Bandingkan dengan `for _, link := range links { link.tambahKlik() }`: variabel `link` adalah salinan elemen struct.
5. Agar loop mengubah elemen asli, gunakan `for i := range links { links[i].tambahKlik() }`.
6. Tambahkan kasus pembagian negatif ke test `TestBagi`.

## Jebakan saat live coding

- Simpan file sebelum `go run`; jalankan ulang setelah perubahan.
- Gunakan tanda kutip ganda untuk string, bukan tanda kutip tunggal; tanda kutip tunggal untuk rune seperti `'A'`.
- `nil` tidak berlaku untuk semua tipe: `string` kosong adalah `""`, bukan `nil`.
- Nil slice dapat di-`append`; nil map harus dibuat sebelum ditulisi.
- `for _, link := range links` menyalin elemen struct; perubahan field salinan tidak mengubah elemen asli.
- Utamakan membaca pesan compiler dari error pertama.

Materi ini sengaja tidak membahas goroutine, channel, generics, atau detail interface; tidak diperlukan untuk alur workshop ini.
