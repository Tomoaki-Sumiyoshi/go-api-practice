package todo

import "testing"

func TestNormal(t *testing.T) {
	var todo = Todo{}
	t.Run("TestCreate", func(t *testing.T) {
		firstTitle := "CreateTodo"
		todo = Create("CreateTodo")
		if todo.Title != firstTitle {
			t.Error("Create() is incomplete")
		}
	})

	t.Run("TestFindAll", func(t *testing.T) {
		findAllTodo := FindAll()
		if (len(findAllTodo) != 1) || (todo != findAllTodo[0]) {
			t.Error("FindAll() is incomplete")
		}
	})

	t.Run("TestFindByIDAfterCreate", func(t *testing.T) {
		findTodo, ok := FindByID(todo.ID)
		if (findTodo != todo) || !ok {
			t.Error("FindByID() is incomplete")
		}
	})

	t.Run("TestUpdate", func(t *testing.T) {
		newTodo := Todo{
			ID:        todo.ID,
			Title:     "UpdateTodo",
			Completed: true,
		}
		updateTodo, ok := Update(newTodo.ID, newTodo.Title, newTodo.Completed)
		if (newTodo != updateTodo) || !ok {
			t.Error("Update() is incomplete")
		}
	})

	t.Run("TestDelete", func(t *testing.T) {
		ok := Delete(todo.ID)
		if !ok {
			t.Error("Delete() is incomplete")
		}

		ok = Delete(todo.ID)
		if ok {
			t.Error("Delete() is incomplete")
		}
	})

	t.Run("TestFindByIDAfterDelete", func(t *testing.T) {
		findTodo, ok := FindByID(todo.ID)
		if (findTodo == todo) || ok {
			t.Error("FindByID() is incomplete")
		}
	})
}
