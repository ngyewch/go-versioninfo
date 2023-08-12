package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var (
	rootCmd = &cobra.Command{
		Use:   fmt.Sprintf("%s [flags]", filepath.Base(os.Args[0])),
		Short: filepath.Base(os.Args[0]),
		RunE:  help,
	}
)
