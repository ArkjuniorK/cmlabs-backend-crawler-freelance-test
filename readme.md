## Web Crawler With Go
Aplikasi ini adalah program sederhana yang digunakan untuk melakukan crawling website menggunakan Go berbasis CLI. Program ini berjalan dengan beberapa langkah, yaitu sebagai berikut:
1. Menemukan semua url yang tersedia pada website dan mengunjungi url-url tersebut dengan batasan harus sesuai dengan domain yang telah ditentukan 
2. Membuat direktori sesuai dengan `path` pada url dan mebyimpan kontennya dalam bentuk `html`
3. Memuat seluruh konten pada website dalam satu direktory dengan nama yang sesuai dengan domain dari alamat website yang diberikan. Semua konten akan disimpan pada direktory `result`