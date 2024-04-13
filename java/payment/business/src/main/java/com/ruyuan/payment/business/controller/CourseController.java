package com.ruyuan.payment.business.controller;

import com.ruyuan.payment.server.domain.Course;
import com.ruyuan.payment.server.request.CourseQueryRequest;
import com.ruyuan.payment.server.response.CommonResponse;
import com.ruyuan.payment.server.response.CourseQueryResponse;
import com.ruyuan.payment.server.response.PageResponse;
import com.ruyuan.payment.server.service.CourseService;
import com.ruyuan.payment.server.util.CopyUtil;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;
import java.util.List;

@RestController
@RequestMapping("course")
public class CourseController {
    @Resource
    private CourseService courseService;

    @GetMapping("query-list")
    public CommonResponse<PageResponse<CourseQueryResponse>> queryList(CourseQueryRequest courseQueryRequest) {
        CommonResponse<PageResponse<CourseQueryResponse>> commonResponse = new CommonResponse<>();
        commonResponse.setContent(courseService.queryList(courseQueryRequest));
        return commonResponse;
    }
}
