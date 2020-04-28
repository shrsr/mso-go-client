package models

type TemplateBD struct {
	Ops   string                 `json:",omitempty"`
	Path  string                 `json:",omitempty"`
	Value map[string]interface{} `json:",omitempty"`
}

func NewTemplateBD(ops, path, name, displayName, layer2Unicast string, intersiteBumTrafficAllow, optimizeWanBandwidth, l2Stretch, l3MCast bool, vrfRef, dhcpLabel map[string]interface{}) *TemplateBD {
	var bdMap map[string]interface{}
	bdMap = map[string]interface{}{
		"name":                     name,
		"displayName":              displayName,
		"l2UnknownUnicast":         layer2Unicast,
		"intersiteBumTrafficAllow": intersiteBumTrafficAllow,
		"optimizeWanBandwidth":     optimizeWanBandwidth,
		"l2Stretch":                l2Stretch,
		"l3MCast":                  l3MCast,
		"vrfRef":                   vrfRef,
		"dhcpLabel":                dhcpLabel,
	}

	return &TemplateBD{
		Ops:   ops,
		Path:  path,
		Value: bdMap,
	}

}
