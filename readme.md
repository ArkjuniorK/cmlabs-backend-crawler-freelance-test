## Web Crawler With Go

### Pendahuluan
Aplikasi ini adalah program sederhana yang digunakan untuk melakukan crawling website menggunakan Go berbasis CLI. Program ini berjalan dengan beberapa langkah, yaitu sebagai berikut:
1. Menemukan semua url yang tersedia pada website dan mengunjunginya dengan batasan url harus sesuai dengan domain yang telah ditentukan 
2. Membuat direktori sesuai dengan `path` pada url dan menyimpan kontennya dalam bentuk `html`
3. Memuat seluruh konten pada website dalam satu direktori dengan nama yang sesuai dengan domain dari alamat website yang diberikan. Semua konten akan disimpan pada direktory `result`

### Cara Penggunaan
Untuk menjalankan program ini, pastikan Go telah terpasang pada sistem, jika sudah maka buka terminal pada `root` direktori dan jalankan program dengan menggunakan perintah `go run . --url https://example.com`, dimana flag `url` harus diisi dengan alamat dari website yang akan di crawling