package remotecheck

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

const (
	sourceURL = "https://cdn.jsdelivr.net/gh/ThatGuySam/doesitarm@master/README.md"
)

var (
	pattern, _ = regexp.Compile(`\* \[(.*?)\]\((.*?)\) - (✅|✳️|⏹|🚫|🔶)`)

	infoCache map[string]AppInfo = make(map[string]AppInfo)
)

type AppInfo struct {
	Website    string
	ArmSupport Support
}

// Init loads the list of reported app Arm from Does it ARM.
func Init() error {
	resp, err := http.Get(sourceURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	matches := pattern.FindAllStringSubmatch(string(body), -1)

	for _, match := range matches {
		name := match[1]
		info := AppInfo{
			Website: match[2],
		}
		info.ArmSupport.Parse(match[3])
		infoCache[name] = info
	}

	return nil
}

func GetInfo(name string) (AppInfo, error) {
	if info, ok := infoCache[name]; ok {
		return info, nil
	}
	return AppInfo{}, ErrNotFound
}
