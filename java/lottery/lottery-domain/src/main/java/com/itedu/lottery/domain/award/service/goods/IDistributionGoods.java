package com.itedu.lottery.domain.award.service.goods;

import com.itedu.lottery.domain.award.model.req.GoodsReq;
import com.itedu.lottery.domain.award.model.res.DistributionRes;

public interface IDistributionGoods {

    DistributionRes doDistribution(GoodsReq req);
}
