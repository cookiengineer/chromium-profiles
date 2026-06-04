package extensions

type Rule struct {
	Id        int        `json:"id"`
	Priority  int        `json:"priority"`
	Action    RuleAction `json:"action"`
	Condition RuleCondition `json:"condition"`
}

type RuleAction struct {
	Type           string   `json:"type"`
	RequestHeaders []Header `json:"requestHeaders"`
}

type RuleCondition struct {
	URLFilter     string   `json:"urlFilter"`
	ResourceTypes []string `json:"resourceTypes"`
}
