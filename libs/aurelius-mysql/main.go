package main

import (
	"aurelius/libs/aurelius-mysql/entity"
	"fmt"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "aurelius-mysql",
	Short: "CLI application to recreate the database",
	Run:   recreateDB,
}

func init() {
	// Database flags
	rootCmd.Flags().StringP("host", "", "", "MySQL host")
	rootCmd.Flags().StringP("port", "", "", "MySQL host")
	rootCmd.Flags().StringP("user", "u", "", "MySQL user")
	rootCmd.Flags().StringP("password", "p", "", "MySQL password")
	rootCmd.Flags().StringP("db", "", "", "MySQL db name")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func recreateDB(cmd *cobra.Command, args []string) {
	mysqlHost, _ := cmd.Flags().GetString("host")
	mysqlPort, _ := cmd.Flags().GetString("port")
	mysqlUser, _ := cmd.Flags().GetString("user")
	mysqlPassword, _ := cmd.Flags().GetString("password")
	mysqlDB, _ := cmd.Flags().GetString("db")

	// Connect to the database
	db, err := gorm.Open(mysql.Open(
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDB),
	), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: " + err.Error())
	}

	// drop schema
	db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", mysqlDB))

	// create schema
	db.Exec(fmt.Sprintf("CREATE DATABASE %s", mysqlDB))

	db.Exec(fmt.Sprintf("USE %s", mysqlDB))

	// Auto-migrate tables
	err = db.AutoMigrate(entity.AllModels...)
	if err != nil {
		log.Fatal("Failed to migrate tables: " + err.Error())
	}

	fmt.Println("Database recreated successfully")
}
