package types

type ClientPreference struct {
	AppCode string                  `json:"appCode"`
	Values  []ClientPreferenceValue `json:"values"`
}

type ClientPreferenceValue struct {
	// 前端自定义key规则，通过key规则区分不同业务场景
	Key       string `json:"key"`
	Value     string `json:"value,omitempty"`
	Extension string `json:"extension,omitempty"`
}
