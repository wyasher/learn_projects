package com.itedu.lottery.domain.strategy.service.draw;

import com.itedu.lottery.domain.strategy.service.algorithm.IDrawAlgorithm;

import javax.annotation.PostConstruct;
import javax.annotation.Resource;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

public class DrawConfig {
    @Resource
    private IDrawAlgorithm entiretyRateRandomDrawAlgorithm;

    @Resource
    private IDrawAlgorithm singleRateRandomDrawAlgorithm;

    protected static Map<Integer,IDrawAlgorithm> drawAlgorithmGroup = new ConcurrentHashMap<>();

    @PostConstruct
    public void init(){
        drawAlgorithmGroup.put(1,entiretyRateRandomDrawAlgorithm);
        drawAlgorithmGroup.put(2,singleRateRandomDrawAlgorithm);
    }
}
