# Evermos Technical Test

## TASK 1
Pada kasus stok produk tidak sesuai ataupun stok produk yang negatif dikarenakan adanya lebih dari 1 proses atau request yang mengakses(manipulasi) data di saat yang bersamaan.
Sehingga membuat data akhir tidak sesuai dengan yang diinginkan. Kondisi ini disebut dengan _**Race Condiotion**_.

Pada kasus e-commerce hal ini sangat berbahaya, Sebagai contoh jika ada 5 customer yang melakukan pembalian(checkout) 1 product yang sama pada waktu yang bersamaan padahal stok produk tersebut hanya tersisa 4.
Jumlah stock yang tersedia tidak sesuai dengan jumlah pembalian yang masuk, tapi karena pembelian terjadi diwaktu yang bersamaan bisa saja kelima customer tersebut berhasil membuat transaksi pembelian.

Salah satu cara untuk menghindari race condition, memastikan jika ada proses pembelian yang sedang terjadi maka proses pembelian lain tidak bisa masuk sampai proses yang sebelumnya selesai.
Di GoLang bisa menggunakan ```sync.mutex``` untuk memastikan tidak ada proses masuk sebelum proses sebelumnya selesai.

Pada projek ini ```sync.mutex``` di tambahkan pada service product
```go
type productServiceImpl struct {
	product product.Repository
	order   order.Repository
	mu      sync.Mutex  // race condition handling
}

// Checkout for deduct product stock and create order
func (p *productServiceImpl) Checkout(ctx context.Context, checkoutRequest *request.Checkout) (*response.Order, error) {
	p.mu.Lock()         // lock transaction
	defer p.mu.Unlock() // unlock transaction

	// find the product details
	existingProduct, err := p.product.FindById(ctx, checkoutRequest.ProductId)
	if err != nil {
		return nil, fmt.Errorf("failed to find product id %d. %w", checkoutRequest.ProductId, err)
	}

	// Check available stock and return an error if available stock is less than the requested quantity
	if existingProduct.Stock < checkoutRequest.Quantity {
		return nil, fmt.Errorf("product stock is lower than quantity")
	}

	// Simulating processing time
	time.Sleep(200 * time.Millisecond)

	// deduct stock
	existingProduct.Stock -= checkoutRequest.Quantity

	// store new available stock in the database after deduction
	_, err = p.product.Update(ctx, existingProduct)
	if err != nil {
		return nil, fmt.Errorf("failed to Deduct product %d. %w", checkoutRequest.ProductId, err)
	}

	// create new order
	// count total price
	totalPrice := existingProduct.Price * checkoutRequest.Quantity
	// store new order in the database
	orderSaved, err := p.order.Save(ctx, &entity.Order{ProductId: checkoutRequest.ProductId, Quantity: checkoutRequest.Quantity, TotalPrice: totalPrice})
	if err != nil {
		return nil, fmt.Errorf("failed to save order. %w", err)
	}

	// return response
	respMessage := orderSaved.ToResponse()
	respMessage.ProductName = existingProduct.Name
	return respMessage, nil
}

```

## Teknologi yang digunakan pada Proyek ini adalah :
- MySQL : SQL Database Server
- GoLang : Bahasa Pemrograman Backend (IDE) versi 1.21.3
- Docker : Container
- Git : Sistem Pengontrol Versi (Kode)
- Postman : Dokumentasi API

## Menjalankan Proyek
### 1. Clone proyek dari Github
```
$ mkdir go-project
$ cd go-project
$ git clone https://github.com/MCPutro/evermosTest.git
```

tunggu hingga porses selesai, dan akan muncul seperti berikut :
```
Cloning into 'evermosTest'...
remote: Enumerating objects: 261, done.
remote: Counting objects: 100% (261/261), done.
Receiving objects: 100% (261/261), 53.14 KiB | 1.18 MiB/s, done.
Resolving deltas:  22% (27/122)0% (172/172), done.
Resolving deltas:  27% (33/122)reused 208 (delta 78), pack-reused 0
Resolving deltas: 100% (122/122), done.
```
lalu masuk kedalam folder evermosTest:
```shell
cd evermosTest
```
### 2. Jalankan proyek ini dengan docker compose
pastikan docker sudah berjalan, di komputer atau laptop yang digunakan. Jika sudah, masukan perintah berikut:
```shell
docker compose up -d
```

tunggu hingga project selesai di download dan akan muncul tampilan seperi berikut, coba tunggu beberapa menit :
```
....
[+] Running 3/3
 ✔ Network evermostest_default      Created                  0.9s 
 ✔ Container mysql-local-docker     Started                  2.0s 
 ✔ Container backend                Started                  3.4s 
```

### 3. Akses proyek
untuk melakukan testing terhadap Rest API yang tersedia bisa menggunakan Postman dan untuk melihat/membuka database bisa menggunakan DBeaver.
- ### DBeaver
>1. Buka `DBeaver`.
>2. Klik ikon `Connect to a Database` di Pojok-Kiri-Atas.
>3. Pada Kategori `Popular`, pilih `MySQL` lalu klik Next.
>4. Isikan Konfigurasi sesuai dengan File `.env` pada Proyek ini.
>5. Jika berhasil, akan ada Ikon centang hijau pada daftar Database di sebelah kiri.

- ### Postman
pada proyek ini juga disematkan [Collection Postman (Evermos.postman_collection.json)](https://github.com/MCPutro/evermosTest/blob/master/Evermos.postman_collection.json) yang dapat anda import kepostman untuk mencoba API.
#### Test your API using the Collection Runner
- Setelah import, pilih Collections yang ada di sidebar dan pilih/klik collection atau folder yang akan di test.
- Klik tombol Run, ada di pojok kanan atas
- Beri cek list hanya pada Checkout
- Pada Run configuration, field Iteration isi semisal 100
- Klik Run Evermos


### 4. Mematikan Proyek
```shell
$ docker compose down -v 
```

## TASK 2
untuk task 2 di ada di folder Task2. Untuk menjalankan function test bisa menggunakan script berikut
```shell
go test -v -run TestTask2$ evermosTest/Task2
```
