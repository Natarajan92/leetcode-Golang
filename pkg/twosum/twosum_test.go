package twosum

import (
	"log"
	"reflect"
	"testing"

	. "github.com/franela/goblin"
)

func TestTwoSum(t *testing.T) {
	g := Goblin(t)
	log.Print("Unit Test Started for two sum verification")

	g.Describe("Test TWO SUM output as per Problem Statement- ", func() {

		g.It("With proper input, output and target ", func() {

			nums := []int{3, 2, 4}
			target := 6
			output := []int{1, 2}
			numIndex := twoSum(nums, target)

			g.Assert(reflect.DeepEqual(numIndex, output)).IsTrue()
		})

		g.It("With proper input, output and target", func() {

			nums := []int{3, 2, 3, 4}
			target := 7
			output := []int{2, 3}
			numIndex := twoSum(nums, target)

			g.Assert(reflect.DeepEqual(numIndex, output)).IsTrue()
		})
		g.It("With proper input, output and target - with negative value", func() {

			nums := []int{-3, 2, 4, -1}
			target := -4
			output := []int{0, 3}
			numIndex := twoSum(nums, target)

			g.Assert(reflect.DeepEqual(numIndex, output)).IsTrue()
		})
		g.It("With proper input, output and target outside of input", func() {

			nums := []int{3, 2, 4}
			target := 9
			output := []int{}
			numIndex := twoSum(nums, target)

			g.Assert(reflect.DeepEqual(numIndex, output)).IsTrue()
		})
	})
}
