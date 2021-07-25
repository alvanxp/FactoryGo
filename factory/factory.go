package factory

import (
	"sync"

	".main.go/assemblyspot"
	".main.go/vehicle"
)

const assemblySpots int = 5

type Factory struct {
	AssemblingSpots chan *assemblyspot.AssemblySpot
}

func New() *Factory {
	factory := &Factory{
		AssemblingSpots: make(chan *assemblyspot.AssemblySpot, assemblySpots),
	}

	totalAssemblySpots := 0
	index := 1
	for {
		factory.AssemblingSpots <- &assemblyspot.AssemblySpot{Id: index}

		totalAssemblySpots++
		index++
		if totalAssemblySpots >= assemblySpots {
			break
		}
	}

	return factory
}

//HINT: this function is currently not returning anything, make it return right away every single vehicle once assembled,
//(Do not wait for all of them to be assembled to return them all, send each one ready over to main)
func (f *Factory) StartAssemblingProcess(amountOfVehicles int) <-chan *vehicle.Car {

	in := f.generateVehicleLots(amountOfVehicles)
	out := make(chan *vehicle.Car, assemblySpots)
	go func() {
		defer close(out)
		var wg sync.WaitGroup
		wg.Add(amountOfVehicles)

		for vehicleToAssembly := range in {
			go f.assemblyVehicle(vehicleToAssembly, out, &wg)
		}
		wg.Wait()
	}()
	return out
}

func (f *Factory) assemblyVehicle(vehicle vehicle.Car, out chan *vehicle.Car, wg *sync.WaitGroup) {
	defer wg.Done()
	idleSpot := <-f.AssemblingSpots
	idleSpot.SetVehicle(&vehicle)
	idleSpot.AssembleVehicle()
	vehicle.TestingLog = f.testCar(&vehicle)
	vehicle.AssembleLog = idleSpot.GetAssembledLogs()
	idleSpot.SetVehicle(nil)
	assembledVehicle := &vehicle
	f.AssemblingSpots <- idleSpot
	out <- assembledVehicle
}

func (f *Factory) generateVehicleLots(amountOfVehicles int) <-chan vehicle.Car {

	out := make(chan vehicle.Car)

	go func() {
		for index := 0; index < amountOfVehicles; index++ {
			out <- vehicle.Car{
				Id:            index,
				Chassis:       "NotSet",
				Tires:         "NotSet",
				Engine:        "NotSet",
				Electronics:   "NotSet",
				Dash:          "NotSet",
				Sits:          "NotSet",
				Windows:       "NotSet",
				EngineStarted: false,
			}
		}
		close(out)
	}()

	return out
}

func (f *Factory) testCar(car *vehicle.Car) string {
	logs := ""

	log, err := car.StartEngine()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.MoveForwards(10)
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.MoveForwards(10)
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.TurnLeft()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.TurnRight()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.StopEngine()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	return logs
}
