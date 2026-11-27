#!/bin/sh
# Smoke test: jalankan server, panggil endpoint utama, matikan server.
# Pakai database sementara supaya data lokal tidak tersentuh.
cd "$(dirname "$0")/.."
export DB_DATABASE=/tmp/goravel-link-smoke.db
rm -f "$DB_DATABASE"
go build -o /tmp/goravel-link-bin . || exit 1
/tmp/goravel-link-bin artisan migrate > /dev/null || exit 1
/tmp/goravel-link-bin > /tmp/goravel-link.log 2>&1 &
PID=$!
trap 'kill $PID 2>/dev/null' EXIT
sleep 3

B=http://127.0.0.1:3000
echo "--- GET /";                       curl -s -o /dev/null -w "%{http_code}\n" $B/
echo "--- POST url valid (harap 303)"
LOC=$(curl -s -o /dev/null -w "%{redirect_url}" -d "url=https://example.com/panjang" $B/links)
echo "$LOC"
CODE=${LOC##*created=}
echo "--- GET /$CODE (harap 302 ke example.com)"
curl -s -o /dev/null -w "%{http_code} -> %{redirect_url}\n" $B/$CODE
curl -s -o /dev/null $B/$CODE
echo "--- GET /tidakada (harap 404)";   curl -s -o /dev/null -w "%{http_code}\n" $B/tidakada
echo "--- POST kosong (harap 200 + pesan error)";     curl -s -w "\n%{http_code}\n" -d "url=" $B/links | grep -E "❌|^[0-9]{3}$"
echo "--- POST ftp (harap 200 + pesan error)";        curl -s -w "\n%{http_code}\n" -d "url=ftp://x.com" $B/links | grep -E "❌|^[0-9]{3}$"
echo "--- POST javascript: (harap 200 + pesan error)"; curl -s -o /dev/null -w "%{http_code}\n" -d "url=javascript:alert(1)" $B/links
echo "--- daftar di halaman utama (harap ada /$CODE dengan klik 2)"
curl -s $B/ | grep -A3 "href=\"/$CODE\"" | tr -s ' \n' ' '; echo
echo "--- isi tabel";                   sqlite3 "$DB_DATABASE" "select code,url,clicks from links"
