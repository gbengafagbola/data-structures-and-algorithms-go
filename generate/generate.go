package generate

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	outputFileName string
	maxDepth       int
	ignoredDirs    = map[string]bool{
		".git":        true,
		"node_modules": true,
	}
)

// TreeNode represents a node in the tree structure.
type TreeNode struct {
	Name     string
	Children []*TreeNode
	IsDir    bool
}

// GenerateTree walks through the directory and builds the tree structure.
func GenerateTree(root string, currentDepth int) (*TreeNode, error) {
	info, err := os.Stat(root)
	if err != nil {
		return nil, err
	}

	rootNode := &TreeNode{
		Name:  info.Name(),
		IsDir: info.IsDir(),
	}

	if !info.IsDir() || (maxDepth > 0 && currentDepth >= maxDepth) {
		return rootNode, nil
	}

	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if ignoredDirs[entry.Name()] {
			continue
		}
		childPath := filepath.Join(root, entry.Name())
		childNode, err := GenerateTree(childPath, currentDepth+1)
		if err != nil {
			return nil, err
		}
		rootNode.Children = append(rootNode.Children, childNode)
	}

	return rootNode, nil
}

// RenderTree recursively generates the tree as a string.
func RenderTree(node *TreeNode, prefix string) string {
	var sb strings.Builder
	if node.IsDir {
		sb.WriteString(fmt.Sprintf("%s📂 %s\n", prefix, node.Name))
	} else {
		sb.WriteString(fmt.Sprintf("%s📄 %s\n", prefix, node.Name))
	}

	for i, child := range node.Children {
		newPrefix := prefix + "│   "
		if i == len(node.Children)-1 {
			newPrefix = prefix + "    "
		}
		sb.WriteString(RenderTree(child, newPrefix))
	}

	return sb.String()
}

// UpdateReadme writes the tree structure to README.md.
func UpdateReadme(tree string) error {
	readmeContent := fmt.Sprintf("# Codebase Structure\n\n```\n%s```\n", tree)
	return os.WriteFile(outputFileName, []byte(readmeContent), 0644)
}

// Generate is the entry point that handles flag parsing and tree generation.
func Generate() error {
	// Flags for customization
	flag.StringVar(&outputFileName, "output", "README.md", "Output file name")
	flag.IntVar(&maxDepth, "depth", 0, "Maximum depth of the tree (0 for no limit)")
	flag.Parse()

	rootDir := "." // Root directory of the project
	if flag.NArg() > 0 {
		rootDir = flag.Arg(0) // Use the first argument as the root directory
	}

	tree, err := GenerateTree(rootDir, 0)
	if err != nil {
		return fmt.Errorf("Error generating tree: %v", err)
	}

	treeString := RenderTree(tree, "")
	if err := UpdateReadme(treeString); err != nil {
		return fmt.Errorf("Error updating README.md: %v", err)
	}

	fmt.Println("Tree structure successfully written to", outputFileName)
	return nil
}
