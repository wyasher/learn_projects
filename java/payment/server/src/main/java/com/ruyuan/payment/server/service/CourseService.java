package com.ruyuan.payment.server.service;

import com.github.pagehelper.PageHelper;
import com.github.pagehelper.PageInfo;
import com.ruyuan.payment.server.domain.Course;
import com.ruyuan.payment.server.domain.CourseExample;
import com.ruyuan.payment.server.mapper.CourseMapper;
import com.ruyuan.payment.server.request.CourseQueryRequest;
import com.ruyuan.payment.server.response.CourseQueryResponse;
import com.ruyuan.payment.server.response.PageResponse;
import com.ruyuan.payment.server.util.CopyUtil;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;
import org.springframework.util.ObjectUtils;

import javax.annotation.Resource;
import java.util.List;

@Service
@Slf4j
public class CourseService {
    @Resource
    private CourseMapper courseMapper;

    public int count() {
        return ((int) courseMapper.countByExample(null));
    }


    public PageResponse<CourseQueryResponse> queryList(CourseQueryRequest request) {
        CourseExample courseExample = new CourseExample();
        CourseExample.Criteria criteria = courseExample.createCriteria();
        if (!ObjectUtils.isEmpty(request.getName())){
            criteria.andNameLike("%" + request.getName() + "%");
        }
        PageHelper.startPage(request.getPage(), request.getSize());
        List<Course> courses = courseMapper.selectByExample(courseExample);
        PageInfo<Course> pageInfo = new PageInfo<>(courses);
        List<CourseQueryResponse> courseQueryResponses = CopyUtil.copyList(courses, CourseQueryResponse.class);
        PageResponse<CourseQueryResponse> pageResponse = new PageResponse<>();
        pageResponse.setTotal(pageInfo.getTotal());
        pageResponse.setList(courseQueryResponses);
        return pageResponse;
    }
}
