package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

type user struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Status   string `json:"status"`
	RegDate  string `json:"Reg_date"`
}

func dataTypes() {
	fmt.Println("About basic data types:")
	var vName string = "Derek"
	v1, v2 := 1.2, 3.4
	v3 := "hello"
	v3 = "Hi"
	v4 := 2.4
	var v5 bool
	var v6 string
	var v7 rune
	var v8 int
	var v9 float64
	var v10 float32
	fmt.Println(vName, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10)

	r_array := []rune(vName)

	fmt.Println(reflect.TypeOf(25))
	fmt.Println(reflect.TypeOf(3.14))
	fmt.Println(reflect.TypeOf(25.123456789101112))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("Heey"))
	fmt.Println(reflect.TypeOf(nil))
	fmt.Println(reflect.TypeOf(r_array))
	fmt.Println(reflect.TypeOf(r_array[0]))
}

func casting() {
	fmt.Println("This is basically how casting works:")
	cV1 := 1.5
	cV2 := int(cV1)
	fmt.Println(cV2)
	cV3 := "500000000"
	cV4, err := strconv.Atoi(cV3)
	fmt.Println(cV4, err, reflect.TypeOf(cV4))
	cV5 := strconv.Itoa(cV4)
	fmt.Println(cV5, err, reflect.TypeOf(cV5))

	cV6 := "3.1412"
	if cV8, err := strconv.ParseFloat(cV6, 64); err == nil {
		fmt.Println(cV8)
	}

	cV9 := fmt.Sprintf("%f", 3.14)
	fmt.Println(cV9)

	byteArr := []byte{'a', 'b', 'c'}
	bstr := string(byteArr[:])
	fmt.Println("I'm a string: ", bstr)

	sampleSlice := make([]float64, 6)
	sampleSlice[0] = float64(12)
	sampleSlice[1] = float64(int(13 / 3))
	sampleSlice[2] = 12.01
	fmt.Printf("Casting the len of %v as float64 gives us: %.2f\n", sampleSlice, float64(len(sampleSlice)))
	sampleSlice2 := make([]float64, 0, 6)
	sampleSlice2 = append(sampleSlice2, float64(12))
	sampleSlice2 = append(sampleSlice2, float64(int(13/3)))
	sampleSlice2 = append(sampleSlice2, 12.01)
	fmt.Printf("However, doing this way, %v length as float64 gives us: %.2f\n", sampleSlice2, float64(len(sampleSlice2)))
}

func ifConditional() {
	fmt.Println("How to make an if statemente and conditionals:")
	// Conditional Operators: > < >= <= == !=
	// Logical Operators: && || !
	iAge := 8
	if (iAge >= 1) && (iAge <= 18) {
		fmt.Println("Important BDay")
	} else if (iAge == 21) || (iAge == 50) {
		fmt.Println("Important BDay")
	} else {
		fmt.Println("Not Important BDay")
	}
	fmt.Println("!true ==", !true)
}

func stringManipulation() {
	fmt.Println("This is how to manipulate strings:")
	sV1 := "A word"
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)
	fmt.Println(sV2)                                                   // Another word
	fmt.Println("Length:", len(sV2))                                   // Length: 12
	fmt.Println("Contains Another:", strings.Contains(sV2, "Another")) // Contains Another: true
	fmt.Println("o index:", strings.Index(sV2, "o"))                   // o index: 2
	fmt.Println("Replace:", strings.Replace(sV2, "o", "0", -1))        // Replace: An0ther w0rd
	sV3 := "\nSome Words\n"                                            // \t == tabs, \"
	sV3 = strings.TrimSpace(sV3)
	fmt.Println(sV3)                                             // Some Words
	fmt.Println("Split:", strings.Split("a-b-c-d", "-"))         // Split: [a b c d]
	fmt.Println("Upper:", strings.ToUpper(sV2))                  // Upper: ANOTHER WORD
	fmt.Println("Lower:", strings.ToLower(sV2))                  // Lower: another word
	fmt.Println("Prefix:", strings.HasPrefix("tacocat", "taco")) // Prefix: true
	fmt.Println("Suffix:", strings.HasSuffix("tacocat", "cat"))  // Suffix: true

	// %d : Integer
	// %c : Character
	// %f : Floar (%.2f Float with 2 decimal places)
	// %t : Boolean
	// %s : String
	// %o : Base8
	// %x : Base16
	// %v : Guesses based on data type
	// %T : Type of supplied value
	fmt.Printf("%s %d %c %f %.3f %t %o %x %v\n",
		"Stuff", 1, 'A', 3.141592, 3.141592, true, 1, 1, 2.1)
	fmt.Printf("%9f\n", 3.14)
	fmt.Printf("%.2f\n", 3.141592)
	str1 := fmt.Sprintf("%9.f\n", 3.141592)
	fmt.Println(str1)
}

func runesAndForLoopAndPrintf() {
	fmt.Println("This is how Runes and For loops work:")
	// Details on Printf : https://pkg.go.dev/fmt
	rStr := "abcdef"
	fmt.Println("Rune Count:", utf8.RuneCountInString(rStr))
	for i, runeVal := range rStr {
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}
	// for init; condition; postStatement {<Content>}
	for x := 1; x <= 5; x++ {
		fmt.Println(x) // x only exists inside the scope of this for loop
	}
	fX := 1
	for fX < 5 {
		fmt.Println(fX)
		fX++
	}
	aNums := []int{1, 2, 3}
	for _, v1 := range aNums {
		fmt.Println(v1)
	}
}

func timeAndDate() {
	fmt.Println("Working with time and date:")
	now := time.Now()
	fmt.Println(now.Year(), now.Month(), now.Day())
	fmt.Println(now.Hour(), now.Minute(), now.Day())
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	fmt.Println("Random: ", randNum)
}

func someMathConcepts() {
	fmt.Println("Some math operations and concepts:")
	fmt.Println("5 % 4 =", 5%4)
	mInt := 1
	mInt++ // mInt += 1
	fmt.Println(mInt)
	fmt.Println("Float ", 0.111111111111111111111111+0.111111111111111111111)

	fmt.Println("Abs(-10) =", math.Abs(-10))
	fmt.Println("Pow(4, 2) =", math.Pow(4, 2))
	fmt.Println("Sqrt(16) =", math.Sqrt(16))
	fmt.Println("Cbrt(8) =", math.Cbrt(8))               // Cubic root
	fmt.Println("Ceil(4.43120) =", math.Ceil(4.43120))   // (5) Last integer value greater than or equal to x
	fmt.Println("Floor(4.43120) =", math.Floor(4.43120)) // (4) Leat integer value less than or equal to x
	fmt.Println("Round(4.4) =", math.Round(4.4))         // 4
	fmt.Println("Log2(100) =", math.Log2(100))
	fmt.Println("Log10(100) =", math.Log10(100))
	fmt.Println("Log(100) =", math.Log(100)) // Log of e
	fmt.Println("Max(5, 4) =", math.Max(5, 4))
	fmt.Println("Min(5, 4) =", math.Min(5, 4))

	r90 := 90 * math.Pi / 180
	d90 := r90 * (180 / math.Pi)
	fmt.Printf("%.2f radians == %.2f degrees\n", r90, d90)

	fmt.Println("Sin(90)rad =", math.Sin(r90))
	fmt.Println("Cos(90)rad =", math.Cos(r90))
	fmt.Println("Tan(90)rad =", math.Tan(r90))
	fmt.Println("Acos(90)rad =", math.Acos(r90))
	fmt.Println("Asin(90)rad =", math.Asin(r90))
	fmt.Println("Hypot(3, 4) =", math.Hypot(3, 4))
}

func arrayManipulations() {
	fmt.Println("Manipulating arrays:")
	// Arrays sizes are immutable and have a static type
	var arr1 [5]int
	arr1[0] = 1

	arr2 := []int{1, 2, 3, 4, 5}
	fmt.Println("Index 0:", arr2[0])
	fmt.Println("Arr Length:", len(arr2))

	for i, v := range arr2 {
		fmt.Printf("%d : %d", i, v)
	}
	arr3 := [2][2]int{
		{1, 2},
		{3, 4},
	}
	for i, l := range arr3 {
		for j, c := range l {
			fmt.Printf("Line: %d | Column: %d | Value: %v\n", i, j, c)
		}
	}
}

func slicesAndArrays() {
	fmt.Println("Manipulating slices (more widely used than arrays):")
	// var name []datatype
	sl1 := make([]string, 6)    // == ["","","","","",""]
	sl2 := make([]string, 0, 6) // == [] (len 0, capacity 6)
	sl2 = append(sl2, "asa")
	fmt.Println(sl2)
	sl1[0] = "Society"
	sl1[1] = "of"
	sl1[2] = "the"
	sl1[3] = "Simulated"
	sl1[4] = "Universe"
	fmt.Println("Slice size:", len(sl1))
	for index, piece := range sl1 {
		fmt.Println(index, piece)
	}

	sArr := [5]int{1, 2, 3, 4, 5}
	sl3 := sArr[:3]
	fmt.Println("1st 3:", sArr[:3])
	fmt.Println("Last 3:", sArr[2:])
	sArr[0] = 10
	fmt.Println("sl3:", sl3)
	fmt.Println("sArr:", sArr)
}

func getSum(x int, y int) int {
	return x + y
}

func getDivisionIntegerAndRest(x int, y int) (intPart int, floatPart float32, err error) {
	if y == 0 {
		return 0, 0.0, fmt.Errorf("Unable to divide by zero")
	}
	intPart = int(x / y)
	floatPart = float32(x) / float32(y)
	floatPart = floatPart - float32(intPart)
	return intPart, floatPart, nil
}

func setSumMultiple(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func sumArr(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func functionMunipulations() {
	fmt.Println("This is how functinos work:")
	fmt.Println("5 + 3 =", getSum(5, 3))
	intPart, floatPart, err := getDivisionIntegerAndRest(5, 3)
	fmt.Printf("The integer part of 5/3 is: %d\nThe decimal part of it is: %.3f\n", intPart, floatPart)
	intPart, floatPart, err = getDivisionIntegerAndRest(5, 0)
	fmt.Printf("If you're trying to do 5/0, you'll get the following errro:\n%v\n", err)

	arr1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("The sum of %v is %v\n", arr1, setSumMultiple(arr1...))
	fmt.Printf("That can also be done like: %v\n", setSumMultiple(arr1...))
}

func pointerManipulationMethod(myPtr *int) int {
	*myPtr = 12
	return *myPtr
}

func dbArrValues(arr *[4]int) {
	for x := 0; x < 4; x++ {
		arr[x] *= 2
	}
}

func getAverage(nums ...float64) float64 {
	var sum float64 = 0.0
	var numSize float64 = float64(len(nums))
	for _, val := range nums {
		sum += val
	}
	return sum / numSize
}

func pointersManipulation() {
	fmt.Println("This is how pointers work:")
	f3 := 10
	fmt.Println("f3 is originaly:", f3)
	pointerManipulationMethod(&f3)
	fmt.Println("f3 is now:", f3)

	f4 := 3
	var f4Ptr *int = &f4
	fmt.Println("f4 value is:", *f4Ptr)
	fmt.Println("f4 address is:", f4Ptr)
	pointerManipulationMethod(&f4)
	fmt.Println("f4 value now is:", f4)

	arr := [4]int{1, 2, 3, 4}
	arrBefore := arr
	dbArrValues(&arr)
	fmt.Printf("The array %v when doubled is %v\n", arrBefore, arr)

	iSlice := []float64{11, 13, 14}
	fmt.Printf("The average of %v is %v\n", iSlice, getAverage(iSlice...))
	fmt.Println("Done")
}

func guessingGame() {
	fmt.Println("This is a guessing game:")
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	for true {
		fmt.Println("Guess a number between 0 and 50:")
		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		iGuess, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}
		if iGuess > randNum {
			fmt.Println("You guessed too high")
			continue
		} else if iGuess < randNum {
			fmt.Println("You guessed too low")
			continue
		}
		fmt.Println("you got it right!")
		break
	}
}

func main() {
	dataTypes()
	casting()
	ifConditional()
	stringManipulation()
	runesAndForLoopAndPrintf()
	timeAndDate()
	someMathConcepts()
	arrayManipulations()
	slicesAndArrays()

	functionMunipulations()
	pointersManipulation()

	guessingGame()

	fmt.Println("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	userinfos := []user{
		{
			FullName: name,
			Email:    "blessing@gmail.com",
			Gender:   "Male",
			Status:   "active",
			RegDate:  "20-01-2021",
		},
	}

	jsonBytes, err := json.Marshal(userinfos)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(jsonBytes))
}
