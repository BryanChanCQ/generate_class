package mdd
func GetWrapper(className string) string {
    return className + "Wrapper.java";
}
const Mdd_Wrapper_Template = `package com.yonyou.mm.plmpd.app.wrapper.update;

import org.imeta.orm.base.BizObject;

import com.yonyou.mm.plmpd.app.psdm.util.constant.EntityFullConstants;
import com.yonyou.mm.plmpd.app.wrapper.abs.AbsWrapper;

public class DevPartWrapper extends AbsWrapper<BizObject> {
    @Override
    public String entityName() {
        //  请修改为对应实体的fullname
        return EntityFullConstants;
    }
}
`