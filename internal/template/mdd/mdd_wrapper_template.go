package mdd

import "github.com/BryanChanCQ/generate-class/internal/types"
func GetWrapper(os types.OperateSystem, className string) string {
    if os == types.Linux {
        return  MddMainPath + MainPath + wrapperPackagePath + className + "Wrapper.java";
    } else {
        return  className + "Wrapper.java";
    }
}
const wrapperPackagePath = "com/yonyou/mm/plmpd/app/wrapper/update/"
const Mdd_Wrapper_Template = `package com.yonyou.mm.plmpd.app.wrapper.update;

import org.imeta.orm.base.BizObject;

import com.yonyou.mm.plmpd.app.psdm.util.constant.EntityFullConstants;
import com.yonyou.mm.plmpd.app.wrapper.abs.AbsWrapper;

public class {{.ClassName}}Wrapper extends AbsWrapper<BizObject> {
    @Override
    public String entityName() {
        //  请修改为对应实体的fullname
        // return EntityFullConstants;
    }
}
`