# Goravel Link

Pemendek link sederhana, dibuat dengan [Goravel](https://www.goravel.dev) (framework Go bergaya Laravel) dan SQLite. Project starter untuk workshop **"Dari Nol ke Link Shortener"**.

## Fitur
- Buat link pendek dari URL panjang (hanya `http://` dan `https://`)
- Redirect dari `/{kode}` ke URL asli
- Hitung jumlah klik
- Daftar 20 link terbaru

## Menjalankan

Butuh **Go 1.25+** dan Git.

```bash
cp .env.example .env              # Windows PowerShell: copy .env.example .env
go mod download
go run . artisan key:generate
go run . artisan migrate          # membuat tabel (mulai dari step-2)
go run .
```

Buka http://127.0.0.1:3000

## Branch per tahap

| Branch | Isi |
|---|---|
| `step-1` | Project jalan, halaman hello |
| `step-2` | + migration dan model `Link` |
| `step-3` | + form, validasi, simpan link |
| `step-4` | + redirect dan hitung klik |
| `step-5` / `main` | + daftar link dan tampilan |

Pindah tahap: `git checkout step-3`

Riwayat Git dibangun ulang dari kode hasil akhir dan modul workshop; commit lama tidak dipulihkan. Branch `main` menunjuk ke commit yang sama dengan `step-5`.

Folder `.agents/` dan `.github/`, konfigurasi `.env`, serta database lokal diabaikan oleh Git dan tidak disertakan dalam riwayat branch.

## Struktur penting

| Path | Fungsi |
|---|---|
| `routes/web.go` | Daftar URL |
| `app/http/controllers/link_controller.go` | Logika halaman |
| `app/http/requests/store_link_request.go` | Aturan validasi |
| `app/models/link.go` | Model tabel `links` |
| `database/migrations/` | Struktur tabel |
| `resources/views/home.tmpl` | Tampilan HTML |

## Pengujian

```bash
go test ./app/...                 # unit test pembuat kode, mulai step-3
sh scripts/smoke.sh               # uji endpoint lengkap, pada step-5 atau main
```

`scripts/smoke.sh` memakai database sementara di `/tmp`, jadi data lokalmu aman. Butuh `curl` dan `sqlite3`.

## Catatan keamanan dan batasan
- Kode link dibuat acak 6 karakter dengan `crypto/rand`. Tabrakan kode sangat jarang, tapi belum ada retry.
- Belum ada autentikasi dan rate limit, jadi **jangan dipasang di internet publik apa adanya**.
