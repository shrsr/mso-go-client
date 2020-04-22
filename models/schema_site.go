package models

type SchemaSite struct {
	Ops   string                 `json:",omitempty"`
	Path  string                 `json:",omitempty"`
	Value map[string]interface{} `json:",omitempty"`
}

func NewSchemaSite(ops, siteId, templateName string) *SchemaSite {
	path := "/sites/-"
	siteMap := map[string]interface{}{
		"siteId":          siteId,
		"templateName":    templateName,
		"anps":            []interface{}{},
		"bds":             []interface{}{},
		"contracts":       []interface{}{},
		"externalEpgs":    []interface{}{},
		"intersiteL3outs": []interface{}{},
		"serviceGraphs":   []interface{}{},
		"vrfs":            []interface{}{},
	}

	return &SchemaSite{
		Ops:   ops,
		Path:  path,
		Value: siteMap,
	}

}

func (schemasiteAttributes *SchemaSite) ToMap() (map[string]interface{}, error) {
	schemasiteAttributeMap := make(map[string]interface{})
	A(schemasiteAttributeMap, "op", schemasiteAttributes.Ops)
	A(schemasiteAttributeMap, "path", schemasiteAttributes.Path)
	A(schemasiteAttributeMap, "value", schemasiteAttributes.Value)

	return schemasiteAttributeMap, nil
}
