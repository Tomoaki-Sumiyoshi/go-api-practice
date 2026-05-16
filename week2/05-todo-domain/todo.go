package todo

type Todo struct {
	ID        int
	Title     string
	Completed bool
}

type TodoStore struct {
	todos  map[int]Todo
	nextID int
}
