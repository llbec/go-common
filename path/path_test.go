package path

import (
	"fmt"
	"testing"
)

func Test_HomePath(t *testing.T) {
	home, err := Home()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(home)
}
