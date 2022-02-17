## HashMap实现

```go
type Solution struct {
	n       int
	mapping map[int]int
}

func Constructor(n int, blacklist []int) Solution {
	mapping := map[int]int{}
	sz := n - len(blacklist)
	for _, v := range blacklist {
		mapping[v] = 666
	}

	last := n - 1
	check := func(m map[int]int, key int) bool {
		_, ok := m[key]
		return ok
	}
	for _, v := range blacklist {
		if v >= sz {
			continue
		}
		for check(mapping, last) {
			last--
		}
		mapping[v] = last
		last--
	}
	return Solution{mapping: mapping, n: sz}
}

func (this *Solution) Pick() int {
	index := rand.Intn(this.n)
	if v, ok := this.mapping[index]; ok {
		return v
	}
	return index
}

```