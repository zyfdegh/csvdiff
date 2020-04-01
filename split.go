package main

// Split arr to piece, each len is n
func Split(arr []string, n int) [][]string {
	chunks := make([][]string, 0, len(arr)/n+1)
	for i := 0; i < len(arr); i += n {
		end := i + n
		if end > len(arr) {
			end = len(arr)
		}
		chunks = append(chunks, arr[i:end])
	}
	return chunks
}
