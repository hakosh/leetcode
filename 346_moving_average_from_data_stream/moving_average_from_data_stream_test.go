package moving_average_from_data_stream

import "testing"

func TestMovingAverage_Next(t *testing.T) {
	stream := Constructor(3)

	if v := stream.Next(1); v != 1.0 {
		t.Errorf("Expected %v, got %v", 1.0, v)
	}

	if v := stream.Next(10); v != 5.5 {
		t.Errorf("Expected %v, got %v", 5.5, v)
	}

	if v := stream.Next(3); v != 4.666666666666667 {
		t.Errorf("Expected %v, got %v", 4.666666666666667, v)
	}

	if v := stream.Next(5); v != 6.0 {
		t.Errorf("Expected %v, got %v", 6.0, v)
	}
}

func TestMovingAverage_Next2(t *testing.T) {
	stream := Constructor(1)

	if v := stream.Next(1); v != 1.0 {
		t.Errorf("Expected %v, got %v", 1.0, v)
	}

	if v := stream.Next(10); v != 10.0 {
		t.Errorf("Expected %v, got %v", 10.0, v)
	}

	if v := stream.Next(3); v != 3.0 {
		t.Errorf("Expected %v, got %v", 3.0, v)
	}
}
