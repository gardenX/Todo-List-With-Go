package agung

type Todo struct {
	Content string
	finish  bool
}

func NewList() (list []Todo) {
	return
}

func Add(list *[]Todo, content string) {
	*list = append(*list, Todo{Content: content})
}

func Remove(list *[]Todo, index int) {
	listRef := *list
	*list = append(listRef[:index], listRef[index+1:]...)
}

func Finished(list []Todo) []Todo {
	return filterStatus(list, true)
}

func OnProgress(list []Todo) []Todo {
	return filterStatus(list, false)
}

func ChangeStatus(list []Todo, index int) {
	list[index].finish = !list[index].finish
}

func filterStatus(list []Todo, finish bool) []Todo {
	var newList []Todo

	for _, todo := range list {
		if todo.finish == finish {
			newList = append(newList, todo)
		}
	}

	return newList
}
