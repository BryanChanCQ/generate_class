/*
 * 代码生成工具
 * 用于根据模板生成Java类文件
 * 支持YPd和MDD两种框架
 * 作者: BryanChanCQ
 * 日期: 2026-01-29
 */
package main

import (
	"log"
	"os"
	"runtime"
	"strings"
	"text/template"

	"github.com/BryanChanCQ/generate-class/internal/template/mdd"
	"github.com/BryanChanCQ/generate-class/internal/template/ypd"
	"github.com/BryanChanCQ/generate-class/internal/types"
)



const (
	RepositoryEnum = iota
	QueryWrapperEnum
	WrapperEnum
	mdd_framework = iota
	ypd_framework
)

var ypd_map = map[int]string{
	RepositoryEnum:   ypd.Ypd_Repository_Template,
	QueryWrapperEnum: ypd.Ypd_QueryWrapper_Template,
	WrapperEnum:      ypd.Ypd_Wrapper_Template,
}

var mdd_map = map[int]string{
	RepositoryEnum:   mdd.Mdd_Repository_Template,
	QueryWrapperEnum: mdd.Mdd_QueryWrapper_Template,
	WrapperEnum:      mdd.Mdd_Wrapper_Template,
}


func ConsumerFramework(param string) (map[int]string, int) {
	switch param {
	case "ypd":
		return ypd_map, ypd_framework
	case "mdd":
		return mdd_map, mdd_framework
	default:
		return nil, -1
	}
}
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

func GetNameFunc(name int, framework_type int) types.NameFunc {
	if framework_type == mdd_framework {
		switch name {
		case RepositoryEnum:
			return mdd.GetRepository
		case QueryWrapperEnum:
			return mdd.GetQueryWrapper
		case WrapperEnum:
			return mdd.GetWrapper
		default:
			return nil
		}
	} else {
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
}
// main 函数是程序的入口点
// 解析命令行参数并根据指定的框架（YPd或MDD）生成相应的Java类文件
// 参数格式: ./generc [framework] [entityName1] [entityName2] ...
func main() {
	args := os.Args
	var flag bool
	var param []string
	if flag, param = DefaultParam(args); !flag {
		log.Fatalf("请输入entity类名")
	}
	class_struct := &ClassStruct{}
	var frameworkMap map[int]string
	var framework_type int
	for i, className := range param[1:] {
		if i == 0 {
			if frameworkMap, framework_type = ConsumerFramework(args[1]); frameworkMap == nil {
				log.Fatal("请输入framework参数:ypd/mdd")
				return
			}
			continue
		}
		class_struct.ClassName = className
		for k, v := range frameworkMap {
			tMpl, err := template.New("template").Parse(v)
			if err != nil {
				panic(err)
			}
			// 执行模板并写入到文件
			f, err := os.Create(GetNameFunc(k, framework_type)(runtime.GOOS, className))
			if err != nil {
				panic(err)
			}
			defer f.Close()

			if err = tMpl.Execute(f, class_struct); err != nil {
				panic(err)
			}
		}
	}

}

type ClassStruct struct {
	ClassName string
}
