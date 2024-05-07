package main

import (
	"fmt"
	"slices"
)

func main() {
	/*
		Slice can grow in runtime, as these does not have size in type
	*/

	var s []int            // Nil slice
	fmt.Println(s == nil)  //true
	var ss = []int{}       // Not a Nil slice
	fmt.Println(ss == nil) //false
	fmt.Println(ss)
	var s1 = []int{1, 2} // Slice literal
	fmt.Println(s1)
	// Slice is not comparable expect with nil
	// invalid operation: s1 == ss (slice can only be compared to nil)
	// fmt.Println(s1==ss)

	// New Func added in GO 1.21
	// A slice is equal if len & value of elements both are equal, and comparable
	x := []int{1, 2, 3, 4, 5}
	y := []int{1, 2, 3, 4, 5}
	z := []int{1, 2, 3, 4, 5, 6}
	// s2 := []string{"a", "b", "c"}
	fmt.Println(slices.Equal(x, y)) // prints true
	fmt.Println(slices.Equal(x, z)) // prints false
	// fmt.Println(slices.Equal(x, s2)) // does not compile, as int and string are not comparable

	/*
		Before this equal func, reflect.DeepEqual() was often used for slice comparison
		Don't use in new code, as its slower and less safe
	*/

	/*
		Use len function to find length of slice, beauty of len is that support many kind
		like array, slice, map, channel etc. but reject other types like int, float
	*/
	z = append(z, 20, 80) // append in the array and re-assign it back to same variable, as go is call by value lang
	z = append(z, y...)   //

	/*
		Capacity is larger than len
		Slice can have both diff and equal
		If you try to add additional values when the length equals the
		capacity, the append function uses the Go runtime to allocate a new backing array for
		the slice with a larger capacity. The values in the original backing array are copied
		to the new one, the new values are added to the end of the new backing array, and
		the slice is updated to refer to the new backing array. Finally, the updated slice is
		returned.
		The rule as of Go 1.18 is to double the capacity of a slice when the current capacity is less than 256. A
		bigger slice increases by (current_capacity + 768)/4. This slowly converges at 25%
		growth (a slice with capacity of 512 will grow by 63%, but a slice with capacity 4,096
		will grow by only 30%).
	*/

	var r []int
	fmt.Println(r, len(r), cap(r)) // [], 0,0
	r = append(r, 1)
	fmt.Println(r, len(r), cap(r)) // [1],1,1
	r = append(r, 2)
	fmt.Println(r, len(r), cap(r)) //[1,2],2,2
	r = append(r, 3)
	fmt.Println(r, len(r), cap(r)) //[1,2,3],3,4
	r = append(r, 4)
	fmt.Println(r, len(r), cap(r)) //[1,2,3,4],4,4
	r = append(r, 5)
	fmt.Println(r, len(r), cap(r)) //[1,2,3,4,5],5,8

	// make - to specify the type, length, and, optionally, the capacity
	f := make([]int, 0)
	fmt.Println(f == nil) // false
	ff := make([]int, 2)
	fmt.Println(ff)
	ff = append(ff, 10) //[0,0,10]
	fmt.Println(ff)

	g := make([]int, 2, 10)
	g[0] = 1
	g[1] = 2
	fmt.Println(g) // [1,2]
	clear(g)       // Reset to default value of type
	fmt.Println(g) // [0,0]

	/*
		Decide slice style on the basis of below
		The primary goal is to minimize the number of times the slice needs to grow.

		var x []int // nil slice
		var x = []int{1,2,3} // Slice literal
		var x = make([]int,5) //
		var x = make([]int) or []int{}// after this append values
	*/

	/*
		When you take a slice from a slice, you are not making a copy of the data. Instead,
		you now have two variables that are sharing memory. This means that changes to an
		element in a slice affect all slices that share that element
	*/
	xx := []string{"a", "b", "c", "d"}
	yy := xx[:2]            // [a,b] 0,1
	zz := xx[1:]            //[b,c,d] 1,2,3
	xx[1] = "y"             // [a,y,c,d],[a,y],[y,c,d]
	yy[0] = "x"             // [x,y,c,d],[x,y],[y,c,d]
	zz[1] = "z"             // [x,y,z,d],[x,y],[y,z,d]
	fmt.Println(xx, yy, zz) // [x,y,z,d],[x,y],[y,z,d]

	ax := []string{"a", "b", "c", "d"}
	ay := ax[:2]                  // [a,b],[0,1]
	fmt.Println(cap(ax), cap(ay)) //4,4
	ay = append(ay, "z")          // [a,b,z]
	fmt.Println("ax:", ax)
	fmt.Println("ay:", ay) //[a,b,z,d]

	/*
		Whenever you take a slice from another slice, the subslice’s capacity
		is set to the capacity of the original slice, minus the starting offset of the subslice
		within the original slice. This means elements of the original slice beyond the end of
		the subslice, including unused capacity, are shared by both slices.
	*/

	// Find out start and end index of sub slice as per parent slice
	// Append will add after end index
	// Sub slice have capacity access to parent slice minus its offset.
	bx := make([]string, 0, 5)
	bx = append(bx, "a", "b", "c", "d")    //[a,b,c,d]
	by := bx[:2]                           // [a,b]
	bz := bx[2:]                           // [c,d]
	fmt.Println(cap(bx), cap(by), cap(bz)) //5,5,3

	by = append(by, "i", "j", "k") // [a,b,i,j,k]
	bx = append(bx, "x")           //[a,b,i,j,x]
	bz = append(bz, "y")           //[i,j,y]
	fmt.Println("bx:", bx)         //[a,b,i,j,y]
	fmt.Println("by:", by)         //[]
	fmt.Println("bz:", bz)         //[i,j,y]

	abc := make([]int, 0, 3)
	abc = append(abc, 1, 2, 3)
	fmt.Println(abc)
	abc = append(abc, 10)
	fmt.Println(abc)
	/*
		Be careful when taking a slice of a slice! Both slices share the same
	   memory, and changes to one are reflected in the other. Avoid modifying
	   slices after they have been sliced or if they were produced
	   by slicing. Use a three-part slice expression to prevent append from
	   sharing capacity between slices.
	*/
}
