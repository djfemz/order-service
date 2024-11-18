package apperrors


const(
	ORDER_CREATION_FAILED_MESSAGE="failed to create order, message:: %s"
	INVALID_USER_ID_ERROR_MESSAGE ="user with id %d not found, error:: %s"
)


type AppError struct {
	Message string
}

type OrderCreationFailedError struct {
	*AppError
	Status string
}

func NewOrderCreationFailedError(message string) error {
	return &OrderCreationFailedError{
		AppError: &AppError{Message: message},
	}
}

func (orderCreationFailedError *OrderCreationFailedError) Error() string {
	return orderCreationFailedError.Message
}
