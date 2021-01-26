// This package implements a simple way to retrieve droplet metadata
// from the relevant digitalocean endpoint.
package do_metadata

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Metadata struct {
	DropletID  int      `json:"droplet_id"`
	Hostname   string   `json:"hostname"`
	VendorData string   `json:"vendor_data"`
	PublicKeys []string `json:"public_keys"`
	Region     string   `json:"region"`
	Interfaces struct {
		Private []struct {
			Ipv4 struct {
				IPAddress string `json:"ip_address"`
				Netmask   string `json:"netmask"`
				Gateway   string `json:"gateway"`
			} `json:"ipv4"`
			Mac  string `json:"mac"`
			Type string `json:"type"`
		} `json:"private"`
		Public []struct {
			Ipv4 struct {
				IPAddress string `json:"ip_address"`
				Netmask   string `json:"netmask"`
				Gateway   string `json:"gateway"`
			} `json:"ipv4"`
			Ipv6 struct {
				IPAddress string `json:"ip_address"`
				Cidr      int    `json:"cidr"`
				Gateway   string `json:"gateway"`
			} `json:"ipv6"`
			Mac  string `json:"mac"`
			Type string `json:"type"`
		} `json:"public"`
	} `json:"interfaces"`
	FloatingIP struct {
		Ipv4 struct {
			Active bool `json:"active"`
		} `json:"ipv4"`
	} `json:"floating_ip"`
	DNS struct {
		Nameservers []string `json:"nameservers"`
	} `json:"dns"`
	Features struct {
		DhcpEnabled bool `json:"dhcp_enabled"`
	} `json:"features"`
}

//RetrieveMetadata retrieves the metadata for this droplet from the droplet metadata endpoint
func RetrieveMetadata(timeout time.Duration) (*Metadata, error) {
	var meta Metadata
	cli := &http.Client{Timeout: timeout}

	resp, err := cli.Get("http://169.254.169.254/metadata/v1.json")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 10 * 1024))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &meta); err != nil {
		return nil, err
	}

	return &meta, nil
}