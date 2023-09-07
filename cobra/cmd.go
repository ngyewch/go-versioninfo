package cobra

import (
	"encoding/json"
	"fmt"
	"github.com/goccy/go-yaml"
	"github.com/ngyewch/go-versioninfo/model"
	"github.com/spf13/cobra"
	"time"
)

type versionInfoOutput struct {
	Version         string `json:"Version"`
	Commit          string `json:"Commit"`
	CommitTimestamp string `json:"CommitTimestamp"`
}

func AddVersionCmd(cmd *cobra.Command, versionInfoProvider func() *model.VersionInfo) {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Version",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			short, err := cmd.Flags().GetBool("short")
			if err != nil {
				return err
			}

			format, err := cmd.Flags().GetString("format")
			if err != nil {
				return err
			}

			versionInfo := versionInfoProvider()

			if short {
				fmt.Println(versionInfo.Version)
				return nil
			}

			o := versionInfoOutput{
				Version:         versionInfo.Version,
				Commit:          versionInfo.Commit,
				CommitTimestamp: time.Unix(versionInfo.CommitTimestamp, 0).Local().Format(time.RFC3339),
			}

			switch format {
			case "json":
				b, err := json.MarshalIndent(o, "", "  ")
				if err != nil {
					return err
				}
				fmt.Println(string(b))
			case "yaml":
				b, err := yaml.MarshalWithOptions(o)
				if err != nil {
					return err
				}
				fmt.Println(string(b))
			default:
				return fmt.Errorf("unknown format")
			}

			return nil
		},
	}
	versionCmd.Flags().Bool("short", false, "Print just the version number.")
	versionCmd.Flags().String("format", "yaml", "Output format. One of 'yaml' or 'json'.")
	versionCmd.MarkFlagsMutuallyExclusive("short", "format")

	cmd.AddCommand(versionCmd)
}
