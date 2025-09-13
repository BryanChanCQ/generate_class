package main

import (
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/BryanChanCQ/generate-class/internal/template/ypd"
)

const (
	RepositoryEnum = iota
	QueryWrapperEnum
	WrapperEnum
)

var ypd_map = map[int]string{
	RepositoryEnum:   ypd.Ypd_Repository_Template,
	QueryWrapperEnum: ypd.Ypd_QueryWrapper_Template,
	WrapperEnum:      ypd.Ypd_Wrapper_Template,
}

type NameFunc = func(string) string

func DefaultParam(param []string) (bool, []string) {
	if len(param) == 1 && strings.Contains(param[0], "debug") {
		return true, []string{param[0], "Debug"}
	}
	if len(param) == 1 && !strings.Contains(param[0], "debug") {
		return false, param
	}
	if len(param) > 1 {
		return true, param
	}
	return false, nil
}

func GetNameFunc(name int) NameFunc {
	switch name {
	case RepositoryEnum:
		return ypd.GetRepository
	case QueryWrapperEnum:
		return ypd.GetQueryWrapper
	case WrapperEnum:
		return ypd.GetWrapper
	default:
		return nil
	}
}
func main() {
	args := os.Args
	var flag bool
	var param []string
	if flag, param = DefaultParam(args); !flag {
		log.Fatalf("请输入entity类名")
	}
	ypd_struct := &YpdStruct{}
	for _, className := range param[1:] {
		ypd_struct.ClassName = className
		for k, v := range ypd_map {
			tMpl, err := template.New("111").Parse(v)
			if err != nil {
				panic(err)
			}
			// 执行模板并写入到文件
			f, err := os.Create(GetNameFunc(k)(className))
			if err != nil {
				panic(err)
			}
			defer f.Close()

			if err = tMpl.Execute(f, ypd_struct); err != nil {
				panic(err)
			}
		}
	}

}

type YpdStruct struct {
	ClassName string
}
