package models

type SchemaSiteAnpEpgDomain struct {
	Ops   string                 `json:",omitempty"`
	Path  string                 `json:",omitempty"`
	Value map[string]interface{} `json:",omitempty"`
}

func NewSchemaSiteAnpEpgDomain(ops, path, domainType, dn, deploymentImmediacy, resolutionImmediacy, microSegVlanType, portEncapVlanType, vlanEncapMode, switchingMode, switchType, enhancedLagpolicyName, enhancedLagpolicyDn string, allowMicroSegmentation bool, microSegVlan, portEncapVlan int) *SchemaSiteAnpEpgDomain {
	var siteAnpEpgDomainMap map[string]interface{}
	siteAnpEpgDomainMap = map[string]interface{}{
		"domainType":             domainType,
		"dn":                     dn,
		"deploymentImmediacy":    deploymentImmediacy,
		"resolutionImmediacy":    resolutionImmediacy,
		"microSegVlanType":       microSegVlanType,
		"portEncapVlanType":      portEncapVlanType,
		"vlanEncapMode":          vlanEncapMode,
		"switchingMode":          switchingMode,
		"switchType":             switchType,
		"enhancedLagpolicyName":  enhancedLagpolicyName,
		"enhancedLagpolicyDn":    enhancedLagpolicyDn,
		"allowMicroSegmentation": allowMicroSegmentation,
		"microSegVlan":           microSegVlan,
		"portEncapVlan":          portEncapVlan,
	}

	return &SchemaSiteAnpEpgDomain{
		Ops:   ops,
		Path:  path,
		Value: siteAnpEpgDomainMap,
	}

}

func (siteAnpEpgDomainAttributes *SchemaSiteAnpEpgDomain) ToMap() (map[string]interface{}, error) {
	siteAnpEpgDomainAttributesMap := make(map[string]interface{})
	A(siteAnpEpgDomainAttributesMap, "op", siteAnpEpgDomainAttributes.Ops)
	A(siteAnpEpgDomainAttributesMap, "path", siteAnpEpgDomainAttributes.Path)
	if siteAnpEpgDomainAttributes.Value != nil {
		A(siteAnpEpgDomainAttributesMap, "value", siteAnpEpgDomainAttributes.Value)
	}

	return siteAnpEpgDomainAttributesMap, nil
}
