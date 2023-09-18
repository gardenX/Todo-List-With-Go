package agung

import "testing"

func TestNewList(t *testing.T) {
	list := NewList()

	if len(list) != 0 {
		t.Error("list not empty")
	}
}

func TestAdd(t *testing.T) {
	var list []Todo
	inputs := []string{"Foo", "Bar", "Baz"}

	for _, input := range inputs {
		Add(&list, input)
	}

	if len(list) != len(inputs) {
		t.Error("failed to add item to a list")
	}
}

func TestFinished(t *testing.T) {
	list := []Todo{
		{Content: "Foo"},
		{Content: "Bar", finish: true},
	}

	finished := Finished(list)

	if len(finished) != 1 || finished[0].Content == "Foo" {
		t.Error("failed to filter finished list")
	}
}

func TestOnProgress(t *testing.T) {
	list := []Todo{
		{Content: "Foo"},
		{Content: "Bar", finish: true},
	}

	onProgress := OnProgress(list)

	if len(onProgress) != 1 || onProgress[0].Content == "Bar" {
		t.Error("failed to filter on progress list")
	}
}

func TestChangeStatus(t *testing.T) {
	todo := []Todo{
		{Content: "Foo"},
		{Content: "Bar", finish: true},
	}

	ChangeStatus(todo, 0)
	ChangeStatus(todo, 1)

	if !todo[0].finish {
		t.Error("failed to set true on todo item")
	}

	if todo[1].finish {
		t.Error("failed to set false on todo item")
	}
}

func TestRemove(t *testing.T) {
	list := []Todo{
		{Content: "Foo"},
		{Content: "Bar"},
		{Content: "Baz"},
	}

	Remove(&list, 1)

	if len(list) != 2 || list[0].Content != "Foo" || list[1].Content != "Baz" {
		t.Error("unable to remove item from the list")
	}
}
