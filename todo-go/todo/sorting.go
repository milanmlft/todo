package todo

// Implement the sort.Interface for the Todos struct,
// based on the Priority and position fields
// The sort.Interface requires the Len(), Less() and Swap() meethods
// https://pkg.go.dev/sort#Interface

func (t Todos) Len() int      { return len(t) }
func (t Todos) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t Todos) Less(i, j int) bool {
	// If priorities are equal, use position to determine order
	if t[i].Priority == t[j].Priority {
		return t[i].position < t[j].position
	}
	return t[i].Priority < t[j].Priority
}
