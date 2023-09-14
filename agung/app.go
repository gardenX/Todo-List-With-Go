package agung

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	add        = "A"
	remove     = "R"
	check      = "C"
	list       = "M"
	finished   = "F"
	onProgress = "O"
	help       = "H"
	finish     = "X"
)

func Main() {
	var (
		command string
	)

	fmt.Println("Simple Todo")
	fmt.Println("-------------------")

	showHelp()

	reader := bufio.NewReader(os.Stdin)
	myTask := NewList()

	for {
		fmt.Print("Command : ")
		command, _ = reader.ReadString('\n')

		trimInput(&command)
		filterCommand(&command)

		switch command {
		case list:
			showTasks(myTask)
		case help:
			showHelp()
		case add:
			addOperation(reader, &myTask)
		case remove:
			removeOperation(reader, &myTask)
		case check:
			setStatus(reader, &myTask)
		case finished:
			showTasks(Finished(myTask))
		case onProgress:
			showTasks(OnProgress(myTask))
		case finish:
			closeOperation()
		default:
			invalidMessage("invalid command given")
		}
	}
}

func setStatus(reader *bufio.Reader, myTask *[]Todo) {
	showTasks(*myTask)

	index, err := getIntInput(reader, "Task number you want to check : ")
	target := index - 1

	if err != nil || outOfBonds(*myTask, target) {
		invalidMessage("invalid task number given")
		return
	}

	ChangeStatus(*myTask, target)
	showTasks(*myTask)
}

func removeOperation(reader *bufio.Reader, myTask *[]Todo) {
	showTasks(*myTask)
	index, err := getIntInput(reader, "Task number you want to remove : ")
	target := index - 1

	if err != nil || outOfBonds(*myTask, target) {
		invalidMessage("invalid task number given")
		return
	}

	Remove(myTask, target)
	showTasks(*myTask)
}

func outOfBonds(task []Todo, i int) bool {
	return len(task)-1 < i
}

func invalidMessage(messages ...interface{}) {
	fmt.Println(messages...)
	fmt.Println()
}

func getIntInput(reader *bufio.Reader, message string) (index int, err error) {
	fmt.Print(message)
	target, _ := reader.ReadString('\n')
	trimInput(&target)
	index, err = strconv.Atoi(target)

	return
}

func closeOperation() {
	fmt.Println("Program closed")
	os.Exit(0)
}

func addOperation(reader *bufio.Reader, list *[]Todo) {
	listRef := *list

	if len(listRef) >= 10 {
		fmt.Println("unable to add more task")
		return
	}

	fmt.Print("New task : ")
	task, _ := reader.ReadString('\n')
	trimInput(&task)
	Add(list, task)
	showTasks(*list)
}

func showTasks(list []Todo) {
	if len(list) == 0 {
		fmt.Println("Task is empty")
		fmt.Println()
		return
	}

	fmt.Println("My tasks")

	for i, item := range list {
		if item.finish {
			fmt.Println(fmt.Sprintf("%d. [X] %s", i+1, item.Content))
		} else {
			fmt.Println(fmt.Sprintf("%d. [ ] %s", i+1, item.Content))
		}
	}

	fmt.Println()
}

func filterCommand(command *string) {
	*command = strings.ToUpper(*command)
}

func trimInput(command *string) {
	*command = strings.Trim(strings.Trim(*command, "\n"), " ")
}

func showHelp() {
	commands := []string{
		"Supported commands :",
		"(M) my task",
		"(A) add task",
		"(R) remove task",
		"(C) change task status",
		"(F) filter finished task",
		"(O) filter on progress task",
		"(H) show command",
		"(X) close program",
		"",
	}

	for _, command := range commands {
		fmt.Println(command)
	}
}
