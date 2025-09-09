package main

import (
	"github.com/BryanChanCQ/generate-class/internal/template/ypd"
	"log"
	"os"
	"strings"
	"text/template"
)

const (
	RepositoryEnum = iota
	QuerywrapperEnum
	WrapperEnum
)

var ypd_map = map[int]string{
	RepositoryEnum:   ypd.Ypd_Repository_Template,
	QuerywrapperEnum: ypd.Ypd_Querywrapper_Template,
	WrapperEnum:      ypd.Ypd_Wrapper_Template,
}

type NameFunc = func(string) string

func GetNameFunc(name int) NameFunc {
	switch name {
	case RepositoryEnum:
		return ypd.GetRepository
	case QuerywrapperEnum:
		return ypd.GetQuerywrapper
	case WrapperEnum:
		return ypd.GetWrapper
	default:
		return nil
	}
}
func main() {
	args := os.Args
	if len(args) == 1 {
		log.Fatalf("请输入entity类名")
	}
	ypd_struct := &YpdStrct{}
	funcMap := template.FuncMap{
		"title": strings.Title, // 注册 title 函数
	}
	for _, className := range args[1:] {
		ypd_struct.ClassName = className
		for k, v := range ypd_map {
			tmpl, err := template.New("111").Funcs(funcMap).Parse(v)
			if err != nil {
				panic(err)
			}
			// 执行模板并写入到文件
			f, err := os.Create(GetNameFunc(k)(className))
			if err != nil {
				panic(err)
			}
			defer f.Close()

			if err = tmpl.Execute(f, ypd_struct); err != nil {
				panic(err)
			}
		}
	}

}

type YpdStrct struct {
	ClassName string
}
