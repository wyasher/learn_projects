package com.wy.whitelist;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication(scanBasePackages = {"com.wy.*"})
public class WhitelistTestApplication {
    public static void main(String[] args)
    {
        // 启动SpringBoot项目
        SpringApplication.run(WhitelistTestApplication.class, args);
    }



}
