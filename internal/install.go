package internal

import (
	"fmt"
	"os"
	"path"
	"strings"
)

// Install the given modules
func (repo *Repo) Install(modules []string) (err error) {
	allModules, err := repo.ModuleSet()
	if err != nil {
		return err
	}

	for _, module := range modules {
		_, ok := allModules[module]
		if !ok {
			return fmt.Errorf("Module %q doesn't exist")
		}
	}

	modulePaths := make([]string, len(modules))
	for i, module := range modules {
		modulePaths[i] = repo.ModulePath(module)
	}
	return repo.install(modulePaths, repo.targetDir, repo.Config.NoFold)
}

func (repo *Repo) install(sources []string, target string, nofold PathTree) (err error) {
	if len(sources) == 0 {
		return nil
	}

	if len(sources) == 1 {
		return repo.installSingle(sources[0], target, nofold)
	}

	for _, source := range sources {
		stat, err := os.Stat(source)
		if err != nil {
			return err
		}
		if !stat.IsDir() {
			return MultipleFiles{
				files: sources,
			}
		}
	}

	targetStat, err := os.Lstat(target)
	switch {
	case err != nil && os.IsNotExist(err):
		err := os.Mkdir(target, 0777)
		if err != nil {
			return err
		}
		fallthrough
	case targetStat.IsDir(): // Target is directory
		return repo.installDirsToDir(sources, target, nofold)

	case err != nil:
		return err

	case targetStat.Mode()&os.ModeSymlink != 0: // Target is symlink

		location, err := readLinkAbs(target)
		if err != nil {
			return err
		}
		if !strings.HasPrefix(location, repo.repoDir) {
			return fmt.Errorf("file %s already exists", target)
		}

		// Could probably be made easier using a call to Repo.Unfold
		err = os.Remove(target)
		if err != nil {
			return err
		}
		err = os.Mkdir(target, 0777)
		if err != nil {
			return err
		}
		return repo.install(append(sources, location), target, nofold)

	default:
		return TargetExists(target)
	}
}

// Assumes both target, and all sources are paths to directories
func (repo *Repo) installDirsToDir(sources []string, target string, nofold PathTree) (err error) {
	newPaths := map[string][]string{}
	for _, source := range sources {
		content, err := os.ReadDir(source)
		if err != nil {
			return err
		}
		for _, entry := range content {
			newPaths[entry.Name()] = append(newPaths[entry.Name()], path.Join(source, entry.Name()))
		}
	}
	for key, val := range newPaths {
		if err = repo.install(val, path.Join(target, key), nofold[key]); err != nil {
			return err
		}
	}
	return nil
}

func (repo *Repo) installSingle(source string, target string, nofold PathTree) (err error) {
	targetStat, err := os.Lstat(target)
	if os.IsNotExist(err) {
		if nofold == nil {
			return os.Symlink(source, target)
		} else {
			err = os.Mkdir(target, 0777)
			if err != nil {
				return err
			}
			targetStat, err = os.Lstat(target)
		}
	}
	if err != nil {
		return err
	}

	sourceStat, err := os.Stat(source)
	if err != nil {
		return err
	}

	if targetStat.IsDir() && sourceStat.IsDir() {
		entries, err := os.ReadDir(source)
		if err != nil {
			return err
		}
		for _, entry := range entries {
			repo.installSingle(
				path.Join(source, entry.Name()),
				path.Join(target, entry.Name()),
				nofold[entry.Name()],
			)
		}
		return nil
	}

	if targetStat.Mode()&os.ModeSymlink != 0 {
		location, err := readLinkAbs(target)
		if err != nil {
			return err
		}
		if location == source {
			return nil // Already installed
		}
		if !strings.HasPrefix(location, repo.repoDir) {
			return TargetExists(target)
		}

		err = os.Remove(target)
		if err != nil {
			return err
		}
		err = os.Mkdir(target, 0777)
		if err != nil {
			return err
		}

		return repo.install([]string{location, source}, target, nofold)
	}

	return TargetExists(target)
}

func (repo *Repo) installSingleToNonexistentTarget(source, target string) (err error) {
	return os.Symlink(source, target)
}
