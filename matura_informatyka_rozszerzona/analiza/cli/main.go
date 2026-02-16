package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	dbDir  string
	mainDB *sql.DB
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "matura",
		Short: "CLI for matura informatyka exam preparation",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Resolve dbDir
			if dbDir == "" {
				if d := os.Getenv("MATURA_DB_DIR"); d != "" {
					dbDir = d
				} else {
					exe, err := os.Executable()
					if err == nil {
						dbDir = filepath.Dir(exe)
					} else {
						dbDir = "."
					}
				}
			}

			// Skip DB open for "data import" — it opens its own DB
			if cmd.Name() == "import" && cmd.Parent() != nil && cmd.Parent().Name() == "data" {
				return nil
			}

			db, err := OpenDB(dbDir)
			if err != nil {
				return fmt.Errorf("open DB: %w", err)
			}
			mainDB = db
			return nil
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			if mainDB != nil {
				mainDB.Close()
			}
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	rootCmd.PersistentFlags().StringVar(&dbDir, "db-dir", "", "Directory for DB files (default: binary dir or MATURA_DB_DIR)")

	// === exercise ===
	exerciseCmd := &cobra.Command{Use: "exercise", Short: "Exercise operations"}
	exerciseCmd.AddCommand(exerciseGetCmd(), exerciseReviewCmd())

	// === progress ===
	progressCmd := &cobra.Command{Use: "progress", Short: "Progress tracking"}
	progressCmd.AddCommand(progressUpdateCmd(), progressStatusCmd())

	// === cke ===
	ckeCmd := &cobra.Command{Use: "cke", Short: "CKE exam subtasks"}
	ckeCmd.AddCommand(ckeGetCmd(), ckeSaveCmd())

	// === exam ===
	examCmd := &cobra.Command{Use: "exam", Short: "Mock exam operations"}
	examCmd.AddCommand(examMetaCmd(), examTaskCmd(), examSaveCmd())

	// === trap ===
	trapCmd := &cobra.Command{Use: "trap", Short: "Exam traps/gotchas"}
	trapCmd.AddCommand(trapListCmd(), trapSaveCmd())

	// === cheatsheet ===
	cheatsheetCmd := &cobra.Command{Use: "cheatsheet", Short: "Study cheatsheets"}
	cheatsheetCmd.AddCommand(cheatsheetGetCmd())

	// === data ===
	dataCmd := &cobra.Command{Use: "data", Short: "Data import and stats"}
	dataCmd.AddCommand(dataImportCmd(), dataStatsCmd())

	rootCmd.AddCommand(exerciseCmd, progressCmd, ckeCmd, examCmd, trapCmd, cheatsheetCmd, dataCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}
