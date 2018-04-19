type WaitStatus uint32

func (w WaitStatus) ExitStatus() int {
	if !w.Exited() {
		return -1
	}
	return int(w>>shift) & 0xFF
}