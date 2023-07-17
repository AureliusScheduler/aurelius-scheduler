package main

import (
	"aurelius/libs/aurelius-mysql/db_context"
	"aurelius/libs/aurelius-mysql/entity"
	"aurelius/libs/aurelius-mysql/seeder"
	"fmt"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "aurelius-mysql",
	Short: "CLI application to recreate the database",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var recreateDBCmd = &cobra.Command{
	Use:   "recreate",
	Short: "Recreate the database",
	Run:   recreateDB,
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed the database",
	Run:   seed,
}

func init() {
	rootCmd.AddCommand(recreateDBCmd)
	rootCmd.AddCommand(seedCmd)

	// Database flags
	recreateDBCmd.Flags().StringP("host", "", "", "MySQL host")
	recreateDBCmd.Flags().StringP("port", "", "", "MySQL host")
	recreateDBCmd.Flags().StringP("user", "u", "", "MySQL user")
	recreateDBCmd.Flags().StringP("password", "p", "", "MySQL password")
	recreateDBCmd.Flags().StringP("db", "", "", "MySQL db name")

	seedCmd.Flags().StringP("host", "", "", "MySQL host")
	seedCmd.Flags().StringP("port", "", "", "MySQL host")
	seedCmd.Flags().StringP("user", "u", "", "MySQL user")
	seedCmd.Flags().StringP("password", "p", "", "MySQL password")
	seedCmd.Flags().StringP("db", "", "", "MySQL db name")
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

func seed(cmd *cobra.Command, args []string) {
	mysqlHost, _ := cmd.Flags().GetString("host")
	mysqlPort, _ := cmd.Flags().GetString("port")
	mysqlUser, _ := cmd.Flags().GetString("user")
	mysqlPassword, _ := cmd.Flags().GetString("password")
	mysqlDB, _ := cmd.Flags().GetString("db")

	ctx := db_context.NewDbContext(&db_context.DbOptions{
		Host:     mysqlHost,
		Port:     mysqlPort,
		User:     mysqlUser,
		Password: mysqlPassword,
		Database: mysqlDB,
	})

	s := seeder.NewSeeder()
	s.Seed(ctx)

	fmt.Println("Database seeded successfully")
}
