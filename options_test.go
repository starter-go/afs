package afs

import "testing"

func TestOptionsMaker(t *testing.T) {

	maker := OptionsMaker{}
	maker.SetFlags().Create().Excl()
	maker.SetPermissions().SetMode(7, 5, 5)
	opt := maker.Options()

	flag := opt.Flag
	perm := opt.Permission

	t.Logf("flags = %v", flag)
	t.Logf("perm = %v", perm.String())
}
