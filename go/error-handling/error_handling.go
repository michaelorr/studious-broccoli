package erratum

const testVersion = 2

func Use(opener ResourceOpener, input string) (retErr error) {
	resource, retErr := openResource(opener)
	if retErr != nil {
		return retErr
	}

	defer func() {
		if err := recover(); err != nil {
			if frobErr, ok := err.(FrobError); ok {
				resource.Defrob(frobErr.defrobTag)
			}
			retErr = err.(error)
		}
		resource.Close()
	}()

	resource.Frob(input)
	return retErr
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
