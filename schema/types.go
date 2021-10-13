package schema

import (
	"time"
)

type Schema struct {
	Class     *Class     `json:",omitempty"`
	Function  *Func      `json:",omitempty"`
	Variable  *Variable  `json:",omitempty"`
	Enum      *Enum      `json:",omitempty"`
	Struct    *Struct    `json:",omitempty"`
	TypeAlias *TypeAlias `json:",omitempty"`

	Kind     string
	PullDate time.Time
	Version  int
}

type Identifier struct {
	Name        string `json:",omitempty"`
	Description string `json:",omitempty"`
	Declaration string `json:",omitempty"`

	Frameworks []string `json:",omitempty"`
	Platforms  []string `json:",omitempty"`

	Deprecated bool   `json:",omitempty"`
	TopicURL   string `json:",omitempty"`
}

type Class struct {
	Identifier

	InstanceMethods    []Method   `json:",omitempty"`
	InstanceProperties []Property `json:",omitempty"`

	TypeMethods    []Method   `json:",omitempty"`
	TypeProperties []Property `json:",omitempty"`
}

type DataType struct {
	Name        string     `json:",omitempty"`
	IsPtr       bool       `json:",omitempty"`
	IsPtrPtr    bool       `json:",omitempty"`
	Annotations []string   `json:",omitempty"`
	FuncPtr     *Func      `json:",omitempty"`
	Block       *Func      `json:",omitempty"`
	Params      []DataType `json:",omitempty"`
}

type Func struct {
	Identifier

	Return DataType
	Args   []Arg
}

type Arg struct {
	Name string `json:",omitempty"`
	Type DataType
}

type Variable struct {
	Identifier

	Type  DataType
	Value string `json:",omitempty"`
}

type Enum struct {
	Identifier

	Type  DataType
	Cases []Variable
}

type Struct struct {
	Identifier

	Fields []Variable
}

type TypeAlias struct {
	Identifier

	Type   DataType
	Values []Variable `json:",omitempty"`
}

type Property struct {
	Name        string
	Description string
	Declaration string
	Type        DataType
	Attrs       map[string]interface{}
	IsOutlet    bool   `json:",omitempty"`
	Deprecated  bool   `json:",omitempty"`
	TopicURL    string `json:",omitempty"`
}

type Method struct {
	Name        string
	Description string
	Declaration string
	Return      DataType
	Args        []Arg
	Deprecated  bool   `json:",omitempty"`
	TopicURL    string `json:",omitempty"`
}

type Topic struct {
	Path        string
	Title       string
	Type        string
	Description string
	Declaration string
	Frameworks  []string
	Platforms   []string
	Topics      []Link
	LastFetch   time.Time
	LastVersion int
}

type Link struct {
	Section string
	Name    string
	Path    string
}
