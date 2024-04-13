package com.itedu.lottery.application.process.res;

import com.itedu.lottery.common.Result;

public class RuleQuantificationCrowdResult extends Result {
    /** 活动ID */
    private Long activityId;

    public RuleQuantificationCrowdResult(String code, String info) {
        super(code, info);
    }

    public Long getActivityId() {
        return activityId;
    }

    public void setActivityId(Long activityId) {
        this.activityId = activityId;
    }
}

