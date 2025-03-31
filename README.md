# Penugasan Netics 1 2025 (Implementasi CI/CD)

## Teknis Pengerjaan
Implementasikan modul CI/CD ini pada sebuah sistem server sederhana, dengan detail sebagai berikut
1. Buatlah API publik dengan endpoint /health yang menampilkan informasi sebagai berikut:
    CONTOH (value disesuaikan)
  {
    "nama": "Tunas Bimatara Chrisnanta Budiman",
    "nrp": "5025231999",
    "status": "UP",
    “timestamp”: time	// Current time
    "uptime": time		// Server uptime
  }
  Bahasa pemrograman dan teknologi yang digunakan dibebaskan kepada peserta.
3. Lakukan deployment API tersebut dalam bentuk container (Docker Multi-stage) pada VPS publik.
4. Lakukan proses CI/CD menggunakan GitHub Actions untuk melakukan otomasi proses deployment API. Terapkan juga best practices untuk menjaga kualitas environment CI/CD
5. Dokumentasikan pengerjaan di sebuah laporan berbentuk Markdown pada repositori peserta masing-masing.

## Laporan Pengerjaan
### 1. API Publik 
Pembuatan API Publik sederhana (backend) untuk menampilkan health server dengan bahasa Golang

Inisiasi package yang diperlukan untuk membuat API sederhana
  ```
  package main

  import (
  	"encoding/json"
  	"net/http"
  	"time"
  )
  ```

Membuat struct untuk menampilkan endpoint yang terdiri dari Nama, NRP, Status, Timestamp, dan Uptime
  ```
  type HealthResp struct {
  	Nama      string    `json:"nama"`
  	NRP       string    `json:"nrp"`
  	Status    string    `json:"status"`
  	Timestamp time.Time `json:"timestamp"`
  	Uptime    string    `json:"uptime"`
  }
  ```

Memulai menyatakan variabel startTime sebagai waktu
  ```
  var startTime time.Time
  ```

Membuat fungsi untuk mengubah variabel startTime ke waktu sekarang (waktu mulai server dijalankan)
  ```
  func init() {
  	startTime = time.Now()
  }
  ```

Membuat fungsi untuk menghandle request ke /health dengan mengisi uptime dengan menghitung waktu sejak startTime serta currtime adalah waktu pada saat user mengakses web server
  ```
  func handlerHealth(w http.ResponseWriter, r *http.Request) {
  	uptime := time.Since(startTime) 
  
  	currtime := time.Now()
  
  	response := HealthResp{
  		Nama:      "Arkananta Masarief",
  		NRP:       "5025231115",
  		Status:    "UP",
  		Timestamp: currtime,
  		Uptime:    uptime.String(),
  	}
  
  	w.Header().Set("Content-Type", "application/json")
  	json.NewEncoder(w).Encode(response)
  }
  ```

Fungsi main untuk menjalankan fungsi /health dan listen ke port 8080 (default)
  ```
  func main() {
  	http.HandleFunc("/health", handlerHealth)
  
  	http.ListenAndServe(":8080", nil)
  }
  ```

#### Tampilan endpoint /health pada localhost saat server dijalankan
<img src="/media/server-running.png">

### 2. Deployment Docker dan VPS
#### Docker

Membuat Dockerfile untuk build docker dengan mengambil dari image docker 1.24 sebagai builder
  ```
  FROM golang:1.24 AS builder
  ```

Menerapkan /app sebagai directory pekerjaan
  ```
  WORKDIR /app
  ```

Copy go.mod dari file github ke dalam docker container serta download go mod di docker
  ```
  COPY go.mod ./
  RUN go mod download
  ```

Copy semua file yang ada dari directory pekerjaan ke dalam docker container lalu compile golang 
  ```
  COPY . .
  RUN go build -o main .
  ```

Gunakan port 8080 dan jalankan file main untuk menjalankan server golang
  ```
  EXPOSE 8080

  CMD ["./main"]
  ```

#### Build Docker

Login ke akun docker lalu ke website untuk autentikasi akun docker
```
docker login
```
<img src="/media/login-docker.png">

Tag dengan nama repo docker sesuai nama yang kita inginkan serta push untuk menguploadnya ke docker hub
```
docker tag tugas-netics-1-2025 arkanantaaa/tugas-netics-1-2025
docker push arkanantaaa/tugas-netics-1-2025
```
<img src="/media/docker-push.png">

> [!NOTE] 
> Karena saya lupa screenshot saat push docker pertama kali, maka tulisan pada screenshotannya layer already exists

#### Publish dengan VPS Publik
Pada penugasan kali ini, saya menggunakan Railway sebagai VPS untuk mempublish API. 

Buka Railway lalu pilih add new service, lalu pilih Docker dan ketik repo docker yang ingin dibuild
<img src="/media/vps-docker.png">

Click deploy dan tunggu hingga berhasil
<img src="/media/vps-deploy.png">
<img src="/media/vps-deploying.png">
<img src="/media/vps-success.png">

Buka settings dan generate domain untuk membuatnya diakses publik
<img src="/media/public-vps.png">


