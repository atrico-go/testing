package assert

func PanicCatcher(action func()) (thePanic interface{}) {
	func() {
		defer func() { thePanic = recover() }()
		action()
	}()
	return thePanic
}
