CARA RUN

1.  Set environment variable pada terminal
    
    untuk windows :

    set MONGO_HOST=localhost

    set MONGO_PORT=27017

    set MONGO_DB=enigma

    set MONGO_USER=maulana

    set MONGO_PASSWORD=1212

    set API_PORT=8888

    set API_HOST=localhost

    sesuaikan dengan settingan laptop 

2.  go run main.go pada terminal

3.  Untuk mencoba REST API, gunakan file postman.json dan import ke postman. Bisa juga dengan link berikut:
    https://www.getpostman.com/collections/392a97a35ce231a79ee5


Keterangan REST API
1.  Insert Product -> POST

    localhost:8888/product

    body : json 

    contoh body :

    {
    "name":"Gula Aren",
    "price":10000,
    "category":"food"
    }

2.  Find All Product With Pagination -> GET

    localhost:8888/product/<page>/<total document>

    ganti <page> dengan halaman yg ingin ditampilkan dan <total document> dengan banyak document yang ditampilkan per halaman

3.  Update Product -> PATCH

    localhost:8888/product/<id> -> ganti <id> dengan id produk yg ingin di update

    body : json -> field yang mau di update

    contoh body :
    {
    "price":15000,
    }

4.  Delete Product -> DELETE

    localhost:8888/product/<id> -> ganti <id> dengan id produk yg ingin di hapus

5.  Get By Id -> GET

    localhost:8888/product/id/<id> -> ganti <id> dengan id produk yang ingin ditampilkan

6.  Get By Category -> GET

    localhost:8888/product/category/<category> -> ganti <category> dengan category produk yang ingin ditampilkan
