package com.wy.ratelimiter;

import com.google.common.util.concurrent.RateLimiter;

import java.util.Collections;
import java.util.HashMap;
import java.util.Map;

public class Constants {
    public static final Map<String, RateLimiter> rateLimiterMap = Collections.synchronizedMap(new HashMap<>());
}
