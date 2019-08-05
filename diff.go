package diff

import (
	"fmt"
	"strings"
)

type path struct {
	i int
	j int
}

// Diff computes difference between a and b.
func Diff(a, b string) []Edit {
	x := strings.Split(a, "\n")
	y := strings.Split(b, "\n")

	return diff(x, y)
}

// diff computes Shortest Edit Script by O(ND) algorithm.
// http://www.xmailserver.org/diff2.pdf
func diff(x, y []string) []Edit {

	m := len(x)
	n := len(y)
	v := make([]int, m+n+1)

	var p []path

	for d := 0; d <= m+n; d++ {
		min, max := getBoundary(d, m, n)

		for k := max; k >= min; k -= 2 {
			i := advance(v, d, k, m)

			p = append(p, path{i: i, j: i + k})

			for i < m && i+k < n && x[i] == y[i+k] {
				i++
				p = append(p, path{i: i, j: i + k})
			}
			if k == n-m && i == m {
				return tracePath(p, x, y)
			}
			v[m+k] = i
		}
	}
	// Never reach here.
	panic(fmt.Errorf("found a bug: x = %v, y = %v ", x, y))
}

func tracePath(path []path, x []string, y []string) []Edit {
	i := len(path) - 1
	k := 1

	var ses []Edit
	for i > 0 {
		p := path[i]
		q := path[i-k]

		di := p.i - q.i
		dj := p.j - q.j

		switch {
		case di == 1 && dj == 1:
			ses = append([]Edit{{Action: 0, Value: x[q.i]}}, ses...)
		case di == 0 && dj == 1:
			ses = append([]Edit{{Action: 1, Value: y[q.j]}}, ses...)
		case di == 1 && dj == 0:
			if len(ses) > 0 && ses[0].Action == 1 {
				added := ses[0]
				ses[0] = Edit{Action: 2, From: x[q.i], Value: added.Value}
			} else {
				ses = append([]Edit{{Action: -1, Value: x[q.i]}}, ses...)
			}
		default:
			k++
			continue
		}
		i -= k
		k = 1
	}

	return ses
}

func getBoundary(d int, m int, n int) (int, int) {
	var min, max int
	if d <= m {
		min = -d
	} else {
		min = d - (2 * m)
	}
	if d <= n {
		max = d
	} else {
		max = -d + (2 * n)
	}
	return min, max
}

func advance(v []int, d int, k int, offset int) int {
	positive := offset + k + 1
	negative := offset + k - 1
	if d == 0 {
		return 0
	}
	if k == -d {
		return v[positive] + 1
	}
	if k == d {
		return v[negative]
	}
	if v[positive]+1 > v[negative] {
		return v[positive] + 1
	}
	return v[negative]
}
