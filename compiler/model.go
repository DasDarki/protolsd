package compiler

type script struct {
	Path              string              `json:"path"`
	AreImportsPrivate bool                `json:"areImportsPrivate"`
	UsedEmptyResponse bool                `json:"usedEmptyResponse"`
	UsedEmptyRequest  bool                `json:"usedEmptyRequest"`
	Imports           []*importStatement  `json:"imports"`
	TypeAliases       map[string]string   `json:"typeAliases"`
	Enums             map[string]*enum    `json:"enums"`
	Messages          map[string]*message `json:"messages"`
	Services          map[string]*service `json:"services"`
}

type importStatement struct {
	Path    string  `json:"path"`
	Private bool    `json:"private"`
	Script  *script `json:"script"`
}

type enum struct {
	Name      string         `json:"name"`
	IsClass   bool           `json:"isClass"`
	Constants map[string]int `json:"constants"`
}

type message struct {
	Name     string                   `json:"name"`
	Fields   map[string]*messageField `json:"fields"`
	Children map[string]*message      `json:"children"`
}

type messageField struct {
	Name     string    `json:"name"`
	DataType *dataType `json:"dataType"`
	Order    int       `json:"order"`
	Modifier *string   `json:"modifier"`
}

type dataType struct {
	Text       string `json:"identifier"`
	IsOptional bool   `json:"isOptional"`
	IsArray    bool   `json:"isArray"`
}

type service struct {
	Name string                `json:"name"`
	Rpcs map[string]*rpcMethod `json:"rpcs"`
}

type rpcMethod struct {
	Name    string            `json:"name"`
	Returns *rpcMethodMessage `json:"returns"` // if nil, empty response should be used
	Input   *rpcMethodMessage `json:"input"`   // if nil, empty request should be used
}

type rpcMethodMessage struct {
	NeedsAutoWire bool     `json:"needsAutoWire"`
	Name          *string  `json:"name"`
	Message       *message `json:"message"`
}
