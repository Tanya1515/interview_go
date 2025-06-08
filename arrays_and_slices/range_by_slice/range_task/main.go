package main

func main() {
	var x = []string{"A", "M", "C"}

	for i, s := range x {
		// 1 - А, 2 - М, 3 - С, поскольку 
		// срез s будет скопирован в range. 
		println(i, s) 

		x[i+1] = "M"
		x = append(x, "Z")
		x[i+1] = "Z"
	}

	// Однако, исходный срез будет модифицирован
	// [ "A", "Z", "Z", "Z", "Z", "Z" ]
}
