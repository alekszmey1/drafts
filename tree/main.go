package main

import (
	"fmt"
	"log"
)

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}

func main() {
	k := 11
	treeNode := InitTree(k)
	fmt.Println(treeNode)
}
func InitTree(k int) (root *TreeNode) {
	value := 1
	rootNode := TreeNode{Value: value}
	for i := 2; i <= k; i++ {
		node := TreeNode{}
		fmt.Printf("вставляем %v значение \n", i)
		InsertNodeToTree(&rootNode, &node, i)
	}
	return &rootNode
}

func InsertNodeToTree(root *TreeNode, node *TreeNode, i int) {
	if root.Right != nil && root.Left != nil {
		fmt.Println("левая и правая ветка не пусты")
		if i%4 == 0 || i%4 == 1 {
			fmt.Println("спускаемся по левой ветке")
			InsertNodeToTree(root.Left, node, i)
		} else {
			fmt.Println("спускаемся по правой ветке")
			InsertNodeToTree(root.Right, node, i)
		}
	} else if root.Right == nil && root.Left == nil {
		fmt.Printf("левая ветка пуста, ставим значение %v \n", root.Value*2)
		node.Value = root.Value * 2
		root.Left = node
		return
	} else if root.Left != nil && root.Right == nil {
		fmt.Println("сработало иначе")
		fmt.Printf("левая ветка не пуста, ставим значение в правую ветку %v \n", root.Left.Value+1)
		node.Value = root.Left.Value + 1
		root.Right = node
		return
	}
}

func change(parent *TreeNode, i int) *TreeNode {
	if parent.Left.Value == i {
		log.Println("нашли совпадение слева ")
		fmt.Println(parent.Value)
		//fmt.Println(parent.Left.Value)
		child := TreeNode{
			Left:   parent.Left.Left,
			Right:  parent.Right,
			Value:  parent.Value,
			Parent: parent.Left,
		}
		parent.Left = &child
		fmt.Println(child.Value)
		fmt.Println(parent.Value)
		//fmt.Println(parent.Left.Value)
	} else if parent.Right.Value == i {
		log.Println("нашли совпадение справа ")
		fmt.Println(parent.Value)
		//fmt.Println(parent.Right.Value)
		child := TreeNode{
			Left:   parent.Left,
			Right:  parent.Right.Right,
			Value:  parent.Value,
			Parent: parent.Right,
		}
		parent.Right = &child
		fmt.Println(child.Value)
		fmt.Println(parent.Value)
		//fmt.Println(parent.Right.Value)
	}
	return parent
}
func (root *TreeNode) searchAndChange(i int) {
	if root != nil {
		root.Left.searchAndChange(i)
		root.Right.searchAndChange(i)
	}

}
func (root *TreeNode) search(i int) {
	fmt.Println("взяли число ", +i)
	if root != nil {
		if root.Left != nil {
			// print left tree
			if root.Left.Value == i {
				change(root, i)
				return
			}
			root.Left.search(i)
		} else if root.Right != nil {
			if root.Left.Value == i {
				change(root, i)
				return
			}
			root.Right.search(i)
		}
	}
}
func SearchBool(t *TreeNode, i int) bool {
	b := false
	fmt.Println(i, t.Value)
	if t.Right != nil && t.Right.Value == i || t.Left != nil && i == t.Left.Value {
		fmt.Println("нашли совпадение")
		b = true
	}
	return b
}
func search2(t *TreeNode, k int) *TreeNode {
	var b bool
	b = SearchBool(t, k)
	if b == true {
		return t
	} else if t != nil {
		if t.Right != nil {
			b = SearchBool(t.Right, k)
			if b == true {
				return t.Right
			} else if t.Left != nil {
				fmt.Println("пошли налево")
				return search2(t.Left, k)
			} else if t.Parent.Parent.Right != nil {
				fmt.Println("поднялись наверх и пошли направо")
				return search2(t.Parent.Right, k)
			}
		}
	}
	return nil
}
