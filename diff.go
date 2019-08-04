package diff

// Diff computes difference between a and b.
func Diff(a, b string) []Edit {
	x := []rune(a)
	y := []rune(b)

	return diff(x, y)
}

func diff(x, y []rune) []Edit {
	m := len(x)
	n := len(y)

	g := initGraph(m, n)

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if x[i-1] == y[j-1] {
				g[i][j] = g[i-1][j-1]
			} else {
				g[i][j] = min(g[i][j-1]+1, g[i-1][j]+1)
			}
		}
	}

	r := traceGraph([]Edit{}, g, x, y, m, n)

	var d []Edit
	for i := len(r) - 1; i >= 0; i-- {
		d = append(d, r[i])
	}
	return d
}

func traceGraph(ses []Edit, g [][]int, x []rune, y []rune, i int, j int) []Edit {

	if i == 0 && j == 0 {
		return ses
	}

	if i == 0 {
		return markAdded(ses, g, x, y, i, j)
	}

	if j == 0 {
		return markDeleted(ses, g, x, y, i, j)
	}

	return stepGraph(ses, g, x, y, i, j)
}

func stepGraph(ses []Edit, g [][]int, x []rune, y []rune, i int, j int) []Edit {
	now := g[i][j]
	unchanged := g[i-1][j-1]
	added := g[i][j-1]
	deleted := g[i-1][j]

	if unchanged < added && unchanged < deleted && now == unchanged {
		return markUnChanged(ses, g, x, y, i, j)
	}

	if added <= deleted {
		return markAdded(ses, g, x, y, i, j)
	}

	return markDeleted(ses, g, x, y, i, j)
}

func markUnChanged(ses []Edit, g [][]int, x []rune, y []rune, i int, j int) []Edit {
	edit := Edit{Action: 0, Value: string(y[j-1])}
	ses = append(ses, edit)
	return traceGraph(ses, g, x, y, i-1, j-1)
}

func markAdded(ses []Edit, g [][]int, x []rune, y []rune, i int, j int) []Edit {
	edit := Edit{Action: 1,  Value: string(y[j-1])}
	ses = append(ses, edit)
	return traceGraph(ses, g, x, y, i, j-1)
}

func markDeleted(ses []Edit, g [][]int, x []rune, y []rune, i int, j int) []Edit {
	edit := Edit{Action: -1, Value: string(x[i-1])}
	ses = append(ses, edit)
	return traceGraph(ses, g, x, y, i-1, j)
}

func initGraph(m int, n int) [][]int {
	g := make([][]int, m+1)

	for i := range g {
		g[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 || j == 0 {
				g[i][j] = i + j
			}
		}
	}
	return g
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
