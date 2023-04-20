package types

import (
	"database/sql/driver"
	"encoding/json"
)

type (
	// PolicyRule is for describing a policy.
	PolicyRule struct {
		Effect     PolicyEffect `json:"effect"`
		Actions    []string     `json:"actions"`
		Resources  []string     `json:"resources,omitempty"`
		Conditions []string     `json:"conditions,omitempty"`
	}
)

func (t *PolicyRule) Scan(v interface{}) (err error) {
	switch v := v.(type) {
	case string:
		err = json.Unmarshal([]byte(v), t)
	case []byte:
		err = json.Unmarshal(v, t)
	}
	return
}

func (t *PolicyRule) Value() (driver.Value, error) {
	return json.Marshal(t)
}
