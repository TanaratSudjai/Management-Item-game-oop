package main

import (
	

	"github.com/TanaratSudjai/project-golang-api-shop/config"
	"github.com/TanaratSudjai/project-golang-api-shop/databases"
	"github.com/TanaratSudjai/project-golang-api-shop/entities"
)

func main() {
	
	conf := config.ConfigGetting() 
	db := databases.NewPostgresDatabase(conf.Database)

	// tx := db.ConnectionGetting().Begin()


	// fmt.Println(db.ConnectionGetting()) //connection error ! ..


	playerMigration(db) //tx paramitor
	adminMigration(db)
	itemMigration(db)
	playerCoinMigration(db)
	incentoryMigration(db)
	purchaseHistoryMigration(db)

}

// func playerMigration(tx *gorm.DB){
// 	tx.Migrator().CreateTable(&entities.Player{})

// }

func playerMigration(db databases.Database){
	db.ConnectionGetting().Migrator().CreateTable(&entities.Player{})

}

func adminMigration(db databases.Database){
	db.ConnectionGetting().Migrator().CreateTable(&entities.Admin{})
	
}

func itemMigration(db databases.Database){
	db.ConnectionGetting().Migrator().CreateTable(&entities.Item{})
	
}

func playerCoinMigration(db databases.Database){
	db.ConnectionGetting().Migrator().CreateTable(&entities.PlayerCoin{})
	
}

func incentoryMigration(db databases.Database){
	db.ConnectionGetting().Migrator().CreateTable(&entities.Inventory{})
	
}

func purchaseHistoryMigration(db databases.Database){
	db.ConnectionGetting().Migrator().CreateTable(&entities.PurchaseHistory{})
	
}

