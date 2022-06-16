package main


import("fmt")


type Vehicle struct{
    NumberOfWheels int
    NumberOfPassengers int
}

type Car struct{
    Make string
    Model string
    Year int
    IsEletric bool
    IsHybrid bool
    Vehicle Vehicle
}

func (v Vehicle) showDetails() {
    fmt.Println(" Show the number of passengers ",v.NumberOfPassengers )
    fmt.Println(" Show the number of wheels",v.NumberOfWheels )
}

func (c Car) show() {
    fmt.Println(" Make ",c.Make )
    fmt.Println(" Model",c.Model )
    fmt.Println(" Year",c.Year )
    fmt.Println(" Is eletric",c.IsEletric )
    fmt.Println(" Is hybrid",c.IsHybrid )

    c.Vehicle.showDetails() 
}


func main(){

    suv := Vehicle{
        NumberOfWheels:4,
        NumberOfPassengers:6,
    }

    volvoXC90 := Car{
        Make: "Volvo",
        Model: "xc90",
        Year: 2021,
        IsHybrid:  true ,
        IsEletric:  false ,
        Vehicle: suv,

    }

    teslaModelX := Car{
        Make: "Tesla",
        Model: "X",
        Year: 2022,
        IsHybrid:  false ,
        IsEletric:  true ,
        Vehicle: suv,

    }

    volvoXC90.show()

    fmt.Println("--------")

    teslaModelX.show()

    teslaModelX.Vehicle.NumberOfPassengers=7
    
    fmt.Println("--------")

    teslaModelX.show()

}



