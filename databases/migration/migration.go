package main

import (
	"fmt"

	"github.com/TanaratSudjai/project-golang-api-shop/config"
	"github.com/TanaratSudjai/project-golang-api-shop/databases"
	"github.com/TanaratSudjai/project-golang-api-shop/entities"
)

func main() {
	
	conf := config.ConfigGetting() 
	db := databases.NewPostgresDatabase(conf.Database)


	fmt.Println(db.ConnectionGetting()) //connection error ! ..


	// playerMigration(db)
	// adminMigration(db)
	// itemMigration(db)
	// playerCoinMigration(db)
	// incentoryMigration(db)
	// purchaseHistoryMigration(db)
	
}

func playerMigration(db databases.Database){
	db.ConnectionGetting().Migrator().TableType(&entities.Player{})

}

func adminMigration(db databases.Database){
	db.ConnectionGetting().Migrator().TableType(&entities.Admin{})
	
}

func itemMigration(db databases.Database){
	db.ConnectionGetting().Migrator().TableType(&entities.Item{})
	
}

func playerCoinMigration(db databases.Database){
	db.ConnectionGetting().Migrator().TableType(&entities.PlayerCoin{})
	
}

func incentoryMigration(db databases.Database){
	db.ConnectionGetting().Migrator().TableType(&entities.Inventory{})
	
}

func purchaseHistoryMigration(db databases.Database){
	db.ConnectionGetting().Migrator().TableType(&entities.PurchaseHistory{})
	
}

