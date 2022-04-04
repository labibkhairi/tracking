package main

import (
	"fmt"
	"log"

	"github.com/cengsin/oracle"
	"gorm.io/gorm"
)

func main() {
	_, err := gorm.Open(oracle.Open("demo/demo@127.0.0.1:1521/XEPDB1"), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
		// panic error or log error info
	}

	fmt.Println("Koneksi ke db oracle berhasil")
}
