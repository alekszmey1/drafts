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

var DefaultValue = -1024

func main() {
	treeNode := InitTree(5, 4, 6, 8, 9, 7, 1, 3, 2)
	fmt.Println(treeNode)
}
func InitTree(values ...int) (root *TreeNode) {
	rootNode := TreeNode{Value: DefaultValue, Right: nil, Left: nil}
	for _, value := range values {
		node := TreeNode{Value: value}
		InsertNodeToTree(&rootNode, &node)
	}
	return &rootNode
}

/*func InsertNodeToTree(tree *TreeNode, node *TreeNode) {
	if tree == nil {
		return
	}
	if tree.Value == DefaultValue {
		tree.Value = node.Value
		return
	}
	if node.Value > tree.Value {
		log.Printf("вставляемое значение %v больше значения в дереве %v", node.Value, tree.Value)
		if tree.Right == nil {
			log.Printf("правая ветка пуста, вставляем значение %v ", node.Value)
			tree.Right = &TreeNode{Value: DefaultValue}
		}
		InsertNodeToTree(tree.Right, node)
	}
	if node.Value < tree.Value {
		log.Printf("вставляемое значение %v меньше значения в дереве %v", node.Value, tree.Value)
		if tree.Left == nil {
			log.Printf("левая ветка пуста, вставляем значение %v ", node.Value)
			tree.Left = &TreeNode{Value: DefaultValue}
		}
		InsertNodeToTree(tree.Left, node)
	}
}*/

func InsertNodeToTree(tree *TreeNode, node *TreeNode) {
	if tree == nil {
		panic("cannot insert into nil root")
	}
	if node.Value > tree.Value {
		log.Printf("вставляемое значение %v больше значения в дереве %v", node.Value, tree.Value)
		if tree.Right == nil {
			log.Printf("правая ветка пуста, вставляем значение %v ", node.Value)
			tree.Right = node
		} else {
			log.Printf("правая ветка не пуста, спускаемся на уровень ниже")
			InsertNodeToTree(tree.Right, node)
		}
	}
	if node.Value < tree.Value {
		log.Printf("вставляемое значение %v меньше значения в дереве %v", node.Value, tree.Value)
		if tree.Left == nil {
			log.Printf("левая ветка пуста, вставляем значение %v ", node.Value)
			tree.Left = node
		} else {
			log.Printf("правая ветка не пуста, спускаемся на уровень ниже")
			InsertNodeToTree(tree.Left, node)
		}
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
