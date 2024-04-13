package com.ruyuan.payment.server.util;

public class NumUtil {

    public static Integer generate3() {
        return (int)(((Math.random() * 9) + 1) * 100);
    }

    public static Integer generate6() {
        return (int)(((Math.random() * 9) + 1) * 100000);
    }
}
