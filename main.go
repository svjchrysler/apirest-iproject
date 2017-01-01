package main

import (
	"fmt"
	"log"

	"./packages"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kataras/iris"
)

const (
	//HOST database
	HOST = "localhost"
	//USER database
	USER = "postgres"
	//DB database
	DB = "ApiIProject"
	//SSL database
	SSL = "disable"
	//PASSWORD database
	PASSWORD = "k3yl0gg3r"
)

func conection() *gorm.DB {
	conexion := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s", HOST, USER, DB, SSL, PASSWORD)
	db, e := gorm.Open("postgres", conexion)
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
	return db
}

func migrations() {

	db := conection()

	if !db.HasTable(&packages.Country{}) {
		db.CreateTable(&packages.Country{})
	}

	if !db.HasTable(&packages.City{}) {
		db.CreateTable(&packages.City{}).
			AddForeignKey("country_id", "countries(id)", "CASCADE", "CASCADE")
	}

	if !db.HasTable(&packages.Product{}) {
		db.CreateTable(&packages.Product{})
	}

	defer db.Close()
}

//GetAPIContries get data contries
func GetAPIContries(ctx *iris.Context) {
	ctx.SetHeader("Access-Control-Allow-Origin", "*")
	var countries []packages.Country
	db := conection()
	db.Find(&countries)

	ctx.JSON(iris.StatusOK, countries)

	defer db.Close()

}

//GetAPICities get data cities
func GetAPICities(ctx *iris.Context) {
	ctx.SetHeader("Access-Control-Allow-Origin", "*")
	var cities []packages.City
	id := ctx.Param("id")
	db := conection()
	db.Where("country_id = ?", id).Find(&cities)
	ctx.JSON(iris.StatusOK, cities)

	defer db.Close()
}

//GetAPIProducts get data products
func GetAPIProducts(ctx *iris.Context) {
	ctx.SetHeader("Access-Control-Allow-Origin", "*")
	var products []packages.Product
	db := conection()
	db.Find(&products)
	ctx.JSON(iris.StatusOK, products)
	defer db.Close()
}

func main() {
	migrations()

	iris.Static("/images", "./public/images", 1)

	iris.Get("/api/countries", GetAPIContries)
	iris.Get("/api/cities/:id", GetAPICities)
	iris.Get("/api/products", GetAPIProducts)

	iris.Listen(":8080")

}
