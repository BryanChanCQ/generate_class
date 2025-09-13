package ypd

func GetQueryWrapper(className string) string {
    return "Query" + className + "Wrapper.java";
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