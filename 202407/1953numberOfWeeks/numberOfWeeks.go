package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(numberOfWeeks([]int{1, 2, 3}))
}

func numberOfWeeks(milestones []int) int64 {
	// 	我们考虑什么情况下不能完成所有阶段任务。如果存在一个项目 i，它的阶段任务数大于其余所有项目的阶段任务数之和再加 1，
	// 	那么就不能完成所有阶段任务。否则，我们一定可以通过不同项目之间来回穿插的方式完成所有阶段任务
	mx := slices.Max(milestones)
	sum := 0
	for _, ch := range milestones {
		sum += ch
	}
	rest := sum - mx

	if mx > rest+1 {
		return int64(2*rest + 1)
	}
	return int64(sum)
}

/*

select round(avg(a.event_date is not null), 2) fraction
from (select player_id, min(event_date) as first_login
      from activity
      group by player_id) p
         left join activity a
                   on p.player_id = a.player_id and datediff(a.event_date, p.first_login) = 1;

*/
