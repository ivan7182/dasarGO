package main

import "fmt"

type Validationerr struct {
	Message string
}

type NotFoundErr struct {
	Message string
}

func (v *Validationerr) Error() string {
	return v.Message
}

func (n *NotFoundErr) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &Validationerr{Message: "Validator err"}
	}
	if id != "ivan" {
		return &NotFoundErr{Message: "not found err"}
	}

	return nil
}

func main() {

	err := SaveData("", nil)
	if err != nil {
		switch FinalErr := err.(type) {
		case *Validationerr:
			fmt.Println("validation Err :", FinalErr.Error())
		case *NotFoundErr:
			fmt.Println("Not Found err :", FinalErr.Error())
		default:
			fmt.Println("unknown err", err.Error())

		}

	} else {
		fmt.Println("sukses")
	}

}
