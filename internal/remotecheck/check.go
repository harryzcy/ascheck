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

	infoCache map[string]appInfo = make(map[string]appInfo)
)

type appInfo struct {
	Website string
	Support string
}

// getInfo loads the list of reported app support from Does it ARM.
func getInfo() (map[string]appInfo, error) {
	resp, err := http.Get(sourceURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	matches := pattern.FindAllStringSubmatch(string(body), -1)

	for _, match := range matches {
		name := match[1]
		info := appInfo{
			Website: match[2],
			Support: match[3],
		}
		infoCache[name] = info
	}

	return infoCache, nil
}
