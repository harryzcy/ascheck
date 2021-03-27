package macapp

import (
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"
)

const (
	applicationPath = "/Applications"
)

type Application struct {
	Name          string
	Path          string
	Architectures []string
}

// GetAllApplications returns all applications
func GetAllApplications(dirs []string) ([]Application, error) {
	var (
		applications []Application
	)

	if dirs == nil {
		dirs = []string{applicationPath}
	}

	for _, dir := range dirs {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			return nil, err
		}

		for _, f := range files {
			if strings.HasSuffix(f.Name(), ".app") {
				app, err := resolveApplication(dir, f)
				if err != nil {
					return nil, err
				}
				applications = append(applications, app)
			}
		}
	}

	return applications, nil
}

func resolveApplication(dir string, f fs.FileInfo) (Application, error) {
	app := Application{
		Name: strings.TrimSuffix(f.Name(), ".app"),
		Path: filepath.Join(dir, f.Name()),
	}

	return app, nil
}
