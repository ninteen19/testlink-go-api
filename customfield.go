package testlink

type CustomField struct {
	Id                     int    `json:"id"`
	Name                   string `json:"name"`
	Label                  string `json:"label"`
	Type                   int    `json:"type"`
	PossibleValues         string `json:"possibleValues"`
	DefaultValue           string `json:"defaultValue"`
	ValidRegexp            string `json:"validRegexp"`
	LengthMin              int    `json:"lengthMin"`
	LengthMax              int    `json:"lengthMax"`
	ShowOnDesign           bool   `json:"showOnDesign"`
	EnableOnDesign         bool   `json:"enableOnDesign"`
	ShowEOnExecution       bool   `json:"showEOnExecution"`
	EnableOnExecution      bool   `json:"enableOnExecution"`
	ShowOnTestPlanDesign   bool   `json:"showOnTestPlanDesign"`
	EnableOnTestPlanDesign bool   `json:"enableOnTestPlanDesign"`
	DisplayOrder           int    `json:"displayOrder"`
	Location               int    `json:"location"`
	Value                  string `json:"value"`
}
