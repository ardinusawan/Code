// https://www.hackerrank.com/challenges/time-conversion/problem

func timeConversion(s string) string {
    suffix := s[len(s)-2:]
    var convertedTime string
    if suffix == "PM" {
        hour := s[:2]
        h, _ := strconv.Atoi(hour)
        if h != 12 {
            h += 12    
        }
        if h == 24 {
            convertedTime = fmt.Sprintf("00%s", s[2:len(s)-2])
        } else {
            convertedTime = fmt.Sprintf("%02d%s", h, s[2:len(s)-2])
        }
    } else {
        hour := s[:2]
        h, _ := strconv.Atoi(hour)
        if h == 12 {
            convertedTime = fmt.Sprintf("00%s", s[2:len(s)-2])
        } else {
            convertedTime = fmt.Sprintf("%02d%s", h, s[2:len(s)-2])
        }
    }
    return convertedTime
}
