package com.wy.leetcode;

import java.util.Arrays;

public class MergeSortedArray_88 {

    public void merge(int[] nums1, int m, int[] nums2, int n) {
        for (int i = 0; i < n; i++) {
            nums1[m + i] = nums2[i];
        }
        Arrays.sort(nums1);
        System.out.println(Arrays.toString(nums1));
    }

    // 双指针
    // https://assets.leetcode-cn.com/solution-static/88/1.gif
    public void mergeDoublePoint(int[] nums1, int m, int[] nums2, int n) {
        int p1 = 0, p2 = 0;
        int[] sorted = new int[m + n];
        int cur;
        while (p1 < m || p2 < n) {
            if (p1 == m) {
                cur = nums2[p2++];
            } else if (p2 == n) {
                cur = nums1[p1++];
            } else if (nums1[p1] < nums2[p2]) {
                cur = nums1[p1++];
            } else {
                cur = nums2[p2++];
            }
            sorted[p1 + p2 - 1] = cur;
        }
        for (int i = 0; i < m + n; i++) {
            nums1[i] = sorted[i];
        }
        System.out.println(Arrays.toString(nums1));

    }

    // 逆双指针
    public void mergeDoublePointBack(int[] nums1, int m, int[] nums2, int n) {
        int p1 = m - 1, p2 = n - 1;
        int tail = m + n - 1;
        int cur;
        while (tail > 0) {
            if (p1 == 0) {
                cur = nums2[p2--];
            } else if (p2 == 0) {
                cur = nums1[p1--];
            } else if (nums1[p1] > nums2[p2]) {
                cur = nums1[p1--];
            } else {
                cur = nums2[p2--];
            }
            nums1[tail--] = cur;
        }
        System.out.println(Arrays.toString(nums1));
    }

    public static void main(String[] args) {
        final MergeSortedArray_88 mergeSortedArray88 = new MergeSortedArray_88();
        int[] nums1 = {1, 2, 5, 5, 0, 0};
        int[] nums2 = {2, 4};
        mergeSortedArray88.merge(nums1.clone(), nums1.length - nums2.length, nums2.clone(), nums2.length);
        mergeSortedArray88.mergeDoublePoint(nums1.clone(), nums1.length - nums2.length, nums2.clone(), nums2.length);
        mergeSortedArray88.mergeDoublePointBack(nums1.clone(), nums1.length - nums2.length, nums2.clone(), nums2.length);


    }
}
