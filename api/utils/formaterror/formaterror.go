package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {

	if strings.Contains(err, "OrderBookName") {
		return errors.New("OrderBookName Already Taken")
	}

	return errors.New("An error happened that hasn't been defined yet")
}
