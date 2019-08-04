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
		min, max := getBoundary(D, m, n)

		for k := max; k >= min; k -= 2 {
			i := advance(v, D, k, offset)

			p = append(p, path{i: i, j: i + k})

			for i < m && i+k < n && x[i] == y[i+k] {
				i++
				p = append(p, path{i: i, j: i + k})
			}
			if k == n-m && i == m {
				return tracePath(p, x, y)
			}
			v[offset+k] = i
		}
	}
	// Never reach here.
	panic(fmt.Errorf("found a bug: x = %v, y = %v ", x, y))
}

func tracePath(path []path, x []rune, y []rune) []Edit {
	i := len(path) - 1
	k := 1

	var reversedSes []Edit
	for i > 0 {
		p := path[i]
		q := path[i-k]
		switch {
		case p.i-q.i == 1 && p.j-q.j == 1:
			edit := Edit{Action: 0, Value: string(x[q.i])}
			reversedSes = append(reversedSes, edit)
		case p.j-q.j == 1 && p.i == q.i:
			edit := Edit{Action: 1, Value: string(y[q.j])}
			reversedSes = append(reversedSes, edit)
		case p.i-q.i == 1 && p.j == q.j:
			edit := Edit{Action: -1, Value: string(x[q.i])}
			reversedSes = append(reversedSes, edit)
		default:
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
	if d == 0 {
		return 0
	}
	if k == -d {
		return v[offset+k+1] + 1
	}
	if k == d {
		return v[offset+k-1]
	}
	if v[offset+k+1]+1 > v[offset+k-1] {
		return v[offset+k+1] + 1
	}
	return v[offset+k-1]
}
