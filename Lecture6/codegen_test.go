package main

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

func TestReadJSONFromFile(t *testing.T) {

	file, err := os.Open("C:/Users/aruke/OneDrive/Рабочий стол/justCode/Lecture6/user_data.json")
	if err != nil {
		t.Fatalf("Failed to open JSON file: %v", err)
	}
	defer file.Close()

	// Читаем данные из файла и десериализуем их
	var user User
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&user)
	if err != nil {
		t.Fatalf("Failed to decode JSON: %v", err)
	}

	// Проверяем, что данные были правильно десериализованы
	expectedUser := User{
		ID:      222,
		Name:    "Aruzhan",
		Surname: "Keulenzhanova",
		Age:     19,
		Address: UserAddress{
			City: City{
				Name:    "Almaty",
				ZipCode: "f12",
			},
			Region: Region{
				Name: "Center",
				Code: "f22",
			},
			Street: "Zhiber zholy",
		},
	}

	if !reflect.DeepEqual(user, expectedUser) {
		t.Errorf("Unexpected user data: %+v", user)
	}
}
