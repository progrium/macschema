package declparse

import (
	"github.com/progrium/macschema/declparse/keywords"
	"github.com/progrium/macschema/lexer"
)

func parseInterface(p *Parser) (next stateFn, node Node, err error) {
	decl := &InterfaceDecl{}

	if err := p.expectToken(keywords.INTERFACE); err != nil {
		return nil, nil, err
	}

	decl.Name, err = p.expectIdent()
	if err != nil {
		return nil, nil, err
	}

	if tok, _, _ := p.tb.Scan(); tok == lexer.COLON {
		if decl.SuperName, err = p.expectIdent(); err != nil {
			return nil, nil, err
		}
	} else {
		p.tb.Unscan()
	}

	return nil, decl, nil
}
