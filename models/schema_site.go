package models

type SchemaSiteAttributes struct {
	Schema string `json:",omitempty"`

	Site string `json:",omitempty"`

	Template string `json:",omitempty"`
}

func NewSchemaSite(schemasiteattr SchemaSiteAttributes) *SchemaSiteAttributes {

	SchemaSiteAttributes := schemasiteattr
	return &SchemaSiteAttributes

}

func (schemasiteAttributes *SchemaSiteAttributes) ToMap() (map[string]interface{}, error) {
	schemasiteAttributeMap := make(map[string]interface{})
	A(schemasiteAttributeMap, "schema", schemasiteAttributes.Schema)
	A(schemasiteAttributeMap, "templates", schemasiteAttributes.Template)
	A(schemasiteAttributeMap, "sites", schemasiteAttributes.Site)

	return schemasiteAttributeMap, nil
}

// func TenantFromContainerList(cont *container.Container, index int) *Tenant {

// 	TenantCont := cont.S("imdata").Index(index).S(FvtenantClassName, "attributes")
// 	return &Tenant{
// 		BaseAttributes{
// 			DistinguishedName: G(TenantCont, "dn"),
// 			Description:       G(TenantCont, "descr"),
// 			Status:            G(TenantCont, "status"),
// 			ClassName:         FvtenantClassName,
// 			Rn:                G(TenantCont, "rn"),
// 		},

// 		TenantAttributes{

// 			Name: G(TenantCont, "name"),

// 			Annotation: G(TenantCont, "annotation"),

// 			NameAlias: G(TenantCont, "nameAlias"),
// 		},
// 	}
// }

// func TenantFromContainer(cont *container.Container) *Tenant {

// 	return TenantFromContainerList(cont, 0)
// }

// func TenantListFromContainer(cont *container.Container) []*Tenant {
// 	length, _ := strconv.Atoi(G(cont, "totalCount"))

// 	arr := make([]*Tenant, length)

// 	for i := 0; i < length; i++ {

// 		arr[i] = TenantFromContainerList(cont, i)
// 	}

// 	return arr
// }
