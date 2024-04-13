package com.wy.whitelist.annotation;

import java.lang.annotation.*;

@Retention(RetentionPolicy.RUNTIME)
@Target(ElementType.METHOD)
@Inherited
public @interface DoWhiteList {
    String key() default "";
    String returnJson() default "";
}
