package git

import (
	"errors"
	"github.com/go-git/go-git/v5"
	"path/filepath"
)

func FindRepository(dir string) (*git.Repository, error) {
	absDir, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}
	dir = absDir
	for {
		repo, err := git.PlainOpen(dir)
		if err != nil {
			if errors.Is(err, git.ErrRepositoryNotExists) {
				if dir == "/" {
					return nil, git.ErrRepositoryNotExists
				}
				dir = filepath.Dir(dir)
				if dir == "" {
					return nil, git.ErrRepositoryNotExists
				}
			} else {
				return nil, err
			}
		} else {
			return repo, nil
		}
	}
}
