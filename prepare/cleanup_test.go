package st

import "testing"

func Test_CleanUp(t *testing.T) {
	str1 := "   Hello,    World! 124  2 This 23is Go project,		It is amazing  . I am excited!    Are you too?    "
	r1 := CleanUp(str1)

	if r1 != "Hello, World! 124 2 This 23is Go project, It is amazing . I am excited! Are you too?" {
		t.Errorf("CleanUp(str1) should return \"Hello, World! 124 2 This 23is Go project, It is amazing . I am excited! Are you too?\": %#v", r1)
	}

	str2 := "		Hello World!	This is Go 		"
	r2 := CleanUp(str2)

	if r2 != "Hello World! This is Go" {
		t.Errorf("CleanUp(str2) should return \"Hello World! This is Go\": %#v", r2)
	}
}
