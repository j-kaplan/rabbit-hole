package rabbithole

import (
	"encoding/json"
	"net/http"
)

type DefinitionsInfo struct {
	RabbitVersion    string                   `json:"rabbit_version"`
	Users            []UserInfo               `json:"users"`
	VHosts           []VhostInfo              `json:"vhosts"`
	Permissions      []PermissionInfo         `json:"permissions"`
	Policies         []Policy                 `json:"policies"`
	Queues           []QueueInfo              `json:"queues"`
	Exchanges        []ExchangeInfo           `json:"exchanges"`
	Bindings         []BindingInfo            `json:"bindings"`
	Parameters       []map[string]interface{} `json:"parameters"`
	GlobalParameters []map[string]interface{} `json:"global_parameters"`
}

//
// GET /api/definitions/
//

func (c *Client) GetDefinitions() (rec DefinitionsInfo, err error) {
	req, err := newGETRequest(c, "definitions")
	if err != nil {
		return DefinitionsInfo{}, err
	}

	if err = executeAndParseRequest(c, req, &rec); err != nil {
		return DefinitionsInfo{}, err
	}

	return rec, nil
}

//
// GET /api/definitions/{vhost}
//

func (c *Client) GetDefinitionsIn(vhost string) (rec DefinitionsInfo, err error) {
	req, err := newGETRequest(c, "definitions/"+PathEscape(vhost))
	if err != nil {
		return DefinitionsInfo{}, err
	}

	if err = executeAndParseRequest(c, req, &rec); err != nil {
		return DefinitionsInfo{}, err
	}

	return rec, nil
}

//
// POST /api/definitions/
//

// Updates definitions.
func (c *Client) PostDefinitions(info DefinitionsInfo) (res *http.Response, err error) {
	body, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}

	req, err := newRequestWithBody(c, "POST", "definitions", body)
	if err != nil {
		return nil, err
	}

	res, err = executeRequest(c, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

//
// POST /api/definitions/{vhost}
//

// Updates definitions about individual vhost.
func (c *Client) PostDefinitionsIn(vhost string, info DefinitionsInfo) (res *http.Response, err error) {
	body, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}

	req, err := newRequestWithBody(c, "POST", "definitions/"+PathEscape(vhost), body)
	if err != nil {
		return nil, err
	}

	res, err = executeRequest(c, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
