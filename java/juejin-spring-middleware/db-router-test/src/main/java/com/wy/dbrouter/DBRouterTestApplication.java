package com.wy.dbrouter;

import org.mybatis.spring.annotation.MapperScan;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Configuration;

@SpringBootApplication
@Configuration
@MapperScan(basePackages = {"com.wy.dbrouter.infrastructure.dao"})
public class DBRouterTestApplication {

    public static void main(String[] args) {

        // 启动项目
        SpringApplication.run(DBRouterTestApplication.class, args);
    }
}
