package unionfind_test

import (
	"fmt"
	"testing"

	"github.com/zzzFelix/dsa/unionfind"
)

func TestUnionFind(t *testing.T) {
	t.Run("find if path exists in graph", func(t *testing.T) {
		// given
		source := 5
		destination := 9
		edges := [][]int{{4, 3}, {1, 4}, {4, 8}, {1, 7}, {6, 4}, {4, 2}, {7, 4}, {4, 0}, {0, 9}, {5, 4}}
		expected := true

		// when
		unions := map[int]*unionfind.Member{}

		for _, edge := range edges {
			aMem, aOk := unions[edge[0]]
			bMem, bOk := unions[edge[1]]
			if !aOk {
				aMem = unionfind.NewMember(edge[0])
			}
			if !bOk {
				bMem = unionfind.NewMember(edge[1])
			}
			member := unionfind.Union(aMem, bMem)
			unions[edge[0]] = member
			unions[edge[1]] = member
		}

		// then
		srcMem, srcOk := unions[source]
		destMem, destOk := unions[destination]
		if srcOk && destOk {
			actual := unionfind.Find(srcMem) == unionfind.Find(destMem)
			if expected != actual {
				t.Fatalf("expected: %v, actual: %v", expected, actual)
			}
		}
	})

	t.Run("Most Stones Removed With Same Row or Column", func(t *testing.T) {
		rows := map[int]*unionfind.Member{}
		cols := map[int]*unionfind.Member{}
		stones := [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}

		for _, stn := range stones {
			member := unionfind.NewMember(stn)
			val, rowOk := rows[stn[0]]
			if !rowOk {
				rows[stn[0]] = member
			} else {
				rows[stn[0]] = unionfind.Union(member, val)
			}

			val, colOk := cols[stn[1]]
			if !colOk {
				cols[stn[1]] = unionfind.Find(member)
			} else {
				cols[stn[1]] = unionfind.Union(member, val)
			}
		}

		unions := map[*unionfind.Member]interface{}{}
		for _, val := range rows {
			parent := unionfind.Find(val)
			unions[parent] = nil
		}
		for _, val := range cols {
			parent := unionfind.Find(val)
			unions[parent] = nil
		}
		result := 0
		for union := range unions {
			fmt.Println(union)
			result += union.UnionSize - 1
		}

		if result != 5 {
			t.Fatalf("expected: %d, actual: %d", 5, result)
		}
	})
}
