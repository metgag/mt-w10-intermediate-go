# 📦 Backend

**Backend** adalah bagian dari aplikasi yang bekerja di balik layar (server). Bagian ini tidak terlihat langsung oleh pengguna, tapi sangat penting karena:

- Menangani permintaan dari tampilan depan (frontend)
- Mengolah data sesuai kebutuhan aplikasi
- Menyimpan atau mengambil data dari database
- Mengirimkan hasilnya kembali ke frontend

---

## 🌐 Apa Itu API?

**API** (Application Programming Interface) adalah jembatan penghubung antar aplikasi, yang memungkinkan frontend dan backend untuk saling berkomunikasi.

---

## 🔁 REST / RESTful API

**REST** (Representational State Transfer) adalah aturan atau standar yang biasa digunakan untuk membuat API, terutama di aplikasi web.

REST API menggunakan protokol HTTP (seperti GET dan POST) untuk mengirim dan menerima data antara client (frontend) dan server (backend).

---

### ✨ Ciri-ciri REST API

| Ciri-ciri         | Penjelasan singkat |
|-------------------|--------------------|
| **Stateless**     | Setiap permintaan harus lengkap, server tidak menyimpan informasi dari permintaan sebelumnya. |
| **Client-Server** | Pemisahan antara frontend dan backend. |
| **Cacheable**     | Hasil dari permintaan bisa disimpan untuk mempercepat akses. |
| **Layered System**| Sistem bisa terdiri dari beberapa lapisan, seperti proxy, firewall, atau load balancer. |
| **Uniform Interface** | Semua permintaan dan respons mengikuti format yang konsisten. |

---

## 📬 Cara Kerja Request dan Response

### 📤 Permintaan (Request) dari Frontend

Saat pengguna berinteraksi dengan aplikasi (misalnya klik tombol), frontend akan mengirim request ke backend.

Struktur request biasanya terdiri dari:

- **Method**: Jenis aksi (GET, POST, PUT, DELETE, dll)
- **Endpoint**: Alamat URL tujuan di backend (misalnya `/users`)
- **Header**: Informasi tambahan seperti token login
- **Body (opsional)**: Data yang dikirim, seperti isi form

### 📥 Jawaban (Response) dari Backend

Backend akan merespons permintaan tersebut dengan:

- **Status Code**: Kode angka yang menunjukkan hasil (contoh: `200` berarti berhasil)
- **Header**: Info tambahan, seperti format data (JSON)
- **Body**: Data yang dikirim balik ke frontend (biasanya dalam format JSON)

---

## 🛠️ HTTP Method (Jenis Aksi)

| Method   | Fungsi                                                                 |
|----------|------------------------------------------------------------------------|
| **GET**     | Mengambil atau membaca data                                           |
| **POST**    | Mengirim atau menambahkan data baru ke server                        |
| **PUT**     | Mengganti seluruh data yang ada                                      |
| **PATCH**   | Mengubah sebagian dari data yang ada                                 |
| **DELETE**  | Menghapus data                                                       |

> 💡 Analogi:
> - PUT = ganti seluruh rumah  
> - PATCH = ganti pintunya saja  

---

## 📍 Apa Itu Endpoint?

**Endpoint** adalah alamat URL yang digunakan untuk mengakses fitur tertentu di backend. Setiap endpoint biasanya mewakili suatu data atau fungsi, seperti daftar buku, detail pengguna, dll.

Contoh endpoint untuk fitur manajemen buku:

| Endpoint         | Method | Fungsi                                           |
|------------------|--------|--------------------------------------------------|
| `/books`         | GET    | Mengambil semua data buku                        |
| `/books/:id`     | GET    | Mengambil satu buku berdasarkan ID               |
| `/books`         | POST   | Menambahkan buku baru                            |
| `/books/:id`     | PUT    | Mengganti seluruh data buku berdasarkan ID       |
| `/books/:id`     | DELETE | Menghapus buku tertentu berdasarkan ID           |

> 📌 Catatan: Endpoint boleh sama, tapi hasilnya bisa berbeda tergantung **HTTP Method** yang digunakan.

---

## 🚦 HTTP Status Code

**Status code** adalah kode angka dari server yang menunjukkan hasil dari permintaan client. Kode ini penting untuk mengetahui apakah permintaan berhasil atau ada error.

| Kategori | Kode Umum                   | Makna                                                    |
|----------|-----------------------------|-----------------------------------------------------------|
| 1xx      | `100 Continue`              | Informasi, permintaan sedang diproses                    |
| 2xx      | `200 OK`, `201 Created`, `204 No Content` | Permintaan berhasil diproses                             |
| 3xx      | `301 Moved Permanently`, `304 Not Modified` | Pengalihan (redirect), halaman sudah pindah              |
| 4xx      | `400 Bad Request`, `401 Unauthorized`, `404 Not Found` | Kesalahan dari sisi client, seperti salah input atau belum login |
| 5xx      | `500 Internal Server Error`, `503 Service Unavailable` | Kesalahan dari sisi server                               |

---

## 🔄 CRUD (Create, Read, Update, Delete)

**CRUD** adalah empat operasi dasar yang sering digunakan saat bekerja dengan data di aplikasi.

| Fungsi  | HTTP Method   | SQL (Database) Setara     |
|---------|---------------|---------------------------|
| Create  | POST          | `INSERT INTO`             |
| Read    | GET           | `SELECT`                  |
| Update  | PUT / PATCH   | `UPDATE`                  |
| Delete  | DELETE        | `DELETE FROM`             |

---

## 📚 Contoh Alur Kerja REST API

Misalnya pengguna ingin menambahkan buku baru ke aplikasi:

1. Pengguna isi form "Tambah Buku" di frontend
2. Frontend kirim permintaan `POST /books` ke backend
3. Backend memproses data dan menyimpannya ke database
4. Backend kirim balasan dengan status `201 Created`
5. Frontend menampilkan pesan: "Buku berhasil ditambahkan"

---