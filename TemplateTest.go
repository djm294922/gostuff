package main
	import "text/template"
	import "os"
	import "io/ioutil"

func main() {
	type Inventory struct {
		Material string
		Count    uint
	}

	println("TemplateTest starts")

	contents, err := ioutil.ReadFile("jeff.html.template")
	if err != nil { panic(err) }
	println("contents len=", len(contents))

	sweaters := Inventory{"wool", 17}
	//tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	tmpl, err := template.New("test").Parse(string(contents))
	if err != nil { panic(err) }
		err = tmpl.Execute(os.Stdout, sweaters)
	println("");
	if err != nil { panic(err) }

}
