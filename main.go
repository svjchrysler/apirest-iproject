package main

import (
	"fmt"
	"log"

	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kataras/iris"
	"github.com/salguero/apirest-iproject/models"
)

const (
	//HOST database
	HOST = "ec2-54-221-217-158.compute-1.amazonaws.com"
	//USER database
	USER = "kaytqbatuadino"
	//DB database
	DB = "d1471skb605131"
	//SSL database
	SSL = "disable"
	//PASSWORD database
	PASSWORD = "2eb0d490c52f84554882caf3b8acd2dc4613cd76702410bf23e47a5a5c962ead"
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

	if !db.HasTable(&models.Country{}) {
		db.CreateTable(&models.Country{})
	}

	if !db.HasTable(&models.City{}) {
		db.CreateTable(&models.City{}).
			AddForeignKey("country_id", "countries(id)", "CASCADE", "CASCADE")
	}

	if !db.HasTable(&models.Product{}) {
		db.CreateTable(&models.Product{})
	}

	defer db.Close()
}

//GetAPIContries get data contries
func GetAPIContries(ctx *iris.Context) {
	ctx.SetHeader("Access-Control-Allow-Origin", "*")
	var countries []models.Country
	db := conection()
	db.Find(&countries)

	ctx.JSON(iris.StatusOK, countries)

	defer db.Close()

}

//GetAPICities get data cities
func GetAPICities(ctx *iris.Context) {
	ctx.SetHeader("Access-Control-Allow-Origin", "*")
	var cities []models.City
	id := ctx.Param("id")
	db := conection()
	db.Where("country_id = ?", id).Find(&cities)
	ctx.JSON(iris.StatusOK, cities)

	defer db.Close()
}

//GetAPIProducts get data products
func GetAPIProducts(ctx *iris.Context) {
	ctx.SetHeader("Access-Control-Allow-Origin", "*")
	var products []models.Product
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

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	iris.Listen(":" + port)

}
