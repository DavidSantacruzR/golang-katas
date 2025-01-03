package error_handling

func UnwrapOrElse[T any](result T, err error, handleError func(error) T) T {
	if err != nil {
		return handleError(err)
	}
	return result
}
