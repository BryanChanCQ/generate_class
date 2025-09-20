package ypd

import "github.com/BryanChanCQ/generate-class/internal/types"
func GetWrapper(os types.OperateSystem, className string) string {
    if os == types.Linux {
        return YpdMainPath + MainPath + wrapperPackagePath + className + "Wrapper.java";
    } else {
        return className + "Wrapper.java";
    }
}
const wrapperPackagePath = "com/yonyou/ucf/mdf/wrapper/updateOrInsert/"
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