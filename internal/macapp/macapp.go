package macapp

import (
	"debug/macho"
	"io/fs"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"howett.net/plist"
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
	Architectures Architectures
}

type executableDecoded struct {
	CFBundleExecutable string
}

func (a *Application) GetExecutableName() (string, error) {
	plistFile := filepath.Join(a.Path, "Contents", "Info.plist")

	f, err := os.Open(plistFile)
	if err != nil {
		return "", err
	}
	defer f.Close()

	decoder := plist.NewDecoder(f)
	var plistDecoded executableDecoded
	err = decoder.Decode(&plistDecoded)
	if err != nil {
		return "", err
	}

	return plistDecoded.CFBundleExecutable, err
}

func (a *Application) GetArchitectures() (Architectures, error) {
	executable, err := a.GetExecutableName()
	if err != nil {
		return a.Architectures, err
	}

	// binary file path
	binary := filepath.Join(a.Path, "Contents", "MacOS", executable)

	fat, err := macho.OpenFat(binary)
	if err == nil {
		// file is Mach-O universal
		a.Architectures.LoadFromFat(fat.Arches)
	} else {
		// file is Mach-O
		f, err := macho.Open(binary)
		if err != nil {
			return a.Architectures, err
		}
		a.Architectures.Load(f.Cpu)
	}

	return a.Architectures, nil
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
	app.GetArchitectures()

	return app, nil
}
