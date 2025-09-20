package mdd

import "github.com/BryanChanCQ/generate-class/internal/types"
const MddMainPath = "/home/bryanchancq/work_code/yonbip-mm-plmpd-orgl/"
const MainPath = "/yonbip-mm-plmpd-be/yonbip-mm-plmpd-app/src/main/java/"
const queryPackagePath = "com/yonyou/mm/plmpd/app/wrapper/query/"
func GetQueryWrapper(os types.OperateSystem, className string) string {
    if os == types.Linux {
        return MddMainPath + MainPath + queryPackagePath + className + "Query" + "Wrapper.java";
    } else {
        return className + "Query" + "Wrapper.java";
    }
}
const Mdd_QueryWrapper_Template = `package com.yonyou.mm.plmpd.app.wrapper.query;

import com.yonyou.mm.plmpd.app.psdm.util.constant.EntityFullConstants;
import com.yonyou.mm.plmpd.app.wrapper.abs.AbsQueryWrapper;

public class {{.ClassName}}QueryWrapper extends AbsQueryWrapper {

    @Override
    public String getEntityName() {
        //  请修改为对应实体的fullname
        // return EntityFullConstants;
    }

    @Override
    public String getGroupName() {
        //  请修改为对应实体的groupname
        // return "IMP-PLM-PSDM";
    }

}
`