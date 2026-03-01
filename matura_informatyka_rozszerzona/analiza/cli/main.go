package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var dbDir string

func main() {
	var openedDB *sql.DB

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
			if cmd.Parent() != nil && cmd.Parent().Name() == "test-report" {
				return nil
			}

			db, attached, err := OpenDB(dbDir)
			if err != nil {
				return fmt.Errorf("open DB: %w", err)
			}
			openedDB = db

			if !attached {
				return fmt.Errorf("matura.db not found in %s — run: matura data import --source <path>", dbDir)
			}

			cmd.SetContext(context.WithValue(cmd.Context(), ctxKey{}, db))
			return nil
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	rootCmd.PersistentFlags().StringVar(&dbDir, "db-dir", "", "Directory for DB files (default: binary dir or MATURA_DB_DIR)")

	// === exercise ===
	exerciseCmd := &cobra.Command{Use: "exercise", Short: "Exercise operations"}
	exerciseCmd.AddCommand(exerciseQuestionCmd(), exerciseHintsCmd(), exerciseAnswerCmd(), exerciseReviewCmd(), exerciseNextCmd(), exerciseRubricCmd(), exerciseCountCmd(), exerciseSuggestErrorCmd(), exerciseCheckAnswerCmd())

	// === progress ===
	progressCmd := &cobra.Command{Use: "progress", Short: "Progress tracking"}
	progressCmd.AddCommand(progressUpdateCmd(), progressStatusCmd(), progressBladCmd(), progressDiagnoseCmd())

	// === cke ===
	ckeCmd := &cobra.Command{Use: "cke", Short: "CKE exam subtasks"}
	ckeCmd.AddCommand(ckeGetCmd(), ckeWorkedExampleCmd(), ckeSaveCmd(), ckeStatusCmd())

	// === exam ===
	examCmd := &cobra.Command{Use: "exam", Short: "Mock exam operations"}
	examCmd.AddCommand(examMetaCmd(), examTaskCmd(), examSaveCmd(), examListCmd())

	// === typ ===
	typCmd := &cobra.Command{Use: "typ", Short: "Type information"}
	typCmd.AddCommand(typIntroCmd())

	// === trap ===
	trapCmd := &cobra.Command{Use: "trap", Short: "Exam traps/gotchas"}
	trapCmd.AddCommand(trapListCmd(), trapSaveCmd())

	// === cheatsheet ===
	cheatsheetCmd := &cobra.Command{Use: "cheatsheet", Short: "Study cheatsheets"}
	cheatsheetCmd.AddCommand(cheatsheetGetCmd())

	// === data ===
	dataCmd := &cobra.Command{Use: "data", Short: "Data import and stats"}
	dataCmd.AddCommand(dataImportCmd(), dataStatsCmd(), dataVerifyCmd())

	// === test-report ===
	testReportCmd := &cobra.Command{Use: "test-report", Short: "Test-tutor report analysis"}
	testReportCmd.AddCommand(testReportSummaryCmd())

	rootCmd.AddCommand(exerciseCmd, progressCmd, ckeCmd, examCmd, typCmd, trapCmd, cheatsheetCmd, dataCmd, testReportCmd)

	err := rootCmd.Execute()
	if openedDB != nil {
		openedDB.Close()
	}
	if err != nil {
		var nf errNotFound
		if errors.As(err, &nf) {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}
