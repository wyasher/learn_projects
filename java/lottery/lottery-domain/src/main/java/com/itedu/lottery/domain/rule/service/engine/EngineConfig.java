package com.itedu.lottery.domain.rule.service.engine;

import com.itedu.lottery.domain.rule.service.logic.LogicFilter;
import com.itedu.lottery.domain.rule.service.logic.impl.UserAgeFilter;
import com.itedu.lottery.domain.rule.service.logic.impl.UserGenderFilter;

import javax.annotation.PostConstruct;
import javax.annotation.Resource;
import java.util.HashMap;
import java.util.Map;

public class EngineConfig {
    protected static Map<String, LogicFilter> logicFilterMap = new HashMap<>();

    @Resource
    private UserAgeFilter userAgeFilter;
    @Resource
    private UserGenderFilter userGenderFilter;

    @PostConstruct
    public void init() {
        logicFilterMap.put("userAge", userAgeFilter);
        logicFilterMap.put("userGender", userGenderFilter);
    }
}
