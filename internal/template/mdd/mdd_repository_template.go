package mdd

import "github.com/BryanChanCQ/generate-class/internal/types"
func GetRepository(os types.OperateSystem, className string) string{
    if os == types.Linux {
        return  MddMainPath + MainPath + repoPackagePath + className + "Repository.java";
    } else {
        return className + "Repository.java";
    }
}
const repoPackagePath = "com/yonyou/mm/plmpd/app/psdm/repo/"
const Mdd_Repository_Template = `package com.yonyou.mm.plmpd.app.psdm.repo;

import org.imeta.orm.base.BizObject;
import org.springframework.stereotype.Repository;

import com.yonyou.mm.plmpd.app.psdm.repo.abs.AbsRepository;
import com.yonyou.mm.plmpd.app.wrapper.abs.AbsQueryWrapper;
import com.yonyou.mm.plmpd.app.wrapper.abs.AbsWrapper;
import com.yonyou.mm.plmpd.app.wrapper.query.{{.ClassName}}QueryWrapper;
import com.yonyou.mm.plmpd.app.wrapper.update.{{.ClassName}}Wrapper;

@Repository
public class {{.ClassName}}Repository extends AbsRepository {

    @Override
    protected AbsQueryWrapper getQueryWrapper() {
        return new {{.ClassName}}QueryWrapper();
    }

    @Override
    protected AbsWrapper<BizObject> getChangeWrapper() {
        return new {{.ClassName}}Wrapper();
    }
        
    @Override
    protected boolean ownDomain() {
        return true;
    }
}
`