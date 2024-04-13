package com.wy.leetcode;

/**
 * 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
 * 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
 * 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
 */
public class BestTimeToBuyAndSellStock_121 {

    public void maxProfit1(int[] nums){
        int length = nums.length;
        int currentIndex = 0;
        int maxProfit = 0;
        for (int i = 1; i < length; i++) {
            if (nums[currentIndex] > nums[i]){
                currentIndex = i;
            }else {
                maxProfit = Math.max(maxProfit, nums[i] - nums[currentIndex]);
            }
        }
        System.out.println(maxProfit);

    }
    public static void main(String[] args) {
        int[] prices = {7, 20,1, 5, 3, 6, 4,11};
        BestTimeToBuyAndSellStock_121 app = new BestTimeToBuyAndSellStock_121();
        app.maxProfit1(prices.clone());

    }
}
