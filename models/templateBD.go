package models

type TemplateBD struct {
	Ops   string                 `json:",omitempty"`
	Path  string                 `json:",omitempty"`
	Value map[string]interface{} `json:",omitempty"`
}

func NewTemplateBD(ops, path, name, displayName, layer2Unicast, unkMcastAct, multiDstPktAct, v6unkMcastAct, vmac string, intersiteBumTrafficAllow, optimizeWanBandwidth, l2Stretch, l3MCast, arpFlood, unicastRouting bool, vrfRef, dhcpLabel map[string]interface{}) *TemplateBD {
	var bdMap map[string]interface{}
	bdMap = map[string]interface{}{
		"name":                     name,
		"displayName":              displayName,
		"l2UnknownUnicast":         layer2Unicast,
		"unkMcastAct":              unkMcastAct,
		"multiDstPktAct":           multiDstPktAct,
		"v6unkMcastAct":            v6unkMcastAct,
		"vmac":                     vmac,
		"arpFlood":                 arpFlood,
		"unicastRouting":           unicastRouting,
		"intersiteBumTrafficAllow": intersiteBumTrafficAllow,
		"optimizeWanBandwidth":     optimizeWanBandwidth,
		"l2Stretch":                l2Stretch,
		"l3MCast":                  l3MCast,
		"vrfRef":                   vrfRef,
		"dhcpLabel":                dhcpLabel,
	}

	if bdMap["l2UnknownUnicast"] == "" {
		bdMap["l2UnknownUnicast"] = "flood"
	}

	if bdMap["dhcpLabel"] == nil {
		delete(bdMap, "dhcpLabel")
	}

	return &TemplateBD{
		Ops:   ops,
		Path:  path,
		Value: bdMap,
	}

}

func (bdAttributes *TemplateBD) ToMap() (map[string]interface{}, error) {
	bdAttributesMap := make(map[string]interface{})
	A(bdAttributesMap, "op", bdAttributes.Ops)
	A(bdAttributesMap, "path", bdAttributes.Path)
	if bdAttributes.Value != nil {
		A(bdAttributesMap, "value", bdAttributes.Value)
	}

	return bdAttributesMap, nil
}
