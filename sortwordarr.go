package piscine

func SortWordArr(a []string) {
	quickSort(a, 0, len(a)-1)
}

func quickSort(table []string, begin int, end int) {
	if begin < end {
		locus := swap(table, begin, end)
		quickSort(table, begin, locus-1)
		quickSort(table, locus+1, end)
	}
}

func swap(table []string, begin int, end int) int {
	last := table[end]
	i := begin - 1
	for j := begin; j < end; j++ {
		if table[j] <= last {
			i++
			temp := table[j]
			table[j] = table[i]
			table[i] = temp
		}
	}
	temp := table[end]
	table[end] = table[i+1]
	table[i+1] = temp
	return i + 1
}
