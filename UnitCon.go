package main

import ("fmt"
        "strconv"
        "strings"
        "regexp"
        "bufio"
        "os"
        )


func main(){
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter conversion: e.g 35cm to mm")
    userInput,_ := reader.ReadString('\n')

    inputValue, inputUnit, outputUnit, unitType := separate(userInput)

    var outputValue float64

    //check for to the corresponding function to call
    if unitType == "length" {
        outputValue = length(inputValue,inputUnit, outputUnit)

    }else if unitType == "temp" {
        outputValue = temperature(inputValue,inputUnit, outputUnit)
    }

    fmt.Println(fmt.Sprint(outputValue) + strings.ToUpper(outputUnit))
}

func separate(userInput string)(float64, string, string, string){
    var midIndex = strings.Index(userInput, "to")
    
    var inputUnit string = strings.Replace(userInput[0:midIndex]," ","",-1)
    var outputUnit string = strings.Replace(userInput[midIndex+2:len(userInput)-1]," ","",-1)
    re := regexp.MustCompile("[0-9]+")
    inputValueArray := re.FindAllString(inputUnit,-1)
    inputValue :=inputValueArray[0]
    
    inputUnit = strings.Replace(inputUnit,inputValue,"",-1)
    
    floatInputValue, err := strconv.ParseFloat(inputValue, 64)
    if err != nil{
        panic(err)
    }

    //convert all units to uppercase for precise checking in "if" statement
    inputUnit = strings.ToLower(inputUnit)
    outputUnit = strings.ToLower(outputUnit)

    unitsArr := [2]string{inputUnit,outputUnit}

    var unitType string

    //move separated input to supported units
    for i := 0;i<len(unitsArr);i++{
        switch unitsArr[i] {
            case "mm","millimetre" : unitsArr[i], unitType = "mm", "length"
            case "cm","centimetre" : unitsArr[i], unitType = "cm", "length"
            case "km","kilometre"  : unitsArr[i], unitType = "km", "length"
            case "inch"            : unitsArr[i], unitType = "inch", "length"
            case "c","celsius"     : unitsArr[i], unitType = "c", "temp"
            case "f","fahrenheit"  : unitsArr[i], unitType = "f", "temp"
            case "k","kelvin"      : unitsArr[i], unitType = "k", "temp"
        }
    }
    
    return floatInputValue, unitsArr[0], unitsArr[1], unitType
}

//handle type length conversions
func length(inputValue float64, fromUnit string, toUnit string) float64{
    var newValue float64

    if fromUnit == "cm" {
        switch toUnit {
            case "mm"  : newValue = inputValue*10
            case "inch": newValue = inputValue*0.393701
            case "m"   : newValue = inputValue/100
            case "km"  : newValue = inputValue/100000
        }
    }else if fromUnit == "mm" {
        switch toUnit {
            case "cm"  : newValue = inputValue/10
            case "inch": newValue = inputValue*0.0393701
            case "m"   : newValue = inputValue/1000
            case "km"  : newValue = inputValue/1000000
        }
    }
    return newValue
}

//handle type temperature conversions
func temperature(inputValue float64, fromUnit string, toUnit string) float64{
    var newValue float64

    if fromUnit == "c" {
        switch toUnit {
            case "f" : newValue = inputValue*9/5 + 32
            case "k" : newValue = inputValue + 273.15
        }
    } else if fromUnit == "f" {
        switch toUnit {
            case "c" : newValue = (inputValue - 32)*5/9
            case "k" : newValue = (inputValue - 32)*5/9 +273.15
        }
    }else if fromUnit == "k" {
        switch toUnit {
            case "c" : newValue = inputValue - 273.15
        }
    }

    return newValue
}

