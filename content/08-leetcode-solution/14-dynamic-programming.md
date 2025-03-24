---
title: "动态规划"
date: 2022-09-03T00:43:42+08:00
tags: ["leetcode", "dp", "c++"]
image: /covers/leetcode-solution.jpg
cover: false
weight: 814
---

## Longest Palindromic Substring
[LeetCode](https://leetcode.com/problems/longest-palindromic-substring)/[力扣](https://leetcode-cn.com/problems/longest-palindromic-substring)

代码实现

```cpp
string longestPalindrome(string s) {
    if (s.size() <= 1){
        return s;
    }

    int left = 0;
    int maxLen = 0;
    for (int k = 0; k < 2 * s.size() - 1; k++){
        int i = k / 2;
        int j = i + k % 2;
        int len = 0;
        while (i >= 0 && j < s.size() && s[i] == s[j]){
            len++;
            i--;
            j++;
        }
        // len = 2 * len + 1;
        if (len > maxLen){
            left = i + 1;
            maxLen = len;
        
        }
    }

    return s.substr(left, maxLen + 1);
}
```

## Maximum Subarray
[LeetCode](https://leetcode.com/problems/maximum-subarray)/[力扣](https://leetcode-cn.com/problems/maximum-subarray)

代码实现

```cpp
int maxSubArray(vector<int>& nums) {
    int sum = INT_MIN,csum = 0, n = nums.size();
    for(int i=0;i<n;i++)
    {
        csum = max(csum+nums[i],nums[i]);
        sum = max(csum,sum);
    }
    return sum;
}
```

## Jump Game
[LeetCode](https://leetcode.com/problems/jump-game)/[力扣](https://leetcode-cn.com/problems/jump-game)

代码实现

```cpp
bool canJump(vector<int>& nums) {
    int n = nums.size();
    int dp = nums[0]; // 当前能达到的最远距离
    for(int i = 1; i < n; i++) {
        if(dp >= n - 1) return true;
        if(dp >= i && i + nums[i] > dp) dp = i + nums[i];
    }
    return dp >= n - 1;
}
```

## Unique Paths
[LeetCode](https://leetcode.com/problems/unique-paths)/[力扣](https://leetcode-cn.com/problems/unique-paths)

代码实现

```cpp
int uniquePaths(int m, int n) {
    vector<vector<int>> dp(m,vector(n,0));
    for(int i = 0; i < m; i++) {
        dp[i][0] = 1;
    }
    for(int i = 0; i < n; i++) {
        dp[0][i] = 1;
    }
    for(int i = 1; i < m; i++) {
        for(int j = 1; j < n; j++) {
            dp[i][j] = dp[i - 1][j] + dp[i][j - 1];
        }
    }
    return dp[m - 1][n - 1];
}
```

## Climbing Stairs
[LeetCode](https://leetcode.com/problems/climbing-stairs)/[力扣](https://leetcode-cn.com/problems/climbing-stairs)

代码实现

```cpp
int climbStairs(int n) {
    if(n <= 2) return n;
    vector<int> dp(n,0);
    dp[0] = 1; dp[1] = 2;
    for(int i = 2; i < n; i++) {
        dp[i] = dp[i - 1] + dp[i-2];
    }
    return dp[n-1];
}
```

## Decode Ways
[LeetCode](https://leetcode.com/problems/decode-ways)/[力扣](https://leetcode-cn.com/problems/decode-ways)

代码实现

```cpp
int numDecodings(string s) {
    int ways_1 = 0, ways_2 = 0, ways = 0; // read as ways minus 1, ways minus 2 and ways
    
    if (s.size() == 0 || s[0] == '0')
        return 0;
    
    ways_1 = ways_2 = ways = 1;
    
    for (int i = 1; i < s.size(); ++i) {
        if (s[i] == '0') {
            ways = (s[i-1] == '1' || s[i-1] == '2') ? ways_2 : 0;
        } else {
            ways = (s[i-1] == '1' || (s[i-1] == '2' && s[i] <'7')) ?
                    (ways_1 + ways_2) : ways_1;
        }
        ways_2 = ways_1;
        ways_1 = ways;
    }
    return ways;
}
```

## Best Time to Buy and Sell Stock
[LeetCode](https://leetcode.com/problems/best-time-to-buy-and-sell-stock)/[力扣](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock)

代码实现

```cpp
int maxProfit(vector<int>& prices) {
    if(prices.size() <= 1) return 0;
    int mini = 0;
    int res = 0;
    for(int i = 0; i < prices.size(); i++){
        if(prices[i] < prices[mini]) {
            mini = i;
        } else {
            res = max(res, prices[i] - prices[mini]);
        }
    }
    return res;
}
```

## Word Break
[LeetCode](https://leetcode.com/problems/word-break)/[力扣](https://leetcode-cn.com/problems/word-break)

代码实现

```cpp
bool wordBreak(string s, vector<string>& wordDict) {
    unordered_set<string> dict(wordDict.begin(), wordDict.end());
    int n = s.size();
    vector<int> dp(n + 1,0);
    s = " " + s;
    dp[0] = 1;
    for(int i = 1; i <= n; i++) {
        for(int j = 0; j < i; j++){
            if(dp[j] == 1){
                string new_str = s.substr(j+1, i - j);
                if(dict.count(new_str)){
                    dp[i] = 1;
                    break;
                }
            }
        }
    }
    return dp[n];
}
```

## Maximum Product Subarray
[LeetCode](https://leetcode.com/problems/maximum-product-subarray)/[力扣](https://leetcode-cn.com/problems/maximum-product-subarray)

除了保存当前最大值，还要保存负数的最小值，因为乘法负负得正
代码实现

```cpp
int maxProduct(vector<int>& nums) {
    int ans=nums[0];
    int imax=ans,imin=ans;
    
    for(int i=1;i<nums.size();i++){
        
        if(nums[i]<0)
            swap(imin,imax);
        
        imin=min(nums[i],imin*nums[i]);
        imax=max(nums[i],imax*nums[i]);
        
        ans=max(ans,imax);
    }
    
    return ans;
}
```

## House Robber
[LeetCode](https://leetcode.com/problems/house-robber)/[力扣](https://leetcode-cn.com/problems/house-robber)

代码实现

```cpp
int rob(vector<int>& nums) {
    int n = nums.size();
    if(n == 0) return 0;
    if(n == 1) return nums[0];
    if(n == 2) return max(nums[0],nums[1]);
    int res = 0;
    int res_1 = max(nums[0],nums[1]);
    int res_2 = nums[0];
    for(int i = 2; i < nums.size(); i++){
        res = max(res_2 + nums[i], res_1);
        res_2 = res_1;
        res_1 = res;  
    }
    return res;
}
```

## House Robber II
[LeetCode](https://leetcode.com/problems/house-robber-ii)/[力扣](https://leetcode-cn.com/problems/house-robber-ii)

代码实现

```cpp
int rob(vector<int>& nums) {
    int n = nums.size();
    if(n==0) return 0;
    if(n==1) return nums[0];
    vector<int> dp1(n+1,0), dp2(n+1,0);
    for(int i=n-1;i>0;i--){
        if(i==n-1) dp2[i]=nums[i];
        else dp2[i] = max(nums[i]+dp2[i+2],dp2[i+1]);
    }
    for(int i=n-2;i>=0;i--){
        if(i==n-2) dp1[i]=nums[i];
        else dp1[i] = max(nums[i]+dp1[i+2],dp1[i+1]);
    }
    return max(dp1[0],dp2[1]);
}
```

## Longest Increasing Subsequence
[LeetCode](https://leetcode.com/problems/longest-increasing-subsequence)/[力扣](https://leetcode-cn.com/problems/longest-increasing-subsequence)

代码实现

```cpp
int lengthOfLIS(vector<int>& nums) {
    int n = nums.size();
    if(n <= 1) return n;
    vector<int> dp(n,0);
    dp[0] = 1;
    int res = 1;
    for(int i = 1; i < n; i++){
        int temp = 0;
        for(int j = 0; j < i; j++) {
            if(nums[i] > nums[j]) {
                temp = max(dp[j],temp);
            }
            
        }
        dp[i] = temp + 1;
        res = max(res, dp[i]);
    }
    return res;
}
```

## Range Sum Query - Immutable
[LeetCode](https://leetcode.com/problems/range-sum-query-immutable)/[力扣](https://leetcode-cn.com/problems/range-sum-query-immutable)

代码实现

```cpp
class NumArray {
vector<int> sum;
NumArray(vector<int>& nums) {
    sum = calSum(nums);
}
vector<int> calSum(vector<int>& nums) {
    int n = nums.size();
    vector<int> res(n,0);
    if(nums.size() == 0) return res;
    res[0] = nums[0];
    
    for(int i = 1; i < nums.size(); i++) {
        res[i] = res[i - 1] + nums[i];
    }
    return res;
}

int sumRange(int i, int j) {
    return i == 0 ? sum[j] : sum[j] - sum[i - 1];
}
```

## Best Time to Buy and Sell Stock with Cooldown
[LeetCode](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown)/[力扣](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown)

代码实现

```cpp
int maxProfit(vector<int>& prices) {
     if(prices.size()<=1) return 0;
     int len=prices.size();
     vector<int>stock(len);
     vector<int>nostock(len);
     //Initializing
     nostock[0]=0; stock[0]=-prices[0];
     nostock[1]=max(stock[0]+prices[1],nostock[0]);  stock[1]=max(stock[0],-prices[1]);

     for(int i=2;i<len;i++)
     {
         stock[i]=max(stock[i-1],nostock[i-2]-prices[i]);
         nostock[i]=max(nostock[i-1],stock[i-1]+prices[i]);
     }
    return max(stock[len-1],nostock[len-1]);
}
```

## Coin Change
[LeetCode](https://leetcode.com/problems/coin-change)/[力扣](https://leetcode-cn.com/problems/coin-change)

代码实现

```cpp
int coinChange(vector<int>& coins, int amount) {
    vector<int> dp(amount + 1, amount + 1);
    dp[0] = 0;
    for(int i = 1; i <= amount; i++) {
        for(int j = 0; j < coins.size(); j++) {
            if(coins[j] <= i) {
                dp[i] = min(dp[i],dp[i - coins[j]] + 1);
            }
        }
    }
    return dp[amount] > amount ? -1 : dp[amount];
}
```

## Counting Bits
[LeetCode](https://leetcode.com/problems/counting-bits)/[力扣](https://leetcode-cn.com/problems/counting-bits)

代码实现

```cpp
vector<int> countBits(int num) {
    vector<int> dp(num + 1, 0);
    // int cur = 0;
    for(int i = 1; i <= num; i++) {
        if(pow(2,dp[i-1]) == i){
            dp[i] = 1;
            // cur = i;
        }else {
            dp[i] = dp[i /2 ] + dp[i - i/2];
        }
    }
    return dp;
}
```

## Combination Sum IV
[LeetCode](https://leetcode.com/problems/combination-sum-iv)/[力扣](https://leetcode-cn.com/problems/combination-sum-iv)

代码实现

```cpp
int combinationSum4(vector<int>& nums, int target) {
    vector<int> dp(target+1);
    dp[0] = 1;
    sort(nums.begin(),nums.end());
    for(int i = 1;i<=target;i++) {
        for(int num : nums) {
            if(i<num)break;
            dp[i] = ((unsigned int)dp[i]+dp[i-num]); 
        }
    }
    return dp[target];
}
```

## Partition Equal Subset Sum
[LeetCode](https://leetcode.com/problems/partition-equal-subset-sum)/[力扣](https://leetcode-cn.com/problems/partition-equal-subset-sum)

代码实现

```cpp
// If you know how to solve any of the combination sum problems (such as https://leetcode.com/problems/combination-sum-ii/),
// you'll be able to solve this very easily as well, because it's basically the same problem (only worded differently).
bool canPartition(vector<int>& nums)
{
    int sum = 0;
    for (int num : nums)
        sum += num;

	// Because we have to find two subsets with equal sums, we can only do so if the total sum of the array is even.
	// If the total sum is not even, we return false.
    if ((sum & 1))
        return false;
    
	// If the sum is even, we have to find in the array a combination of numbers that add up to its half.
	// If we find such a combination, we return true, else we return false.
    std::unordered_map<std::string, bool> memo;
    return helper(nums, 0, sum >> 1, memo);
 }

bool helper(const std::vector<int>& nums, int index, int sum, std::unordered_map<std::string, bool>& memo)
{
    if (sum == 0)
        return true;

    if (sum < 0)
        return false;

    bool possible = false;

    for (int i = index; i < nums.size(); ++i)
    {
        std::string key = std::to_string(i) + "-" + std::to_string(sum);

        if (memo.find(key) != memo.end())
            return memo[key];

        bool res = helper(nums, i + 1, sum - nums[i], memo);
       
        memo.emplace(key, res);
        
        possible |= res;

        if (possible)
            break;
    }

    return possible;
}
```

## Palindromic Substrings
[LeetCode](https://leetcode.com/problems/palindromic-substrings)/[力扣](https://leetcode-cn.com/problems/palindromic-substrings)

代码实现

```cpp
int countSubstrings(string s) {
    int n = s.size(), res = 0;
    for(int i = 0; i <= 2*n - 1; i++) {
        int left = i / 2;
        int right = left + i % 2;
        while(left >= 0 && right < n && s[left] == s[right]) {
            res++;
            left--;
            right++;
        }
    }
    return res;
}
```

## Number of Longest Increasing Subsequence
[LeetCode](https://leetcode.com/problems/number-of-longest-increasing-subsequence)/[力扣](https://leetcode-cn.com/problems/number-of-longest-increasing-subsequence)

代码实现

```cpp
int findNumberOfLIS(vector<int>& nums) {
    int n = nums.size();
    if(n <= 1) return n;
    vector<int> lengths(n,0);
    vector<int> counts(n,1);
    
    for(int j = 0; j < n; j++){
        for(int i = 0; i < j; i++) {
            if(nums[i] < nums[j])
            if(lengths[i] >= lengths[j]){
                lengths[j] = lengths[i] + 1;
                counts[j] = counts[i];
            }else if(lengths[i] + 1 == lengths[j]){
                counts[j] += counts[i];
            }
        }
    }
    
    int longest = 0, ans = 0;
    for(int length : lengths) {
        longest = max(length,longest);
    }
    
    for(int i = 0; i < n; i++) {
        if(lengths[i] == longest) {
            ans += counts[i];
        }
    }
    return ans;
}
```

## Partition to K Equal Sum Subsets
[LeetCode](https://leetcode.com/problems/partition-to-k-equal-sum-subsets)/[力扣](https://leetcode-cn.com/problems/partition-to-k-equal-sum-subsets)

代码实现

```cpp
//currsum is the current sum of subset we are working on and targetsum is the sum we require for each subset
    
bool dfs(vector<int>& nums,vector<int>visited,int idx,int k,int currsum,int targetsum)
{
    if(k==1) return true;                                               //if k==1 then all the subsets have been found so return true.
    if(currsum==targetsum) return dfs(nums,visited,0,k-1,0,targetsum);  //this condition means you have found one suset so start from begining for another subset.
    for(int i=idx ; i<nums.size() ; i++)
    {
        if(!visited[i])                                                //if the index is not visited then it can be used in the current subset or bucket.
        {
            visited[i]=true;                                                          //set this index as used to avoid redundancy.
            if(dfs(nums,visited,i+1,k,currsum+nums[i],targetsum)) return true;        //explore the choices
            visited[i]=false;                                                         //for backtrack.
        }
    }
    return false;
}
bool canPartitionKSubsets(vector<int>& nums, int k) {
    vector<int>visited(nums.size(),false);
    int sum=0;
    for(auto x:nums) sum+=x;
    if(sum%k!=0) return false;
    int targetsum=sum/k;
    return dfs(nums,visited,0,k,0,targetsum);
    
}
```