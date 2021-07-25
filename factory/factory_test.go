package factory

import (
	"testing"

	".main.go/assemblyspot"
	"github.com/stretchr/testify/suite"
)

type factoryUnitTestSuite struct {
	suite.Suite
	adapter *Factory
}

func (s *factoryUnitTestSuite) SetupSuite() {
	const assemblySpots int = 5
	factory := &Factory{
		AssemblingSpots: make(chan *assemblyspot.AssemblySpot, assemblySpots),
	}

	totalAssemblySpots := 0
	Id := 1
	for {
		factory.AssemblingSpots <- &assemblyspot.AssemblySpot{Id: Id}

		totalAssemblySpots++
		Id++
		if totalAssemblySpots >= assemblySpots {
			break
		}
	}

	s.adapter = factory

}

func TestFactoryUnitTestSuite(t *testing.T) {
	suite.Run(t, &factoryUnitTestSuite{})
}

func (s *factoryUnitTestSuite) TestSamble() {
	//code here
	// Assert
	expected := "Assembled"
	count := 0
	numberOfVehicles := 3

	out := s.adapter.StartAssemblingProcess(numberOfVehicles)

	for vehicle := range out {

		s.Assert().Equal(vehicle.Chassis, expected)
		s.Assert().Equal(vehicle.Dash, expected)
		s.Assert().Equal(vehicle.Electronics, expected)
		s.Assert().Equal(vehicle.Engine, expected)
		s.Assert().Equal(vehicle.Sits, expected)
		s.Assert().Equal(vehicle.Tires, expected)
		s.Assert().Equal(vehicle.Windows, expected)
		count++
	}
	s.Assert().Equal(count, numberOfVehicles)
}
