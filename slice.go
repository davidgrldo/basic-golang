package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",   // 0
		"Februari",  // 1
		"Maret",     // 2
		"April",     // 3
		"Mei",       // 4
		"Juni",      // 5
		"Juli",      // 6
		"Agustus",   // 7
		"September", // 8
		"Oktober",   // 9
		"November",  // 10
		"Desember",  // 11
	}

	// Create a slice from array (start from index low until index before high (-1))
	var firstSlice = months[4:7] // 4 to 6 (7-1)
	fmt.Println(firstSlice)      // [Mei Juni Juli]

	// Get the length
	fmt.Println(len(firstSlice)) // 3

	// Get the capacity
	fmt.Println(cap(firstSlice)) // 8

	var secondSlice = months[2:4] // 2 (low) to 3 (4 (high) -1)
	fmt.Println(secondSlice)      // [Maret April]

	/**
	* Create a new slice (add data in the last position of slice)
	* if the capacuty full? create new array
	 */
	var thirdSlice = append(secondSlice, "Bulan baru")
	fmt.Println(thirdSlice)         // [Maret April Bulan baru] => create new data
	thirdSlice[1] = "Bulan purnama" // Update the index data of slice (April)
	fmt.Println(thirdSlice)         // [Maret Bulan purnama Bulan baru] => updated data (April changed)

	fmt.Println(secondSlice) // [Maret Bulan purnama] => updated data (April changed), because the update from thirdSlice[1]

	/**
	 * April chaged to Bulan purnama because the update thirdSlice[1]
	 * Mei changed to Bulan baru because the append of secondSlice (index data is still in the array / does not exceed capacity)
	 * [Januari Februari Maret Bulan purnama Bulan baru Juni Juli Agustus September Oktober November Desember]
	 */
	fmt.Println(months)

	var days = [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	var slice1 = days[1:3]
	fmt.Println(slice1)

	var slice2 = days[6:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Hari ulang tahun")
	fmt.Println(slice3)
	slice3[1] = "Hari libur"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(days)

	// Make new slice (length, capacity)
	newSlice := make([]string, 2, 5)

	newSlice[0] = "David"
	newSlice[1] = "Geraldo"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// Copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice) // parameter(target, source)
	fmt.Println(copySlice)

	isArray := [5]int{1, 2, 3, 4, 5}
	isSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(isArray)
	fmt.Println(isSlice)

}
