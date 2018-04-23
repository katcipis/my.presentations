// Sys returns system-dependent exit information about

// the process. Convert it to the appropriate underlying

// type, such as syscall.WaitStatus on Unix, to access its contents.

func (p *ProcessState) Sys() interface{} {

	return p.sys()

}