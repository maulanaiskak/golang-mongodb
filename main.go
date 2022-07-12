package main

import "golang-mongodb/delivery"

func main() {
	delivery.NewServer().Run()
}

/*
set MONGO_HOST=localhost
set MONGO_PORT=27017
set MONGO_DB=enigma
set MONGO_USER=maulana
set MONGO_PASSWORD=1212
set API_PORT=8888
set API_HOST=localhost
go run .

*/
