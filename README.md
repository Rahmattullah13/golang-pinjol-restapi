# Golang Pinjaman Online(PINJOL)

## Deskripsi
Proyek ini adalah sebuah aplikasi yang dibangun dengan menggunakan framework Gin dan GORM dengan database PostgreSQL. Aplikasi ini menyediakan layanan pinjaman uang yang dapat dibayar per bulan.

## Teknologi yang Digunakan

- Gin: Framework web yang ringan dan cepat untuk Go.
- GORM: Library ORM (Object-Relational Mapping) untuk Go.
- PostgreSQL: Sistem manajemen basis data relasional yang digunakan sebagai database.

## Penggunaan

Berikut adalah langkah-langkah untuk menjalankan proyek ini:

1. Pastikan Go dan PostgreSQL telah terinstal di sistem Anda.
2. Unduh atau klon proyek ini ke dalam direktori lokal Anda.
3. Buat database baru di PostgreSQL yang akan digunakan untuk proyek ini.
4. Konfigurasi koneksi database di file config/config.go.
5. Sesuaikan konfigurasinya di file .env
6. Jalankan perintah berikut untuk menginstal dependensi proyek:

```bash
go mod tidy
```
7. Jalankan perintah berikut untuk menjalankan proyek:

```bash
go run main.go
```
8. Aplikasi akan berjalan di http://localhost:3000.

## Dokumentasi Swagger
Dokumentasi Swagger untuk proyek ini dapat ditemukan di http://localhost:3000/swagger/index.html. Gunakan dokumentasi Swagger untuk melihat detail dan penggunaan dari setiap endpoint API.