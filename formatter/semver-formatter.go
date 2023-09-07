package formatter

import (
	"fmt"
	"github.com/Masterminds/semver/v3"
	"github.com/ngyewch/go-versioninfo/model"
	"strings"
)

type SemVerFormatter struct {
	config       Config
	semVerConfig SemVerConfig
}

type SemVerConfig struct {
	PrereleasePrefix string
}

func NewSemVerFormatter(config Config, semVerConfig SemVerConfig) *SemVerFormatter {
	return &SemVerFormatter{
		config:       config,
		semVerConfig: semVerConfig,
	}
}

func (formatter *SemVerFormatter) Format(info *SimpleDescribeInfo) *model.VersionInfo {
	versionInfo := &model.VersionInfo{
		Commit:          info.Commit,
		CommitTimestamp: info.CommitTimestamp,
	}

	processedTag := formatter.config.ProcessTag(info.Tag)
	resolvedVersion, err := formatter.resolveVersion(info, processedTag)
	if err != nil {
		panic(err) // TODO
	}

	versionInfo.Version = resolvedVersion

	return versionInfo
}

func (formatter *SemVerFormatter) resolveVersion(info *SimpleDescribeInfo, processedTag string) (string, error) {
	v, err := semver.NewVersion(processedTag)
	if err != nil {
		return "", err
	}

	metadata := v.Metadata()
	v2, err := v.SetMetadata("")
	if err != nil {
		return "", err
	}

	s := v2.String()
	if (info.AdditionalCommits > 0) || info.Dirty {
		if v2.Prerelease() != "" {
			s += "."
		} else {
			s += "-"
		}
		var parts []string
		if formatter.semVerConfig.PrereleasePrefix != "" {
			parts = append(parts, formatter.semVerConfig.PrereleasePrefix)
		}
		if info.AdditionalCommits > 0 {
			parts = append(parts,
				fmt.Sprintf("%d", info.AdditionalCommits),
				fmt.Sprintf("g%s", info.Commit[:7]),
			)
		}
		if info.Dirty {
			parts = append(parts, "dirty")
		}
		s += strings.Join(parts, ".")
	}
	if metadata != "" {
		s += fmt.Sprintf("+%s", metadata)
	}

	return s, nil
}
