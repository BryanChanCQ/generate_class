package ypd

import "github.com/BryanChanCQ/generate-class/internal/types"
const YpdMainPath = "/home/bryanchancq/work_code/yonbip-zskpd/"
const MainPath = "/zskpd-be/dev-zskpd-service/src/main/java/"
const queryPackagePath = "com/yonyou/ucf/mdf/wrapper/query/"
func GetQueryWrapper(os types.OperateSystem, className string) string {
    if os == types.Linux {
        return YpdMainPath + MainPath + queryPackagePath + "Query" + className + "Wrapper.java";
    } else {
        return "Query" + className + "Wrapper.java";
    }
}
const Ypd_QueryWrapper_Template = `package com.yonyou.ucf.mdf.wrapper.query;
import com.yonyou.ucf.mdf.bill.entity.{{ .ClassName }};
import com.yonyou.ucf.mdf.wrapper.abs.AbsQueryWrapper;

public class Query{{ .ClassName }}Wrapper extends AbsQueryWrapper<{{ .ClassName }}> {
    @Override
    public String getEntityName() {
        return {{ .ClassName }}.ENTITY_NAME;
    }
}
`