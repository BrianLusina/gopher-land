package erratum

func Use(opener ResourceOpener, input string) (err error) {
	var resource Resource

	for {
		resource, err = opener()
		if err == nil {
			break
		}
		if _, ok := err.(TransientError); !ok {
			return
		}
	}

	// another implementation of the above flow control
	// for resource, err = opener(); err != nil; resource, err = opener() {
	// 	if _, ok := err.(TransientError); !ok {
	// 		return
	// 	}
	// }

	defer resource.Close()

	defer func() {
		if r := recover(); r != nil {
			switch r := r.(type) {
			case FrobError:
				resource.Defrob(r.defrobTag)
				err = r
			case error:
				err = r
			default:
				err = NewUnexpectedError(r)
			}
			// or succintly:
			// if frobErr, ok := r.(FrobError); ok {
			// 	resource.Defrob(frobErr.defrobTag)
			// }

			// err = r.(error)
		}
	}()

	resource.Frob(input)
	return
}
