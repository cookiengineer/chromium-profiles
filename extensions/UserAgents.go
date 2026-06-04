package extensions

import "encoding/json"
import _ "embed"

var UserAgents map[string]*UserAgent

//go:embed UserAgents.json
var embedded_user_agents []byte

func init() {

	UserAgents = make(map[string]*UserAgent)

	tmp := make([]UserAgent, 0)
	err := json.Unmarshal(embedded_user_agents, &tmp)

	if err == nil {

		for _, user_agent := range tmp {
			UserAgents[user_agent.Name] = &user_agent
		}

	}

}
