package migrations

import (
    "address/models"

    "fmt"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

var AddressList []models.Address

func init() {
    address := models.Address{"81543", "Konradinstr", "Munich"}
    address2 := models.Address{"80636", "Konradinstr", "Munich"}
    AddressList = make([]models.Address, 2)
    AddressList[0] = address
    AddressList[1] = address2
}

func Migrate(connectionString string) {
    fmt.Println(connectionString)
    db, err := gorm.Open("postgres", connectionString)

    if err != nil {
        fmt.Println("We have an error")
        fmt.Println(err)
    }

    if db.HasTable(&models.Address{}) {
        db.DropTable(&models.Address{})
    }

    db.CreateTable(&models.Address{})

    for _, address := range AddressList {
        db.Create(&address)
    }
}
