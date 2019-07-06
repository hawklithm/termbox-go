package termbox

// On macOS, enable behavior which will wait before deciding that the escape
// key was pressed, to account for partially send escape sequences, especially
// with regard to lengthy mouse sequences.
// See https://github.com/hawklithm/termbox-go/issues/132
func enable_wait_for_escape_sequence() bool {
	return true
}
