package ipcheck

import (
	"fmt"
	"testing"
)

func TestIPCheck(t *testing.T) {
	i := NewIPCheck(false)
	i.Logger(`testing logger output`)
	i.Start()
	i.Wait()
	fmt.Printf("IPCheck returned: %s\n", i.Response)
}
