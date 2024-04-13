package com.wy.leetcode;

import java.util.Arrays;

public class RemoveDuplicatesSortedArray_26 {

    public void removeDuplicates1(int[] nums) {
        int length = nums.length;
        int current = 0;
        int cursor = 1;
        while (cursor < length) {
            if (nums[current] != nums[cursor]) {
                nums[++current] = nums[cursor];
            }
            cursor++;
        }
        System.out.println(current);
        System.out.println(Arrays.toString(nums));
    }

    public static void main(String[] args) {
        int[] nums = new int[]{1, 1, 2, 3, 3, 4, 5,5,5,6};
        RemoveDuplicatesSortedArray_26 app = new RemoveDuplicatesSortedArray_26();
        app.removeDuplicates1(nums.clone());
    }
}
