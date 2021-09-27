// +build !darwin

package direct

import (
	"fmt"

	"github.com/trevor403/gostream/pkg/input"
)

func HandlePtr(ev input.MouseEvent) error {
	fmt.Println("sending ptr", ev)
	return nil
}

func HandleKey(ev input.KeyEvent) error {
	fmt.Println("sending key", ev)
	return nil
}
