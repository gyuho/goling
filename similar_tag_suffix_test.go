package goling

import "testing"

func Test_TagSimilarity(t *testing.T) {
	// similar tags (154.30334996209191)
	tag1 := []string{"go", "golang"}
	tag2 := []string{"go"}

	// similar tags (8301.898583861128)
	tag3 := []string{"asp.net-mvc", "asp.net-mvc-3", "asp.net-mvc-4"}
	tag4 := []string{"asp.net-mvc-3", "asp.net-mvc-4", "asp.net-mvc-5", "html-helper"}
	if TagSimilarity(tag1, tag2) > TagSimilarity(tag3, tag4) {
		t.Errorf("Error: tag3, tag4 should be more similar than tag1, tag2")
	}

	// un-similar tags
	tag5 := []string{"regression", "suppressor", "control"}
	tag6 := []string{"php", "html", "forms", "post", "if-statement"}
	if TagSimilarity(tag5, tag6) != 39.28371006591931 {
		t.Errorf("Error: tag5, tag6, TagSimilarity")
	}
}
