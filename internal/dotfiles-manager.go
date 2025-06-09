package internal

import (
	"encoding/json"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strings"
)

func NewRepo(repoDir, targetDir string) (repo *Repo) {
	repo = &Repo{
		repoDir:   repoDir,
		targetDir: targetDir,
	}
	return repo
}

type Repo struct {
	repoDir   string
	targetDir string
	Config    *Config
}

type Config struct {
	NoFold PathTree `json:"no fold"`
}

type PathTree map[string]PathTree

func (t PathTree) Add(path string) {
	path = strings.TrimPrefix(filepath.Clean(path), "/")
	firstSlash := strings.IndexRune(path, '/')

	var parts []string
	for len(path) != firstSlash+1 {
		var part string
		path, part = filepath.Split(path)
		parts = append(parts, part)
		if part == "" {
			break
		}
	}
	if len(path) > 0 {
		parts = append(parts, path)
	}
	slices.Reverse(parts)
	for _, name := range parts {
		if t[name] == nil {
			t[name] = PathTree{}
		}
		t = t[name]
	}
}

func (repo *Repo) SaveConfig() (err error) {
	if repo.Config == nil {
		return // No need to save
	}
	file, err := os.Create(path.Join(repo.repoDir, "config.json"))
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(file)
	return encoder.Encode(repo.Config)
}

func (repo *Repo) LoadConfig(create bool) (err error) {
	var config Config

	configFile, err := os.Open(path.Join(repo.repoDir, "config.json"))
	if os.IsNotExist(err) {
		if create {
			repo.Config = &config
		}
		return nil // No config is not an error
	}

	decoder := json.NewDecoder(configFile)
	decoder.Decode(&config)
	repo.Config = &config
	return nil
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
