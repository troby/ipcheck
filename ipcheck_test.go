package lib

import "testing"

func TestIPCheck(t *testing.T) {
	i := NewIPCheck(false)
	i.Start()
	i.Wait()
	i.Logger(`testing logger output`)
}
