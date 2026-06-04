package extensions

import "cmp"
import "fmt"
import "slices"
import "strings"
import "math/rand/v2"

type UserAgent struct {
	Name     string             `json:"name"`
	Platform UserAgentPlatform  `json:"platform"`
	Engine   UserAgentEngine    `json:"engine"`
	Browsers []UserAgentBrowser `json:"browsers"`
}

type UserAgentPlatform struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Arch    string `json:"arch"`
}

type UserAgentEngine struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type UserAgentBrowser struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func (user_agent *UserAgent) sortBrowsers() {

	order := map[string]int{
		"chrome":   0,
		"chromium": 0,
		"safari":   1,
		"edg":      2,
		"edge":     2,
	}

	slices.SortFunc(user_agent.Browsers, func(a, b UserAgentBrowser) int {

		priority_a, ok_a := order[strings.ToLower(a.Name)]
		priority_b, ok_b := order[strings.ToLower(b.Name)]

		switch {
		case ok_a && ok_b:
			return cmp.Compare(priority_a, priority_b)
		case ok_a:
			return -1
		case ok_b:
			return 1
		default:
			return cmp.Compare(a.Name, b.Name)
		}

	})

}

func (user_agent *UserAgent) String() string {

	result := ""

	if len(user_agent.Browsers) == 1 {

		result = fmt.Sprintf(
			"Mozilla/5.0 (%s %s; %s) %s/%s (KHTML, like Gecko) %s/%s",
			user_agent.Platform.Name,
			user_agent.Platform.Version,
			user_agent.Platform.Arch,
			user_agent.Engine.Name,
			user_agent.Engine.Version,
			user_agent.Browsers[0].Name,
			user_agent.Browsers[0].Version,
		)

	} else if len(user_agent.Browsers) == 2 {

		result = fmt.Sprintf(
			"Mozilla/5.0 (%s %s; %s) %s/%s (KHTML, like Gecko) %s/%s %s/%s",
			user_agent.Platform.Name,
			user_agent.Platform.Version,
			user_agent.Platform.Arch,
			user_agent.Engine.Name,
			user_agent.Engine.Version,
			user_agent.Browsers[0].Name,
			user_agent.Browsers[0].Version,
			user_agent.Browsers[1].Name,
			user_agent.Browsers[1].Version,
		)

	} else if len(user_agent.Browsers) == 3 {

		result = fmt.Sprintf(
			"Mozilla/5.0 (%s %s; %s) %s/%s (KHTML, like Gecko) %s/%s %s/%s %s/%s",
			user_agent.Platform.Name,
			user_agent.Platform.Version,
			user_agent.Platform.Arch,
			user_agent.Engine.Name,
			user_agent.Engine.Version,
			user_agent.Browsers[0].Name,
			user_agent.Browsers[0].Version,
			user_agent.Browsers[1].Name,
			user_agent.Browsers[1].Version,
			user_agent.Browsers[2].Name,
			user_agent.Browsers[2].Version,
		)

	}

	return result

}

func (user_agent *UserAgent) ClientHintString() string {

	result := ""

	if strings.HasSuffix(user_agent.Name, "-chrome") {

		version := strings.TrimSpace(user_agent.Browsers[0].Version)
		major   := strings.TrimSpace(version[0:strings.Index(version, ".")])
		random  := rand.IntN(99) + 1

		result = strings.Join([]string{
			fmt.Sprintf("\"Google Chrome\";v=\"%s\"", major),
			fmt.Sprintf("\"Chromium\";v=\"%s\"", major),
			fmt.Sprintf("\"Not/A)Brand\";v=\"%d\"", random),
		}, ", ")

	} else if strings.HasSuffix(user_agent.Name, "-edge") {

		version := strings.TrimSpace(user_agent.Browsers[0].Version)
		major   := strings.TrimSpace(version[0:strings.Index(version, ".")])
		random  := rand.IntN(99) + 1

		result = strings.Join([]string{
			fmt.Sprintf("\"Chromium\";v=\"%s\"", major),
			fmt.Sprintf("\"Not_A Brand\";v=\"%d\"", random),
			fmt.Sprintf("\"Microsoft Edge\";v=\"%s\"", major),
		}, ", ")

	}

	return result

}

func (user_agent *UserAgent) Headers() []Header {

	result := make([]Header, 0)

	result = append(result, Header{
		Header: "User-Agent",
		Operation: "set",
		Value: user_agent.String(),
	})

	result = append(result, Header{
		Header: "Sec-CH-UA",
		Operation: "set",
		Value: user_agent.ClientHintString(),
	})

	result = append(result, Header{
		Header: "Sec-CH-UA-Mobile",
		Operation: "set",
		Value: "?0",
	})

	result = append(result, Header{
		Header: "Sec-CH-UA-Platform",
		Operation: "set",
		Value: "\"Windows\"",
	})

	if strings.HasPrefix(user_agent.Name, "win10-") {

		result = append(result, Header{
			Header: "Sec-CH-UA-Platform-Version",
			Operation: "set",
			Value: "\"10.0.0\"",
		})

	} else if strings.HasPrefix(user_agent.Name, "win11-") {

		result = append(result, Header{
			Header: "Sec-CH-UA-Platform-Version",
			Operation: "set",
			Value: "\"15.0.0\"",
		})

	}

	result = append(result, Header{
		Header: "Sec-CH-UA-Arch",
		Operation: "set",
		Value: "\"x86\"",
	})

	result = append(result, Header{
		Header: "Sec-CH-UA-Bitness",
		Operation: "set",
		Value: "\"64\"",
	})

	result = append(result, Header{
		Header: "Sec-CH-UA-Model",
		Operation: "set",
		Value: "\"\"",
	})

	return result

}
