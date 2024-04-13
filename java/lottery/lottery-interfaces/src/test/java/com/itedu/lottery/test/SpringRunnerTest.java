package com.itedu.lottery.test;

import com.alibaba.fastjson.JSON;
import com.itedu.lottery.common.Constants;
import com.itedu.lottery.domain.award.model.req.GoodsReq;
import com.itedu.lottery.domain.award.model.res.DistributionRes;
import com.itedu.lottery.domain.award.service.factory.DistributionGoodsFactory;
import com.itedu.lottery.domain.award.service.goods.IDistributionGoods;
import com.itedu.lottery.domain.strategy.model.req.DrawReq;
import com.itedu.lottery.domain.strategy.model.res.DrawResult;
import com.itedu.lottery.domain.strategy.model.vo.DrawAwardVO;
import com.itedu.lottery.domain.strategy.service.draw.IDrawExec;
import com.itedu.lottery.infrastructure.dao.IActivityDao;
import com.itedu.lottery.infrastructure.po.Activity;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

import javax.annotation.Resource;
import java.util.Date;

@RunWith(SpringRunner.class)
@SpringBootTest
public class SpringRunnerTest {
    private Logger logger = LoggerFactory.getLogger(SpringRunnerTest.class);

    @Resource
    private IActivityDao activityDao;
    @Resource
    private IDrawExec drawExec;

    @Resource
    private DistributionGoodsFactory distributionGoodsFactory;

    @Test
    public void test_drawExec() {
        drawExec.doDrawExec(new DrawReq("小傅哥", 10001L));
        drawExec.doDrawExec(new DrawReq("小佳佳", 10001L));
        drawExec.doDrawExec(new DrawReq("小蜗牛", 10001L));
        drawExec.doDrawExec(new DrawReq("八杯水", 10001L));
    }

    @Test
    public void test_insert() {
        Activity activity = new Activity();
        activity.setActivityId(100001L);
        activity.setActivityName("测试活动");
        activity.setActivityDesc("仅用于插入数据测试");
        activity.setBeginDateTime(new Date());
        activity.setEndDateTime(new Date());
        activity.setStockCount(100);
        activity.setTakeCount(10);
        activity.setState(0);
        activity.setCreator("xiaofuge");
        activityDao.insert(activity);
    }


}
