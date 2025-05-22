package db

// MakeArgsList takes a prefix (like domainID) and a slice of strings (like URLs or reports),
// and builds a [][]any that can be used with ExecBulkInsert.
func MakeArgsList(prefix any, values []string) [][]any {
	argsList := make([][]any, len(values))
	for i, val := range values {
		argsList[i] = []any{prefix, val}
	}
	return argsList
}
