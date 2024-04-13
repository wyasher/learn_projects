package com.itedu.lottery.domain.award.service.goods;

import com.itedu.lottery.domain.award.repository.IOrderRepository;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import javax.annotation.Resource;

public class DistributionBase {
    protected Logger logger = LoggerFactory.getLogger(DistributionBase.class.getName());

    @Resource
    protected IOrderRepository orderRepository;

    protected void updateUserAwardState(String uId, Long orderId, String awardId, Integer grantState) {
        orderRepository.updateUserAwardState(uId, orderId, awardId, grantState);
    }

}
