package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (repo *Repo) Unfold(path string) (err error) {
	if strings.HasSuffix(path, "/") {
		path = path[:len(path)-1]
	}

	abspath, err := filepath.Abs(path)
	if !strings.HasPrefix(abspath, repo.targetDir) {
		return fmt.Errorf("Path not in repo: %q", path)
	}
	pathRelativeToRepo := strings.TrimPrefix(abspath, repo.targetDir)

	if repo.Config.NoFold == nil {
		repo.Config.NoFold = PathTree{}
	}
	repo.Config.NoFold.Add(pathRelativeToRepo)

	stat, err := os.Lstat(abspath)
	if err != nil {
		return err
	}
	if stat.Mode() & os.ModeSymlink == 0 {
		return
	}

	location, err := readLinkAbs(abspath)
	if err != nil {
		return err
	}
	if !strings.HasPrefix(location, repo.repoDir) {
		return fmt.Errorf("Not a symlink to file in repo")
	}
	err = os.Remove(path)
	if err != nil {
		return err
	}
	err = os.Mkdir(path, 0777)
	if err != nil {
		return err
	}

	return repo.install([]string{location}, path, nil)
}
