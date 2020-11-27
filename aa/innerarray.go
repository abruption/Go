package main


import "fmt"

func main(){
	floatSlice := make([]float64, 10)
	boolSlice := make([]bool, 10)
	fmt.Println(floatSlice[8], boolSlice[6])
	fmt.Println()

	myMusic := []string{"Brown Eyes", "Exo", "AKMU", "BIG BANG"}
	urMusic := myMusic[:]
	urMusic[3] = "Sun"
	fmt.Printf("%x\n", &urMusic[0])
	fmt.Printf("%x\n", &myMusic[0])		// Shallow Copy
	myMusic = append(myMusic, "Irin")
	fmt.Println(urMusic)
	fmt.Println(myMusic)
	fmt.Printf("%x\n", &urMusic[0])
	fmt.Printf("%x\n", &myMusic[0])		// Deep Copy
}
