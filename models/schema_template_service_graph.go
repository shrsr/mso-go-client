package models

type TemplateServiceGraph struct {
	Ops   string `json:",omitempty"`
	Path  string `json:",omitempty"`
	Value map[string]interface{}
}

func NewTemplateServiceGraph(ops, path string, graphRef map[string]interface{}) *TemplateServiceGraph {

	return &TemplateServiceGraph{
		Ops:   ops,
		Path:  path,
		Value: graphRef,
	}

}

func (graphAttributes *TemplateServiceGraph) ToMap() (map[string]interface{}, error) {
	graphAttributesMap := make(map[string]interface{})
	A(graphAttributesMap, "op", graphAttributes.Ops)
	A(graphAttributesMap, "path", graphAttributes.Path)
	if graphAttributes.Value != nil {
		A(graphAttributesMap, "value", graphAttributes.Value)
	}

	return graphAttributesMap, nil
}
