Pahami use case di bawah untuk menentukan design dan architecture codes yang bisa antisipasi perubahan melalui abstraksi terutama dengan menerapkan basic [SOLID](https://en.wikipedia.org/wiki/SOLID) principle 

Stage 1:

1. Design sebuah pelabuhan ferry yang punya 2 jenis ferry:
   - ferry dengan kapasitas kecil untuk 8 kendaraan kecil (mobil bayar Rp 30.000,- / pickup bayar Rp 40.000,-)
   - ferry dengan kapasitas besar untuk 6 kendaraan besar (bis bayar Rp 50.000,- / truk bayar Rp 60.000,-)
   
   Program harus bisa menampilkan seluruh pendapatan dari tiket.


2. Kendaraan secara konstan datang ke pelabuhan secara acak.
   Seorang staff pelabuhan dibayar untuk memarkir kendaraan ke dalam ferry, dengan bayaran 10% dari tiket.
   Aplikasi harus bisa menampilkan pendapatan si staff pelabuhan.


3. Setiap kendaraan yang datang ke pelabuhan punya sisa bensin yang berbeda-beda secara acak.
   Jika kendaraan punya sisa bensin kurang dari 10%, staff pelabuhan harus mengisi bensin kendaraan tersebut ke pom bensin.

   Aplikasi harus bisa menampilkan jumlah sisa bensin untuk kendaraan yg sedang diperiksa.


4. Setiap pickup dan truck harus melalui pemeriksaan bea cukai. 
   Staff pelabuhan harus memeriksa muatan dengan membuka pintu kargo.
   Aplikasi harus bisa menampikan status dari kendaraan yang sedang diperiksa, pintu terbuka atau tertutup.


5. Lacak kendaraan mulai dari kedatangan sampai masuk ke dalam ferry.
   Aplikasi harus mampu menampilkan lokasi yg dikunjungi kendaraan, baik bea cukai atau pom bensin. (perempatan tidak perlu dihitung)

            A
            |
        G - 1 - S
        |   |
        C - 2 - L
    
        A - Gerbang Masuk
        G - Pom Bensin
        1 - perempatan 1
        S - Ferry kecil
        C - Bea Cukai
        2 - perempatan 2
        L - Ferry besar


Round 2:

6. Tambahkan staff pelabuhan dengan pekerjaan yang sama.
   Gaji yang diterima 11% dari pendapatan ticket.
   Aplikasi harus mampu menampilkan total pendapatannya


7. Tambahkan statiun pengisian baterei untuk mobil listrik.
   Tiap mobil listrik yang datang ke pelabukan dengan sisa daya 10%, staff terminal harus mengisi ulang daya mobil tersebut di statiun pengisian baterei.

   Aplikasi harus bisa menampikan sisa daya baterei mobil yang sedang diperiksa.


8. Tambahkan armada ferry terbaru khusus untuk mobil ramah lingkungan (mobil listrik bayar Rp 10.000,-)

            A
            |
        G - 1 - S
        |   |
        C - 2 - L
        |   |
        B - 3 - E
	
        A - Gerbang Masuk
        G - Pom Bensin
        1 - perempatan 1
        S - Ferry kecil
        C - Bea Cukai
        2 - perempatan 2
        L - Ferry besar
        B - Station pengisian baterei
        3 - perempatan 3
        E - Eco ferry

Jika di Stage pertama abstraksi cukup baik, maka di Stage kedua akan cukup mudah.