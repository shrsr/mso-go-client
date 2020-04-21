package models

<<<<<<< HEAD
type SchemaSite struct {
	Ops   string                 `json:",omitempty"`
	Path  string                 `json:",omitempty"`
	Value map[string]interface{} `json:",omitempty"`
}

func NewSchemaSite(ops, path, siteId, templateName string) *SchemaSite {
	var siteMap map[string]interface{}
	if ops != "remove" {
		siteMap = map[string]interface{}{
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
	} else {
		siteMap = nil
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
	if schemasiteAttributes.Value != nil {
		A(schemasiteAttributeMap, "value", schemasiteAttributes.Value)
	}

	return schemasiteAttributeMap, nil
}
=======

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/mso-go-client/container"
)


  
type SchemaSiteAttributes struct {
	
	
	Schema string `json:",omitempty"`
	
	
    
	Site       string `json:",omitempty"`
	
    
	Template       string `json:",omitempty"`
	
    
}
   

func NewSchemaSite(schemasiteattr SchemaSiteAttributes) *SchemaSiteAttributes {
        
		SchemaSiteAttributes := schemasiteattr
		return &SchemaSiteAttributes
         
	}


func (schemasite *Tenant) ToMap() (map[string]interface{}, error) {
// 	fvTenantMap, err := fvTenant.BaseAttributes.ToMap()
// 	if err != nil {
// 		return nil, err
// 	}

	
	
// 	A(fvTenantMap, "name",fvTenant.Name)
	
	
    
// 	A(fvTenantMap, "annotation",fvTenant.Annotation)
	
    
// 	A(fvTenantMap, "nameAlias",fvTenant.NameAlias)
	
    
	

// 	return fvTenantMap, err
// }

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
		
		
// 			Name : G(TenantCont, "name"),
		
		
        
// 	        Annotation : G(TenantCont, "annotation"),
		
        
// 	        NameAlias : G(TenantCont, "nameAlias"),
		
        		
//         },
        
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
>>>>>>> TPM-8 incomplete schema_site model file
