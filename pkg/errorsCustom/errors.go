package errorsCustom

import "fmt"

type NotFoundError struct {
	ID       any    // Любой тип
	Resource string // Описание ошибки
}

type UnauthorizedError struct {
	ID     any
	Reason string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("%s With err ID %v not found", e.Resource, e.ID)
}

func (e UnauthorizedError) Error() string {
	if e.Reason == "" {
		return "Unauthorized"
	}

	return "Unauthorized: " + e.Reason
}
