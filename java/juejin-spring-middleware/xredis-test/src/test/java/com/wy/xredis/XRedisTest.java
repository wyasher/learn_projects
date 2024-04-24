package com.wy.xredis;

import com.wy.xredis.service.IRedisService;
import com.wy.xredis.service.XRedisTestApplication;
import org.junit.jupiter.api.Test;
import org.springframework.boot.test.context.SpringBootTest;

import javax.annotation.Resource;
@SpringBootTest(classes = XRedisTestApplication.class)
public class XRedisTest {

    @Resource
    private IRedisService iRedisService;

    @Test
    public void test_set(){
        iRedisService.set("hello","world");
    }
    @Test
    public void test_get(){
        String result = iRedisService.get("hello");
        System.out.println(result);
    }

}
