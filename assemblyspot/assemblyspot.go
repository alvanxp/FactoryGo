package assemblyspot

import (
	"errors"
	"fmt"
	"sync"
	"time"

	".main.go/vehicle"
)

type AssemblySpot struct {
	Id                int
	vehicleToAssemble *vehicle.Car
	assemblyLog       string
	mutex             sync.RWMutex
}

func (s *AssemblySpot) SetVehicle(v *vehicle.Car) {
	s.vehicleToAssemble = v
}

func (s *AssemblySpot) GetAssembledVehicle() *vehicle.Car {
	return s.vehicleToAssemble
}

func (s *AssemblySpot) GetAssembledLogs() string {
	s.mutex.RLock()
	log := s.assemblyLog
	s.mutex.RUnlock()
	return log
}

//hint: improve this function to execute this process concurrenlty
func (s *AssemblySpot) AssembleVehicle() (*vehicle.Car, error) {
	if s.vehicleToAssemble == nil {
		return nil, errors.New("no vehicle set to start assembling")
	}

	var wg sync.WaitGroup
	wg.Add(7)
	go s.assembleChassis(&wg)
	go s.assembleTires(&wg)
	go s.assembleEngine(&wg)
	go s.assembleElectronics(&wg)
	go s.assembleDash(&wg)
	go s.assembleSeats(&wg)
	go s.assembleWindows(&wg)
	wg.Wait()
	return s.vehicleToAssemble, nil
}

func (s *AssemblySpot) assembleChassis(wg *sync.WaitGroup) {
	defer wg.Done()
	s.vehicleToAssemble.Chassis = "Assembled"
	time.Sleep(1 * time.Second)
	s.mutex.Lock()
	s.assemblyLog += fmt.Sprintf("Chassis at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
	s.mutex.Unlock()
}

func (s *AssemblySpot) assembleTires(wg *sync.WaitGroup) {
	defer wg.Done()
	s.vehicleToAssemble.Tires = "Assembled"
	time.Sleep(1 * time.Second)
	s.mutex.Lock()
	s.assemblyLog += fmt.Sprintf("Tires at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
	s.mutex.Unlock()
}

func (s *AssemblySpot) assembleEngine(wg *sync.WaitGroup) {
	defer wg.Done()
	s.vehicleToAssemble.Engine = "Assembled"
	time.Sleep(1 * time.Second)
	s.mutex.Lock()
	s.assemblyLog += fmt.Sprintf("Engine at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
	s.mutex.Unlock()
}

func (s *AssemblySpot) assembleElectronics(wg *sync.WaitGroup) {
	defer wg.Done()
	s.vehicleToAssemble.Electronics = "Assembled"
	time.Sleep(1 * time.Second)
	s.mutex.Lock()
	s.assemblyLog += fmt.Sprintf("Electronics at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
	s.mutex.Unlock()
}

func (s *AssemblySpot) assembleDash(wg *sync.WaitGroup) {
	defer wg.Done()
	s.vehicleToAssemble.Dash = "Assembled"
	time.Sleep(1 * time.Second)
	s.mutex.Lock()
	s.assemblyLog += fmt.Sprintf("Dash at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
	s.mutex.Unlock()
}

func (s *AssemblySpot) assembleSeats(wg *sync.WaitGroup) {
	defer wg.Done()
	s.vehicleToAssemble.Sits = "Assembled"
	time.Sleep(1 * time.Second)
	s.mutex.Lock()
	s.assemblyLog += fmt.Sprintf("Sits at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
	s.mutex.Unlock()
}

func (s *AssemblySpot) assembleWindows(wg *sync.WaitGroup) {
	defer wg.Done()
	s.vehicleToAssemble.Windows = "Assembled"
	time.Sleep(1 * time.Second)
	s.mutex.Lock()
	s.assemblyLog += fmt.Sprintf("Windows at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
	s.mutex.Unlock()
}
