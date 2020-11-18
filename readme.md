# MSA Harbor - Rute Lalu Lintas pada Pelabuhan

## Use Case

Aplikasi ini berjalan sesuai dengan use-case pada file [use-case.md](use-case.md). berikut cuplikannya:

    Pahami use case di bawah untuk menentukan design dan architecture codes yang bisa antisipasi perubahan melalui abstraksi terutama dengan menerapkan basic [SOLID](https://en.wikipedia.org/wiki/SOLID) principle 

    Stage 1:

    1. Design sebuah pelabuhan ferry yang punya 2 jenis ferry:
    - ferry dengan kapasitas kecil untuk 8 kendaraan kecil (mobil bayar Rp 30.000,- / pickup bayar Rp 40.000,-)
    - ferry dengan kapasitas besar untuk 6 kendaraan besar (bis bayar Rp 50.000,- / truk bayar Rp 60.000,-)
   
    Program harus bisa menampilkan seluruh pendapatan dari tiket.


    2. Kendaraan secara konstan datang ke pelabuhan secara acak.
    Seorang staff pelabuhan dibayar untuk memarkir kendaraan ke dalam ferry, dengan bayaran 10% dari tiket.
    Aplikasi harus bisa menampilkan pendapatan si staff pelabuhan.

...

[Lihat Selengkapnya](use-case.md)

## Versi

Implementasi stage 1 tersedia pada branch v1. branch master saat ini adalah implementasi dari stage 2

## Instalasi

1. Clone repository ini
2. Jalankan `make build` untuk melakukan proses build

## Penggunaan

Setelah melakukan proses build, anda dapat melakukan eksekusi pada file binary `msa-harbor`. Aplikasi ini berbasis CLI dan hanya mempunyai satu flag yakni `-interval` yang berguna untuk menentukan kepadatan lalu lintas masuk ke dalam pelabuhan

Ketik dan jalankan perintah berikut:

```sh
./msa-harbor -interval=4
```

output untuk satu kendaraan :

```sh
tekan CTRL + C untuk menampilkan laporan
[G - Gerbang Masuk] mobil masuk: 
jenis=truck 
ukuran=kecil 
muatan=penumpang 
sisa_bahan_bakar=6.4%

[hariadi]: menerima jasa parkir 
tiket mobil diparkir=60000.00 
diterima=6000.00 (10.00%) 
total penghasilan=6000.00

[G - Pom Bensin]: bahan bakar saat ini: 7.0L dari 110.0L (6.4%)
[C - Bea Cukai]: pemeriksaan isi kargo: kargo terbuka

Rute yang telah dilewati:
|
A - Gerbang Masuk
|
G - Pom Bensin
|
C - Bea Cukai
|
L - Ferry Besar
--------------------------------------------------------
```

Aplikasi ini akan terus berjalan dan membuat tumpukan output seperti diatas, sesuai dengan jumlah kendaraan yang masuk setiap x detik.

untuk keluar dari aplikasi, silahkan tekan `CTRL + C`.

Aplikasi akan mencetak laporan pendapatan tiket setelah tombol `CTRL + C` ditekan.

output laporan:

```sh
^C
[Laporan Pendapatan Tiket]

ferry=[S - Ferry Kecil]

kind=car 
qty=0 
total_pendapatan=0.00

kind=pickup 
qty=0 
total_pendapatan=0.00

--------------------------------------------------------

ferry=[L - Ferry Besar]

kind=truck 
qty=1 
total_pendapatan=60000.00

kind=bus 
qty=0 
total_pendapatan=0.00

--------------------------------------------------------

TOTAL SELURUH PENDAPATAN KAPAL=Rp. 60000.00
```