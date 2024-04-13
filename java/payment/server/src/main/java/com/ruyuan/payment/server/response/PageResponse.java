package com.ruyuan.payment.server.response;

import lombok.Data;

import java.util.List;

@Data
public class PageResponse<T> {
    private Long total;
    private List<T> list;

}
