package config

import (
    "log"
    "os"
    "fmt"
    "io/ioutil"
    "path/filepath"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "github.com/joho/godotenv"
    "go-api/app/models"
)

var DB *gorm.DB

func InitDatabase() {
    if _, err := os.Stat(".env"); err == nil {
        err := godotenv.Load(".env")
        if err != nil {
            log.Fatal("Erro")
        }
    }

    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    if dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" || dbName == "" {
        log.Fatal("Erro ao definir credenciais")
    }

    dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
    var err error
    DB, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
    if err != nil {
        log.Fatalf("Erro ao abrir a conexão com o banco de dados: %v\n", err)
    }

    Migrate()

    fmt.Println("Conexão bem sucedida!")
}

func Migrate() {
    err := DB.AutoMigrate(&models.Client{}, &models.BenefitEntity{}, &models.Partner{}, &models.Product{}, &models.Billing{}, &models.ResourceUsage{}, &models.ServiceInfo{}, &models.Subscription{}, &models.Tag{})
    if err != nil {
        log.Fatalf("Erro ao executar migrations: %v\n", err)
    }

    var count int64
    DB.Table("clients").Count(&count)
    if count == 0 {

        insertionScripts := []string{
            "0010_insert_clients.sql",
            "0011_insert_billings.sql",
            "0012_insert_partners.up.sql",
            "0013_insert_products.up.sql",
            "0014_insert_subscriptions.up.sql",
            "0015_insert_benefit_entities.up.sql",
            "0016_insert_services.up.sql",
            "0017_insert_resources.up.sql",
            "0018_insert_tag.up.sql",
        }

        for _, script := range insertionScripts {
            filePath := filepath.Join("migration", script)
            sqlScript, err := ioutil.ReadFile(filePath)
            if err != nil {
                log.Fatalf("Erro ao ler arquivo de script SQL: %v\n", err)
            }

            if err := DB.Exec(string(sqlScript)).Error; err != nil {
                log.Fatalf("Erro ao executar script SQL: %v\n", err)
            }
        }
    } else {
        fmt.Println("Dados já inseridos, pulando scripts de inserção.")
    }
}