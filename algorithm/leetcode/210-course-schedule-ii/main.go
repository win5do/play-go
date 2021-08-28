package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		edges   = make([][]int, numCourses)
		visited = make([]int, numCourses)
		result  []int
		valid   bool = true
		dfs     func(u int)
	)

	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[u] = 2
		result = append(result, u)
	}

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	if !valid {
		return []int{}
	}
	for i := 0; i < len(result)/2; i++ {
		result[i], result[numCourses-i-1] = result[numCourses-i-1], result[i]
	}
	return result
}
