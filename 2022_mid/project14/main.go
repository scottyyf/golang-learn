package main

import "fmt"

type userInfo struct {
	name string
	password string
}

func modifyUser(users map[string]map[string]string, name string){
	value, ret := users[name]
	if ret{
		value["password"] = "888888"
	}else {
		users[name] = map[string]string{"alias":"nil",
			"password": "",
		}
	}
}


func main() {
	users := make(map[string]map[string]string)
	info := make(map[string]string)
	info["alias"] = "xiaomi"
	info["password"] = "77777777"
	users["no1"] = info

	modifyUser(users, "no1")
	modifyUser(users, "no2")
	modifyUser(users, "no3")
	fmt.Println(users)
}