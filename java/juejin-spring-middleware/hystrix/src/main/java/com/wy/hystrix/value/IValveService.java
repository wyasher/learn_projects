package com.wy.hystrix.value;

import com.wy.hystrix.annotation.DoHystrix;
import org.aspectj.lang.ProceedingJoinPoint;

import java.lang.reflect.Method;

public interface IValveService {
    Object access(ProceedingJoinPoint jp, Method method,
                  DoHystrix doHystrix, Object[] args) throws Throwable;
}
