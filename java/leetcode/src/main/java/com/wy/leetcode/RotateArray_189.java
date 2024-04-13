package com.wy.leetcode;

import java.util.Arrays;

/**
 * 给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
 */
public class RotateArray_189 {
    private void rotate1(int[] nums, int k) {
        int n = nums.length;
        int[] newArr = new int[n];
        for (int i = 0; i < n; ++i) {
            newArr[(i + k) % n] = nums[i];
        }
        System.arraycopy(newArr, 0, nums, 0, n);
        System.out.println(Arrays.toString(newArr));
    }
    private void rotate2(int[] nums, int k){
        int n = nums.length;
        k = k % n;
        int count = gcd(k, n);
        for (int start = 0; start < count; ++start) {
            int current = start;
            int prev = nums[start];
            do {
                int next = (current + k) % n;
                int temp = nums[next];
                nums[next] = prev;
                prev = temp;
                current = next;
            } while (start != current);
        }
        System.out.println(Arrays.toString(nums));
    }
    private void rotate3(int[] nums, int k) {
        k %= nums.length;
        reverse(nums,0,nums.length-1);
        reverse(nums,0,k-1);
        reverse(nums,k,nums.length-1);
        System.out.println(Arrays.toString(nums));
    }
    public void reverse(int[] nums,int start,int end){
        while (start<end){
            int temp = nums[start];
            nums[start] = nums[end];
            nums[end] = temp;
            start++;
            end--;
        }
    }
    private int gcd(int x, int y) {
        return y > 0 ? gcd(y, x % y) : x;
    }
    public static void main(String[] args) {

        int[] nums = new int[]{1,2,3,4,5,6,7,8,9};
        int k = 3;
        RotateArray_189 rotateArray169 = new RotateArray_189();
        rotateArray169.rotate1(nums.clone(), k);
        rotateArray169.rotate2(nums.clone(), k);
        rotateArray169.rotate3(nums.clone(), k);
    }


}
