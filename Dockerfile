# Menggunakan golang:1.22 sebagai base image
FROM golang:1.22 AS builder

# Mengatur working directory di dalam container
WORKDIR /app

# Menyalin file Go mod dan sum untuk mengambil dependencies
COPY go.mod .
COPY go.sum .

# Mengambil dependencies menggunakan go mod
RUN go mod download

# Menyalin seluruh file dari direktori proyek Anda ke dalam container
COPY . .

# Menjalankan aplikasi Golang tanpa membangun binary terlebih dahulu
CMD ["go", "run", "."]
