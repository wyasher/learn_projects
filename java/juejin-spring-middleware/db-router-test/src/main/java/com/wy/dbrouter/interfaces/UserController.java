package com.wy.dbrouter.interfaces;

import com.wy.dbrouter.infrastructure.dao.IUserDao;
import com.wy.dbrouter.infrastructure.po.User;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;

@RestController
public class UserController {
    private Logger logger = LoggerFactory.getLogger(UserController.class);

    @Resource
    private IUserDao userDao;

    @RequestMapping(path = "/api/queryUserInfoById", method = RequestMethod.GET)
    public User queryUserInfoById(String userId){
        return userDao.queryUserInfoByUserId(new User(userId));
    }


}
