package com.itedu.lottery.domain.strategy.service.draw;

import com.itedu.lottery.domain.strategy.model.req.DrawReq;
import com.itedu.lottery.domain.strategy.model.res.DrawResult;

public interface IDrawExec {
    DrawResult doDrawExec(DrawReq req);
}
