package modules

import (
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// NodeDetails Type to store the Response Node Details
type NodeDetails struct {
	IP   string
	PORT uint16
}

type Hashing struct {
	address string
	Result  NodeDetails
}

// Connect Connects to the consistent hashing Address
func (t *Hashing) Connect(address, apiVersion string) (err error) {
	if apiVersion == "v1" {
		t.address = address + "/v1/db/node/ip?hash="
	} else {
		return errors.New("invalid apiVersion")
	}

	return
}

// parseData Method for parsing byte data and store response into NodeDetails format
func (t *NodeDetails) parseData(data []byte) (err error) {
	var new_map map[string]any
	err = json.Unmarshal(data, &new_map)
	if err != nil {
		return err
	}

	result, ok := new_map["result"].(bool)
	if !ok {
		return errors.New("consistent hashing API error")
	}
	if !result {
		return errors.New("API error, Please make sure you are using correct version")
	}

	content, ok := new_map["content"].(map[string]any)
	if !ok {
		return errors.New("content JSON object parsing failed")
	}

	ip, ok := content["ip"].(string)
	if !ok {
		return errors.New("invalid IP format in JSON response")
	}
	portFloat, ok := content["port"].(float64)
	if !ok {
		return errors.New("invalid port format in JSON response")
	}
	t.IP = ip
	t.PORT = uint16(portFloat)

	return
}

// GetNode64 Function returns the Node details according to the input hash
func (t *Hashing) GetNode64(hash uint64) (err error) {
	resp, err := FetchData(t.address + strconv.FormatUint(hash, 10))
	if err != nil {
		return err
	}
	if resp == nil {
		return errors.New("empty response received from server")
	}
	defer resp.Close()

	body, err := io.ReadAll(resp)
	if err != nil {
		return err
	}

	err = t.Result.parseData(body)
	if err != nil {
		return err
	}

	return
}
