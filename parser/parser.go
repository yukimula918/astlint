package parser

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/c"
	"github.com/smacker/go-tree-sitter/cpp"
	"github.com/smacker/go-tree-sitter/golang"
	"github.com/smacker/go-tree-sitter/java"
	"github.com/smacker/go-tree-sitter/javascript"
	"github.com/smacker/go-tree-sitter/kotlin"
	"github.com/smacker/go-tree-sitter/python"
	"github.com/smacker/go-tree-sitter/ruby"
	"github.com/smacker/go-tree-sitter/rust"
	"github.com/smacker/go-tree-sitter/scala"
	"github.com/smacker/go-tree-sitter/swift"
	"github.com/smacker/go-tree-sitter/typescript/typescript"

	"github.com/yukimula918/astlint/domain/astree"
)

// Parser defines the interface for parsing code to unified model of AST.
type Parser interface {
	// Parse implements syntax parsing from source code.
	Parse(ctx context.Context, src *astree.Source) (*astree.ASTree, error)
}

// ParseTreeSitter parses the source code to treesitter AST.
func ParseTreeSitter(ctx context.Context, source *astree.Source) (*sitter.Tree, error) {
	if source == nil {
		return nil, fmt.Errorf("cannot parse treesitter AST from nil Source")
	}

	var p = sitter.NewParser()
	switch source.Language {
	case astree.Language_C:
		p.SetLanguage(c.GetLanguage())
	case astree.Language_Cpp:
		p.SetLanguage(cpp.GetLanguage())
	case astree.Language_Java:
		p.SetLanguage(java.GetLanguage())
	case astree.Language_Kotlin:
		p.SetLanguage(kotlin.GetLanguage())
	case astree.Language_Scala:
		p.SetLanguage(scala.GetLanguage())
	case astree.Language_Go:
		p.SetLanguage(golang.GetLanguage())
	case astree.Language_Python:
		p.SetLanguage(python.GetLanguage())
	case astree.Language_Swift:
		p.SetLanguage(swift.GetLanguage())
	case astree.Language_Javascript:
		p.SetLanguage(javascript.GetLanguage())
	case astree.Language_Typescript:
		p.SetLanguage(typescript.GetLanguage())
	case astree.Language_Ruby:
		p.SetLanguage(ruby.GetLanguage())
	case astree.Language_Rust:
		p.SetLanguage(rust.GetLanguage())
	default:
		return nil, fmt.Errorf("unsupported language: %s", source.Language.Name())
	}

	tree, err := p.ParseCtx(ctx, nil, []byte(source.Content))
	if err != nil {
		return nil, errors.Wrapf(err, "parse source code to treesitter AST error")
	} else if tree == nil || tree.RootNode() == nil {
		return nil, fmt.Errorf("parse source code to treesitter AST error: nil tree")
	}
	return tree, nil
}
