# 981. Time Based Key-Value Store

Create a timebased key-value store class TimeMap, that supports two operations.

I. set(string key, string value, int timestamp)

- Stores the key and value, along with the given timestamp.

II. get(string key, int timestamp)

- Returns a value such that set(key, value, timestamp_prev) was called previously, with timestamp_prev <= timestamp.
- If there are multiple such values, it returns the one with the largest timestamp_prev.
- If there are no values, it returns the empty string ("").
 
```
Example 1:

Input: inputs = ["TimeMap","set","get","get","set","get","get"], inputs = [[],["foo","bar",1],["foo",1],["foo",3],["foo","bar2",4],["foo",4],["foo",5]]
Output: [null,null,"bar","bar",null,"bar2","bar2"]

Explanation:   
TimeMap kv;   
kv.set("foo", "bar", 1); // store the key "foo" and value "bar" along with timestamp = 1   
kv.get("foo", 1);  // output "bar"   
kv.get("foo", 3); // output "bar" since there is no value corresponding to foo at timestamp 3 and timestamp 2, then the only value is at timestamp 1 ie "bar"   
kv.set("foo", "bar2", 4);   
kv.get("foo", 4); // output "bar2"   
kv.get("foo", 5); //output "bar2"   

Example 2:

Input: inputs = ["TimeMap","set","set","get","get","get","get","get"], inputs = [[],["love","high",10],["love","low",20],["love",5],["love",10],["love",15],["love",20],["love",25]]
Output: [null,null,null,"","high","high","low","low"]
 

Note:

All key/value strings are lowercase.
All key/value strings have length in the range [1, 100]
The timestamps for all TimeMap.set operations are strictly increasing.
1 <= timestamp <= 10^7
TimeMap.set and TimeMap.get functions will be called a total of 120000 times (combined) per test case.e of each element in nums will be in the range [-9999, 9999].
```

[981. Time Based Key-Value Store](https://leetcode.com/problems/time-based-key-value-store/)


```
Time complexity: set => O(logn)/O(1), get => O(logn)
Space comlexity: O(n)
```

## hashtable + vector + binary search

```golang
import "sort"

type data struct {
    time  int
    value string
}

// TimeMap is a timebased key-value store
type TimeMap map[string][]data

// Constructor Initialize your data structure here.
func Constructor() TimeMap {
    return make(map[string][]data, 1024)
}

// Set the key and the value
// Note 3:
// The timestamps for all TimeMap.set operations are strictly increasing.
func (m TimeMap) Set(key string, value string, timestamp int) {
    if _, ok := m[key]; !ok {
        m[key] = make([]data, 1, 1024)
    }
    m[key] = append(m[key], data{
        time:  timestamp,
        value: value,
    })
}

// Get the value of key
func (m TimeMap) Get(key string, timestamp int) string {
    // The timestamps for all TimeMap.set operations are strictly increasing.
    d := m[key]
    // binary search
    i := sort.Search(len(d), func(i int) bool {
        return timestamp < d[i].time
    })
    i--
    return m[key][i].value
}


/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
```
