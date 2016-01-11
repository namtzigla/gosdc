package cloudapi

import (
	"net/http"

	"github.com/joyent/gocommon/client"
	"github.com/joyent/gocommon/errors"
)

// AddMachineTags adds additional tags to the specified machine.
// This API lets you append new tags, not overwrite existing tags.
// See API docs: http://apidocs.joyent.com/cloudapi/#AddMachineTags
func (c *Client) AddMachineTags(machineId string, tags map[string]string) (map[string]string, error) {
	var resp map[string]string
	req := request{
		method:   client.POST,
		url:      makeURL(apiMachines, machineId, apiTags),
		reqValue: tags,
		resp:     &resp,
	}
	if _, err := c.sendRequest(req); err != nil {
		return nil, errors.Newf(err, "failed to add tags for machine with id %s", machineId)
	}
	return resp, nil
}

// ReplaceMachineTags replaces existing tags for the specified machine.
// This API lets you overwrite existing tags, not append to existing tags.
// See API docs: http://apidocs.joyent.com/cloudapi/#ReplaceMachineTags
func (c *Client) ReplaceMachineTags(machineId string, tags map[string]string) (map[string]string, error) {
	var resp map[string]string
	req := request{
		method:   client.PUT,
		url:      makeURL(apiMachines, machineId, apiTags),
		reqValue: tags,
		resp:     &resp,
	}
	if _, err := c.sendRequest(req); err != nil {
		return nil, errors.Newf(err, "failed to replace tags for machine with id %s", machineId)
	}
	return resp, nil
}

// ListMachineTags returns the complete set of tags associated with the specified machine.
// See API docs: http://apidocs.joyent.com/cloudapi/#ListMachineTags
func (c *Client) ListMachineTags(machineId string) (map[string]string, error) {
	var resp map[string]string
	req := request{
		method: client.GET,
		url:    makeURL(apiMachines, machineId, apiTags),
		resp:   &resp,
	}
	if _, err := c.sendRequest(req); err != nil {
		return nil, errors.Newf(err, "failed to get list of tags for machine with id %s", machineId)
	}
	return resp, nil
}

// GetMachineTag returns the value for a single tag on the specified machine.
// See API docs: http://apidocs.joyent.com/cloudapi/#GetMachineTag
func (c *Client) GetMachineTag(machineId, tagKey string) (string, error) {
	var resp string
	requestHeaders := make(http.Header)
	requestHeaders.Set("Accept", "text/plain")
	req := request{
		method:    client.GET,
		url:       makeURL(apiMachines, machineId, apiTags, tagKey),
		resp:      &resp,
		reqHeader: requestHeaders,
	}
	if _, err := c.sendRequest(req); err != nil {
		return "", errors.Newf(err, "failed to get tag %s for machine with id %s", tagKey, machineId)
	}
	return resp, nil
}

// DeleteMachineTag deletes a single tag from the specified machine.
// See API docs: http://apidocs.joyent.com/cloudapi/#DeleteMachineTag
func (c *Client) DeleteMachineTag(machineId, tagKey string) error {
	req := request{
		method:         client.DELETE,
		url:            makeURL(apiMachines, machineId, apiTags, tagKey),
		expectedStatus: http.StatusNoContent,
	}
	if _, err := c.sendRequest(req); err != nil {
		return errors.Newf(err, "failed to delete tag with key %s for machine with id %s", tagKey, machineId)
	}
	return nil
}

// DeleteMachineTags deletes all tags from the specified machine.
// See API docs: http://apidocs.joyent.com/cloudapi/#DeleteMachineTags
func (c *Client) DeleteMachineTags(machineId string) error {
	req := request{
		method:         client.DELETE,
		url:            makeURL(apiMachines, machineId, apiTags),
		expectedStatus: http.StatusNoContent,
	}
	if _, err := c.sendRequest(req); err != nil {
		return errors.Newf(err, "failed to delete tags for machine with id %s", machineId)
	}
	return nil
}
