package main

import "testing"

func TestProcessTruck(t *testing.T) {
	t.Run("NormalTruck should increase cargo by 2", func(t *testing.T) {
		truck := &NormalTruck{id: "1"}

		err := processTruck(truck)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if truck.cargo != 2 {
			t.Errorf("expected cargo to be 2, got %d", truck.cargo)
		}
	})

	t.Run("ElectricTruck should increase cargo by 1 and reduce battery", func(t *testing.T) {
		truck := &ElectricTruck{id: "2", battery: 10}

		err := processTruck(truck)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if truck.cargo != 1 {
			t.Errorf("expected cargo to be 1, got %d", truck.cargo)
		}

		if truck.battery != -1 {
			t.Errorf("expected battery to be -1, got %f", truck.battery)
		}
	})
}
