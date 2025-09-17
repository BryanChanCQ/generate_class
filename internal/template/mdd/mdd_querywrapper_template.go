package mdd

func GetQueryWrapper(className string) string {
    return className + "Query" + "Wrapper.java";
}
const Mdd_QueryWrapper_Template = `package com.yonyou.mm.plmpd.app.wrapper.query;

import com.yonyou.mm.plmpd.app.psdm.util.constant.EntityFullConstants;
import com.yonyou.mm.plmpd.app.wrapper.abs.AbsQueryWrapper;

public class {{.ClassName}}QueryWrapper extends AbsQueryWrapper {

    @Override
    public String getEntityName() {
        //  请修改为对应实体的fullname
        return EntityFullConstants;
    }

    @Override
    public String getGroupName() {
        //  请修改为对应实体的groupname
        // return "IMP-PLM-PSDM";
    }

}
`