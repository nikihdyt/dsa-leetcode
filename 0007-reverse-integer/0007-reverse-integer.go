import (
    "strconv"
    "math"
)

func reverse(x int) int {
    sNum := strconv.Itoa(x)

	var temp string

    if x < 0{
        temp += "-"
        for i, _ := range sNum[1:]{
		    temp += string(sNum[len(sNum)-i-1])
        }
    } else{
        for i, _ := range sNum{
		    temp += string(sNum[len(sNum)-i-1])
        }
    }

	res, _ := strconv.Atoi(temp)

    if res > math.MaxInt32 || res < math.MinInt32{
        return 0
    }
	return res
}