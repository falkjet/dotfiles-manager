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
	pathRelativeToRepo = strings.TrimPrefix(pathRelativeToRepo, "/")
	repo.addUnfold(path)

	location, err := readLinkAbs(path)
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

func (repo *Repo) addUnfold(path string) {

}
