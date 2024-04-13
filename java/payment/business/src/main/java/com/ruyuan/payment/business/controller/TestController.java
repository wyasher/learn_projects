package com.ruyuan.payment.business.controller;

import com.ruyuan.payment.server.service.CourseService;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;

@RestController
@RequestMapping("test")
public class TestController {
    @Resource
    private CourseService courseService;
    @GetMapping("hello")
    public String test() {
        return "test" + courseService.count();
    }
}
