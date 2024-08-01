package utils

import (
	"encoding/json"
	"fmt"
	"lexium-utility/datahandler"
	"os"
	"strconv"
	"strings"
)

func CreateTree(input []datahandler.Element) *datahandler.TreeNode {
	data, _ := json.Marshal(input)
	os.WriteFile("inputData.json", data, 0777)
	node := new(datahandler.TreeNode)
	node.Children = make([]*datahandler.TreeNode, 0)
	tm := make(map[string]*datahandler.TreeNode)
	checkEls := make(map[string]bool)
	for _, v := range input {
		checkEls[v.ElementID] = true
		if v.DataType == "object" || v.DataType == "array" {
			tm[v.ElementID] = &datahandler.TreeNode{NodeData: v, Children: make([]*datahandler.TreeNode, 0)}
		} else {
			tm[v.ElementID] = &datahandler.TreeNode{NodeData: v}
		}
	}
	for _, v := range tm {
		if v.NodeData.ElementID == "" || v.NodeData.ElementID == "Home" {
			continue
		}
		p := getParent(v.NodeData.Parent, v.NodeData.Depth, v.NodeData.GroupID)
		if (v.NodeData.Parent != p || (strings.Contains(v.NodeData.ElementID, p) && v.NodeData.JsonTagName == "") || (!strings.Contains(p, "Gr"))) && v.NodeData.GroupID != 0 {
			checkEls[v.NodeData.ElementID] = false
		}
		if p == "" || p == "Home" {

			node.Children = append(node.Children, v)
			// sort.Slice(node.Children, func(a, b int) bool {
			// 	return node.Children[a].NodeData.SeqID < node.Children[b].NodeData.SeqID
			// })
			v.Parent = node
			insertionSort(node.Children)
		} else {

			if tm[p] == nil || tm[p].Children == nil {
				fmt.Println(p)
			} else {

				tm[p].Children = append(tm[p].Children, v)
				v.Parent = tm[p]
				// sort.Slice(tm[p].Children, func(a, b int) bool {
				// 	return tm[p].Children[a].NodeData.SeqID < tm[p].Children[b].NodeData.SeqID
				// })
				insertionSort(tm[p].Children)
			}

		}
	}
	f, _ := os.Create("./checkEls.txt")
	for k, v := range checkEls {
		if !v {
			fmt.Fprintln(f, k)
		}
	}
	// f, _ := os.Create("./checkEls.txt")
	// for k, v := range checkEls {
	// 	if !v {
	// 		fmt.Fprintln(f, k)
	// 	}
	// }
	return node
}

func getParent(p string, d string, g int) string {

	arr := strings.Split(d, "-")
	if len(arr) == 1 || g != 0 {
		return p
	} else {
		val2, _ := strconv.Atoi(arr[1])
		val1, _ := strconv.Atoi(arr[0][1:])
		parr := strings.Split(p, ".")
		if len(parr) < val1-val2 {
			fmt.Println(len(parr), p, d)
			return p
		}
		return strings.Join(parr[:val1-val2], ".")
	}
}

func insertionSort(arr []*datahandler.TreeNode) []*datahandler.TreeNode {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1].NodeData.SeqID > arr[j].NodeData.SeqID; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}
