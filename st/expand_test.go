package st

import "testing"

func Test_ExpandApostrophe(t *testing.T) {
	str := "was't weren't aren't isn't I'd you'd I'll he doesn't Don't I won't I'm You're i've what's up it's I didn't"
	r := ExpandApostrophe(str)
	rc := "was not were not are not is not i would you would i will he does not do not i will not i am you are i have what is up it is i did not"

	if r != rc {
		t.Errorf("ExpandApostrophe(str) should return\n%v\nbut\n%v", rc, r)
	}
}
