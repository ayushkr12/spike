package utils

func RemoveDuplicatesAndEmptyStrings(slice []string) []string {
	unique := make(map[string]struct{})
	result := []string{}

	for _, item := range slice {
		if _, exists := unique[item]; !exists {
			unique[item] = struct{}{}
			if item != "" {
				result = append(result, item)
			}
		}
	}

	return result
}
