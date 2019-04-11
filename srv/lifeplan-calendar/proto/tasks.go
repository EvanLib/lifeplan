package calendar

// StateMap create a map of tree tasks and there states
func StateMap(taskNodes []*TreeNode, m map[string]int32) {
	for _, treeNode := range taskNodes {
		m[treeNode.Task.Id] = treeNode.Task.State
		if len(treeNode.Subtasks) > 0 {
			StateMap(treeNode.Subtasks, m)
		}
	}
}
