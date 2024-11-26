package internal

import (
	"os"
	"path"
	"strings"
)

func NewRepo(repoDir, targetDir string) *Repo {
	return &Repo{
		repoDir:   repoDir,
		targetDir: targetDir,
	}
}

type Repo struct {
	repoDir   string
	targetDir string
}

func (repo *Repo) ModulePath(module string) string {
	return path.Join(repo.repoDir, module)
}

func (repo *Repo) ListModules() (modules []string, err error) {
	modules = []string{}
	entries, err := os.ReadDir(repo.repoDir)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		if !strings.HasPrefix(entry.Name(), ".") && entry.IsDir() {
			modules = append(modules, entry.Name())
		}
	}
	return modules, nil
}

func (repo *Repo) ModuleSet() (modules map[string]struct{}, err error) {
	allModules, err := repo.ListModules()
	if err != nil {
		return nil, err
	}

	modules = make(map[string]struct{})

	for _, module := range allModules {
		modules[module] = struct{}{}
	}
	return modules, nil
}
