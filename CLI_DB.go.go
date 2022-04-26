package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

type employee struct {
	Id     int    `json:"int"`
	Name   string `json:"Name"`
	Age    int    `json:"Age"`
	Salary int    `json:"Salary"`
}

var Employees map[int][]employee //map of slice for work of exporting data to the file

type storage interface {
	insert(id int, name string, age, salary int) error
	get(id int) error
	delete(id int) error
	write(id int) error
}

type memoryStorage struct { //DB
	data map[int]employee
}

func newMemoryStorage() *memoryStorage { //function to initialize DB and map of slice
	Employees = make(map[int][]employee)
	return &memoryStorage{
		data: make(map[int]employee),
	}
}

func (s *memoryStorage) insert(id int, name string, age, salary int) error {
	var e employee
	e = employee{id, name, age, salary}
	s.data[e.Id] = e
	Employees[e.Id] = append(Employees[e.Id], e)

	return nil
}

func (s *memoryStorage) get(id int) error {
	e, exists := s.data[id]
	if !exists {
		return errors.New("\n[Error] employee with such id doesn't exist")
	}
	fmt.Printf("\nID: %[1]d\nname: %s\nage: %[3]d\nsalary: %[4]d$", e.Id, e.Name, e.Age, e.Salary)
	return errors.New("")
}

func (s *memoryStorage) delete(id int) error {
	_, exists := s.data[id]
	if !exists {
		return errors.New("\n[Error] employee with such id doesn't exist")
	}
	delete(s.data, id)
	delete(Employees, id)

	return nil
}

func newFile(fileName string) error { //Creates json file with eror handling
	fileName += ".json"
	file, err := os.Create(fileName)
	if err != nil {
		log.Print(err)
	}
	defer file.Close()

	return nil
}

func exportData(fileName string) error { //Marshal all employees data to the json format, then writes it to the pre-creating file
	jsonData, err := json.MarshalIndent(Employees, "", " ")
	if err != nil {
		return err
	}
	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		log.Print(err)
	}
	return nil
}

func main() {
	DB := newMemoryStorage() //declares new DB

	chekcer := bufio.NewScanner(os.Stdin)

	//Program cycle
	for {
		fmt.Print("\n[Menu] choice option \n1. Add an employee to the database\n2. Get employee details\n3. Delete an employee from data base\n4. Export employee data to the pre-created file\n5. Exit\n\n- ")
		chekcer.Scan()
		option, err := strconv.Atoi(chekcer.Text())

		if err != nil {
			err = errors.New("\n[Error] invalid input")
			fmt.Println(err)
		} else {
			switch option {
			case 1:
				var (
					id, age, salary int
					name            string
				)
				fmt.Print("\nEnter emplyee ID: ")
				fmt.Scanln(&id)
				fmt.Print("Enter emplyee name: ")
				fmt.Scanln(&name)
				fmt.Print("Enter emplyee age: ")
				fmt.Scanln(&age)
				fmt.Print("Enter emplyee salary: ")
				fmt.Scanln(&salary)
				DB.insert(id, name, age, salary)
				fmt.Println("\n[Info] employee was successfully added to the database")
			case 2:
				fmt.Print("\nEnter emplyee ID: ")
				var id int
				fmt.Scanln(&id)
				fmt.Println(DB.get(id))
			case 3:
				fmt.Print("\nEnter emplyee ID: ")
				var id int
				fmt.Scanln(&id)
				DB.delete(id)
				fmt.Println("\n[Info] employee was successfully delete from database")
			case 4:
				fmt.Print("\nEnter the file name: ")
				var fileName string
				fmt.Scanln(&fileName)
				newFile(fileName)
				exportData(fileName + ".json")
				fmt.Println("\n[Info] all employees were successfully exported to a file")
			case 5:
				fmt.Print("\n[Info] enter any button to exit\n\n- ")
				fmt.Scanln()
				break
			default:
				fmt.Println("\n[Error] not an option")
			}
		}

	}
}
