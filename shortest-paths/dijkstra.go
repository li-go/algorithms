// https://leetcode.com/problems/word-ladder-ii/
package main

type DisRecord struct {
	Val     int
	Paths   [][]int
	Visited bool
}

func (d *DisRecord) AddEle(ele int) {
	for i := range d.Paths {
		d.Paths[i] = append(d.Paths[i], ele)
	}
}

func (d *DisRecord) AddPaths(paths [][]int) {
	for _, p := range paths {
		tmpP := make([]int, len(p))
		copy(tmpP, p)
		d.Paths = append(d.Paths, tmpP)
	}
}

func (d *DisRecord) ReplacePaths(paths [][]int) {
	d.Paths = make([][]int, len(paths))
	for i, p := range paths {
		d.Paths[i] = make([]int, len(p))
		copy(d.Paths[i], p)
	}
}

var (
	g [][]int
)

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	list := append([]string{beginWord}, wordList...)
	list = unique(list)
	var end int
	for i, s := range list {
		if s == endWord {
			end = i
			break
		}
	}
	if end == 0 {
		return nil
	}

	initGraph(list)
	dis := make([]*DisRecord, len(list))
	dijkstra(0, dis)

	if !dis[end].Visited {
		return nil
	}

	paths := dis[end].Paths
	r := make([][]string, len(paths))
	for i := 0; i < len(paths); i++ {
		r[i] = make([]string, len(paths[i]))
		for j := 0; j < len(paths[i]); j++ {
			r[i][j] = list[paths[i][j]]
		}
	}
	return r
}

func dijkstra(begin int, dis []*DisRecord) {
	for i := 0; i < len(dis); i++ {
		dis[i] = &DisRecord{g[begin][i], [][]int{{begin}}, false}
	}
	dis[begin].Visited = true

	count := 1
	for count < len(dis) {
		var minI int
		minD := &DisRecord{Val: 1 << 32}
		for i, d := range dis {
			if !d.Visited && d.Val < minD.Val {
				minI = i
				minD = d
			}
		}
		minD.AddEle(minI)
		minD.Visited = true
		count++
		// re-calculate dis left
		for i, d := range dis {
			if d.Visited {
				continue
			}
			if minD.Val+g[i][minI] < d.Val {
				d.Val = minD.Val + g[i][minI]
				d.ReplacePaths(minD.Paths)
			} else if minD.Val+g[i][minI] == d.Val {
				d.AddPaths(minD.Paths)
			}
		}
	}
}

func initGraph(list []string) {
	g = make([][]int, len(list))
	for i := 0; i < len(list); i++ {
		g[i] = make([]int, len(list))
	}
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if match(list[i], list[j]) {
				g[i][j] = 1
				g[j][i] = 1
			} else {
				g[i][j] = 1 << 32
				g[j][i] = 1 << 32
			}
		}
	}
}

func unique(list []string) []string {
	m := make(map[string]bool)
	var nl []string
	for _, s := range list {
		if _, ok := m[s]; ok {
			continue
		}
		m[s] = true
		nl = append(nl, s)
	}
	return nl
}

func match(w1, w2 string) bool {
	var diff int
	for i := 0; i < len(w1); i++ {
		if w1[i] != w2[i] {
			diff++
			if diff > 1 {
				return false
			}
		}
	}
	return true
}
