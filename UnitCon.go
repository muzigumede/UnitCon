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
    fmt.Println(inputValue)
    fmt.Println(inputUnit)
    fmt.Println(outputUnit)
}

type conversion struct {
    value float64
    fromUnit string
    toUnit string
}

func separate(userInput string)(float64, string, string){
    var midIndex = strings.Index(userInput, "to")
    
    var inputUnit string = strings.Replace(userInput[0:midIndex]," ","",-1)
    var outputUnit string = strings.Replace(userInput[midIndex+2:len(userInput)]," ","",-1)
    re := regexp.MustCompile("[0-9]+")
    inputValueArray := re.FindAllString(inputUnit,-1)
    inputValue :=inputValueArray[0]
    
    inputUnit = strings.Replace(inputUnit,inputValue,"",-1)
    
    floatInputValue, err := strconv.ParseFloat(inputValue, 64)
    if err != nil{
        panic(err)
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

func cmConvert(cm string, toUnit string) float64{
    var newUnit float64 = 0.0

    if toUnit == "mm" {
        mm, err := strconv.ParseFloat(cm, 64)
        if err != nil {
            panic(err)
        }
        mm = mm*10
        newUnit = mm
    }
    return newUnit
}


