package algorithm

func reverse(arr []rune) {
	i := 0
	j := len(arr) - 1

	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

// NextPermutation return the next permutation of str
func NextPermutation(str *string) bool {
	if len(*str) <= 1 {
		return false
	}

	arr := []rune(*str)

	last := len(arr) - 1
	i := last

	for {
		var i1, i2 int

		i1 = i
		i--
		if arr[i] < arr[i1] {
			i2 = last
			for !(arr[i] < arr[i2]) {
				i2--
			}

			arr[i], arr[i2] = arr[i2], arr[i]

			for i1 < last {
				arr[i1], arr[last] = arr[last], arr[i1]
				i1++
				last--
			}

			*str = string(arr)

			return true
		}

		if i == 0 {
			reverse(arr)
			*str = string(arr)
			return false
		}
	}
}
