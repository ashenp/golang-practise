package LC1

func jump(nums []int) int {
	step := 0
	n := len(nums)
	for i := 0; i < n; {
		farest := 0
		next := i + 1
		step++
		if i+nums[i] >= n-1 {
			break
		}
		for j := i + 1; j <= i+nums[i] && j < n; j++ {
			if j+nums[j] > farest {
				farest = j + nums[j]
				next = j
			}
			if farest >= n-1 {
				break
			}
		}
		i = next
	}
	return step

}

/**
作者：Alice
链接：https://zhuanlan.zhihu.com/p/369899951
来源：知乎
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

class Solution {
     public int jump(int[] N) {
        int len = N.length - 1;

        // previous的下标是根据每次跳跃后到达的下标值来更新
        int previous = -1;

        // 可到达的最远下标位置
        // 每次 拿 上次能到达的位置 和这次如果跳跃 能到达的位置, 比较，选取大值赋给reach
        int reach = 0;

        // 目前跳跃的次数
        int ans = 0;

        for (int i = 0; reach < len; i++) {

            // 需要跳跃，所以if里面ans自增
            if (i > previous) {
                ans++;
                previous = reach;
            }

            // 目前到达的最远下标位置reach
            // 但在i <= previous的情况下，让i继续for循环，来不断尝试更新reach的最大值
            // 因为题意说，你可以选择 跳跃 不大于这个跳跃数的任何数
            reach = Math.max(reach, N[i] + i);

        }
        return ans;
    }
}
*/
