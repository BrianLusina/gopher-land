package longestsubstring

import "math"

func lengthOfLongestSubstring(s string) int {
    start := 0
    result := 0    
    seen := map[uint8]int{}
    
    for end := 0; end < len(s); end++ {
        char := s[end]
        
        duplicateIndex, ok:= seen[char]
        if ok {
            result = int(math.Max(float64(result), float64(end-start)))
            for i:=start;i <duplicateIndex; i++ {
                delete(seen, s[i])
            }
            
            start = duplicateIndex + 1
        }
        
        seen[s[end]] = end
    }
    
    result = int(math.Max(float64(result), float64(len(s) - start)))
    return result
}
