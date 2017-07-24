package erratum

const testVersion = 2

func Use(opener ResourceOpener, input string) (ret_err error) {
	resource, ret_err := openResource(opener)
	if ret_err != nil {
		return ret_err
	}

	defer func() {
		if err := recover(); err != nil {
			if frobErr, ok := err.(FrobError); ok {
				resource.Defrob(frobErr.defrobTag)
			}
			ret_err = err.(error)
		}
		resource.Close()
	}()

	resource.Frob(input)
	return ret_err
}

func openResource(opener ResourceOpener) (Resource, error) {
AttemptOpen:
	resource, err := opener()
	if err != nil {
		if _, ok := err.(TransientError); ok {
			goto AttemptOpen
		}
		return nil, err
	}
	return resource, nil
}
