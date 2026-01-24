package main

import ("fmt")
// maps in golang do not give o/p acc to the indexes but, the speed is far better than that of arrays, 
// so at the time of fast retrieval maps would be preferred if we do not care about the sequence otherwise arrays
// but remember arrays are slower than that of arrays
func maps(){
	var mp map[string]int = map[string] int {
		"apple" : 1,
		"banana" : 2,
		"orange" : 3,
	}
	fmt.Println(mp)
	mp["apple"] = 17    // change
	// fmt.Println(mp["apple"])   // access
	//delete(mp, "apple")  


	// for accessing the val if it exists
     val, ok := mp["apple"]
      fmt.Println(val, ok)
	
}