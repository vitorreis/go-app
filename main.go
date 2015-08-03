package main

import(
  "fmt"
  "strings"
)



func main(){
  
  plants := []PowerPlant{
    PowerPlant{hydro, 300, active},
    PowerPlant{wind, 30, active},
    PowerPlant{wind, 25, inactive},
    PowerPlant{wind, 35, active},
    PowerPlant{solar, 45, unavailable},
    PowerPlant{solar, 40, inactive},
  }

  grid := PowerGrid {300, plants}

  fmt.Println("1) Generate Power Plant Report")
  fmt.Println("2) Generate Power Grid Report")
  fmt.Println("Choose something")

  var option string

  fmt.Scanln(&option)

  switch option {
  case "1":
    grid.generatePlantReport()
  case "2":
    grid.generateGridReport()
  }

}

func generatePlantCapacitiesReport(plantCapacities... float64){
    for idx, cap := range plantCapacities {
        fmt.Printf("Plan %d capacity: %0.f\n", idx, cap)
    }
}

func generatePowerGridReport(activePlants []int, plantCapacities []float64, gridLoad float64){
    capacity := 0.
    for _, plantId := range activePlants {
        capacity += plantCapacities[plantId]
    }

    fmt.Printf("%-20s%.0f\n", "Capacity: ", capacity)
    fmt.Printf("%-20s%.0f\n", "Load: ", gridLoad)
    fmt.Printf("%-20s%.1f\n", "Utilization: ", gridLoad/capacity * 100)
}

type PlantType string

const(
    hydro PlantType = "Hydro"
    wind PlantType = "Wind"
    solar PlantType = "Solar"
)

type PlantStatus string

const(
    active PlantStatus = "Active"
    inactive PlantStatus = "Inactive"
    unavailable PlantStatus = "Unavailable"
)

type PowerPlant struct {
    plantType PlantType
    capacity float64
    status PlantStatus
}

type PowerGrid struct {
    load float64
    plants []PowerPlant
}

func (pg *PowerGrid) generatePlantReport() {
    for idx, p := range pg.plants {
        label := fmt.Sprintf("%s%d", "Plant X", idx)
        fmt.Println(label)
        fmt.Println(strings.Repeat("-", len(label)))
        fmt.Printf("%-20s%s\n", "Type: ", p.plantType)
        fmt.Printf("%-20s%.0f\n", "Capacity: ", p.capacity)
        fmt.Printf("%-20s%s\n", "Status: ", p.status)
        fmt.Println()
    }
}

func (pg *PowerGrid) generateGridReport() {
    capacity := 0.
    for _, p := range pg.plants {
        if p.status == active {
            capacity += p.capacity
        }
    }

    label := "Power Grid Report"
    fmt.Println(label)
    fmt.Println(strings.Repeat("-", len(label)))

    fmt.Printf("%-20s%.0f\n", "Capacity: ", capacity)
    fmt.Printf("%-20s%.0f\n", "Load: ", pg.load)
    fmt.Printf("%-20s%.1f\n", "Utilization: ", pg.load/capacity * 100)
    fmt.Println()

}