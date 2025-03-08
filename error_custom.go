package main

import "fmt"

type validationError struct {
	Message string
}

type notFoundError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"validation err"}
	}

	if id != "Ivan" {
		return &notFoundError{"not found err"}
	}

	return nil
}

func main() {
	err := SaveData("ivan", nil)
	if err != nil {
		switch finalErr := err.(type) {
		case *validationError:
			fmt.Println("validation err", finalErr.Error())
		case *notFoundError:
			fmt.Println("not found err", finalErr.Error())
		default:
			fmt.Println("unknown err", err.Error())

		}
	} else {
		fmt.Println("sukses")
	}
}
