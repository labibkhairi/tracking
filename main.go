package main

import (
	"Tracking/cnote"
	"log"
	"net/http"

	"github.com/cengsin/oracle"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	// db, err := gorm.Open(oracle.Open("system/SysPassword1@127.0.0.1:1521/XEPDB1"), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal(err.Error())
	// 	// panic error or log error info
	// }

	// fmt.Println("Koneksi ke db oracle berhasil")

	// var cnotes []cnote.Cnote
	// db.Find(&cnotes)

	// for _, cnote := range cnotes {
	// 	fmt.Println(cnote.CnoteNo)
	// 	fmt.Println(cnote.CnoteReceiverName)
	// 	fmt.Println("=====")
	// }

	router := gin.Default()
	router.GET("/handler", handler)
	router.Run()
}

func handler(c *gin.Context) {
	db, err := gorm.Open(oracle.Open("system/SysPassword1@127.0.0.1:1521/XEPDB1"), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
		// panic error or log error info
	}

	var cnotes []cnote.Cnote
	db.Find(&cnotes)

	c.JSON(http.StatusOK, cnotes)
}
