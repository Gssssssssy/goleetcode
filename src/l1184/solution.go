package l1184

/*
环形公交路线上有 n 个站，按次序从 0 到 n - 1 进行编号。我们已知每一对相邻公交站之间的距离，distance[i] 表示编号为 i 的车站和编号为 (i + 1) % n 的车站之间的距离。

环线上的公交车都可以按顺时针和逆时针的方向行驶。

返回乘客从出发点 start 到目的地 destination 之间的最短距离。

Tips:
>
> - 1 <= n <= 10^4
> - distance.length == n
> - 0 <= start, destination < n
> - 0 <= distance[i] <= 10^4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/distance-between-bus-stops
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	optimalDst := 0

	clockWiseIndex := start
	clockWiseDst := 0

	counterClockWiseDst := 0

	for {
		if clockWiseIndex == destination {
			break
		}
		if clockWiseIndex == len(distance) {
			clockWiseIndex = 0
		}
		clockWiseDst += distance[clockWiseIndex]
		clockWiseIndex++
	}

	//TODO: 待完成

	if counterClockWiseDst < clockWiseDst {
		optimalDst = counterClockWiseDst
	} else {
		optimalDst = clockWiseDst
	}
	return optimalDst
}
