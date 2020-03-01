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

    inputValue, inputUnit, outputUnit := separate(userInput)

    var outputValue float64

    //check for to the corresponding function to call
    if inputUnit == "cm" {
        outputValue = convertCM(inputValue, outputUnit)
    }else if inputUnit == "mm" {
        outputValue = convertMM(inputValue, outputUnit)
    }else if inputUnit == "c" {
        outputValue = convertCelsius(inputValue, outputUnit)
    }

    fmt.Println(fmt.Sprint(outputValue) + strings.ToUpper(outputUnit))
}

type conversion struct {
    value float64
    fromUnit string
    toUnit string
}

func separate(userInput string)(float64, string, string){
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

    //move separated input to supported units
    if(strings.Contains(inputUnit, "cm") || strings.Contains(inputUnit, "centimet")){
        inputUnit = "cm"    
    }else if(strings.Contains(inputUnit, "mm") || strings.Contains(inputUnit, "millimet")){
        inputUnit = "mm"
    }else if(strings.Contains(inputUnit, "c") || strings.Contains(inputUnit, "celsius")){
        inputUnit = "c"
    }

    return floatInputValue, inputUnit, outputUnit
}

//convert user input string into an fixed array with [1]=number, [2]=unit, [3]=to unit
/*func makeStruct(formatedInput [3]string) conversion{
    //stuff to be done

    conversion1 := conversion{
        value: strconv.ParseFloat(formatedInput[0],
        fromUnit: formatedInput[1],
        toUnit: formatedInput[2]}

    return conversion1
}*/

func convertCM(cm float64, toUnit string) float64{
    var newUnit float64 = 0.0

    if toUnit == "mm" {
        mm := cm*10
        newUnit = mm
    }else if toUnit == "m" {
        m := cm/100
        newUnit = m
    }else if toUnit == "km" {
        km := cm/100000
        newUnit = km
    }
    return newUnit
}

func convertMM(mm float64, toUnit string) float64{
    var newUnit float64 = 0.0

    if strings.Contains(toUnit, "cm") {
        cm := mm/10
        newUnit = cm
    }
    return newUnit
}

func convertCelsius(c float64, toUnit string) float64{
    var newUnit float64 

    if strings.Contains(toUnit, "f") {
        f := c*9/5 + 32
        newUnit = f
    }
    return newUnit
}

