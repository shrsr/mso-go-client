package client

import (
	"github.com/ciscoecosystem/mso-go-client/container"
	"github.com/ciscoecosystem/mso-go-client/models"
)

func (client *Client) CreateDHCPOptionPolicy(obj *models.DHCPOptionPolicy) (*container.Container, error) {
	path := "api/v1/policies/dhcp/option"
	cont, err := client.Save(path, obj)
	if err != nil {
		return nil, err
	}
	return cont, nil
}

func (client *Client) ReadDHCPOptionPolicy(id string) (*container.Container, error) {
	path := "api/v1/policies/dhcp/option/" + id
	cont, err := client.GetViaURL(path)
	if err != nil {
		return nil, err
	}
	return cont, nil
}

func (client *Client) UpdateDHCPOptionPolicy(id string, obj *models.DHCPOptionPolicy) (*container.Container, error) {
	path := "api/v1/policies/dhcp/option/" + id
	cont, err := client.Put(path, obj)
	if err != nil {
		return nil, err
	}
	return cont, nil
}

func (client *Client) DeleteDHCPOptionPolicy(id string) error {
	path := "api/v1/policies/dhcp/option/" + id
	err := client.DeletebyId(path)
	if err != nil {
		return err
	}
	return nil
}
