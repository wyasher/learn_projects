package com.itedu.lottery.domain.award.service.factory;

import com.itedu.lottery.common.Constants;
import com.itedu.lottery.domain.award.service.goods.IDistributionGoods;
import com.itedu.lottery.domain.award.service.goods.impl.CouponGoods;
import com.itedu.lottery.domain.award.service.goods.impl.DescGoods;
import com.itedu.lottery.domain.award.service.goods.impl.PhysicalGoods;
import com.itedu.lottery.domain.award.service.goods.impl.RedeemCodeGoods;

import javax.annotation.PostConstruct;
import javax.annotation.Resource;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

public class GoodsConfig {
    protected static Map<Integer, IDistributionGoods> goodsMap = new ConcurrentHashMap<>();


    @Resource
    private DescGoods descGoods;

    @Resource
    private RedeemCodeGoods redeemCodeGoods;

    @Resource
    private CouponGoods couponGoods;

    @Resource
    private PhysicalGoods physicalGoods;

    @PostConstruct
    public void init() {
        goodsMap.put(Constants.AwardType.DESC.getCode(), descGoods);
        goodsMap.put(Constants.AwardType.RedeemCodeGoods.getCode(), redeemCodeGoods);
        goodsMap.put(Constants.AwardType.CouponGoods.getCode(), couponGoods);
        goodsMap.put(Constants.AwardType.PhysicalGoods.getCode(), physicalGoods);
    }
}
