# scheduler
[Tugas Progif]

Cara pengujian:
1. Buka browser
2. Masukkan "http://localhost:8080/index/" ke dalam url
3. Klik salah satu tombol:
	- "Lihat": lihat semua jadwal kegiatan
	- "Tambah": tambah 1 kegiatan baru
	- "Hapus": hapus 1 kegiatan yang sudah terjadwal

Cara pengujian "Lihat":
1. Klik tombol "Lihat Semua"
2. GET request dikirim ke web service
3. Web service mengirim data jadwal dalam json dari tabel "jadwal"
4. Klik tombol "Kembali" untuk kembali ke halaman index

Cara pengujian "Tambah":
1. Masukkan data:
	- Hari: format H
	- Bulan: format B
	- Tahun: format TT
	- Jam: format J
	- Tempat: teks string
	- Kegiatan: teks string
	- Keterangan: teks string
   ke dalam form yang tersedia
2. Klik tombol "Tambah"
3. POST request dikirim beserta data dalam json
4. Web service memasukkan data json itu ke tabel "jadwal"
5. Data di tabel "jadwal" ditampilkan
6. Klik tombol "Kembali" untuk kembali ke halaman index

Cara Pengujian "Hapus":
1. Masukkan ID jadwal yang ingin dihapus
2. Klik tombol "Hapus"
3. DELETE Request dikirimkan ke web service
4. Web service menghapus data terkait
5. Data di tabel "jadwal" akan ditampilkan
6. Klik tombol "Kembali" untuk kembali ke halaman index