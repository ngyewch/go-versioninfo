package v

import (
	"github.com/ngyewch/go-versioninfo/formatter"
	"github.com/ngyewch/go-versioninfo/resolver"
	"github.com/ngyewch/go-versioninfo/resolver/env"
	"github.com/ngyewch/go-versioninfo/resolver/git"
	"github.com/ngyewch/go-versioninfo/resolver/github"
)

var (
	defaultFallbackTag      = "v0.0.0"
	defaultTagPrefix        = "v"
	defaultPrereleasePrefix = "dev"

	defaultFormatter = formatter.NewDefaultFormatter(formatter.Config{
		FallbackTag: defaultFallbackTag,
		TagPrefix:   defaultTagPrefix,
	})
	semVerFormatter = formatter.NewSemVerFormatter(formatter.Config{
		FallbackTag: defaultFallbackTag,
		TagPrefix:   defaultTagPrefix,
	}, formatter.SemVerConfig{
		PrereleasePrefix: defaultPrereleasePrefix,
	})
)

func DefaultFormatter() formatter.Formatter {
	return defaultFormatter
}

func SemVerFormatter() formatter.Formatter {
	return semVerFormatter
}

func DefaultResolvers() (resolver.Resolvers, error) {
	return []resolver.Resolver{
		env.New(""),
		github.New(github.Config{}, SemVerFormatter()),
		git.New(git.Config{
			TagPrefix:  defaultTagPrefix,
			CheckDirty: true,
		}, SemVerFormatter()),
	}, nil
}
