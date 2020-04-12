package main

import ("fmt"
        "strconv"
        "strings"
        "regexp"
        "bufio"
        "os"
        )


type conversion struct {
    value       float64
    inputUnit   string
    outputUnit  string
    unitType    string
}

/*still to implement:
        -Base converter
        -Mass converter
        -Volume converter
*/

func main(){

    //ask user for input
    fmt.Println("Enter conversion: e.g 35cm to mm")

    //store input from the user
    reader       := bufio.NewReader(os.Stdin)
    userInput, _ := reader.ReadString('\n')

    convObj      := conversion{}

    separate(&convObj, userInput)

    //check for to the corresponding function to call
    if convObj.unitType == "length" {
        convertLength(&convObj.value, convObj.inputUnit, convObj.outputUnit)

    } else if convObj.unitType == "temp" {
        convertTemperature(&convObj.value, convObj.inputUnit, convObj.outputUnit)
    }

    fmt.Println(fmt.Sprint(convObj.value) + strings.ToUpper(convObj.outputUnit))
}

func separate(conv *conversion, userInput string){
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

    unitsArr := []string{inputUnit, outputUnit}

    
    //move separated input to supported measure units
    for i := 0;i<len(unitsArr);i++{
        switch unitsArr[i] {
            case "mm","millimetre" : unitsArr[i], conv.unitType = "mm", "length"
            case "cm","centimetre" : unitsArr[i], conv.unitType = "cm", "length"
            case "km","kilometre"  : unitsArr[i], conv.unitType = "km", "length"
            case "inch"            : unitsArr[i], conv.unitType = "inch", "length"
            case "c","celsius"     : unitsArr[i], conv.unitType = "c", "temp"
            case "f","fahrenheit"  : unitsArr[i], conv.unitType = "f", "temp"
            case "k","kelvin"      : unitsArr[i], conv.unitType = "k", "temp"
        }
    }
    
    conv.value      = floatInputValue
    conv.inputUnit  = inputUnit
    conv.outputUnit = outputUnit
}

//handle type length conversions
func convertLength(value *float64, fromUnit string, toUnit string) {

    if fromUnit == "cm" {
        switch toUnit {
            case "mm"  : *value = *value*10
            case "inch": *value = *value*0.393701
            case "m"   : *value = *value/100
            case "km"  : *value = *value/100000
        }

    } else if fromUnit == "mm" {
        switch toUnit {
            case "cm"  : *value = *value/10
            case "inch": *value = *value*0.0393701
            case "m"   : *value = *value/1000
            case "km"  : *value = *value/1000000
        }
    }
}

//handle type temperature conversions
func convertTemperature(value *float64, fromUnit string, toUnit string) {

    if fromUnit == "c" {
        switch toUnit {
            case "f" : *value = *value*9/5 + 32
            case "k" : *value = *value + 273.15
        }

    } else if fromUnit == "f" {
        switch toUnit {
            case "c" : *value = (*value - 32)*5/9
            case "k" : *value = (*value - 32)*5/9 +273.15
        }

    } else if fromUnit == "k" {
        switch toUnit {
            case "c" : *value = *value - 273.15
        }
    }
}


