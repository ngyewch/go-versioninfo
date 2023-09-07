package git

import (
	"errors"
	"github.com/Masterminds/semver/v3"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"os/exec"
	"strings"
)

type Helper struct {
	repo *git.Repository
}

type DescribeInfo struct {
	Commit            *object.Commit
	Tag               *object.Tag
	AdditionalCommits int
	Dirty             bool
}

func NewHelper(repo *git.Repository) *Helper {
	return &Helper{
		repo: repo,
	}
}

func (helper *Helper) GetTags() ([]*object.Tag, error) {
	var tags []*object.Tag
	tagsIter, err := helper.repo.Tags()
	if err != nil {
		return nil, err
	}
	err = tagsIter.ForEach(func(ref *plumbing.Reference) error {
		obj, err := helper.repo.TagObject(ref.Hash())
		if err == nil {
			if strings.HasPrefix(obj.Name, "v") {
				_, err := semver.NewVersion(obj.Name[1:])
				if err == nil {
					tags = append(tags, obj)
				}
			}
		} else if errors.Is(err, plumbing.ErrObjectNotFound) {
			// not a tag object
		} else {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (helper *Helper) Describe(tagPrefix string, checkDirty bool) (*DescribeInfo, error) {
	var describe DescribeInfo

	tags, err := helper.GetTags()
	if err != nil {
		return nil, err
	}

	head, err := helper.repo.Head()
	if err != nil {
		return nil, err
	}

	cIter, err := helper.repo.Log(&git.LogOptions{
		From:  head.Hash(),
		Order: git.LogOrderCommitterTime,
	})
	if err != nil {
		return nil, err
	}

	additionalCommits := 0
	err = cIter.ForEach(func(c *object.Commit) error {
		if additionalCommits == 0 {
			describe.Commit = c
		}
		for _, tag := range tags {
			if strings.HasPrefix(tag.Name, tagPrefix) && (c.Hash.String() == tag.Target.String()) {
				describe.Tag = tag
				describe.AdditionalCommits = additionalCommits
				return storer.ErrStop
			}
		}
		additionalCommits++
		return nil
	})

	worktree, err := helper.repo.Worktree()
	if err != nil {
		return nil, err
	}
	rootPath := worktree.Filesystem.Root()

	if checkDirty {
		cmd := exec.Command("git", "status", "--short")
		cmd.Dir = rootPath
		gitStatusOutputBytes, err := cmd.Output()
		if err != nil {
			return nil, err
		}
		gitStatusOutput := string(gitStatusOutputBytes)
		if strings.TrimSpace(gitStatusOutput) != "" {
			describe.Dirty = true
		}
	}

	return &describe, nil
}
