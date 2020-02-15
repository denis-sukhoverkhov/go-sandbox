package main

import "fmt"

func climbingLeaderboard(scores []int32, alice []int32) []int32 {

	var total_positions []int32

	var ranks []int32

	prev := int32(-1)
	cur_rank := int32(0)
	for _, v := range scores {
		if prev != v {
			prev = v
			cur_rank++
		}
		ranks = append(ranks, cur_rank)
	}

	cur_limit := len(scores) - 1
	for i := 0; i < len(alice); i++ {
		for cur_limit >= 0 {
			if alice[i] == scores[cur_limit] {
				total_positions = append(total_positions, ranks[cur_limit])
				break
			} else if alice[i] > scores[cur_limit] {
				if cur_limit == 0 {
					total_positions = append(total_positions, 1)
					break
				} else {
					cur_limit--
				}
			} else if alice[i] < scores[cur_limit] {
				var rank int32
				if cur_limit == len(ranks)-1 {
					rank = ranks[len(ranks)-1] + 1
				} else {
					rank = ranks[cur_limit+1]
				}
				total_positions = append(total_positions, rank)
				break
			}
		}
	}

	return total_positions
}

func main() {
	//scores := []int32{100, 90, 90, 80, 75, 60}
	//alice := []int32{50, 65, 77, 90, 102}

	scores := []int32{100, 100, 50, 40, 40, 20, 10}
	alice := []int32{5, 25, 50, 120}

	fmt.Println(climbingLeaderboard(scores, alice))
}
