package com.itedu.lottery.domain.activity.service.partake;

import com.itedu.lottery.domain.activity.model.req.PartakeReq;
import com.itedu.lottery.domain.activity.model.vo.ActivityBillVO;
import com.itedu.lottery.domain.activity.repository.IActivityRepository;

import javax.annotation.Resource;

/**
 * 活动领取模操作，一些通用的数据服务
 */
public class ActivityPartakeSupport {
    @Resource
    public IActivityRepository activityRepository;

    public ActivityBillVO queryActivityBill(PartakeReq req){
        return activityRepository.queryActivityBill(req);
    }
}
