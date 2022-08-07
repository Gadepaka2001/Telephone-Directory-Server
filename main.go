package main

import (
	"bytes"
	"fmt"
	"net/http"
	"rest-api/backend"
	"strings"
)

func Main(w http.ResponseWriter, r *http.Request) {
	HTMLcontent := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<title>Telephone Directory</title>
	</head>
	<body>
	<h1>Fill The Contact Details</h1>
	<form action="/done" method="POST" id="nameform">
	<label for="fname">First name:</label>
	<input type="text" id="fname" name="fname"><br><br>
	<label for="lname">Last name:</label>
	<input type="text" id="lname" name="lname">
	<label for="email">Email:</label>
	<input type="text" id="Emain" name="email">
	<label for="mobile">Mobile number:</label>
	<input type="text" id="mobile" name="mobile">
	</form>
	<button type="submit" form="nameform" value="Submit">Submit</button>
	<h1>SEARCH FOR A CONTACT</h1>
	<form action="/search" method="GET" id="search">
		<lable for="fname">Name: </label>
		<input type="text" id="name" name="name"><br>
	</form>
	<button type="submit" form="search" value="Submit">Submit</button>
	<h1>MODIFICATION</h1>
	<form action="/edit" method="put" id="edit">
		<lable for="fname">Name: </label>
		<input type="text" id="name" name="name"><br>
		<lable for="mobile">Mobile Number: </label>
		<input type="text" id="mobile" name="mobile"><br>
	</form>
	<button type="submit" form="edit" value="Submit">Submit</button>
	<h1>DELETION</h1>
	<form action="/delete" method="delete" id="delete">
		<lable for="fname">Name: </label>
		<input type="text" id="name" name="name"><br>
	</form>
	<button type="submit" form="delete" value="Submit">Submit</button>
	</body>
	</html>`
	w.Write([]byte(HTMLcontent))
}

func POST(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	newStr := buf.String()
	Params := strings.Split(newStr, "&")
	fname := strings.Split(Params[0], "=")[1]
	lname := strings.Split(Params[1], "=")[1]
	email := strings.Replace(strings.Split(Params[2], "=")[1], "%40", "@", 1)
	mobile := strings.Split(Params[3], "=")[1]
	contact := fmt.Sprintf("%v&%v&%v&%v", fname, lname, email, mobile)
	backend.AppendIntoFile(contact)

	HTMLcontent := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<title>Movies</title>
	</head>
	<body>
	<h1>Successfully Added </h1>
	<a href="http://localhost:3333/main">Main</a>
	</body>
	</html>`
	w.Write([]byte(HTMLcontent))
}

func GET(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	data := backend.ReadFile()
	contactsDeatailsInList := strings.Split(data, "\n")
	i := 0
	for ; i < len(contactsDeatailsInList); i++ {
		contactDetails := strings.Split(contactsDeatailsInList[i], "&")
		if contactDetails[0] == name {
			HTMLcontent := `<!DOCTYPE html>
			<html lang="en">
			<head>
				<title>Contact Details</title>
			</head>
			<body>
			<h4>First Name: </h4>{fstname}
			<h4>Last Name: </h4>{lstname}
			<h4>Email: </h4>{email}
			<h4>Moible: </h4>{mobile}
			<a href="http://localhost:3333/main">Main</a>
			</body>
			</html>`
			HTMLcontent = strings.Replace(HTMLcontent, "{fstname}", contactDetails[0], 1)
			HTMLcontent = strings.Replace(HTMLcontent, "{lstname}", contactDetails[1], 1)
			HTMLcontent = strings.Replace(HTMLcontent, "{email}", contactDetails[2], 1)
			HTMLcontent = strings.Replace(HTMLcontent, "{mobile}", contactDetails[3], 1)
			w.Write([]byte(HTMLcontent))
			break
		}
	}
	if i == len(contactsDeatailsInList) {
		HTMLcontent := `<!DOCTYPE html>
			<html lang="en">
			<head>
				<title>Contact Details</title>
			</head>
			<body>
			<h1>Contact Not Found</h1>
			<a href="http://localhost:3333/main">Main</a>
			</body>
			</html>`
		w.Write([]byte(HTMLcontent))
	}
}

func PUT(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	mobile := r.URL.Query().Get("mobile")
	data := backend.ReadFile()
	contactsDeatailsInList := strings.Split(data, "\n")
	i := 0
	for i < len(contactsDeatailsInList) {
		contactDetails := strings.Split(contactsDeatailsInList[i], "&")
		if contactDetails[0] == name {
			contactDetails[3] = mobile
			contactDetailsString := strings.Join(contactDetails, "&")
			contactsDeatailsInList[i] = contactDetailsString
			contactsDetailsInListStrings := strings.Join(contactsDeatailsInList, "\n")
			backend.WriteIntoFile(contactsDetailsInListStrings, false)
			HTMLcontent := `<!DOCTYPE html>
			<html lang="en">
			<head>
				<title>Contact Details</title>
			</head>
			<body>
			<h1>Contact {name} Mobile Updated to {mobile}</h1>
			<a href="http://localhost:3333/main">Main</a>
			</body>
			</html>`
			HTMLcontent = strings.Replace(HTMLcontent, "{mobile}", mobile, 1)
			HTMLcontent = strings.Replace(HTMLcontent, "{name}", name, 1)
			w.Write([]byte(HTMLcontent))
			break
		}
		i++
	}
	if i == len(contactsDeatailsInList) {
		HTMLcontent := `<!DOCTYPE html>
			<html lang="en">
			<head>
				<title>Contact Details</title>
			</head>
			<body>
			<h1>Contact Not Found</h1>
			<a href="http://localhost:3333/main">Main</a>
			</body>
			</html>`
		w.Write([]byte(HTMLcontent))
	}
}

func DELETE(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	data := backend.ReadFile()
	contactsDeatailsInList := strings.Split(data, "\n")
	i := 0
	for i < len(contactsDeatailsInList) {
		contactDetails := strings.Split(contactsDeatailsInList[i], "&")
		if contactDetails[0] == name {
			contactsDeatailsInList = append(contactsDeatailsInList[:i], contactsDeatailsInList[i+1:]...)
			contactsDetailsInListStrings := strings.Join(contactsDeatailsInList, "\n")
			backend.WriteIntoFile(contactsDetailsInListStrings, false)
			HTMLcontent := `<!DOCTYPE html>
			<html lang="en">
			<head>
				<title>Contact Details</title>
			</head>
			<body>
			<h1>Contact {name} Deleted</h1>
			<a href="http://localhost:3333/main">Main</a>
			</body>
			</html>`
			HTMLcontent = strings.Replace(HTMLcontent, "{name}", name, 1)
			w.Write([]byte(HTMLcontent))
			break
		}
		i++
	}
	if i == len(contactsDeatailsInList) {
		HTMLcontent := `<!DOCTYPE html>
			<html lang="en">
			<head>
				<title>Contact Details</title>
			</head>
			<body>
			<h1>Contact Not Found</h1>
			<a href="http://localhost:3333/main">Main</a>
			</body>
			</html>`
		w.Write([]byte(HTMLcontent))
	}
}

func main() {
	http.HandleFunc("/", Main)
	http.HandleFunc("/done", POST)
	http.HandleFunc("/search", GET)
	http.HandleFunc("/edit", PUT)
	http.HandleFunc("/delete", DELETE)
	println("Listening...")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		println(err.Error())
	}
}
