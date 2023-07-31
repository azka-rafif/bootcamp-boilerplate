package pagination

import "strconv"

func ConvertToInt(s string) (int, error) {
	return strconv.Atoi(s)
}
