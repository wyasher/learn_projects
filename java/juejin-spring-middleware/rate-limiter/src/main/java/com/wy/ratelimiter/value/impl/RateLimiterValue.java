package com.wy.ratelimiter.value.impl;

import com.alibaba.fastjson.JSON;
import com.google.common.util.concurrent.RateLimiter;
import com.wy.ratelimiter.Constants;
import com.wy.ratelimiter.annotation.DoRateLimiter;
import com.wy.ratelimiter.value.IValueService;
import org.aspectj.lang.ProceedingJoinPoint;

import java.lang.reflect.Method;

public class RateLimiterValue implements IValueService {
    @Override
    public Object access(ProceedingJoinPoint jp, Method method, DoRateLimiter doRateLimiter, Object[] args) throws Throwable {
//        0 为不开启限流
        if (0D == doRateLimiter.permitsPerSecond()) {
            return jp.proceed();
        }
        String className = jp.getTarget().getClass().getName();
        String methodName = method.getName();

        String key = className + "." + methodName;
        if (null == Constants.rateLimiterMap.get(key)) {
            Constants.rateLimiterMap.put(key, RateLimiter.create(doRateLimiter.permitsPerSecond()));
        }
        RateLimiter rateLimiter = Constants.rateLimiterMap.get(key);
//        未被限流 执行
        if (rateLimiter.tryAcquire()) {
            return jp.proceed();
        }
//        限流
        return JSON.parseObject(doRateLimiter.returnJson(), method.getReturnType());
    }
}
