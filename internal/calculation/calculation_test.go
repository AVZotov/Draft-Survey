package calculation

import (
	"testing"
)

func getFreshWaterTank() FreshWaterTank {
	return FreshWaterTank{
		Name:     "test",
		Sounding: 3.5,
		Volume:   3.5,
	}
}

func getBallastWaterTank() BallastWaterTank {
	return BallastWaterTank{
		Name:     "test",
		Sounding: 3.5,
		Volume:   3.5,
		Density:  1.025,
	}
}

func TestFreshWaterTank_GetWeight(t *testing.T) {
	const weight = 3.5
	tank := getFreshWaterTank()
	tankWeight := tank.GetWeight()
	if tankWeight != weight {
		t.Errorf("Expected %f, got %f", weight, tankWeight)
	}
}

func TestBallastWaterTank_GetWeight(t *testing.T) {
	const weight = 3.587
	tank := getBallastWaterTank()
	tankWeight := round3(tank.GetWeight())
	if tankWeight != weight {
		t.Errorf("Expected %f, got %f", weight, tankWeight)
	}
}

func TestTotalFreshWater(t *testing.T) {
	const weight = 17.5
	tank := getFreshWaterTank()
	var freshWaterTanks []FreshWaterTank

	for i := 0; i < 5; i++ {
		freshWaterTanks = append(freshWaterTanks, tank)
	}
	totalWeight := TotalFreshWater(freshWaterTanks)
	if totalWeight != weight {
		t.Errorf("Expected %f, got %f", weight, totalWeight)
	}
}

func TestTotalBallastWater(t *testing.T) {
	const weight = 17.935
	tank := getBallastWaterTank()
	var ballastWaterTanks []BallastWaterTank

	for i := 0; i < 5; i++ {
		ballastWaterTanks = append(ballastWaterTanks, tank)
	}
	totalWeight := round3(TotalBallastWater(ballastWaterTanks))
	if totalWeight != weight {
		t.Errorf("Expected %f, got %f", weight, totalWeight)
	}
}
