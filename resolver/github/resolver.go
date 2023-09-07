package github

import (
	"github.com/go-git/go-git/v5"
	"github.com/ngyewch/go-versioninfo/formatter"
	"github.com/ngyewch/go-versioninfo/model"
	gitResolver "github.com/ngyewch/go-versioninfo/resolver/git"
	"os"
	"strings"
)

const tagPrefix = "refs/tags/"

type Resolver struct {
	config    Config
	formatter formatter.Formatter
	repo      *git.Repository
}

type Config struct {
	CheckDirty bool
}

func New(config Config, formatter formatter.Formatter) (*Resolver, error) {
	repo, err := gitResolver.FindRepository(".")
	if err != nil {
		return nil, err
	}

	return &Resolver{
		config:    config,
		formatter: formatter,
		repo:      repo,
	}, nil
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

	// TODO Find other ways to get commit timestamp
	_ = func() error {
		head, err := resolver.repo.Head()
		if err != nil {
			return err
		}

		cIter, err := resolver.repo.Log(&git.LogOptions{
			From:  head.Hash(),
			Order: git.LogOrderCommitterTime,
		})
		if err != nil {
			return err
		}

		c, err := cIter.Next()
		if err != nil {
			return err
		}

		info.CommitTimestamp = c.Committer.When.Unix()

		return nil
	}()

	return resolver.formatter.Format(info), nil
}
