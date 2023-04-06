package commands

import (
	"context"
	"fmt"
	"strings"

	"github.com/config-source/cli/client"
	"github.com/spf13/cobra"
)

var envTreeCmd = &cobra.Command{
	Use:   "tree",
	Short: "Print the promotion tree of your environments",
	RunE: func(cmd *cobra.Command, args []string) error {
		envs, err := api.Environments().All(context.Background())
		if err != nil {
			return err
		}

		type treeNode struct {
			element  client.Environment
			parent   *treeNode
			children []*treeNode
		}

		root := treeNode{}
		lookupTable := make(map[int64]*treeNode)

		// First find the root environment and update our root node.
		for _, env := range envs {
			if env.PromotesToId == 0 {
				root.element = env
				lookupTable[env.Id] = &root
			}
		}

		// Now we build the tree using the lookupTable to determine if a node
		// already exists
		for _, env := range envs {
			var node *treeNode

			existingNode, ok := lookupTable[env.Id]
			if ok && existingNode.element.Id == env.Id {
				continue
			} else if ok {
				node = existingNode
				node.element = env
			} else {
				node = &treeNode{element: env}
			}

			lookupTable[env.Id] = node
			parent, ok := lookupTable[env.PromotesToId]
			if !ok {
				parent = &treeNode{
					children: []*treeNode{node},
				}
				lookupTable[env.PromotesToId] = parent
			} else {
				parent.children = append(parent.children, node)
			}
		}

		depth := 0
		nodes := []*treeNode{&root}
		for len(nodes) > 0 {
			children := make([]*treeNode, 0)

			for _, node := range nodes {
				indent := ""
				leading := strings.Repeat("─", depth)
				parentMarker := ""
				if depth != 0 {
					parentMarker = "└"
					indent = strings.Repeat("   ", depth)
				}

				fmt.Printf("%s%s%s%s\n", indent, parentMarker, leading, node.element.Name)
				children = append(children, node.children...)
			}

			depth += 1
			nodes = children
		}

		return nil
	},
}
