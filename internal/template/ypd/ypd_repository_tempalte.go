package ypd

import "github.com/BryanChanCQ/generate-class/internal/types"

func GetRepository(os types.OperateSystem, className string) string{
    if os == types.Linux {
        return  YpdMainPath + MainPath + repoPackagePath + className + "Repository.java";
    } else {
        return className + "Repository.java";
    }
}
const repoPackagePath = "/com/yonyou/ucf/mdf/repository/"
const Ypd_Repository_Template = `package com.yonyou.ucf.mdf.repository;
import org.springframework.stereotype.Repository;

import com.yonyou.ucf.mdf.bill.entity.{{ .ClassName }};
import com.yonyou.ucf.mdf.repository.abs.AbsRepository;
import com.yonyou.ucf.mdf.wrapper.abs.AbsQueryWrapper;
import com.yonyou.ucf.mdf.wrapper.abs.AbsWrapper;
import com.yonyou.ucf.mdf.wrapper.query.Query{{ .ClassName }}Wrapper;
import com.yonyou.ucf.mdf.wrapper.updateOrInsert.{{ .ClassName }}Wrapper;
@Repository
public class {{ .ClassName }}Repository extends AbsRepository<{{ .ClassName }}> {
    @Override
	protected AbsQueryWrapper<{{ .ClassName }}> queryWrapper() {
		return new Query{{ .ClassName }}Wrapper();
	}
    
    @Override
    protected AbsWrapper<{{ .ClassName }}> changeWrapper() {
        return new {{ .ClassName }}Wrapper();
    }
}
`