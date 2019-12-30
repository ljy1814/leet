/*
 * Author : yajin
 * Email : yajin160305@gmail.com
 * File : 140.cpp
 * CreateDate : 2019-12-28 16:34:23
 * */
#include <iostream>
#include <map>
#include <vector>
#include <string>
#include <list>
#include <algorithm>
#include <unordered_set>


using namespace std;

class Solution {
public:
    struct TrieNode {
        bool flag;
        map<char, TrieNode*> next;
        TrieNode() : flag(false) {}
    };
    TrieNode* root;
    void init(const vector<string>& words) {
        root = new TrieNode;
        for (auto& w : words) {
            auto node = root;
            for (auto c : w) {
                if (!node->next.count(c)) {
                    node->next[c] = new TrieNode;
                }
                node = node->next[c];
            }
            node->flag = true;
        }
    }
    
    void dfs(const string& s, const vector<vector<int> >& dp, int i, vector<string>& v, vector<string>& res) {
        if (i == 0) {
            string t;
            for (auto it = v.rbegin(); it != v.rend(); ++it) {
                t += *it + " ";
            }
            t.pop_back();
            res.push_back(t);
            return;
        }
        for (auto j : dp[i]) {
            v.push_back(s.substr(j, i - j));
            dfs(s, dp, j, v, res);
            v.pop_back();
        }
    }
    vector<string> wordBreak(string s, vector<string>& wordDict) {
        // 构建字典树
        init(wordDict);
        unordered_set<char> cs;
        for (auto& w : wordDict) {
            for (auto c : w) {
                cs.insert(c);
            }
        }
        for (auto c : s) {
            if (cs.count(c) == 0) {
                return {};
            }
        }
        // 动态规划部分
        int N = s.size();
        vector<vector<int> > dp(N + 1);
        dp[0] = vector<int>{-1};
        for (int i = 0; i < N; ++i) {
            if (dp[i].empty()) continue;
            int j = i;
            auto node = root;
            while (j < N && node->next.count(s[j])) {
                node = node->next[s[j++]];
                if (node->flag) {
                    dp[j].push_back(i);
                }
            }
        }

        for (int i = 0; i <= N; ++i) {
            vector<int> d = dp[i];
            cout << i << "|" ;
            for (auto& p : d ) {
                cout << p << " ";
            }
            cout << endl;
        }


        // 路径回溯构建所有可能的结果字符串
        vector<string> v;
        vector<string> res;
        dfs(s, dp, N, v, res);
        return res;
    }
};

class Solution1 {
public:
  vector<string> wordBreak(string s, vector<string>& wordDict) {
    map<string, vector<string>> map;
    return search(s, wordDict, map);
  }
  vector<string> search(string s, const vector<string>& wordDict,
                        map<string, vector<string>>& map) {
    if (map.count(s) > 0) return map[s];
    if (s.empty()) { return { "" }; }
    vector<string> ans;
    for (auto word : wordDict) {
      int len = word.size();
      if (s.substr(0, len) != word) continue;
      vector<string> subAns = search(s.substr(len), wordDict, map);
      for (auto& item : subAns) {
        string t = word + (item == "" ? "" : " " + item);
        ans.push_back(t);
      }
    }
    map[s] = ans;
    return ans;
  }
};

/*
class Solution2 {
public:
    vector<string> wordBreak(string s, vector<string>& wordDict) {
        set<string> strSet(wordDict.begin(),wordDict.end());
        int strlen = s.length();
        vector<vector<string>> result(strlen+1);
        if(!canWordBreak(s,wordDict,strSet)) return result[0];
        vector<int> memo(strlen+1,0);
        memo[strlen] = 1;
        for(int start = strlen-1; start>=0;--start){
            for(int len = 1;len<=strlen-start;++len){
                if(memo[start+len]==1&&strSet.find(s.substr(start,len))!=strSet.end()){
                    memo[start] = 1;
                    if(result[start+len].size()==0) result[start].emplace_back(s.substr(start,len));
                    for(int i = 0;i<result[start+len].size();++i){
                        string temp = s.substr(start,len);
                        temp.append(" "+result[start+len][i]);
                        result[start].emplace_back(temp);
                    }
                }                
            }
        }
        return result[0];
    }
private:
    bool canWordBreak(string s, vector<string>& wordDict, set<string>& strSet) {
        int strlen = s.length(), cur = strlen-1;
        if(strlen<=0) return 0;
        vector<int> memo(strlen+1,0);
        memo[strlen] = 1;
        for(;cur>=0;--cur){
            for(int i = cur+1;i<=strlen;++i){
                if(memo[i]==1){
                    if(strSet.find(s.substr(cur,i-cur))!=strSet.end()){
                        memo[cur] = 1;
                        break;
                    }
                }
            }
        }
        return memo[0];
    }
};
*/


class Solution2 {
public:
    vector<string> wordBreak(string s, const vector<string>& wordDict) {
        vector<bool> dp(s.size() + 10, false);
        dp[0] = true;
        for (int i = 1; i<=s.size(); i++) {
            for (int j = 0; j<=i; j++) {
                if (dp[j] && count(wordDict.begin(), wordDict.end(), s.substr(j, i-j))>=1) {
                    dp[i] = true;
                    break;
                }
            }
        }
        vector<string> res;
        if (dp[s.size()]) {
            dfs(s, "", wordDict, res, 0, 0);
        }
        return res;
    }
    void dfs(string& s, string ss, const vector<string>& wordDict, vector<string>& res, int pre, int cur) {
        if (s.size() == cur && count(wordDict.begin(), wordDict.end(), s.substr(pre, cur-pre)) > 0) {
            ss += s.substr(pre, cur-pre);
            res.push_back(ss);
            return ;
        }
        if (cur >= s.size())    return ;
        string sss = ss;
        if (count(wordDict.begin(), wordDict.end(), s.substr(pre, cur-pre)) > 0) {
            sss += s.substr(pre, cur-pre) + " ";
            //cout << ss << endl << pre << cur << endl;
            dfs(s, sss, wordDict, res, cur, cur);
        }
        dfs(s, ss, wordDict, res, pre, cur+1);
    }
};


class Solution3 {
public:
    vector<string> wordBreak(string s, vector<string>& wordDict) {
        int n = s.length();
        set<string> st;
        
        // dp[j], 记录长度为 j 的串，以 j 为结尾的哪些后缀在字典中。
        vector<vector<int>> dp(n+1);
        
        dp[0] = {0};
        
        int minLength = INT_MAX;
        int maxLength = 0;
        
        // Step 1. 追踪字典中最短的单词和最长的单词，并插入 set
        for (auto& word : wordDict) {
            if (word.length() < minLength) minLength = word.length();
            if (word.length() > maxLength) maxLength = word.length();
            st.insert(word);
        }
        
        // m[i] 表示长度为 i 的串构造的句子
        vector<vector<string>> m(n + 1);
        
        // Step 2.计算长度为 i 的串是否能被拆分
        for (int j = 1; j <= n; ++j) {
            for (int i = max(0, j - maxLength); i <= j - minLength; ++i) {
                if (!dp[i].empty() && st.count(s.substr(i, j - i))) {
                    dp[j].push_back(i);
                }
            }
        }
        
        // 提前剪枝
        if (dp[n].empty()) return vector<string>();
        
        // Step 3. 下面这段不能合并到 Step 2 中，否则会超时。这里正好可以用到 Step 2 结果。当然你也可以使用 dfs，那就得再遍历一遍串，并且查找 dict，费事。
        for (int j = 1; j <= n; ++j) {
            if (dp[j].empty()) continue;
            for (auto i : dp[j]) {
                if (m[i].empty()) {
                    m[j].push_back(s.substr(i, j - i));
                } else {
                    for (auto& ss : m[i]) {
                        m[j].push_back(ss + ' ' + s.substr(i, j - i));
                    }
                }
            }
        }
        
        return m[n];
    }
};


Solution3 obj;
int main(int argc, char** argv)
{
    string s = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
    vector<string> words = { "a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa" };
    //string s = "catsanddog";
    //vector<string> words = {"cat", "cats", "and", "sand", "dog"};
    vector<string> result = obj.wordBreak(s, words);
    //cout << result << endl;
    for (auto& w : result) {
        cout << w << endl;
    }
    return 0;
}


/* vim: set tabstop=4 set shiftwidth=4 */

