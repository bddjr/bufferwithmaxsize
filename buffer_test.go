package bufferwithmaxsize_test

import (
	"fmt"
	"testing"

	"github.com/bddjr/bufferwithmaxsize"
)

func TestXxx(t *testing.T) {
	b := bufferwithmaxsize.NewBuffer(4)
	b.Write([]byte{1, 2, 3, 4, 5, 6})
	fmt.Println(b.Bytes())
	b.Write([]byte{7, 8, 9})
	fmt.Println(b.Bytes())
}
