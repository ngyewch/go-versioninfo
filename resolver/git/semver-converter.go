package git

import (
	"fmt"
	"github.com/Masterminds/semver/v3"
	"github.com/ngyewch/go-versioninfo/model"
	"strings"
)

type SemVerVersionInfoConverter struct {
	config       Config
	semVerConfig SemVerVersionInfoConverterConfig
}

type SemVerVersionInfoConverterConfig struct {
	PrereleasePrefix string
}

func NewSemVerVersionInfoConverter(config Config, semVerConfig SemVerVersionInfoConverterConfig) *SemVerVersionInfoConverter {
	return &SemVerVersionInfoConverter{
		config:       config,
		semVerConfig: semVerConfig,
	}
}

func (converter *SemVerVersionInfoConverter) Convert(describeInfo *DescribeInfo) *model.VersionInfo {
	var versionInfo model.VersionInfo

	if describeInfo.Commit != nil {
		versionInfo.Commit = describeInfo.Commit.Hash.String()
		versionInfo.CommitTimestamp = describeInfo.Commit.Committer.When.Unix()
	}

	tagName := converter.config.FallbackTag
	if describeInfo.Tag != nil {
		tagName = describeInfo.Tag.Name
	}
	versionFromTagName := tagName
	if converter.config.DropTagPrefix && strings.HasPrefix(versionFromTagName, converter.config.TagPrefix) {
		versionFromTagName = versionFromTagName[len(converter.config.TagPrefix):]
	}

	resolvedVersion, err := converter.resolveVersion(describeInfo, versionFromTagName)
	if err != nil {
		panic(err) // TODO
	}

	versionInfo.Version = resolvedVersion

	return &versionInfo
}

func (converter *SemVerVersionInfoConverter) resolveVersion(describeInfo *DescribeInfo, versionFromTagName string) (string, error) {
	v, err := semver.NewVersion(versionFromTagName)
	if err != nil {
		return "", err
	}

	metadata := v.Metadata()
	v2, err := v.SetMetadata("")
	if err != nil {
		return "", err
	}

	s := ""
	if !converter.config.DropTagPrefix {
		s += converter.config.TagPrefix
	}
	s += v2.String()
	if (describeInfo.AdditionalCommits > 0) || describeInfo.Dirty {
		if v2.Prerelease() != "" {
			s += "."
		} else {
			s += "-"
		}
		var parts []string
		if converter.semVerConfig.PrereleasePrefix != "" {
			parts = append(parts, converter.semVerConfig.PrereleasePrefix)
		}
		if describeInfo.AdditionalCommits > 0 {
			parts = append(parts,
				fmt.Sprintf("%d", describeInfo.AdditionalCommits),
				fmt.Sprintf("g%s", describeInfo.Commit.Hash.String()[:7]),
			)
		}
		if describeInfo.Dirty {
			parts = append(parts, "dirty")
		}
		s += strings.Join(parts, ".")
	}
	if metadata != "" {
		s += fmt.Sprintf("+%s", metadata)
	}

	return s, nil
}
