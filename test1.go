package main

import (
	"os"
	"strings" // 导入 strings 包以使用 strings.Title
	"text/template"
)

// 定义一个结构体来存储模板中的变量
type JavaClass struct {
	PackageName string
	ClassName   string
	Fields      []Field
}

// Field 结构体表示类中的字段
type Field struct {
	Type string
	Name string
}

func main1() {
	// 准备数据模型
	fields := []Field{
		{Type: "String", Name: "name"},
		{Type: "int", Name: "age"},
	}
	javaClass := JavaClass{
		PackageName: "com.example",
		ClassName:   "Person",
		Fields:      fields,
	}

	// 定义模板，并注册自定义函数
	funcMap := template.FuncMap{
		"title": strings.Title, // 注册 title 函数
	}
	tmpl, err := template.New("111").Funcs(funcMap).Parse(javaTemplate)
	if err != nil {
		panic(err)
	}

	// 执行模板并写入到文件
	f, err := os.Create("Person.java")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err = tmpl.Execute(f, javaClass); err != nil {
		panic(err)
	}
}

const javaTemplate = `
package {{.PackageName}};
@Data
public class {{.ClassName}} {
{{range .Fields}}
    private {{.Type}} {{.Name}};
{{end}}
{{range .Fields}}
    public {{.Type}} get{{title .Name}}() {
        return this.{{.Name}};
    }
    public void set{{title .Name}}({{.Type}} {{.Name}}) {
        this.{{.Name}} = {{.Name}};
    }
{{end}}
}
`
