# https://leetcode.com/problems/insert-interval/

class Solution:
    def insert(self, intervals: List[List[int]], newInterval: List[int]) -> List[List[int]]:
        result = []
        pushed = False
        for i,v in enumerate(intervals):
            # if newInterval not merged yet and can be merged
            if pushed==False and self.can_merge(v, newInterval):
                merged_data = self.merge_list(v, newInterval)
                result.append(merged_data)  
                pushed=True
                
            # try to merge next interval data with latest result
            if(i+1!=len(intervals) and len(result)>0):
                last_data = result[-1]
                if self.can_merge(last_data, intervals[i+1]):
                    merged_data = self.merge_list(last_data, intervals[i+1])
                    result.pop()
                    result.append(merged_data)
            
            # if current interval inside range last result ignore it
            if len(result)>0 and self.inside_slice(v, result[-1]):
                continue
            else:
                result.append(v)
            
            # if until end of loop newInterval not yet pushed
            if pushed==False and i==len(intervals)-1:
                result.append(newInterval)
                pushed=True
        
        # if intervals is empty, push newInterval
        if len(intervals)==0:
            result.append(newInterval)

        # sort result based on first index
        return sorted(result, key=lambda data: data[0])
    
    def can_merge(self, l1, l2):
        return not (l2[1] < l1[0] or l2[0] > l1[1])
        
    def merge_list(self, l1, l2):
        lowest = min(l1[0], l2[0])
        maxest = max(l1[1], l2[1])
        return [lowest, maxest]
    
    def inside_slice(self, l1, l2):
        return l1[0]>=l2[0] and l1[1]<=l2[1]
