package intslice_test

// . "github.com/onsi/ginkgo"
// . "github.com/onsi/gomega"

var _ = Describe("IntSlice", func() {

	var (
		emptySlice              IntSlice
		oneElementSlice         IntSlice
		twoElementSlice         IntSlice
		moreThanTwoElementSlice IntSlice
	)

	BeforeEach(func() {

		emptySlice = IntSlice{}
		oneElementSlice = IntSlice{1}
		twoElementSlice = IntSlice{2, 1}
		moreThanTwoElementSlice = IntSlice{5, 7, 4, 9, 1}
	})

	Describe("test 1", func() {
		Context("with x1", func() {
			It("should be y1", func() {
				Expect(emptySlice).To(Equal(IntSlice{}))
			})
		})

		Context("with x2", func() {
			It("should be y2", func() {
				Expect(oneElementSlice).To(Equal(IntSlice{1}))
			})
		})

		Context("with x3", func() {
			It("should be y3", func() {
				Expect(twoElementSlice).To(Equal(IntSlice{1, 2}))
			})
		})

		Context("with x4", func() {
			It("should be y4", func() {
				Expect(moreThanTwoElementSlice).To(Equal(IntSlice{5, 7, 4, 9, 1}))
			})
		})
	})
})
