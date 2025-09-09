package ypd

func GetRepository(className string) string{
    return className + "Repository.java";
}

const Ypd_Repository_Template = `package com.yonyou.ucf.mdf.repository;
import java.util.List;
import java.util.Map;

import org.imeta.orm.schema.QueryCondition;
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
	public AbsQueryWrapper queryWrapper() {
		return new Query{{ .ClassName }}Wrapper();
	}
    
    @Override
    public AbsWrapper<{{ .ClassName }}> changeWrapper() {
        return new {{ .ClassName }}Wrapper();
    }
}
`