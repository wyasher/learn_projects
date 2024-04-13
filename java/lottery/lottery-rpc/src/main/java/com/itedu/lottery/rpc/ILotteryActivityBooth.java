package com.itedu.lottery.rpc;

import com.itedu.lottery.rpc.req.DrawReq;
import com.itedu.lottery.rpc.req.QuantificationDrawReq;
import com.itedu.lottery.rpc.res.DrawRes;

/**
 * 抽奖活动展台接口
 */
public interface ILotteryActivityBooth {
    /**
     * 指定活动抽奖
     * @param drawReq 请求参数
     * @return        抽奖结果
     */
    DrawRes doDraw(DrawReq drawReq);

    /**
     * 量化人群抽奖
     * @param quantificationDrawReq 请求参数
     * @return                      抽奖结果
     */
    DrawRes doQuantificationDraw(QuantificationDrawReq quantificationDrawReq);

}
