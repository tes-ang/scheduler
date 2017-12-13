# scheduler
[Tugas Progif]

Cara pengujian GET:
1. Buka browser
2. Masukkan "http://localhost:8080/jadwal/GET" ke dalam url
3. GET request dikirim ke web service
4. Web service mengirim data jadwal dalam json dari tabel "jadwal"

Cara pengujian POST:
1. Buka command prompt
2. Masukkan "curl -v -H "Content-Type: application/json" -X POST -d '{"Hari":[H], "Bulan":[B], "Tahun":[T]}, "Jam":[J], "Tempat":"[Tempat]", "Kegiatan":"[Kegiatan]", "Keterangan":"[Keterangan]"}' http://localhost:8080/jadwal/"
3. POST request dikirim beserta data dalam json
4. Web service memasukkan data json itu ke tabel "jadwal"