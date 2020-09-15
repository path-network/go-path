package path

type Filters struct {
	Filters []Filter `json:"filters"`
}

type Filter struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type FiltersOptions struct {
	Filters []FilterOption `json:"filters"`
}

type FilterOption struct {
	Name        string              `json:"name"`
	Label       string              `json:"label"`
	Description string              `json:"description"`
	Fields      []FilterOptionField `json:"fields"`
}

type FilterOptionField struct {
	Name        string      `json:"name"`
	Label       string      `json:"label"`
	Description string      `json:"description"`
	Value       interface{} `json:"value"`
}

type FilterOptionFieldArray struct {
	Type      string `json:"type"`
	Subtype   string `json:"subtype"`
	MinLength int    `json:"min_length"`
	MaxLength int    `json:"max_length"`
}

type FilterOptionFieldBool struct {
	Type string `json:"type"`
}

type FilterOptionFieldCIDR struct {
	Type string `json:"type"`
}

type FilterOptionFieldIP struct {
	Type string `json:"type"`
}

type FilterOptionFieldInteger struct {
	Type string `json:"type"`
	Min  int    `json:"min"`
	Max  int    `json:"max"`
}

type FilterOptionFieldPortRange struct {
	Type string `json:"type"`
}

type FilterOptionFieldSelect struct {
	Type string `json:"type"`
}

type FilterOptionFieldSelectOption struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type FilterOptionFieldString struct {
	Type      string `json:"type"`
	MinLength int    `json:"min_length"`
	MaxLength int    `json:"max_length"`
}
