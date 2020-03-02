package main

import ("fmt"
        "strconv"
        "strings"
        "regexp"
        "bufio"
        "os"
        //"reflect"
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

type conversion struct {
    value float64
    fromUnit string
    toUnit string
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

    var unitType string
    //move separated input to supported units
    if(strings.Contains(inputUnit, "cm") || strings.Contains(inputUnit, "centimet")){
        inputUnit = "cm" 
        unitType = "length"
    }else if(strings.Contains(inputUnit, "mm") || strings.Contains(inputUnit, "millimet")){
        inputUnit = "mm"
        unitType = "length"
    }else if(strings.Contains(inputUnit, "c") || strings.Contains(inputUnit, "celsius")){
        inputUnit = "c"
        unitType = "temp"
    }else if(strings.Contains(inputUnit, "f") || strings.Contains(inputUnit, "fahrenheit")){
        inputUnit = "f"
        unitType = "temp"
    }else if(strings.Contains(inputUnit, "k") || strings.Contains(inputUnit, "kelvin")){
        inputUnit = "k"
        unitType = "temp"
    }

    return floatInputValue, inputUnit, outputUnit, unitType
}

func length(inputValue float64, fromUnit string, toUnit string) float64{
    var newValue float64
    var i string = toUnit

    if fromUnit == "cm" {
        switch i {
            case "mm"  : newValue = inputValue*10
            case "inch": newValue = inputValue*0.393701
            case "m"   : newValue = inputValue/100
            case "km"  : newValue = inputValue/100000
        }
    }else if fromUnit == "mm" {
        switch i {
            case "cm"  : newValue = inputValue/10
            case "inch": newValue = inputValue*0.0393701
            case "m"   : newValue = inputValue/1000
            case "km"  : newValue = inputValue/1000000
        }
    }
    return newValue
}

func temperature(inputValue float64, fromUnit string, toUnit string) float64{
    var newValue float64
    var i string = toUnit

    if fromUnit == "c" {
        switch i {
            case "f" : newValue = inputValue*9/5 + 32
            case "k" : newValue = inputValue + 273.15
        }
    } else if fromUnit == "f" {
        switch i {
            case "c" : newValue = (inputValue - 32)*5/9
            case "k" : newValue = (inputValue - 32)*5/9 +273.15
        }
    }else if fromUnit == "k" {
        switch i {
            case "c" : newValue = inputValue - 273.15
        }
    }
    return newValue
}

