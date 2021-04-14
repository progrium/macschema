package declparse

import (
	"fmt"
	"sort"
	"strings"
)

type PropAttr int

const (
	PropAttrClass PropAttr = iota
	PropAttrWeak
	PropAttrStrong
	PropAttrReadonly
	PropAttrReadwrite
	PropAttrNonatomic
	PropAttrAtomic
	PropAttrCopy
	PropAttrAssign
	PropAttrRetain
	PropAttrNullable
	PropAttrNonnull
	PropAttrGetter
	PropAttrSetter
)

var propAttrs = map[PropAttr]string{
	PropAttrClass:     "class",
	PropAttrWeak:      "weak",
	PropAttrStrong:    "strong",
	PropAttrReadonly:  "readonly",
	PropAttrReadwrite: "readwrite",
	PropAttrNonatomic: "nonatomic",
	PropAttrAtomic:    "atomic",
	PropAttrCopy:      "copy",
	PropAttrAssign:    "assign",
	PropAttrRetain:    "retain",
	PropAttrNullable:  "nullable",
	PropAttrNonnull:   "nonnull",
	PropAttrGetter:    "getter",
	PropAttrSetter:    "setter",
}

func (attr PropAttr) String() string {
	return propAttrs[attr]
}

func PropAttrs() []PropAttr {
	keys := make([]int, len(propAttrs))
	props := make([]PropAttr, len(propAttrs))
	i := 0
	for k := range propAttrs {
		keys[i] = int(k)
		i++
	}
	sort.Ints(keys)
	i = 0
	for _, prop := range keys {
		props[i] = PropAttr(prop)
		i++
	}
	return props
}

type TypeAnnotation int

const (
	TypeAnnotConst TypeAnnotation = iota
	TypeAnnotOneway
	TypeAnnotSigned
	TypeAnnotUnsigned
	TypeAnnotKindOf

	annonatedType

	TypeAnnotNullable
	TypeAnnotNonnull
	TypeAnnotNullUnspecified
)

var typeAnnots = map[TypeAnnotation]string{
	TypeAnnotConst:           "const %s",
	TypeAnnotOneway:          "oneway %s",
	TypeAnnotSigned:          "signed %s",
	TypeAnnotUnsigned:        "unsigned %s",
	TypeAnnotKindOf:          "__kindof %s",
	TypeAnnotNullable:        "%s _Nullable",
	TypeAnnotNonnull:         "%s _Nonnull",
	TypeAnnotNullUnspecified: "%s _Null_unspecified",
}

func (annot TypeAnnotation) Format() string {
	return typeAnnots[annot]
}

func (annot TypeAnnotation) String() string {
	return strings.Trim(fmt.Sprintf(typeAnnots[annot], ""), " ")
}

func TypeAnnotations() []TypeAnnotation {
	keys := make([]int, len(typeAnnots))
	annots := make([]TypeAnnotation, len(typeAnnots))
	i := 0
	for k := range typeAnnots {
		keys[i] = int(k)
		i++
	}
	sort.Ints(keys)
	i = 0
	for _, prop := range keys {
		annots[i] = TypeAnnotation(prop)
		i++
	}
	return annots
}

func isTypeAnnot(s string, beforeType bool) (TypeAnnotation, bool) {
	for _, annot := range TypeAnnotations() {
		if beforeType && annot > annonatedType {
			continue
		}
		if !beforeType && annot < annonatedType {
			continue
		}
		if s == annot.String() {
			return annot, true
		}
	}
	return annonatedType, false
}