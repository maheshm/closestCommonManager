package lca

import "testing"

func TestSearch(t *testing.T) {
	var one, two, three, four, five, six, seven *Data
	one = &Data{Name: "1"}
	two = &Data{Name: "2"}
	three = &Data{Name: "3"}
	four = &Data{Name: "4"}
	five = &Data{Name: "5"}
	six = &Data{Name: "6"}
	seven = &Data{Name: "7"}

	one.Children = append(one.Children, two)
	one.Children = append(one.Children, three)
	two.Children = append(two.Children, four)
	two.Children = append(two.Children, five)
	three.Children = append(three.Children, six)
	three.Children = append(three.Children, seven)

	Init(one)
	ancestor := Search("3", "4")
	if ancestor != "1" {
		t.Errorf("Search was incorrect, got: %s, want %s", ancestor, "1")
	}

	Init(one)
	ancestor = Search("5", "4")
	if ancestor != "2" {
		t.Errorf("Search was incorrect, got: %s, want %s", ancestor, "2")
	}

	Init(one)
	ancestor = Search("7", "1")
	if ancestor != "1" {
		t.Errorf("Search was incorrect, got: %s, want %s", ancestor, "1")
	}
}
