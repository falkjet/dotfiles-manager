package internal

import (
	"fmt"
	"os"
	"strings"
)

func (repo *Repo) Unfold(path string) (err error) {
	if strings.HasSuffix(path, "/") {
		path = path[:len(path)-1]
	}
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

	return repo.install([]string{location}, path)
}
