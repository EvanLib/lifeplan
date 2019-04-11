package calendar

import fmt "fmt"

// StateMap create a map of tree tasks and there states
func StateMap(taskNodes []*TreeNode, m map[string]int32) {
	for _, treeNode := range taskNodes {
		m[treeNode.Task.Id] = treeNode.Task.State
		if len(treeNode.Subtasks) > 0 {
			StateMap(treeNode.Subtasks, m)
		}
	}
}

// StateMap create a map of tree tasks and there states
func ApplyStateMap(taskNodes []*TreeNode, stateMap map[string]int32) {
	for _, treeNode := range taskNodes {
		if val, ok := stateMap[treeNode.Task.Id]; ok {
			treeNode.Task.State = val
		}
		if len(treeNode.Subtasks) > 0 {
			ApplyStateMap(treeNode.Subtasks, stateMap)
		}
	}
}

// Prints out task tree
func PrintRec(taskNodes []*TreeNode) {
	for _, treeNode := range taskNodes {
		fmt.Println(len(treeNode.Subtasks))
		fmt.Println(treeNode.Task)
		if len(treeNode.Subtasks) > 0 {
			PrintRec(treeNode.Subtasks)
		}
	}
}
