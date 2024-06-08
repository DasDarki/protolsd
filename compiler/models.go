package compiler

type ScriptPackage struct {
	Path              string                    `json:"path"`
	Name              string                    `json:"name"`
	AreImportsPrivate bool                      `json:"areImportsPrivate"`
	Parent            *ScriptPackage            `json:"parent"`
	Scripts           map[string]*Script        `json:"scripts"`
	Packages          map[string]*ScriptPackage `json:"packages"`
}

type Script struct {
	Package           *ScriptPackage      `json:"-"`
	Path              string              `json:"path"`
	AreImportsPrivate bool                `json:"areImportsPrivate"`
	UsedEmptyResponse bool                `json:"usedEmptyResponse"`
	UsedEmptyRequest  bool                `json:"usedEmptyRequest"`
	Imports           []*ImportStatement  `json:"imports"`
	TypeAliases       map[string]string   `json:"typeAliases"`
	DeclaredTypes     map[string]bool     `json:"declaredTypes"`
	Enums             map[string]*Enum    `json:"enums"`
	Messages          map[string]*Message `json:"messages"`
	Services          map[string]*Service `json:"services"`
	AccessibleScripts []*Script           `json:"-"` // gets resolved during resolve phase
}

type ImportStatement struct {
	Path    string `json:"path"`
	Private bool   `json:"private"`
}

type Enum struct {
	Name      string         `json:"name"`
	IsClass   bool           `json:"isClass"`
	Constants map[string]int `json:"constants"`
}

type Message struct {
	Name     string                   `json:"name"`
	Fields   map[string]*messageField `json:"fields"`
	Children map[string]*Message      `json:"children"`
}

type messageField struct {
	Name      string    `json:"name"`
	DataType  *DataType `json:"DataType"`
	Order     int       `json:"order"`
	Modifier  *string   `json:"modifier"`
	Overwrite bool      `json:"overwrite"`
}

type DataType struct {
	Text       string `json:"identifier"`
	IsOptional bool   `json:"isOptional"`
	IsArray    bool   `json:"isArray"`
}

type Service struct {
	Name string                `json:"name"`
	Rpcs map[string]*RpcMethod `json:"rpcs"`
}

type RpcMethod struct {
	Name    string            `json:"name"`
	Returns *RpcMethodMessage `json:"returns"` // if nil, empty response should be used
	Input   *RpcMethodMessage `json:"input"`   // if nil, empty request should be used
}

type RpcMethodMessage struct {
	NeedsAutoWire bool     `json:"needsAutoWire"`
	Name          *string  `json:"name"`
	Message       *Message `json:"Message"`
}
