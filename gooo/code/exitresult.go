type ExitResult interface {
	ExitStatus() int
}

if exiterr, ok := err.(*exec.ExitError); ok {
	if status, ok := exiterr.Sys().(ExitResult); ok {
    } else {
    	fmt.Println("platform does not support ExitResult interface")
    }
} else {
	fmt.Println("not an ExitError")
}