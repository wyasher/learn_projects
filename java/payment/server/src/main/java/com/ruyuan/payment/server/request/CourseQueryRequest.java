package com.ruyuan.payment.server.request;

import lombok.Data;
import lombok.EqualsAndHashCode;

@EqualsAndHashCode(callSuper = true)
@Data
public class CourseQueryRequest extends PageRequest{

    private String name;


}
