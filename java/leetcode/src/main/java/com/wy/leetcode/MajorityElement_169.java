package com.wy.leetcode;

import java.util.*;
import java.util.stream.Collectors;

/**
 * 给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
 * <p>
 * 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
 */
public class MajorityElement_169 {
    public void majorityElement1(int[] nums) {
        int length = nums.length;
        assert length > 0;
        if (length == 1) {
            System.out.println(nums[0]);
            return;
        }
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < length; i++) {
            map.put(nums[i], map.getOrDefault(nums[i], 0) + 1);
        }
        Map.Entry<Integer, Integer> major = null;
        for (Map.Entry<Integer, Integer> entry : map.entrySet()) {
            if (major == null || major.getValue() < entry.getValue()) {
                major = entry;
            }
        }
        System.out.println(major.getKey());

    }

    public void majorityElement2(int[] nums) {
        Arrays.sort(nums);
        System.out.println(nums[nums.length / 2]);
    }

    // 随机找一个元素验证
    public void majorityElement3(int[] nums) {
        Random random = new Random();
        int length = nums.length;
        int majorityCount = length / 2;
        while (true) {
            int randomValue = random.nextInt(nums.length);
            int randomCount = 0;
            for (int i = 0; i < length; i++) {
                if (nums[i] == nums[randomValue]) {
                    randomCount++;
                }
            }
            if (randomCount > majorityCount) {
                System.out.println(nums[randomValue]);
                return;
            }
        }
    }

    // 分治法
    public void majorityElement4(int[] nums) {
        int major = majorityElementRec(nums, 0, nums.length - 1);
        System.out.println(major);

    }

    private int majorityElementRec(int[] nums, int left, int right) {
        if (left == right) {
            return nums[left];
        }
        int mid = (right - left) / 2 + left;
        int leftMajor = majorityElementRec(nums, left, mid);
        int rightMajor = majorityElementRec(nums, mid + 1, right);
        if (leftMajor == rightMajor) {
            return leftMajor;
        }
        int leftCount = countInRange(nums, left, mid, leftMajor);
        int rightCount = countInRange(nums, mid + 1, right, rightMajor);

        return leftCount > rightCount ? leftMajor : rightMajor;
    }

    private int countInRange(int[] nums, int left, int right, int target) {
        int count = 0;
        for (int i = left; i <= right; i++) {
            if (nums[i] == target) {
                count++;
            }
        }
        return count;
    }
    // Boyer-Moore 投票算法
    // 众数 +1 非 众数-1 遍历 总数为0 重新设置众数
    public void majorityElement5(int[] nums) {
        int target = nums[0];
        int count = 0;
        for (int num : nums) {
            // 总数为0 时，代表之前所有数相加等于0 故重新设置众数
            // 因为众数数量比其他数量多 这里设置的众数最终终会时正确的
            if (count == 0){
                target = num;
            }
            count += (num == target)?1:-1;
        }
        System.out.println(target);
    }


    public static void main(String[] args) {
        int[] nums = new int[]{1, 7, 7, 7, 7, 2, 2, 2, 2, 6, 6, 7, 7, 7, 6, 6, 7, 7, 7, 7, 7, 7, 6, 6, 7, 7, 7};
        MajorityElement_169 m = new MajorityElement_169();
        m.majorityElement1(nums.clone());
        m.majorityElement2(nums.clone());
        m.majorityElement3(nums.clone());
        m.majorityElement4(nums.clone());
        m.majorityElement5(nums.clone());
    }
}
