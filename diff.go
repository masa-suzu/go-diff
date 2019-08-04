package diff

import "fmt"

type path struct {
	i int
	j int
}

// Diff computes difference between a and b.
func Diff(a, b string) []Edit {
	x := []rune(a)
	y := []rune(b)

	return diff(x, y)
}

// diff computes Shortest Edit Script by O(ND) algorithm.
// http://www.xmailserver.org/diff2.pdf
func diff(x, y []rune) []Edit {

	m := len(x)
	n := len(y)
	v := make([]int, m+n+1)
	offset := m

	var p []path

	for D := 0; D <= m+n; D++ {
		var min, max int
		if D <= m {
			min = -D
		} else {
			min = D - (2 * m)
		}
		if D <= n {
			max = D
		} else {
			max = -D + (2 * n)
		}

		for k := max; k >= min; k -= 2 {
			i := 0
			if D == 0 {
				i = 0
			} else if k == -D {
				i = v[offset+k+1] + 1
			} else if k == D {
				i = v[offset+k-1]
			} else {
				if v[offset+k+1]+1 > v[offset+k-1] {
					i = v[offset+k+1] + 1
				} else {
					i = v[offset+k-1]
				}
			}

			p = append(p, path{i: i, j: i + k})

			for i < m && i+k < n && x[i] == y[i+k] {
				i++
				p = append(p, path{i: i, j: i + k})
			}
			if k == n-m && i == m {
				return traceGraph(p, x, y)
			}
			v[offset+k] = i
		}
	}
	//
	panic(fmt.Errorf("Found a bug: x = %v, y = %v ", x, y))
}

func traceGraph(path []path, x []rune, y []rune) []Edit {
	i := len(path) - 1
	k := 1

	var reversedSes []Edit
	for i > 0 {
		p := path[i]
		q := path[i-k]
		if p.i-q.i == 1 && p.j-q.j == 1 {
			edit := Edit{Action: 0, Value: string(x[q.i])}
			reversedSes = append(reversedSes, edit)
		} else if p.j-q.j == 1 && p.i == q.i {
			edit := Edit{Action: 1, Value: string(y[q.j])}
			reversedSes = append(reversedSes, edit)
		} else if p.i-q.i == 1 && p.j == q.j {
			edit := Edit{Action: -1, Value: string(x[q.i])}
			reversedSes = append(reversedSes, edit)
		} else {
			k++
			continue
		}
		i -= k
		k = 1
	}

	var ses []Edit
	for i := len(reversedSes) - 1; i >= 0; i-- {
		ses = append(ses, reversedSes[i])
	}
	return ses
}
