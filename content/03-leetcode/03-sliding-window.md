---
title: "滑动窗口"
date: 2022-06-24T15:43:42+08:00
description: "滑动窗口"
tags: ["leetcode", "slide window", "c++"]
image: /covers/leetcode.png
cover: false 
weight: 5 
---

## Longest Substring Without Repeating Characters
[LeetCode](https://leetcode.com/problems/longest-substring-without-repeating-characters)/[力扣](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters)

- 用一个子串保存当前的窗口，查找下一个字符在子串的位置
- 如果为-1，则将字符加入子串，即窗口右移动
- 如果不为-1， 将窗口左边移动到子串中最后一个字符的下一个字符所在位置

{{< highlight cpp >}}
int lengthOfLongestSubString(string s) {
    string subs = "";
    int max = 0;
    for(int i = 0; i < s.size(); i++) {
        int idx = subs.find(s[i]);
        if (idx != -1) {
            subs = subs.substr(idx + 1) + s[i];
        } else {
            subs += s[i];
        }
        max = max > subs.size() ? max : subs.size()
    }
    return max;
}
{{< /highlight  >}}

## Substring with Concatenation of All Words
[LeetCode](https://leetcode.com/problems/substring-with-concatenation-of-all-words)/[力扣](https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words)

- 一层遍历子串 
- 第二层使用定长的滑动窗口，判断值是否在words中出现，并记录出现次数
- 记录words出现的值用map

{{< highlight cpp >}}
bool checkSubstring(string s, vector<string>& words) {
    int width = words[0].size();
    map<string, int> wmap;
    for(size_t i = 0; i < words.size(); i++) {
        wmap[words[i]]++;
    }
    for(size_t i = 0; i <= s.size() - width; i = i + width) {
        string subs = s.substr(i, width);
        if(wmap[subs] <= 0 ) {
            return false;
        }
        wmap[subs]--;
    }
    return true
}

vector<int> findSubstring(string s, vector<string>& words) {
    vector<int> res;
    int size = words.size();
    if(size <= 0 ) return res;

    int width = words[0].size();
    int len = size * width;
    if(len > s.size()) {
        return res;
    }
    for(size_t i = 0; i<= s.size() - len; i++) {
        string subs = s.substr(i, len);
        if(checkSubstring(subs, words)) {
            res.push_back(i);
        }
    }
    return res;
}
{{< /highlight  >}}

## Minimum Window Substring
[LeetCode](https://leetcode.com/problems/minimum-window-substring)/[力扣](https://leetcode-cn.com/problems/minimum-window-substring)

- 左右指针表示滑动窗口大小
- 右指针一直往右边移动，当条件满足时，记录下来
- 开始移动左指针，条件满足时，更新结果，并继续移动，直到窗口大小和T字符串大小相等为止
- 重复上述过程，直到s字符串末尾

{{< highlight cpp >}}
string minWindow(string s, string t) {
    unordered_map<char, int> tcnt;
    unordered_map<char, int> scnt;

    int lans = -1,  rans = -1, left = 0, right = 0;
    int asize = INT_MAX;
    
    for(int i = 0; i < t.size(); i++) tcnt[t[i]]++;
    int flag = 0;
    if(tcnt.find(s[0]) != tcnt.end()) {
        scnt[s[0]]++;
        if(scnt[s[0]] == tcnt[s[0]]) flag++;
    }

    while(1) {
        if(flag == tcnt.size()) {
            if(right - left < asize) {
                lans = left; rans = right; asize = right - left;
            }
            if(tcnt.find(s[left]) != tcnt.end()) {
                if(scnt[s[left]] > 0 ) {
                    scnt[s[left]]--;
                    if(scnt[s[left]] == tcnt[s[left]] - 1) flag--;
                }
            }
            left++;
        } else {
            if (right == s.size() - 1) break;
            right++;
            if(tcnt.find(s[right]) != tcnt.end()) {
                scnt[s[right]]++;
                if(scnt[s[right]] == tcnt[s[right]]) flag++
            }
        }
    }
    if (asize == INT_MAX) return "";
    return s.substr(lans, asize + 1);
}
{{< /highlight  >}}

- 优化后的滑动窗口

{{< highlight cpp >}}
string minWindow(string s, string t) {
    vector<int> hist(128, 0);
    for (char c : t) hist[c]++;
    
    int remaining = t.length();
    int left = 0, right = 0, minStart = 0, minLen = numeric_limits<int>::max();
    while (right < s.length()){
        if (--hist[s[right++]] >= 0) remaining--;
        while (remaining == 0){
            if (right - left < minLen){
                minLen = right - left;
                minStart = left;
            }
            if (++hist[s[left++]] > 0) remaining++;
        }
    }
        
    return minLen < numeric_limits<int>::max() ? s.substr(minStart, minLen) : "";
}
{{< /highlight  >}}

## Minimum Size Subarray Sum
[LeetCode](https://leetcode.com/problems/minimum-size-subarray-sum)/[力扣](https://leetcode-cn.com/problems/minimum-size-subarray-sum)

- 前后两个指针分别指向起始位置和滑动窗口结束位置
- 当和大于结果时，左指针移动
- 当和小于结果时，右指针移动
- 当和等于结果时，两个指针同时移动
- 最终结果，可以用vector把和记录下来

{{< highlight cpp >}}
int minSubArrayLen(int s, vector<int>& nums) {
    if(nums.size() == 0) return 0;
    int sum = 0, begin = 0, end = 0;
    vector<int> sums;
    int res = nums.size() + 1;
    sums.push_back(0);
    for(int i = 0; i < nums.size(); i++) {
        sums.push_back(sums[i] + nums[i]);
    }

    while(end <= nums.size() - 1 && begin <= end) {
        sum = sums[end + 1] - sums[begin + 1] + nums[begin];
        if(sum >= s && begin == end) return 1;
        if(sum > s) {
            res = res > end - begin + 1 ? end - begin + 1 : res;
            begin++;
        } else if (sum < s) {
            end++;
        } else {
            res = res > end - begin + 1 ? end - begin + 1 : res;
            begin++;
            end++;
        }
        end = end > nums.size() ? nums.size() - 1: end;
    }
    return res == nums.size() + 1 ? 0 : res;
}
{{< /highlight  >}}

## Sliding Window Maximum
[LeetCode](https://leetcode.com/problems/sliding-window-maximum)/[力扣](https://leetcode-cn.com/problems/sliding-window-maximum)

- 用队列保存数和数的位置，先把`0~k`中最大的那个数字找到并压入队列
- 从`k - 1`处开始循环，如果当前的数字比队列尾部的数字大，则循环`pop_back`
- 小循环结束后，需要把当前数和位置信息放入队列，并把队列首部数字保存到结果中
- 最后判断一下队列首部数字是否在`k`大小的滑动窗口中，如果不在，则`pop`出来

{{< highlight cpp >}}
vector<int> maxSlidingWindow(vector<int>& nums, int k) {
    int n = nums.size();
    vector<int> res;
    if(n == 0) return res;
    deque<pair<int, int>> maxq;
    for(int i = 0; i < k - 1; i++) {
        while(!maxq.empty() && maxq.back().first < nums[i]) {
            maxq.pop_back();
        }
        maxq.push_back({nums[i], i});
    }

    int j = 0;
    for(int i = k - 1; i < n; i++, j++) {
        while(!maxq.empty() && maxq.back().first < nums[i]) {
            maxq.pop_back();
        }
        maxq.push_back({nums[i], i});
        res.push_back(maxq.front().first);
        if(maxq.front().second == j) {
            maxq.pop_front();
        }
    }
    return res;
}
{{< /highlight  >}}

## Longest Repeating Character Replacement
[LeetCode](https://leetcode.com/problems/longest-repeating-character-replacement)/[力扣](https://leetcode-cn.com/problems/longest-repeating-character-replacement)

- 两个指针分别指向窗口的左边和右边
- 如果当前窗口计算出来的替换值小于等于k，则右边指针移动
- 否则，左边指针移动
- 计算当前窗口的替换值用map来保存，当前窗口中所有字符出现的总次数 - 最大字符长度 = 剩下要替换的次数

{{< highlight cpp >}}
int characterReplacement(string s, int k) {
    if(s == "") return 0;
    int left = 0, right = 0, max = 0;
    map<char, int> cmap;
    do {
        if(left == 0 && right == 0) {
            cmap[s[right]]++;
            right++;
            max = 1;
            continue;
        }
        map<char, it>::iterator cit = cmap.begin();
        int i = 0, tmax = 0;
        while(cit != cmap.end()) {
            t += cit -> second;
            tmax = tmax < cit -> second ? cit -> second : tmax;
            cit++;
        }
        t -= tmax;
        if(t <= k) {
            max = max > right - left ? max : right - left;
            cmap[s[right]]++;
            right++;
        } else {
            cmap[s[left]]--;
            leftt++;
        }
    } while(left <= right && right <= s.size())
    return max;
}
{{< /highlight  >}}

- 上述方法可以改进，用一个长度为26的数组来替代map

{{< highlight cpp >}}
int characterReplacement(string s, int k) {
    int n = s.size();
    if(n == 0 || n == 1 || k > n) {
        return n;
    }
    vector<int> chars = vector<int>(26, 0);
    int start = 0, maxLen = 0, maxCount = 0;
    
    for(int end = 0; end < n; end++) {
        chars[s[end] - 65]++;
        int curCount = chars[s[end] - 65];
        maxCount = max(maxCount, curCount);

        if(end - start - maxCount + 1 > k) {
            chars[s[start] - 65]--;
            start++;
        }

        maxLen = max(maxLen, end - start + 1);
    }
    return maxLen;
}
{{< /highlight  >}}

## Permutation in String
[LeetCode](https://leetcode.com/problems/permutation-in-string)/[力扣](https://leetcode-cn.com/problems/permutation-in-string)

- 固定窗口滑动，需要判断两个字符串是否为组合
- 如果不为组合，则继续向前滑动
- 如果为组合，直接返回true

{{< highlight cpp >}}
bool checkZero(map<char, int>& cmap) {
    map<char, int>::iterator it = cmap.begin();
    while(it != cmap.end()) {
        if(it -> second != 0){
            return false;
        }
        it++;
    }
    return true;
} 

bool checkInclusion(string s1, string s2) {
    int right = s1.size(), len = s2.size();
    bool res = false;
    for(int left = 0; left <= len - right; left++) {
        map<char, int> cmap;
        for(int i = 0; i < right; i++) {
            cmap[s1[i]]++;
            cmap[s2[i + left]]--;
        }
        res = checkZero(cmap);
        if (res) {
            return true;
        }
    }
    return false;
}
{{< /highlight  >}}

- 上述操作无法通过最后一个案例，超时，所以可以改进一下

{{< highlight cpp >}}
bool checkInclusion(string s1, string s2) {
    vector<int> a(26, 0), b(26, 0);
    for(const auto& c : s1) {
        a[c - 'a'] ++;
    } 
    int n1 = s1.length(), n2 = s2.length();
    for(int i = 0; i < n2; i++) {
        ++b[s2[i] - 'a'];
        if(i >= n1) {
            --b[s2[i - n1] - 'a'];
        }
        if(a == b) {
            return true;
        }
    }
    return false;
}
{{< /highlight  >}}

## Count Unique Characters of All Substrings of a Given String
[LeetCode](https://leetcode.com/problems/count-unique-characters-of-all-substrings-of-a-given-string)/[力扣](https://leetcode-cn.com/problems/count-unique-characters-of-all-substrings-of-a-given-string)

{{< highlight cpp >}}
int uniqueLetterString(string s) {
    unordered_map<char, vector<int>> m;
    for(int i = 0; i < s.size(); i++) {
        m[s[i]].push_back(i);
    }
    int res = 0;
    for(auto it : m) {
        vector<int> tv = it.second;
        for(int i = 0; i < tv.size(); i++) {
            long prev = i > 0 ? tv[i - 1] : -1;
            long next = i < tv.size() - 1 ? tv[i + 1] : s.size();
            res += (tv[i] - prev) * (next - tv[i]);
        }
    }
    return (int) (res % 1000000007 )
}
{{< /highlight  >}}

## Fruit Into Baskets
[LeetCode](https://leetcode.com/problems/fruit-into-baskets)/[力扣](https://leetcode-cn.com/problems/fruit-into-baskets)

{{< highlight cpp >}}
int totalFruit(vector<int>& tree) {
    int n = tree.size();
    if(n <= 1)return n;
    int left = 0, right = 0;
    int res = 1;
    unordered_map<int,int> m;
    m[tree[0]]++;
    int count = 1;
    while(right < n - 1 && left <= right) {
        if(count == 1) {
            right++;
            if(m[tree[right]] == 0) count++;
            m[tree[right]]++;
            res = max(res, right - left + 1);
            
        }else if(count == 2) {
            right++;
            if(m[tree[right]] == 0){
                count++;
                m[tree[right]] = 1;
                continue;
            }
            m[tree[right]]++;
            res = max(res, right - left + 1);
        }else {
            m[tree[left]]--;
            if(m[tree[left]] == 0) {
                count--;
                
            }
            left++;
            res = max(res, right - left + 1);
        }
    }
    return res;
}
{{< /highlight  >}}

## Minimum Number of K Consecutive Bit Flips
[LeetCode](https://leetcode.com/problems/minimum-number-of-k-consecutive-bit-flips)/[力扣](https://leetcode-cn.com/problems/minimum-number-of-k-consecutive-bit-flips)

{{< highlight cpp >}}
int minKBitFlips(vector<int>& A, int K) {
    int n = A.size();
    vector<int> hint(n,0);
    int ans = 0, flip = 0;
    
    for(int i = 0; i < n; i++) {
        flip ^= hint[i];
        if(A[i] == flip) {
            ans++;
            if(i + K > n) return -1;
            flip ^= 1;
            if(i+K<n) hint[i+K]^=1;
        }
    }
    return ans;
}
{{< /highlight  >}}