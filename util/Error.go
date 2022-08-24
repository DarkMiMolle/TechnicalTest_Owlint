package util

// PanicErr will panic if err != nil
func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}
