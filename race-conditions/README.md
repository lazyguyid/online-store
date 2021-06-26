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
    Karena ketika proses locking pada suatu baris dilakukan maka transaksi lain yang akan menggunakan baris yang sama harus menunggu sampai transaksi yang sedang berlangsung selesai, maka dengan begitu transaksi selanjutnya mendapatkan selalu mendapatkan data terupdate.
## **Kesimpulan**
    Solusi ini di pilih karena mampu memberikan data yang konsiten pada setiap transaksi yang akan di eksekusi selain itu minim perubahan di codebase yang sudah ada karena perubahan hanya dilakukan di bagian repositori saja.
