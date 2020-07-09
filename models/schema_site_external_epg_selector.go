package models

type SchemaSiteExternalEpgSelector struct {
	Ops   string                 `json:",omitempty"`
	Path  string                 `json:",omitempty"`
	Value map[string]interface{} `json:",omitempty"`
}

type SchemaSiteExternalEpg struct {
	Ops   string                 `json:",omitempty"`
	Path  string                 `json:",omitempty"`
	Value map[string]interface{} `json:",omitempty"`
}

func NewSchemaSiteExternalEpg(ops, path string, epgMap map[string]interface{}) *SchemaSiteExternalEpg {
	return &SchemaSiteExternalEpg{
		Ops:   ops,
		Path:  path,
		Value: epgMap,
	}
}

func NewSchemaSiteExternalEpgSelector(ops, path string, selectorMap map[string]interface{}) *SchemaSiteExternalEpgSelector {
	var temp map[string]interface{}

	if ops != "remove" {
		temp = selectorMap
	} else {
		temp = nil
	}

	return &SchemaSiteExternalEpgSelector{
		Ops:   ops,
		Path:  path,
		Value: temp,
	}
}

func (schemaSiteExternalEpgSelector *SchemaSiteExternalEpgSelector) ToMap() (map[string]interface{}, error) {
	schemaSiteExternalEpgSelectorMap := make(map[string]interface{})

	A(schemaSiteExternalEpgSelectorMap, "op", schemaSiteExternalEpgSelector.Ops)
	A(schemaSiteExternalEpgSelectorMap, "path", schemaSiteExternalEpgSelector.Path)
	if schemaSiteExternalEpgSelector.Value != nil {
		A(schemaSiteExternalEpgSelectorMap, "value", schemaSiteExternalEpgSelector.Value)
	}

	return schemaSiteExternalEpgSelectorMap, nil
}

func (schemaSiteExternalEpg *SchemaSiteExternalEpg) ToMap() (map[string]interface{}, error) {
	schemaSiteExternalEpgMap := make(map[string]interface{})

	A(schemaSiteExternalEpgMap, "op", schemaSiteExternalEpg.Ops)
	A(schemaSiteExternalEpgMap, "path", schemaSiteExternalEpg.Path)
	if schemaSiteExternalEpg.Value != nil {
		A(schemaSiteExternalEpgMap, "value", schemaSiteExternalEpg.Value)
	}

	return schemaSiteExternalEpgMap, nil
}
