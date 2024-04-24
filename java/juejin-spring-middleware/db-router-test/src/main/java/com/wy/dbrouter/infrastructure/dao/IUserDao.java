package com.wy.dbrouter.infrastructure.dao;

import com.wy.dbrouter.annotation.DBRouter;
import com.wy.dbrouter.infrastructure.po.User;
import org.apache.ibatis.annotations.Mapper;

@Mapper
public interface IUserDao {
    @DBRouter(key = "userId")
    User queryUserInfoByUserId(User req);

    @DBRouter(key = "userId")
    void insertUser(User req);
}
