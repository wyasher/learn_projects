package com.wy.whitelist;

import com.alibaba.fastjson.JSON;
import com.wy.whitelist.annotation.DoWhiteList;
import org.apache.commons.beanutils.BeanUtils;
import org.aspectj.lang.JoinPoint;
import org.aspectj.lang.ProceedingJoinPoint;
import org.aspectj.lang.Signature;
import org.aspectj.lang.annotation.Around;
import org.aspectj.lang.annotation.Aspect;
import org.aspectj.lang.annotation.Pointcut;
import org.aspectj.lang.reflect.MethodSignature;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Component;

import javax.annotation.Resource;
import java.lang.reflect.Method;

@Aspect
@Component
public class DoJoinPoint {
    private final Logger logger = LoggerFactory.getLogger(DoJoinPoint.class.getName());

    @Resource
    private String whitelistConfig;

    @Pointcut("@annotation(com.wy.whitelist.annotation.DoWhiteList)")
    public void aopPoint() {
    }

    @Around("aopPoint()")
    public Object doRouter(ProceedingJoinPoint jp) throws Throwable {
        Method method = getMethod(jp);
        DoWhiteList whiteList = method.getAnnotation(DoWhiteList.class);
        String keyValue = getFiledValue(whiteList.key(), jp.getArgs());
        logger.info("middleware whitelist handler method：{} value：{}", method.getName(), keyValue);
        if (null == keyValue || keyValue.isEmpty()) return jp.proceed();
        // 白名单放行
        String[] split = whitelistConfig.split(",");
        for (String str : split) {
            if (keyValue.equals(str)) {
                return jp.proceed();
            }
        }
        // 非白名单，拦截
        return returnObject(whiteList, method);
    }

    private Method getMethod(JoinPoint jp) throws NoSuchMethodException {
        Signature sig = jp.getSignature();
        MethodSignature methodSignature = (MethodSignature) sig;
        return jp.getTarget().getClass().getMethod(methodSignature.getName(), methodSignature.getParameterTypes());
    }

    private String getFiledValue(String filed, Object[] args) {
        String filedValue = null;
        for (Object arg : args) {
            try {
                if (null == filedValue || filedValue.isEmpty()) {
                    filedValue = BeanUtils.getProperty(arg, filed);
                } else {
                    break;
                }

            } catch (Exception e) {
                if (args.length == 1) {
                    return args[0].toString();
                }
            }
        }

        return filedValue;
    }

    private Object returnObject(DoWhiteList whitelist, Method method) throws InstantiationException, IllegalAccessException {
        Class<?> returnType = method.getReturnType();
//        获取返回的json数据
        String returnJson = whitelist.returnJson();
        if (returnJson.isEmpty()){
//            返回空
            return  returnType.newInstance();
        }
        return JSON.parseObject(returnJson, returnType);
    }
}
