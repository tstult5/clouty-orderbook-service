package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {

	if strings.Contains(err, "OrderName") {
		return errors.New("OrderName Already Taken")
	}

	return errors.New("An error happened that hasn't been defined yet")
}
