package cmd

import (
	"fmt"
	"github.com/ngyewch/go-versioninfo/model"
	"github.com/ngyewch/go-versioninfo/resolver"
	"github.com/ngyewch/go-versioninfo/resolver/env"
	"github.com/ngyewch/go-versioninfo/resolver/git"
	"github.com/spf13/cobra"
)

var (
	envCmd = &cobra.Command{
		Use:   fmt.Sprintf("env [flags]"),
		Short: "Environment",
		RunE:  doEnv,
	}
)

func doEnv(cmd *cobra.Command, args []string) error {
	export, err := cmd.Flags().GetBool("export")
	if err != nil {
		return err
	}

	disableEnv, err := cmd.Flags().GetBool("disable-env")
	if err != nil {
		return err
	}

	envPrefix, err := cmd.Flags().GetString("env-prefix")
	if err != nil {
		return err
	}

	disableGit, err := cmd.Flags().GetBool("disable-git")
	if err != nil {
		return err
	}

	gitTagPrefix, err := cmd.Flags().GetString("git-tag-prefix")
	if err != nil {
		return err
	}

	gitDropTagPrefix, err := cmd.Flags().GetBool("git-drop-tag-prefix")
	if err != nil {
		return err
	}

	gitFallbackTag, err := cmd.Flags().GetString("git-fallback-tag")
	if err != nil {
		return err
	}

	gitDescribeMode, err := cmd.Flags().GetString("git-describe-mode")
	if err != nil {
		return err
	}

	gitSemVerPrereleasePrefix, err := cmd.Flags().GetString("git-semver-prerelease-prefix")
	if err != nil {
		return err
	}

	var resolvers []resolver.Resolver

	if !disableEnv {
		envResolver := env.New(envPrefix)
		resolvers = append(resolvers, envResolver)
	}

	if !disableGit {
		gitResolverConfig := git.Config{
			TagPrefix:     gitTagPrefix,
			DropTagPrefix: gitDropTagPrefix,
			FallbackTag:   gitFallbackTag,
		}
		var converter git.VersionInfoConverter
		switch gitDescribeMode {
		case "default":
			converter = git.NewDefaultVersionInfoConverter(gitResolverConfig)
		case "semver":
			converter = git.NewSemVerVersionInfoConverter(gitResolverConfig, git.SemVerVersionInfoConverterConfig{
				PrereleasePrefix: gitSemVerPrereleasePrefix,
			})
		default:
			return fmt.Errorf("unknown git-describe-mode")
		}
		gitResolver, err := git.New(gitResolverConfig, converter)
		if err != nil {
			return err
		}
		resolvers = append(resolvers, gitResolver)
	}

	info, err := func() (*model.VersionInfo, error) {
		for _, r := range resolvers {
			info, err := r.Resolve()
			if err != nil {
				continue
			}
			if info != nil {
				return info, nil
			}
		}
		return nil, nil
	}()
	if err != nil {
		return err
	}

	if info != nil {
		exportPreamble := ""
		if export {
			exportPreamble = "export "
		}
		if info.Version != "" {
			fmt.Printf("%s%sVERSION=%s\n", exportPreamble, envPrefix, info.Version)
		}
		if info.Commit != "" {
			fmt.Printf("%s%sCOMMIT=%s\n", exportPreamble, envPrefix, info.Commit)
		}
		if info.CommitTimestamp > 0 {
			fmt.Printf("%s%sCOMMIT_TIMESTAMP=%d\n", exportPreamble, envPrefix, info.CommitTimestamp)
		}
	}

	return nil
}

func init() {
	rootCmd.AddCommand(envCmd)

	envCmd.Flags().Bool("export", true, "Export.")

	envCmd.Flags().Bool("disable-env", false, "Disable env resolver.")
	envCmd.Flags().String("env-prefix", "", "Environment variable prefix.")

	envCmd.Flags().Bool("disable-git", false, "Disable git resolver.")
	envCmd.Flags().String("git-describe-mode", "default", "Git describe mode (default, semver).")
	envCmd.Flags().String("git-tag-prefix", "v", "Git resolver: Tag prefix.")
	envCmd.Flags().Bool("git-drop-tag-prefix", true, "Git resolver: Drop tag prefix.")
	envCmd.Flags().String("git-fallback-tag", "v0.0.0", "Git resolver: Fallback tag.")
	envCmd.Flags().String("git-semver-prerelease-prefix", "dev", "Git resolver: Semver prerelease prefix.")
}
