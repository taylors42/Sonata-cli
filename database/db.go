package database

import (
	"fmt"
	"log"
	models "sonata-cli/Models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	DB, err := gorm.Open(sqlite.Open("sonata.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Falha ao conectar ao banco de dados:", err)
		return
	}

	err = DB.AutoMigrate(
		&models.Artist{},
		&models.Music{},
		&models.Playlist{},
	)
	if err != nil {
		log.Fatalf("Falha ao migrar o banco de dados: %v", err)
	}

	fmt.Println("Migração executada com sucesso!")

	// Verificar se a tabela Artist foi criada
	if DB.Migrator().HasTable(&models.Artist{}) {
		fmt.Println("A tabela 'Artist' foi criada com sucesso!")
	} else {
		fmt.Println("Falha ao criar a tabela 'Artist'.")
	}
}
