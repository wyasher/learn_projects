package com.wy.hystrix;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class HystrixTestApplication {
    public static void main(String[] args) {

        // 启动应用
        SpringApplication.run(HystrixTestApplication.class, args);
    }
}
