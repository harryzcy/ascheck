package macapp

import (
	"io/fs"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/harryzcy/ascheck/internal/localcheck"
	"github.com/harryzcy/ascheck/internal/remotecheck"
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
	// Name shows the app name
	Name string

	// Path shows the physical location
	Path string
	// Architectures represents the architectures of the currently installed version
	Architectures localcheck.Architectures

	// Website shows the app's website, empty if unknown
	Website string
	// ArmSupport shows the Apple Silicon support based on Does It Arm reports
	ArmSupport remotecheck.Support
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
				app := checkApplication(dir, f)
				applications = append(applications, app)
			}
		}
	}

	return applications, nil
}

func checkApplication(dir string, f fs.FileInfo) Application {
	app := Application{
		Name: strings.TrimSuffix(f.Name(), ".app"),
		Path: filepath.Join(dir, f.Name()),
	}

	app.Architectures, _ = localcheck.GetArchitectures(app.Path)

	info, err := remotecheck.GetInfo(app.Name)
	if err == nil {
		app.Website = info.Website
		app.ArmSupport = info.ArmSupport
	}

	return app
}
