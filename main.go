package main

import "example.com/go_todoapp/app/models"

func main() {
	// u := models.User{}
	// u.Name = "テスト2"
	// u.Email = "テスト2メール"
	// u.Password = "テスト2パス"
	// u.CreateUser()
	// user, _ := models.GetUser(7)
	// fmt.Println(user)

	// u.Name = "テスト1"
	// u.Email = "テスト1メール"
	// u.Password = "テスト1パス"
	// fmt.Println(u)

	// u.CreateUser()

	// u := &models.User{}
	// user, _ := models.GetUser(3)
	// user.CreateTodo("３番目のTodo")
	// user2, _ := models.GetUser(3)
	// todos, _ := user2.GetTodosByUser()

	// for _, v := range todos {
	// 	fmt.Println(v)
	// }
	t, _ := models.GetTodo(1)
	t.Content = "更新済み1"
	t.UpdateTodo()
}
