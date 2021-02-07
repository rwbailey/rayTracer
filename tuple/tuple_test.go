package tuple_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/rwbailey/ray/tuple"
)

var _ = Describe("Tuple", func() {

	Context("point", func() {
		It("is a point when w is 1.0", func() {
			// Given
			t := tuple.New(4.3, -4.2, 3.1, 1.0)

			// Then
			Expect(t.X).To(Equal(4.3))
			Expect(t.Y).To(Equal(-4.2))
			Expect(t.Z).To(Equal(3.1))
			Expect(t.W).To(Equal(1.0))
			Expect(t.IsPoint()).To(BeTrue())
			Expect(t.IsVector()).To(BeFalse())

		})

		It("can be created with the point constructor", func() {
			// Given
			p := tuple.Point(4.3, -4.2, 3.1)

			// Then
			Expect(p.IsPoint()).To(BeTrue())
		})
	})

	Context("vector", func() {
		It("is a point when w is 1.0", func() {
			// Given
			t := tuple.New(4.3, -4.2, 3.1, 0.0)

			// Then
			Expect(t.X).To(Equal(4.3))
			Expect(t.Y).To(Equal(-4.2))
			Expect(t.Z).To(Equal(3.1))
			Expect(t.W).To(Equal(0.0))
			Expect(t.IsPoint()).To(BeFalse())
			Expect(t.IsVector()).To(BeTrue())

		})

		It("can be create with the vector constructor", func() {
			// Given
			p := tuple.Vector(4.3, -4.2, 3.1)

			// Then
			Expect(p.IsVector()).To(BeTrue())
		})
	})

	Context("equality", func() {
		It("can be compared", func() {
			// Given
			p1 := tuple.Point(1.3, 5.8, 8.6)
			p2 := tuple.Point(56, 4.3, 12)

			v1 := tuple.Vector(1.3, 5.8, 8.6)

			// Then
			Expect(p1.Equals(p2)).To(BeFalse())
			Expect(p1.Equals(v1)).To(BeFalse())
		})
	})

	Context("addition", func() {
		It("can be added", func() {
			// Given
			p1 := tuple.Point(1, 2, 3)
			v1 := tuple.Vector(5, 5, 5)

			r := tuple.New(6, 7, 8, 1)

			// Then
			Expect(p1.Add(v1)).To(Equal(r))
		})
	})

	Context("subtraction", func() {
		It("Two points can be subtracted to get the vector from one to the other", func() {
			// Given
			p1 := tuple.Point(1, 2, 3)
			p2 := tuple.Point(5, 5, 5)

			v := tuple.Vector(-4, -3, -2)

			// Then
			Expect(p1.Subtract(p2)).To(Equal(v))
		})

		It("A vector can be subtracted from a point to yield a new point", func() {
			// Given
			p1 := tuple.Point(1, 2, 3)
			v1 := tuple.Point(5, 5, 5)

			p2 := tuple.Vector(-4, -3, -2)

			// Then
			Expect(p1.Subtract(v1)).To(Equal(p2))
		})

		It("Two points can be subtracted to get the vector from one to the other", func() {
			// Given
			p1 := tuple.Point(1, 2, 3)
			p2 := tuple.Point(5, 5, 5)

			v := tuple.Vector(-4, -3, -2)

			// Then
			Expect(p1.Subtract(p2)).To(Equal(v))
		})

		It("Two vectors can be subtracted to yield the change in direction between them", func() {
			// Given
			v1 := tuple.Vector(1, 2, 3)
			v2 := tuple.Vector(5, 5, 5)

			v3 := tuple.Vector(-4, -3, -2)

			// Then
			Expect(v1.Subtract(v2)).To(Equal(v3))
		})
	})

	Context("negation", func() {
		It("can be negated", func() {
			// Given
			v1 := tuple.Vector(5, 5, 5)

			v2 := tuple.Vector(-5, -5, -5)

			// Then
			Expect(v1.Negate()).To(Equal(v2))
		})
	})
})
