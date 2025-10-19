package astree

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlekSi/pointer"
	"github.com/pkg/errors"
	"github.com/yukimula918/astlint/util"
)

// language interface

// Name returns the name of programming language used.
func (l Language) Name() string {
	return langToName[l]
}

var langToName = map[Language]string{
	Language_C:          "C",
	Language_Cpp:        "C++",
	Language_Dart:       "Dart",
	Language_Java:       "Java",
	Language_Kotlin:     "Kotlin",
	Language_Scala:      "Scala",
	Language_Go:         "Go",
	Language_Python:     "Python",
	Language_ObjectiveC: "Objective-C",
	Language_Swift:      "Swift",
	Language_Javascript: "Javascript",
	Language_Typescript: "Typescript",
	Language_Ruby:       "Ruby",
	Language_Rust:       "Rust",
}

// GetLanguageByPath returns the language of code
// content based on its file path's extension.
func GetLanguageByPath(path string) Language {
	ext := filepath.Ext(path)
	if lang, ok := extToLang[ext]; ok {
		return lang
	}
	return Language_NoLang
}

var extToLang = map[string]Language{
	".c":     Language_C,
	".h":     Language_C,
	".cpp":   Language_Cpp,
	".hpp":   Language_Cpp,
	".dart":  Language_Dart,
	".java":  Language_Java,
	".kt":    Language_Kotlin,
	".scala": Language_Scala,
	".go":    Language_Go,
	".py":    Language_Python,
	".m":     Language_ObjectiveC,
	".mm":    Language_ObjectiveC,
	".swift": Language_Swift,
	".js":    Language_Javascript,
	".jsx":   Language_Javascript,
	".ts":    Language_Typescript,
	".tsx":   Language_Typescript,
	".rb":    Language_Ruby,
	".rs":    Language_Rust,
}

// GetLanguageByName returns the language by its name.
func GetLanguageByName(name string) Language {
	if lang, ok := nameToLang[name]; ok {
		return lang
	}
	return Language_NoLang
}

var nameToLang = map[string]Language{
	"c":           Language_C,
	"c++":         Language_Cpp,
	"cpp":         Language_Cpp,
	"dart":        Language_Dart,
	"java":        Language_Java,
	"kotlin":      Language_Kotlin,
	"scala":       Language_Scala,
	"go":          Language_Go,
	"golang":      Language_Go,
	"python":      Language_Python,
	"objc":        Language_ObjectiveC,
	"objective-c": Language_ObjectiveC,
	"swift":       Language_Swift,
	"javascript":  Language_Javascript,
	"typescript":  Language_Typescript,
	"ruby":        Language_Ruby,
	"rust":        Language_Rust,
}

// source file model

// NewSourceFromFile returns a Source object that
// reads the code content from file in local.
//
// The param `filePath` need be relative path of
// the file, to the directory `repoDir` as given.
func NewSourceFromFile(repoDir, filePath string) (*Source, error) {
	content, err := os.ReadFile(filepath.Join(repoDir, filePath))
	if err != nil {
		return nil, errors.Wrapf(err, "no such a file: %s", filepath.Join(repoDir, filePath))
	}
	return &Source{
		RepoDir:  pointer.ToString(repoDir),
		FilePath: pointer.ToString(filePath),
		Language: GetLanguageByPath(filePath),
		Content:  string(content),
	}, nil
}

// NewSourceFromText returns a Source object that
// specifies the input code directly without any
// file path given.
func NewSourceFromText(lang Language, content string) *Source {
	return &Source{
		RepoDir:  nil,
		FilePath: nil,
		Language: lang,
		Content:  content,
	}
}

// syntax tree model

// NewASTree returns an empty syntax tree.
func NewASTree(lang Language) *ASTree {
	return &ASTree{
		Language: lang,
		Root:     nil,
		Nodes:    nil,
	}
}

// NewNode creates a new node with specified type, positions and code text.
func (t *ASTree) NewNode(typ ASTNodeType, pos, end *Position, texts ...string) (*ASTNode, error) {
	if t == nil {
		return nil, fmt.Errorf("cannot new ASTNode with nil ASTree")
	} else if pos == nil {
		return nil, fmt.Errorf("cannot new ASTNode with nil position")
	}

	if end == nil {
		end = pos
	}
	var content *string
	if len(texts) > 0 {
		value := strings.Join(texts, " ")
		value = strings.TrimSpace(value)
		if !util.IsBlankString(value) {
			content = pointer.ToString(value)
		}
	}

	node := &ASTNode{
		ID:       uint32(len(t.Nodes)),
		Type:     typ,
		Pos:      pos,
		End:      end,
		Text:     content,
		Parent:   nil,
		Children: nil,
	}
	t.Nodes = append(t.Nodes, node)
	return node, nil
}

// GetParent returns the parent node of the input.
func (t *ASTree) GetParent(node *ASTNode) *ASTNode {
	if t == nil || node == nil || node.Parent == nil {
		return nil
	}
	if node.Parent.Parent < uint32(len(t.Nodes)) {
		return t.Nodes[node.Parent.Parent]
	}
	return nil
}

// GetChild returns the child node from parent by its index.
func (t *ASTree) GetChild(parent *ASTNode, k uint32) *ASTNode {
	if t == nil || parent == nil {
		return nil
	} else if k >= uint32(len(parent.Children)) {
		return nil
	}

	edge := parent.Children[k]
	if edge != nil && edge.Child < uint32(len(t.Nodes)) {
		return t.Nodes[edge.Child]
	}
	return nil
}

// ListChildren returns the list of child nodes in the parent specified
// by the given edge types.
//
// If types is given empty, it returns all the child nodes in the parent.
func (t *ASTree) ListChildren(parent *ASTNode, types ...ASTEdgeType) []*ASTNode {
	if t == nil || parent == nil {
		return nil
	}

	var children []*ASTNode
	for _, edge := range parent.Children {
		if edge == nil {
			continue
		} else if len(types) > 0 && !hasEdgeTypeInList(edge.Type, types) {
			continue
		} else if edge.Child >= uint32(len(t.Nodes)) {
			continue
		}
		children = append(children, t.Nodes[edge.Child])
	}
	return children
}

func hasEdgeTypeInList(typ ASTEdgeType, types []ASTEdgeType) bool {
	for _, t := range types {
		if t == typ {
			return true
		}
	}
	return false
}

// ParentChild returns the parent and child node of the edge in the ASTree.
func (t *ASTree) ParentChild(edge *ASTEdge) (parent, child *ASTNode) {
	if t == nil || edge == nil {
		return nil, nil
	}
	if edge.Parent < uint32(len(t.Nodes)) {
		parent = t.Nodes[edge.Parent]
	}
	if edge.Child < uint32(len(t.Nodes)) {
		child = t.Nodes[edge.Child]
	}
	return parent, child
}

// LinkTo link the node (as parent) to another (as its child) with specified edge type.
func (n *ASTNode) LinkTo(typ ASTEdgeType, child *ASTNode) (*ASTEdge, error) {
	if n == nil {
		return nil, fmt.Errorf("cannot link ASTEdge with nil parent node")
	} else if child == nil {
		return nil, fmt.Errorf("cannot link ASTEdge with nil child node")
	} else if child.Parent != nil {
		return nil, fmt.Errorf("cannot link ASTEdge with child that has been linked to other parent before")
	}

	edge := &ASTEdge{
		Parent: n.ID,
		Child:  child.ID,
		Type:   typ,
	}
	n.Children = append(n.Children, edge)
	child.Parent = edge
	return edge, nil
}

// Content returns the content of the node point to.
func (n *ASTNode) Content(content []byte) *string {
	if n == nil || n.Pos == nil || n.End == nil {
		return nil
	}
	subText := content[n.Pos.Offset:n.Pos.Offset]
	return pointer.ToString(string(subText))
}
