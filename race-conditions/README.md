# **Mengatasi Masalah Jumlah Stok Item Negatif**
## **Status**
    Diusulkan
## **Isu**
    Pada saat ini terdapat isu jumlah stok negatif pada barang - barang yang banyak di beli saat acara flash sale 12.12.
## **Indikasi**
    Hal ini terjadi oleh beberapa kemungkinan:
        1. Tidak adanya validasi jumlah ketika jumlah itu mencapai angka 0.
        2. Terjadinya race condition yang menyebabkan pesanan tetap di proses.
## **Solusi**
### Melakukan Penguncian Database di Level Baris **(Row Level Database Locks)**
    Karena ketika proses locking pada suatu baris dilakukan maka transaksi lain yang akan menggunakan baris yang sama harus menunggu sampai transaksi yang sedang berlangsung selesai, maka dengan begitu transaksi selanjutnya selalu mendapatkan data terupdate.

## **Cara Kerja**
    Untuk cara kerja cukup sederhana, sebelum memproses update dengan melakukan query SELECT FOR UPDATE maka proses locking baris tertentu akan di eksekusi. sehingga ketika user lain yang akan menggunakan baris tersebut dipaksa menunggu hingga proses locking selesai. untuk lebih detail nya lihat di table di bawah ini.

<p align="center">
  <img src="https://raw.githubusercontent.com/lazyguyid/online-store-problems/main/race-conditions/docs/Row%20Lock%20Diagram.png">
</p>

    Pada table diatas, user 1 melakukan locking pada suatu baris tertentu dan user 2 akan menunggu user 1 hingga selesai terlebih dahulu dikarenakan user 2 akan menggunakan baris yang sama dengan user 1 sedangkan user 3 tidak akan menunggu user 1 karena tidak akan menggunakan data yang sama.

## **Kelebihan**
    - Mampu mengatasi race condition yang menyebabkan konsistensi data hancur.
    - User akan selalu mendapatkan data terbaru sebelum diolah.
## **Kekurangan**
    - Performa sedikit menurun dikarenakan proses menunggu menjadi lebih lama ketika user akan menggunakan data yang sama dalam waktu yang berdekatan.
## **Kesimpulan**
    Solusi ini di pilih karena mampu memberikan data yang konsiten pada setiap transaksi yang akan di eksekusi selain itu minim perubahan di codebase yang sudah ada karena perubahan hanya dilakukan di bagian repositori saja.
