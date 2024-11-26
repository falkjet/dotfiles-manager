package internal

import (
	"os"
	"path"
)

func (repo *Repo) Uninstall(modules []string) (err error) {
	allModules, err := repo.ModuleSet()
	if err != nil {
		return err
	}

	for _, module := range modules {
		_, ok := allModules[module]
		if ok {
			_ = repo.uninstall(repo.ModulePath(module), repo.targetDir)
		}
	}
	return nil
}

func (repo *Repo) uninstall(source, target string) (err error) {
	stat, err := os.Lstat(target)
	switch {
	case err != nil && os.IsNotExist(err):
		return nil
	case err != nil:
		return err
	case stat.IsDir():
		entries, err := os.ReadDir(source)
		if err != nil {
			return err
		}
		for _, entry := range entries {
			repo.uninstall(path.Join(source, entry.Name()), path.Join(target, entry.Name()))
		}
		return nil
	case stat.Mode()&os.ModeSymlink != 0:
		location, err := readLinkAbs(target)
		if err != nil {
			return err
		}
		if source == location {
			return os.Remove(location)
		}
		return nil
	default:
		return nil
	}
}
