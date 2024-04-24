package com.wy.xredis.service;

import com.wy.xredis.annotation.XRedis;

@XRedis
public interface IRedisService {
    String get(String key);

    void set(String key,String val);
}
