package models


import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/mso-go-client/container"
)




type SchemaAttributes struct {
	
	
	Schema string `json:",omitempty"`
	
	
    
	Templates   [] `json:",omitempty"`
	
    
	Sites      [] `json:",omitempty"`
	
    
}



func NewSchemacontainer(schemaAttr SchemaAttributes) *SchemaAttributes {
	
        
		SchemacontainerAttributes:= schemaAttr,
		return &SchemacontainerAttributes
         
	}

   



func (fvTenant *Tenant) ToMap() (map[string]interface{}, error) {
	fvTenantMap, err := fvTenant.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	
	
	A(fvTenantMap, "name",fvTenant.Name)
	
	
    
	A(fvTenantMap, "annotation",fvTenant.Annotation)
	
    
	A(fvTenantMap, "nameAlias",fvTenant.NameAlias)
	
    
	

	return fvTenantMap, err
}

func TenantFromContainerList(cont *container.Container, index int) *Tenant {

	TenantCont := cont.S("imdata").Index(index).S(FvtenantClassName, "attributes")
	return &Tenant{
		BaseAttributes{
			DistinguishedName: G(TenantCont, "dn"),
			Description:       G(TenantCont, "descr"),
			Status:            G(TenantCont, "status"),
			ClassName:         FvtenantClassName,
			Rn:                G(TenantCont, "rn"),
		},
        
		TenantAttributes{
		
		
			Name : G(TenantCont, "name"),
		
		
        
	        Annotation : G(TenantCont, "annotation"),
		
        
	        NameAlias : G(TenantCont, "nameAlias"),
		
        		
        },
        
	}
}

func TenantFromContainer(cont *container.Container) *Tenant {

	return TenantFromContainerList(cont, 0)
}

func TenantListFromContainer(cont *container.Container) []*Tenant {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*Tenant, length)

	for i := 0; i < length; i++ {

		arr[i] = TenantFromContainerList(cont, i)
	}

	return arr
}