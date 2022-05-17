package doctree

import (
	"fmt"
)

type DocTree struct {
	node    string
	subTree []DocTree
}

func (dt DocTree) String() string {
	curStr := fmt.Sprintf("%v\n", dt.node)
	for _, sub := range dt.subTree {
		curStr += fmt.Sprintf("  %v", sub)
		curStr += "\n"
	}
	return curStr
}

// CreateDocTrees 创建文档树
// ori 原始数据
func CreateDocTrees(ori map[string]interface{}) []DocTree {
	top := make([]DocTree, 0, len(ori))
	for k, v := range ori {
		top = append(top, DocTree{
			k, createDocTreeHelper(v),
		})
	}
	return top
}

func createDocTreeHelper(content interface{}) []DocTree {
	res := make([]DocTree, 0, 4)
	switch v := content.(type) {
	case string:
		res = append(res, DocTree{v, nil})
	case bool:
		bStr := fmt.Sprintf("%v", v)
		res = append(res, DocTree{bStr, nil})
	case int:
		iStr := fmt.Sprintf("%d", v)
		res = append(res, DocTree{iStr, nil})
	case float64:
		fStr := fmt.Sprintf("%.2f", v)
		res = append(res, DocTree{fStr, nil})
	case []interface{}:
		for _, nv := range v {
			res = append(res, DocTree{"arrNode", createDocTreeHelper(nv)})
		}
	case map[interface{}]interface{}:
		for nk, nv := range v {
			if sk, ok := nk.(string); ok {
				res = append(res, DocTree{sk, createDocTreeHelper(nv)})
			}
		}
	}
	return res
}
