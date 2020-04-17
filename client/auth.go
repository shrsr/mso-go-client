package client

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Auth struct {
	Token         string
	Expiry        time.Time
	apicCreatedAt time.Time
	realCreatedAt time.Time
	offset        int64
}

func (au *Auth) IsValid() bool {
	if au.Token != "" && au.Expiry.Unix() > au.estimateExpireTime() {
		return true
	}
	return false
}

func (t *Auth) CalculateExpiry(willExpire int64) {
	t.Expiry = time.Unix((t.apicCreatedAt.Unix() + willExpire), 0)
}

func (t *Auth) CaclulateOffset() {
	t.offset = t.apicCreatedAt.Unix() - t.realCreatedAt.Unix()
}

func (t *Auth) estimateExpireTime() int64 {
	return time.Now().Unix() + t.offset
}

func (client *Client) InjectAuthenticationHeader(req *http.Request, path string) (*http.Request, error) {
	log.Printf("[DEBUG] Begin Injection")
	if client.password != "" || client.AuthToken == nil || !client.AuthToken.IsValid() {
		fmt.Println(client)
		err := client.Authenticate()
		fmt.Println(client)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.AuthToken.Token))

		return req, nil
	} else {

		return req, fmt.Errorf("Password is must.")
	}

	return req, nil
}
