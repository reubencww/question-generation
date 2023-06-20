package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"senkawa.moe/haa-chan/app/config"
	database "senkawa.moe/haa-chan/app/db"
	logger "senkawa.moe/haa-chan/app/log"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate the database",
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.NewZapLogger(config.IsDevelopment()).Sugar()
		db := database.NewDB()

		shouldDropTables, _ := cmd.Flags().GetBool("fresh")
		if shouldDropTables && config.IsDevelopment() {
			if err := db.Migrator().DropTable(database.AvailableTables...); err == nil {
				fmt.Println("Dropped all tables.")
				log.Info("Dropped all tables.")
			}
		}

		if err := db.AutoMigrate(database.AvailableTables...); err != nil {
			fmt.Printf("Failed to migrate tables: %v", err)
		}

		fmt.Println("Tables migrated.")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().Bool("seed", false, "Seed the database with a test user")
	migrateCmd.Flags().Bool("fresh", false, "Drop all tables before migrating (no-op in production)")
}
