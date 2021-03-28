package macapp

import (
	"io/fs"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/harryzcy/ascheck/internal/localcheck"
)

var (
	applicationPath []string
)

func init() {
	usr, _ := user.Current()
	userApplication := filepath.Join(usr.HomeDir, "Applications")

	applicationPath = []string{
		"/System/Applications",
		"/Applications",
		userApplication,
	}
}

type Application struct {
	Name          string
	Path          string
	Architectures localcheck.Architectures
}

// GetAllApplications returns all applications
func GetAllApplications(dirs []string) ([]Application, error) {
	var (
		applications []Application
	)

	if dirs == nil {
		dirs = applicationPath
	}

	for _, dir := range dirs {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
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

	var err error
	app.Architectures, err = localcheck.GetArchitectures(app.Path)
	if err != nil {
		return Application{}, err
	}

	return app, nil
}
