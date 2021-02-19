package intslice_test

import (
	"math/rand"
	"sort"
	"time"

	"github.com/alexunjm/golang-study-repo/intslice"
)

var _ = Describe("intslice.MyIntSlice", func() {

	var (
		emptySlice              intslice.MyIntSlice
		oneElementSlice         intslice.MyIntSlice
		twoElementSlice         intslice.MyIntSlice
		moreThanTwoElementSlice intslice.MyIntSlice
		manyElementSlice        intslice.MyIntSlice
	)

	BeforeSuite(func() {
		// BeforeEach(func() {

		emptySlice = intslice.MyIntSlice{}
		oneElementSlice = intslice.MyIntSlice{1}
		twoElementSlice = intslice.MyIntSlice{2, 1}
		moreThanTwoElementSlice = intslice.MyIntSlice{5, 7, 4, 9, 1}
		rand.Seed(time.Now().UnixNano())
		arr := [9999]int{}
		for i := 0; i < 9999; i++ {
			arr[i] = rand.Intn(9999)
		}
		manyElementSlice = arr[:]
	})

	Describe("slice creation", func() {
		Context("with no elements", func() {
			It("should be an empty int slice", func() {
				Expect(emptySlice).Should(BeEquivalentTo([]int{}))
			})
		})

		Context("with one element", func() {
			It("should be have same int value on int slice", func() {
				Expect(oneElementSlice).Should(BeEquivalentTo([]int{1}))
			})
		})

		Context("with two elements", func() {
			It("should be have same int values on int slice", func() {
				Expect(twoElementSlice).Should(BeEquivalentTo([]int{2, 1}))
			})
		})

		Context("with more than two elements", func() {
			It("should be have same int values on int slice with same elements", func() {
				Expect(moreThanTwoElementSlice).Should(BeEquivalentTo([]int{5, 7, 4, 9, 1}))
			})
		})
	})

	Describe("MyIntSlice sort", func() {
		Context("with no elements", func() {
			It("should be an empty int slice", func() {
				Expect(intslice.Sort(emptySlice)).Should(BeEquivalentTo([]int{}))
			})
		})

		Context("with one element", func() {
			It("should be same int value slice", func() {
				Expect(intslice.Sort(oneElementSlice)).Should(BeEquivalentTo([]int{1}))
			})
		})

		Context("with two desc element", func() {
			It("should be same int value slice", func() {
				Expect(intslice.Sort(twoElementSlice)).Should(BeEquivalentTo([]int{1, 2}))
			})
		})

		Context("with more than two unsorted elements", func() {
			It("should be have same int values on int slice with same elements sorted", func() {
				Expect(intslice.Sort(moreThanTwoElementSlice)).Should(BeEquivalentTo([]int{1, 4, 5, 7, 9}))
			})
		})
	})

	Describe("go intSlice sort", func() {
		Context("with no elements", func() {
			It("should be an empty int slice", func() {
				sort.Ints(emptySlice)
				Expect(emptySlice).Should(BeEquivalentTo([]int{}))
			})
		})

		Context("with one element", func() {
			It("should be same int value slice", func() {
				sort.Ints(oneElementSlice)
				Expect(oneElementSlice).Should(BeEquivalentTo([]int{1}))
			})
		})

		Context("with two desc element", func() {
			It("should be same int value slice", func() {
				sort.Ints(twoElementSlice)
				Expect(twoElementSlice).Should(BeEquivalentTo([]int{1, 2}))
			})
		})

		Context("with more than two unsorted elements", func() {
			It("should be have same int values on int slice with same elements sorted", func() {
				sort.Ints(moreThanTwoElementSlice)
				Expect(moreThanTwoElementSlice).Should(BeEquivalentTo([]int{1, 4, 5, 7, 9}))
			})
		})
	})

	Describe("compare sorts", func() {
		Context("mine with many unsorted elements", func() {
			It("should be have same int values on int slice with same elements sorted", func() {
				sort.Ints(manyElementSlice)
				Expect(sort.IntsAreSorted(manyElementSlice)).Should(BeTrue())
			})
		})

		Context("go with many unsorted elements", func() {
			It("should be have same int values on int slice with same elements sorted", func() {
				Expect(sort.IntsAreSorted(intslice.Sort(manyElementSlice))).Should(BeTrue())
			})
		})

		Context("mine with many unsorted elements", func() {
			It("should be have same int values on int slice with same elements sorted", func() {
				sort.Ints(manyElementSlice)
				Expect(sort.IntsAreSorted(manyElementSlice)).Should(BeTrue())
			})
		})
	})
})
