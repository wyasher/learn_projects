package spider

type (
	TaskModel struct {
		Property
		Root  string      `json:"rootScript"`
		Rules []RuleModel `json:"rules"`
	}
	RuleModel struct {
		Name      string `json:"name"`
		ParseFunc string `json:"parseScript"`
	}
)
