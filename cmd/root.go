package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   fmt.Sprintf("%s [flags]", filepath.Base(os.Args[0])),
		Short: filepath.Base(os.Args[0]),
		RunE:  help,
	}
)
