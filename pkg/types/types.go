package types

type Event struct {
	Id       string                 `json:"id"`
	TenantId string                 `json:"tenant_id"`
	Type     string                 `json:"type"`
	Payload  map[string]interface{} `json:"payload"`
}

type Entity struct {
	Predicates map[string]bool
}

type Impact struct {
	Predicates map[string]bool
}

type Changes struct {
	Predicates map[string]bool
}
