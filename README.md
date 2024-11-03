# Evermos Technical Test

Pada kasus stok produk tidak sesuai ataupun stok produk yang negatif dikarenakan adanya lebih dari 1 proses atau request yang mengakses(manipulasi) data di saat yang bersamaan.
Sehingga membuat data akhir tidak sesuai dengan yang diinginkan. Kondisi ini disebut dengan _**Race Condiotion**_.

Pada kasus e-commerce hal ini sangat berbahaya, Sebagai contoh jika ada 5 customer yang melakukan pembalian(checkout) 1 product yang sama pada waktu yang bersamaan padahal stok produk tersebut hanya tersisa 4.
Jumlah stock yang tersedia tidak sesuai dengan jumlah pembalian yang masuk, tapi karena pembelian terjadi diwaktu yang bersamaan bisa saja kelima customer tersebut berhasil membuat transaksi pembelian.

Salah satu cara untuk menghindari race condition, memastikan jika ada proses pembelian yang sedang terjadi maka proses pembelian lain tidak bisa masuk sampai proses yang sebelumnya selesai.
Di GoLang bisa menggunakan ```sync.mutex``` untuk memastikan tidak ada proses masuk sebelum proses sebelumnya selesai.