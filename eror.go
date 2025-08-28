package main

import (
	"errors"
	"fmt"
)

type User struct {
	FirstName string
	LastName  string
	Age       int
}

func CreateUser(firstName, lastName string, age int) (*User, error) {
	// Проверка имени
	if firstName == "" {
		return nil, errors.New("имя не может быть пустым")
	}

	if lastName == "" {
		return nil, errors.New("фамилия не может быть пустой")
	}

	if age < 18 { // строго меньше 18
		return nil, errors.New("пользователь должен быть старше 18 лет")
	}

	user := &User{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}

	return user, nil
}

func main() {
	user, err := CreateUser("Arafat", "Abdukarimov", 16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Пользователь [%s %s] успешно создан\n", user.FirstName, user.LastName)
	}

	user, err2 := CreateUser("Pupkina", "Zalu", 111)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Printf("Пользователь [%s %s] успешно создан\n", user.FirstName, user.LastName)
	}
}
