package schema

import (
	"strings"

	"github.com/progrium/macschema/declparse"
)

func DataTypeFromAst(ti declparse.TypeInfo) (dt DataType) {
	var params []DataType
	for _, param := range ti.Params {
		params = append(params, DataTypeFromAst(param))
	}
	var annots []string
	for annot := range ti.Annots {
		annots = append(annots, strings.ToLower(annot.String()))
	}
	dt = DataType{
		Name:        ti.Name,
		IsPtr:       ti.IsPtr,
		IsPtrPtr:    ti.IsPtrPtr,
		Annotations: annots,
		Params:      params,
	}
	if ti.Func != nil {
		if ti.Func.IsBlock {
			dt.Block = FuncFromAst(ti.Func)
		}
		if ti.Func.IsPtr {
			dt.FuncPtr = FuncFromAst(ti.Func)
		}
	}
	return
}

func FuncFromAst(fn *declparse.FunctionDecl) *Func {
	var args []Arg
	for _, arg := range fn.Args {
		args = append(args, ArgFromAst(arg))
	}
	return &Func{
		Identifier: Identifier{Name: fn.Name},
		Return:     DataTypeFromAst(fn.ReturnType),
		Args:       args,
	}
}

func ArgFromAst(ai declparse.ArgInfo) Arg {
	return Arg{
		Name: ai.Name,
		Type: DataTypeFromAst(ai.Type),
	}
}

func PropertyFromAst(p declparse.PropertyDecl) Property {
	prop := Property{
		Name:     p.Name,
		Type:     DataTypeFromAst(p.Type),
		IsOutlet: p.IsOutlet,
	}
	attrs := make(map[string]interface{})
	for attr, val := range p.Attrs {
		var v interface{}
		if val == "" {
			v = true
		} else {
			v = val
		}
		attrs[attr.String()] = v
	}
	prop.Attrs = attrs
	return prop
}

func MethodFromAst(m declparse.MethodDecl) Method {
	var args []Arg
	for _, arg := range m.Args {
		args = append(args, ArgFromAst(arg))
	}
	return Method{
		Name:   m.Name(),
		Return: DataTypeFromAst(m.ReturnType),
		Args:   args,
	}
}

func VariableFromAst(v declparse.VariableDecl) Variable {
	return Variable{
		Identifier: Identifier{Name: v.Name},
		Value:      v.Value,
		Type:       DataTypeFromAst(v.Type),
	}
}

func EnumFromAst(e declparse.EnumDecl) Enum {
	var cases []Variable
	for _, ecase := range e.Cases {
		cases = append(cases, VariableFromAst(ecase))
	}
	return Enum{
		Identifier: Identifier{Name: e.Name},
		Type:       DataTypeFromAst(e.Type),
		Cases:      cases,
	}
}

func StructFromAst(s declparse.StructDecl) Struct {
	var fields []Variable
	for _, field := range s.Fields {
		fields = append(fields, VariableFromAst(field))
	}
	return Struct{
		Identifier: Identifier{Name: s.Name},
		Fields:     fields,
	}
}
