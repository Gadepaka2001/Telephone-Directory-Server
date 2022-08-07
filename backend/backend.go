package backend

import (
	"io/ioutil"
	"os"
	"strings"
)

func ReadFile() string {
	File := "data.txt"
	path := "./backend/Data/" + File
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	return string(data)
}

func AppendIntoFile(data string) {
	File := "data.txt"
	path := "./backend/Data/" + File
	S := ReadFile() + data + "\n"
	err := ioutil.WriteFile(path, []byte(S), 0644)
	if err != nil {
		panic(err.Error())
	}
}

func Search(name string) (string, bool) {
	S := ReadFile()
	Lst := strings.Split(S, "\n")
	for i := 0; i < len(Lst)-1; i++ {
		list := strings.Split(Lst[i], ":")
		if list[0] == name {
			return "Contact Name: " + list[0] + "\nMobile Number: " + list[1], true
		}
	}
	return " ", false
}

func WriteIntoFile(data string, B bool) {
	File := "data.txt"
	path := "./backend/Data/" + File
	if B == true {
		err := ioutil.WriteFile(path, []byte(data+"\n"), 0644)
		if err != nil {
			panic("Error while opening a file")
		}
	} else {
		err := ioutil.WriteFile(path, []byte(data), 0644)
		if err != nil {
			panic("Error while opening a file")
		}
	}
}

func Delete(C string) bool {
	S := ReadFile()
	Lst := strings.Split(S, "\n")
	for i := 0; i < len(Lst)-1; i++ {
		list := strings.Split(Lst[i], ":")
		if list[0] == C {
			Lst = append(Lst[:i], Lst[i+1:]...)
			S = strings.Join(Lst, "\n")
			WriteIntoFile(S, false)
			return true
		}
	}
	return false
}

func CreateFile() {
	File := "data.txt"
	path := "./Data/" + File
	_, err := os.Create(path)
	if err != nil {
		panic("Error while creating a file")
	}
}

func Check(File string, name string) bool {
	S := ReadFile()
	Lst := strings.Split(S, "\n")
	for i := 0; i < len(Lst)-1; i++ {
		lst := strings.Split(Lst[i], ":")
		if name == lst[0] {
			return true
		}
	}
	return false
}
