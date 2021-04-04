package localcheck

import (
	"os"
	"path/filepath"

	"howett.net/plist"
)

// appInfo contains the app's path on file system and some keys defined in Info.plist file.
type AppInfo struct {
	// path is the app's path on file system
	path string

	// These fields are used for determining localized app name and executable file's name,
	// and they are documented by Apple at
	// https://developer.apple.com/library/archive/documentation/General/Reference/InfoPlistKeyReference/Articles/CoreFoundationKeys.html.
	CFBundleDisplayName string
	CFBundleExecutable  string
}

func (info AppInfo) GetDisplayName() string {
	return info.CFBundleDisplayName
}

func (info AppInfo) GetExecutableName() string {
	return info.CFBundleExecutable
}

// loadLocalizedDisplayName loads CFBundleDisplayName
// using the rule consistent with what macOS's Finder does.
//
// It comapres the value of this key against the actual bundle name on the file system.
// If the two names match, the localized name from the appropriate InfoPlist.strings will be returned.
// Otherwise, the file system name will be returned.
func (info AppInfo) loadLocalizedDisplayName() {

}

func GetAppInfo(appPath string, lang string) (AppInfo, error) {
	var (
		info AppInfo
		err  error
	)

	info.path = appPath
	plistFile := filepath.Join(appPath, "Contents", "Info.plist")

	f, err := os.Open(plistFile)
	if err != nil {
		return info, err
	}
	defer f.Close()

	decoder := plist.NewDecoder(f)
	err = decoder.Decode(&info)
	if err != nil {
		return info, err
	}

	info.loadLocalizedDisplayName()

	return info, nil
}
