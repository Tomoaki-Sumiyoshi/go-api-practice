package todo

var store = TodoStore{
	todos:  map[int]Todo{},
	nextID: 0,
}

func Create(title string) Todo {
	newTodo := Todo{
		ID:        store.nextID,
		Title:     title,
		Completed: false,
	}
	store.todos[store.nextID] = newTodo
	store.nextID++

	return newTodo
}

func FindAll() []Todo {
	todoList := []Todo{}

	for _, value := range store.todos {
		todoList = append(todoList, value)
	}

	return todoList
}

func FindByID(id int) (Todo, bool) {
	todo, ok := store.todos[id]

	return todo, ok
}

func Update(id int, title string, completed bool) (Todo, bool) {
	_, ok := store.todos[id]
	if !ok {
		return Todo{}, false
	}

	newTodo := Todo{
		ID:        id,
		Title:     title,
		Completed: completed,
	}
	store.todos[id] = newTodo

	return newTodo, true
}

func Delete(id int) bool {
	_, ok := store.todos[id]
	if !ok {
		return false
	}

	delete(store.todos, id)
	return true
}
