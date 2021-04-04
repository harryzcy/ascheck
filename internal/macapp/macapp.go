package macapp

import (
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

// Application represents an installed app.
type Application struct {
	// Name represents the app bundle name
	Name string

	// DisplayName represents the localized app name
	DisplayName string

	// Path shows the physical location
	Path string
	// Architectures represents the architectures of the currently installed version
	Architectures localcheck.Architectures

	// Website shows the app's website, empty if unknown
	Website string
	// ArmSupport shows the Apple Silicon support based on Does It Arm reports
	ArmSupport remotecheck.Support
}

// GetAllApplications returns all applications.
func GetAllApplications(dirs []string, lang string) ([]Application, error) {
	var (
		applications []Application
	)

	if dirs == nil {
		dirs = applicationPath
	}

	for _, dir := range dirs {
		entries, err := os.ReadDir(dir)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return nil, err
		}

		for _, entry := range entries {
			if strings.HasSuffix(entry.Name(), ".app") {
				app := checkApplication(dir, entry, lang)
				applications = append(applications, app)
			}
		}
	}

	return applications, nil
}

func checkApplication(dir string, entry os.DirEntry, lang string) Application {
	app := Application{
		Name: strings.TrimSuffix(entry.Name(), ".app"),
		Path: filepath.Join(dir, entry.Name()),
	}

	localInfo, err := localcheck.GetAppInfo(app.Path, lang)
	if err != nil {
		return app
	}
	app.DisplayName = localInfo.GetDisplayName()
	app.Architectures, _ = localInfo.GetArchitectures()

	// mark system apps as natively supported
	if strings.HasPrefix(dir, "/System/") {
		app.ArmSupport = remotecheck.SupportNative
		return app
	}

	info, err := remotecheck.GetInfo(app.Name)
	if err == nil {
		app.Website = info.Website
		app.ArmSupport = info.ArmSupport
	}

	return app
}
