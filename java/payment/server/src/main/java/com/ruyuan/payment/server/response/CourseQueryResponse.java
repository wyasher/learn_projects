package com.ruyuan.payment.server.response;

import lombok.Data;

import java.math.BigDecimal;
@Data
public class CourseQueryResponse {
    private Long id;

    private String name;

    private String level;

    private BigDecimal price;

    private String desc;
}