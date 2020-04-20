package client

import (
	"errors"
	"fmt"
	"log"

	"github.com/ciscoecosystem/mso-go-client/container"
	"github.com/ciscoecosystem/mso-go-client/models"
)

func (c *Client) GetViaURL(url string) (*container.Container, error) {
	req, err := c.MakeRestRequest("GET", url, nil, true)

	if err != nil {
		return nil, err
	}

	obj, _, err := c.Do(req)
	log.Printf("Getvia url %+v", obj)
	if err != nil {
		return nil, err
	}

	if obj == nil {
		return nil, errors.New("Empty response body")
	}
	return obj, CheckForErrors(obj, "GET")

}

func (c *Client) Save(url string, obj models.Model) (*container.Container, error) {

	jsonPayload, err := c.PrepareModel(obj)

	if err != nil {
		return nil, err
	}
	req, err := c.MakeRestRequest("POST", url, jsonPayload, true)
	if err != nil {
		return nil, err
	}

	cont, _, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	return cont, CheckForErrors(cont, "POST")
}

// CheckForErrors parses the response and checks of there is an error attribute in the response
func CheckForErrors(cont *container.Container, method string) error {

	if cont.Exists("code") && cont.Exists("message") {

		return errors.New(fmt.Sprintf("%s", cont.S("message")))
	} else {
		return nil
	}

	return nil
}

func (c *Client) DeletebyId(url string) error {

	req, err := c.MakeRestRequest("DELETE", url, nil, true)
	if err != nil {
		return err
	}

	_, _, err1 := c.Do(req)
	if err1 != nil {
		return err1
	}
	return nil
}

func (c *Client) PatchbyID(url string, obj models.Model) (*container.Container, error) {

	jsonPayload, err := c.PrepareModel(obj)

	if err != nil {
		return nil, err
	}
	req, err := c.MakeRestRequest("PATCH", url, jsonPayload, true)
	if err != nil {
		return nil, err
	}

	cont, _, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	return cont, CheckForErrors(cont, "PATCH")
}

func (c *Client) PrepareModel(obj models.Model) (*container.Container, error) {
	con, err := obj.ToMap()
	if err != nil {
		return nil, err
	}

	payload := &container.Container{}
	if err != nil {
		return nil, err
	}

	for key, value := range con {
		payload.Set(value, key)
	}
	return payload, nil
}
