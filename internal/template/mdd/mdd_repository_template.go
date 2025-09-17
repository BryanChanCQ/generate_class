package mdd
func GetRepository(className string) string{
    return className + "Repository.java";
}

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

}
`