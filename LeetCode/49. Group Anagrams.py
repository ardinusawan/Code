# https://leetcode.com/problems/group-anagrams

class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        hash_map = {}
        for s in strs:
            # sort every element
            d = ''.join(sorted(s))

            # check if element is already in dict
            if d in hash_map:
                # append element to existing dict
                hash_map[d].append(s)
            else:
                # create new element
                hash_map[d] = [s]

        # return only values of hashmap
        # since order is not importance it is ok
        return list(hash_map.values())
