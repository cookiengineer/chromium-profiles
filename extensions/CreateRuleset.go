package extensions

import "encoding/json"
import "fmt"

func CreateRuleset(variant string) ([]byte, error) {

	user_agent, ok := UserAgents[variant]

	if ok == true {

		request_headers    := make([]Header, 0)
		user_agent_headers := user_agent.Headers()

		for _, header := range user_agent_headers {
			request_headers = append(request_headers, header)
		}

		ruleset := make([]Rule, 0)
		ruleset = append(ruleset, Rule{
			Id: 1,
			Priority: 1,
			Action: RuleAction{
				Type: "modifyHeaders",
				RequestHeaders: request_headers,
			},
			Condition: RuleCondition{
				URLFilter: "*",
				ResourceTypes: []string{
					"main_frame",
					"sub_frame",
					"xmlhttprequest",
					"script",
					"image",
					"font",
					"stylesheet",
					"other",
				},
			},
		})

		bytes, err := json.MarshalIndent(ruleset, "", "\t")

		if err == nil {
			return bytes, nil
		} else {
			return []byte{}, err
		}

	} else {
		return []byte{}, fmt.Errorf("Invalid variant \"%s\"", variant)
	}

}
