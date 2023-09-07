package cmd

import (
	"fmt"
	cobra2 "github.com/ngyewch/go-versioninfo/cobra"
	"github.com/ngyewch/go-versioninfo/v"
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

func init() {
	cobra2.AddVersionCmd(rootCmd, v.GetVersionInfo)
}
