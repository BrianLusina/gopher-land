package parkinglot

import "testing"

var largeParkingSpot = NewParkingSpot(1, 1, 1, 5, int(LARGE))
var compactParkingSpot = NewParkingSpot(1, 1, 1, 1, int(COMPACT))
var tourBus = NewBus("GopherTourBus")
var levelOne = NewLevel(0, 10, 10, []ParkingSpot{largeParkingSpot})
var levelTwo = NewLevel(1, 10, 10, []ParkingSpot{compactParkingSpot})
var levelThree = NewLevel(2, 10, 10, []ParkingSpot{compactParkingSpot})
var parkingLot = NewParkingLot(3, []Level{levelOne, levelOne})

var testCasesForParkingSpotAvailability = []struct {
	parkingSpot ParkingSpot
	expected    bool
	description string
}{
	{largeParkingSpot, true, "Should return true when no vehicle is parked"},
	{compactParkingSpot, true, "Should return true when no vehicle is parked"},
}

var testCasesForParkingSpotVehicleFit = []struct {
	parkingSpot ParkingSpot
	vehicle     Bus
	expected    bool
	description string
}{
	{largeParkingSpot, tourBus, true, "Large Parking Spot should be able to fit a bus"},
}

func TestParkingSpotAvailability(t *testing.T) {
	for _, test := range testCasesForParkingSpotAvailability {
		observed := test.parkingSpot.IsAvailable()
		if observed == test.expected {
			t.Logf("PASS: %s", test.description)
		} else {
			t.Errorf("FAIL: %s\nIsAvailable()\nExpected: %t, Actual: %t",
				test.description, test.expected, observed)
		}
	}
}

func TestParkingSpotCanFitVehicles(t *testing.T) {
	for _, test := range testCasesForParkingSpotVehicleFit {
		observed := test.vehicle.CanFitInSpot(&test.parkingSpot)
		if observed == test.expected {
			t.Logf("PASS: %s", test.description)
		} else {
			t.Errorf("FAIL: %s\nCanFitVehicle(%v)\nExpected: %t, Actual: %t",
				test.description, test.vehicle, test.expected, observed)
		}
	}
}
