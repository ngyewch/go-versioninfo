package github

import (
	"github.com/ngyewch/go-versioninfo/formatter"
	"github.com/ngyewch/go-versioninfo/model"
	"os"
	"strings"
)

const tagPrefix = "refs/tags/"

type Resolver struct {
	config    Config
	formatter formatter.Formatter
}

type Config struct {
}

func New(config Config, formatter formatter.Formatter) *Resolver {
	return &Resolver{
		config:    config,
		formatter: formatter,
	}
}

func (resolver *Resolver) Resolve() (*model.VersionInfo, error) {
	if os.Getenv("GITHUB_ACTIONS") != "true" {
		return nil, nil
	}

	info := &formatter.SimpleDescribeInfo{
		Commit: os.Getenv("GITHUB_SHA"),
	}

	ref := os.Getenv("GITHUB_REF")
	if strings.HasPrefix(ref, tagPrefix) {
		info.Tag = ref[len(tagPrefix):]
	}

	return resolver.formatter.Format(info), nil
}
