package binary_tree

import (
	"code_pratice/common"
	"strings"
)

func InvertTree(root *common.Node[int]) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	InvertTree(root.Left)
	InvertTree(root.Right)
}

// 求二叉树的所有路径
func BinaryTreePaths(root *common.Node[string]) []string {
	var (
		path []string
		res  []string
	)
	getPath(root, &path, &res)
	return res
}

func getPath(root *common.Node[string], path *[]string, res *[]string) {

	*path = append(*path, root.Data) // 这里和回溯path的部分呼应

	if root.Left == nil && root.Right == nil {
		tmpStr := strings.Join(*path, "->")
		*res = append(*res, tmpStr)
		return
	}

	if root.Left != nil {
		getPath(root.Left, path, res)
		*path = (*path)[:len(*path)-1] // 这里是回溯
	}

	if root.Right != nil {
		getPath(root.Right, path, res)
		*path = (*path)[:len(*path)-1] // 这里是回溯
	}
}

//不再需要了，不涉及int转string了
//func toStrArr(arr []int) string {
//	strArr := make([]string, 0, len(arr))
//	for _, ele := range arr {
//		strArr = append(strArr, strconv.Itoa(ele))
//	}
//	return strings.Join(strArr, "->")
//}
