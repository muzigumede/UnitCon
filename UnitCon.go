package main

import ("fmt"
        "strconv"
        "strings")

func main(){
    var arr [3]string
    arr[0] = "18"
    arr[1] = "cm"
    arr[2] = "mm"
    fmt.Println(fmt.Sprint(cmConvert(arr[0],arr[2]))+"mm")
}

func separate(userInput string) [3]string{
    var midIndex = strings.Index(userInput, "to")

    if midIndex == 1 {
        fmt.Println(midIndex)
    }

    var arr [3]string
    return arr
}

//convert user input string into an fixed array with [1]=number, [2]=unit, [3]=to unit
/*func makeStruct(formatedInput [3]string) conversion{
    //stuff to be done
    type conversion struct {
        value float64
        fromUnit string
        toUnit string
    }

    conversion1 := conversion{
        value: formatedInput[0],
        fromUnit: formatedInput[1],
        toUnit: formatedInput[2],
        }

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
