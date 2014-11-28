package utils

import "testing"

func TestAllTabIntoSingleSpace(t *testing.T) {
	str := "Hello	World	!	Hello"
	r := AllTabIntoSingleSpace(str)
	if r != "Hello World ! Hello" {
		t.Errorf("AllTabIntoSingleSpace(str) should return \"Hello World ! Hello\": %#v", r)
	}
}

func TestAllSpaceIntoSingleTab(t *testing.T) {
	str := "Hello World! Hello"
	r := AllSpaceIntoSingleTab(str)

	if r != "Hello	World!	Hello" {
		t.Errorf("AllSpaceIntoSingleTab(str) should return \"Hello	World!	Hello\": %#v", r)
	}
}

func TestTabToSpace(t *testing.T) {
	str := "Hello	World	Hello"
	r := TabToSpace(str)

	if r != "Hello World Hello" {
		t.Errorf("TabToSpace(str) should return \"Hello World Hello\": %#v", r)
	}
}

func TestSpaceToTab(t *testing.T) {
	str := "Hello World Hello"
	r := SpaceToTab(str)

	if r != "Hello	World	Hello" {
		t.Errorf("SpaceToTab(str) should return \"Hello	World	Hello\": %#v", r)
	}
}

func TestDeleteSpace(t *testing.T) {
	str := "Hello World	Hello"
	r := DeleteSpace(str)

	if r != "HelloWorldHello" {
		t.Errorf("DeleteSpace(str) should return \"HelloWorldHello\": %#v", r)
	}
}

func TestDeleteTab(t *testing.T) {
	str := "Hello	World Hello"
	r := DeleteTab(str)

	if r != "HelloWorld Hello" {
		t.Errorf("DeleteTab(str) should return \"HelloWorld Hello\": %#v", r)
	}
}
