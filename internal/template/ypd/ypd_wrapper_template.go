package ypd
func GetWrapper(className string) string {
    return className + "Wrapper.java";
}
const Ypd_Wrapper_Template = `package com.yonyou.ucf.mdf.wrapper.updateOrInsert;
import com.yonyou.ucf.mdf.bill.entity.{{ .ClassName }};
import com.yonyou.ucf.mdf.wrapper.abs.AbsWrapper;

public class {{ .ClassName }}Wrapper extends AbsWrapper<{{ .ClassName }}> {
    @Override
    public Class<{{ .ClassName }}> getClazz() {
        return {{ .ClassName }}.class;
    }
}
`