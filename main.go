package main

import (
	"fmt"
	"log"

	db "github.com/q1t/movielist/db"
	"github.com/spf13/viper"
)

func main() {
	// config defaults
	viper.SetDefault("SigningKey", "change me in config/app.toml")
	viper.SetDefault("Port", "8888")
	viper.SetDefault("Database", "postgresql")

	viper.SetConfigType("toml")
	viper.SetConfigName("app")
	viper.AddConfigPath("./config/")

	// TODO better error message
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Print(fmt.Errorf("Fatal error config file: %s \n", err))
		fmt.Println("Config path is ./config/, and name should be app.toml")
		fmt.Println("Using defaults")
	}

	// TODO: check toml for nested propeties
	// TODO: Handle config in a separate module? file?
	dbConfig := viper.GetStringMap("Database")
	dbName := dbConfig["Name"]
	dbUser := dbConfig["User"]
	// fmt.Sprintf("user=%s dbname=%s sslmode=disable", dbUser, dbName)
	dbOptions := db.Options{
		Driver: "postgres",
		Params: fmt.Sprintf("user=%s dbname=%s sslmode=disable", dbUser, dbName),
	}
	dbConnection := db.Connect(dbOptions)
	app := NewApp([]byte(viper.GetString("SigningKey")), dbConnection)

	port := viper.GetString("Port")
	log.Printf("Serving on port: %s\n", port)
	app.E.Run(":" + port)
}
