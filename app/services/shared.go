package services
func perror(err error) {
	if err != nil {
		panic(err)
	}
}
