package com.wy.leetcode;

import java.util.Arrays;

// 有序数组 每一项重复元素最多只有2个
public class RemoveDuplicatesSortedArray_80 {
    public void removeDuplicates(int[] nums) {
        int length = nums.length;
        int cursor = 2;
        int current = 1;
        while (cursor < length) {
            if (nums[cursor] != nums[cursor - 2]) {
                nums[++current] = nums[cursor];
            }
            cursor++;
        }
        System.out.println(current);
        System.out.println(Arrays.toString(nums));

    }

    public static void main(String[] args) {
        int[] nums = {0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 2, 3, 3};
        RemoveDuplicatesSortedArray_80 test = new RemoveDuplicatesSortedArray_80();
        test.removeDuplicates(nums.clone());
    }
}

