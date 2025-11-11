package handlers

import (
    "net/http"
    "sort"
    "strings"

    "github.com/gin-gonic/gin"
)

type TreeNode struct {
    name     string
    methods  []string
    children map[string]*TreeNode
    isLeaf   bool
}

func (h *Handler) RootHandler(c *gin.Context) {
    routes := h.Router.Routes()

    // Build tree
    root := &TreeNode{name: ".", children: make(map[string]*TreeNode)}

    for _, route := range routes {
        parts := strings.Split(strings.Trim(route.Path, "/"), "/")
        current := root

        for i, part := range parts {
            if part == "" {
                continue
            }

            if current.children[part] == nil {
                current.children[part] = &TreeNode{
                    name:     part,
                    children: make(map[string]*TreeNode),
                }
            }
            current = current.children[part]

            // Mark as leaf and add method if last segment
            if i == len(parts)-1 {
                current.isLeaf = true
                current.methods = append(current.methods, route.Method)
            }
        }
    }

    // Render tree
    var output strings.Builder
    output.WriteString(".\n")
    renderTree(root, "", &output)

    c.String(http.StatusOK, output.String())
}

func renderTree(node *TreeNode, prefix string, output *strings.Builder) {
    // Get sorted children
    keys := make([]string, 0, len(node.children))
    for k := range node.children {
        keys = append(keys, k)
    }
    sort.Strings(keys)

    for i, key := range keys {
        child := node.children[key]
        isLastChild := i == len(keys)-1

        // Determine the branch characters	
        branch := "├── "
        if isLastChild {
            branch = "└── "
        }

        // Format output
        output.WriteString(prefix + branch)

        if child.isLeaf && len(child.methods) > 0 {
            sort.Strings(child.methods)
            output.WriteString("[" + strings.Join(child.methods, ", ") + "] ")
        }

        output.WriteString(child.name + "\n")

        // Recursively render children
        if len(child.children) > 0 {
            var newPrefix string
            if isLastChild {
                newPrefix = prefix + "    "
            } else {
                newPrefix = prefix + "│   "
            }
            renderTree(child, newPrefix, output)
        }
    }
}