package com.itedu.lottery.domain.support.ids;

/**
 * 生成id接口
 *      * 1. 雪花算法，用于生成单号
 *      * 2. 日期算法，用于生成活动编号类，特性是生成数字串较短，但指定时间内不能生成太多
 *      * 3. 随机算法，用于生成策略ID
 */
public interface IIdGenerator {
    /**
     * 获取id
     * @return id
     */
    long nextId();
}
