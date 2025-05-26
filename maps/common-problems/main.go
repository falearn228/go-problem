package main

import (
	"fmt"
)

// Задача: Система управления пользователями с reference semantics

// задача для изучения тонкостей работы с ссылками в мапах
// Условие:

// Создайте систему управления пользователями, которая демонстрирует различные аспекты работы с ссылками на ключи и значения в мапах.
// Структуры данных:

type User struct {
	ID         int
	Name       string
	Email      string
	Active     bool
	LoginCount int
}

type UserKey struct {
	Department string
	Role       string
}

func demonstratePointerKeyProblem() {
	// TODO: Создайте map[*UserKey]User
	// Попробуйте добавить пользователя с ключом &UserKey{"IT", "Developer"}
	// Затем попробуйте найти его по "тому же" ключу &UserKey{"IT", "Developer"}
	// Объясните, почему это не работает
	usersmap := make(map[*UserKey]User)
	value := User{
		1,
		"Eugene",
		"r228@gmail.com",
		false,
		3,
	}
	usersmap[&UserKey{"IT", "Developer"}] = value

	// Поиск по ключу: user value -> {0   false 0
	fmt.Printf("user key -> %v\n", &UserKey{"IT", "Developer"})
	fmt.Printf("user value -> %v\n", usersmap[&UserKey{"IT", "Developer"}])
	// поиск не работает, потому что мы храним ссылку на объект, хоть значение
	// в структуре тоже самое, но адрес разные
	// решение -> завести переменную
	// key := &UserKey{"IT", "Developer"}
	// и использовать его
}

func demonstrateStructKeys() {
	// TODO: Создайте map[UserKey]User (без указателя!)
	// Добавьте несколько пользователей
	// Покажите, что поиск работает корректно
	usersmap := make(map[UserKey]User)

	value1 := User{
		1,
		"Eugene",
		"r228@gmail.com",
		false,
		3,
	}
	value2 := User{
		2,
		"Ivan",
		"r229@gmail.com",
		false,
		2,
	}
	value3 := User{
		3,
		"Mike",
		"r230@gmail.com",
		true,
		3,
	}
	usersmap[UserKey{"IT", "Developer"}] = value1
	usersmap[UserKey{"IT-com", "Developer"}] = value2
	usersmap[UserKey{"IT-com", "Head"}] = value3

	fmt.Printf("user key -> %v\n", UserKey{"IT-com", "Developer"})
	fmt.Printf("user value -> %v\n", usersmap[UserKey{"IT-com", "Developer"}])
}

func modifyMapValues(users map[int]*User, userID int) {
	// TODO: Получите указатель на пользователя из мапы
	// Измените его поля (например, увеличьте LoginCount)
	// Покажите, что изменения отражаются в исходной мапе
	userptr := users[userID]

	userptr.LoginCount += 999
}

func tryModifyMapValuesCopy(users map[int]User, userID int) {
	// TODO: Получите пользователя по значению из мапы
	// Попробуйте изменить его поля
	// Покажите, что изменения НЕ отражаются в исходной мапе
	userval := users[userID]

	userval.LoginCount += 999
}

func getUserReference(users map[int]*User, userID int) (*User, bool) {
	// TODO: Безопасно получите ссылку на пользователя
	// Используйте comma-ok syntax для проверки существования
	// Верните указатель и флаг существования
	user, ok := users[userID]
	if !ok {
		return user, ok
	}
	return user, ok
}

func updateAllUsers(users map[int]*User, updateFunc func(*User)) {
	// TODO: Пройдите по всем пользователям в мапе
	// Примените updateFunc к каждому пользователю
	// Убедитесь, что изменения сохраняются

	for _, v := range users {
		updateFunc(v)
	}
}

func demonstrateIterationTrap() {
	// TODO: Создайте мапу пользователей
	// Попробуйте удалить неактивных пользователей во время range-итерации
	// Покажите правильный способ сделать это

	value1 := &User{
		1,
		"Eugene",
		"r228@gmail.com",
		false,
		3,
	}
	value2 := &User{
		2,
		"Ivan",
		"r229@gmail.com",
		false,
		2,
	}
	value3 := &User{
		3,
		"Mike",
		"r230@gmail.com",
		true,
		3,
	}

	usersmap := make(map[int]*User)
	usersmap[0] = value1
	usersmap[1] = value2
	usersmap[2] = value3

	fmt.Println("length users map before ->", len(usersmap))

	for k, v := range usersmap {
		if !v.Active {
			delete(usersmap, k)
		}
	}
	fmt.Println("length users map after ->", len(usersmap))

}

func modifyMapInFunction(users map[int]*User) {
	// TODO: Добавьте нового пользователя в мапу
	// Измените существующего пользователя
	// Покажите, что изменения видны в вызывающей функции
	value1 := &User{
		1,
		"Eugene",
		"r228@gmail.com",
		false,
		3,
	}
	users[0] = value1
}

func replaceMapInFunction(users map[int]*User) {
	// TODO: Попробуйте заменить всю мапу на новую
	// Покажите, что это НЕ отражается в вызывающей функции
	usersmap := make(map[int]*User)

	users = usersmap
}

func main() {
	fmt.Println("=== 0. Проблема с указателями как ключи ===")
	demonstratePointerKeyProblem()
	fmt.Println("\n=== 1. Правильное использование структур ===")
	demonstrateStructKeys()

	fmt.Println("\n=== 2. Изменение значений ===")
	users := map[int]*User{
		1: {ID: 1, Name: "Alice", LoginCount: 0},
		2: {ID: 2, Name: "Bob", LoginCount: 5},
	}

	fmt.Printf("До: %+v\n", users[1])
	modifyMapValues(users, 1)
	fmt.Printf("После: %+v\n", users[1])

	fmt.Println("=== 3. Тестирование tryModifyMapValuesCopy ===")
	usersCopy := map[int]User{
		1: {ID: 1, Name: "Alice", Email: "alice@test.com", Active: true, LoginCount: 5},
		2: {ID: 2, Name: "Bob", Email: "bob@test.com", Active: false, LoginCount: 10},
	}

	fmt.Printf("До модификации: %+v\n", usersCopy[1])
	tryModifyMapValuesCopy(usersCopy, 1) // Должно попытаться увеличить LoginCount
	fmt.Printf("После модификации: %+v\n", usersCopy[1])
	fmt.Println("Ожидание: LoginCount НЕ изменился (копирование по значению)")

	fmt.Println("\n=== 4. Тестирование getUserReference ===")
	usersPtr := map[int]*User{
		1: {ID: 1, Name: "Alice", Email: "alice@test.com", Active: true, LoginCount: 5},
		2: {ID: 2, Name: "Bob", Email: "bob@test.com", Active: false, LoginCount: 10},
	}

	// Тест существующего пользователя
	user, exists := getUserReference(usersPtr, 1)
	fmt.Printf("Пользователь 1 найден: %t, данные: %+v\n", exists, user)

	// Тест несуществующего пользователя
	user, exists = getUserReference(usersPtr, 999)
	fmt.Printf("Пользователь 999 найден: %t, данные: %v\n", exists, user)

	fmt.Println("\n=== 5. Тестирование updateAllUsers ===")
	fmt.Println("До обновления:")
	for id, user := range usersPtr {
		fmt.Printf("  ID %d: LoginCount = %d\n", id, user.LoginCount)
	}

	// Функция для увеличения счетчика логинов
	incrementLogin := func(u *User) {
		u.LoginCount++
		fmt.Printf("  Обновляю пользователя %s: LoginCount = %d\n", u.Name, u.LoginCount)
	}

	updateAllUsers(usersPtr, incrementLogin)

	fmt.Println("После обновления:")
	for id, user := range usersPtr {
		fmt.Printf("  ID %d: LoginCount = %d\n", id, user.LoginCount)
	}

	fmt.Println("\n=== 6. Тестирование demonstrateIterationTrap ===")
	demonstrateIterationTrap()

	fmt.Println("\n=== 7. Тестирование modifyMapInFunction ===")
	testUsers := map[int]*User{
		1: {ID: 1, Name: "Alice", Email: "alice@test.com", Active: true, LoginCount: 5},
	}

	fmt.Printf("До вызова функции: %d пользователей\n", len(testUsers))
	fmt.Printf("Alice LoginCount: %d\n", testUsers[1].LoginCount)

	modifyMapInFunction(testUsers) // Должно добавить пользователя и изменить Alice

	fmt.Printf("После вызова функции: %d пользователей\n", len(testUsers))
	if len(testUsers) > 1 {
		fmt.Printf("Alice LoginCount: %d\n", testUsers[1].LoginCount)
		fmt.Printf("Новый пользователь: %+v\n", testUsers[2]) // Предполагаем ID=2
	}

	fmt.Println("\n=== 8. Тестирование replaceMapInFunction ===")
	// originalMap := testUsers
	originalLen := len(testUsers)

	fmt.Printf("До замены мапы: %d пользователей\n", originalLen)
	fmt.Printf("Адрес мапы до: %p\n", testUsers)

	replaceMapInFunction(testUsers)

	fmt.Printf("После замены мапы: %d пользователей\n", len(testUsers))
	fmt.Printf("Адрес мапы после: %p\n", testUsers)
	// fmt.Printf("Мапа та же? %t\n", originalMap == testUsers)
	fmt.Println("Ожидание: количество пользователей НЕ изменилось (мапа не заменилась)")

	fmt.Println("\n=== 9. Демонстрация правильной замены мапы ===")
	fmt.Println("Для замены мапы нужно передавать указатель на мапу:")

	// Пример того, как ПРАВИЛЬНО заменить мапу
	replaceMapCorrectly := func(users *map[int]*User) {
		*users = map[int]*User{
			100: {ID: 100, Name: "NewUser", Email: "new@test.com", Active: true, LoginCount: 0},
		}
		fmt.Println("  Мапа заменена внутри функции")
	}

	fmt.Printf("До правильной замены: %d пользователей\n", len(testUsers))
	replaceMapCorrectly(&testUsers)
	fmt.Printf("После правильной замены: %d пользователей\n", len(testUsers))
	if len(testUsers) == 1 {
		for id, user := range testUsers {
			fmt.Printf("  Единственный пользователь: ID=%d, Name=%s\n", id, user.Name)
		}
	}
}
