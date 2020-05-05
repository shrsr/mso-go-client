package models

type TemplateContractFilter struct {
	Ops   string                 `json:",omitempty"`
	Path  string                 `json:",omitempty"`
	Value map[string]interface{} `json:",omitempty"`
}

func NewTemplateContractFilter(ops, path, name, displayName, scope, filterType string, filterRelationships,filterRelationshipsProviderToConsumer,filterRelationshipsConsumerToProvider []interface{}) *TemplateContractFilter {
	var contractMap map[string]interface{}
	if ops !="remove" {
	contractMap = map[string]interface{}{
		"name":                                  name,
		"displayName":                           displayName,
		"scope":                                 scope,
		"filterType":                            filterType,
		"filterRelationships":                   filterRelationships,
		"filterRelationshipsProviderToConsumer": filterRelationshipsProviderToConsumer,
		"filterRelationshipsConsumerToProvider": filterRelationshipsConsumerToProvider,
	}

	if contractMap["filterType"] == nil {
		contractMap["filterType"] = "bothway"
	}

	if contractMap["scope"] == "" {
		contractMap["scope"] = "context"
	}
	}else{
		contractMap = nil
	}
	return &TemplateContractFilter{
		Ops:   ops,
		Path:  path,
		Value: contractMap,
	}

}

func (FilterAttributes *TemplateContractFilter) ToMap() (map[string]interface{}, error) {
	FilterAttributesMap := make(map[string]interface{})
	A(FilterAttributesMap, "op", FilterAttributes.Ops)
	A(FilterAttributesMap, "path", FilterAttributes.Path)
	if FilterAttributes.Value != nil {
		A(FilterAttributesMap, "value", FilterAttributes.Value)
	}

	return FilterAttributesMap, nil
}
