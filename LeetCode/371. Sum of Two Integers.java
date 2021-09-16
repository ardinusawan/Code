// https://leetcode.com/problems/sum-of-two-integers/
class Solution {
    public int getSum(int x, int y) {
        while (y != 0)
        {
            int carry = x & y;
            System.out.println(carry);
            x = x ^ y;
            System.out.println(x);
            y = carry << 1;
            System.out.println(y);
            System.out.println("");
        }
        return x;
    }
}
