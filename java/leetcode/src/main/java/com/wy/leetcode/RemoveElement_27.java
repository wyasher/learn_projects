package com.wy.leetcode;

import java.util.Arrays;

public class RemoveElement_27 {
    // 双指针
    void removeElement1(int[] nums, int val) {
        int n = nums.length;
        int left = 0;
        for (int right = 0; right < n; right++) {
            if (nums[right] != val) {
                nums[left++] = nums[right];
            }
        }
        System.out.println(left);
        System.out.println(Arrays.toString(Arrays.stream(nums).limit(left).toArray()));
    }

    // 双指针优化
    void removeElement2(int[] nums, int val) {
        int left = 0;
        int right = nums.length-1;
        while (left < right) {
            if (nums[left] == val) {
                nums[left] = nums[right--];
            }else {
                left++;
            }
        }
        System.out.println(left);
        System.out.println(Arrays.toString(Arrays.stream(nums).limit(left).toArray()));
    }

    public static void main(String[] args) {
        int[] nums = {3, 2, 3, 2, 1, 122,123,5,5,4,4,3,4,4};
        int value = 3;
        RemoveElement_27 removeElement27 = new RemoveElement_27();
        removeElement27.removeElement1(nums.clone(), value);
        removeElement27.removeElement2(nums.clone(), value);

    }
}
