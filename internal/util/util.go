package util

func RemoveFromSlice(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
