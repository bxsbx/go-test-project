package main

import (
	"log"
	"os"
	"text/template"
)

// @Title 查询资源的知识点名(供服务端使用)
// @Description 查询资源的知识点名(供服务端使用)
// @Param  group_id    path  int  true  "Group ID"
// @Param  account_id  path  int  true  "Account ID"
// @router /resources [get]
func main() {
	type P struct {
		Name string
	}

	// Prepare some data to insert into the template.
	type Recipient struct {
		//Arrays     []string
		Arrays     []P
		Name, Gift string
		Attended   bool
		A, B       int
	}
	//list := []string{"12", "csa", "veve", "236788"}
	list := []P{P{"12"}, P{"csa"}, P{"veve"}, P{"236788"}}
	var recipients = []Recipient{
		{list, "effe", "fe", false, 18, 18},
		//{nil, "Aunt Mildred", "bone china tea set", true, 19, 10},
		//{"", "moleskin pants", false},
	}

	// Create a new template and parse the letter into it.
	t := template.Must(template.ParseFiles("H:\\GoProject\\StandardProject\\template\\file\\range.tmpl"))

	// Execute the template for each recipient.
	open, _ := os.Create("H:\\GoProject\\StandardProject\\test.txt")

	for _, r := range recipients {
		err := t.Execute(open, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}
}
