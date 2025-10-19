package test

import (
	"context"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yukimula918/astlint/domain/astree"
	"github.com/yukimula918/astlint/parser"
)

//go:embed testdata/test_treesitter_parser.go
var treesitterParserCode []byte

func Test_Astree_ParseSource(t *testing.T) {
	t.Run("test_astree_parse_source_file", func(t *testing.T) {
		ctx := context.Background()
		source, err := astree.NewSourceFromFile("", "testdata/test_treesitter_parser.go")
		assert.NoError(t, err)
		assert.NotNil(t, source)

		tree, err := parser.ParseTreeSitter(ctx, source)
		assert.NoError(t, err)
		assert.NotNil(t, tree)
	})

	t.Run("test_astree_parse_source_code", func(t *testing.T) {
		ctx := context.Background()
		source := astree.NewSourceFromText(astree.Language_Go, string(treesitterParserCode))
		assert.NotNil(t, source)

		tree, err := parser.ParseTreeSitter(ctx, source)
		assert.NoError(t, err)
		assert.NotNil(t, tree)
	})
}
