package com.wy.ratelimiter;

import com.wy.ratelimiter.annotation.DoRateLimiter;
import com.wy.ratelimiter.value.IValueService;
import com.wy.ratelimiter.value.impl.RateLimiterValue;
import org.aspectj.lang.JoinPoint;
import org.aspectj.lang.ProceedingJoinPoint;
import org.aspectj.lang.Signature;
import org.aspectj.lang.annotation.Around;
import org.aspectj.lang.annotation.Aspect;
import org.aspectj.lang.annotation.Pointcut;
import org.aspectj.lang.reflect.MethodSignature;
import org.springframework.stereotype.Component;

import java.lang.reflect.Method;
@Aspect
@Component
public class DoRateLimiterPoint {
    @Pointcut("@annotation(com.wy.ratelimiter.annotation.DoRateLimiter)")
    public void aopPoint(){

    }

    @Around("aopPoint() && @annotation(doRateLimiter)")
    public Object doRouter(ProceedingJoinPoint jp, DoRateLimiter doRateLimiter) throws Throwable{
        IValueService valueService = new RateLimiterValue();
        return valueService.access(jp,getMethod(jp),doRateLimiter,jp.getArgs());
    }
    private Method getMethod(JoinPoint jp) throws NoSuchMethodException {
        Signature sig = jp.getSignature();
        MethodSignature methodSignature = (MethodSignature) sig;
        return jp.getTarget().getClass().getMethod(methodSignature.getName(), methodSignature.getParameterTypes());
    }
}
