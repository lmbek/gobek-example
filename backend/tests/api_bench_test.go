package tests

import (
	"api"
	"testing"
)

func BenchmarkAPIPath_Links(benchmark *testing.B) {
	// Reset the timer before running the code
	benchmark.ResetTimer()

	// Run the code b.N times
	for i := 0; i < benchmark.N; i++ {
		api.HandleEndpoint("/api/links", nil, nil)
	}
}

func BenchmarkAPIPath_User_NoIdentifier(benchmark *testing.B) {
	// Reset the timer before running the code
	benchmark.ResetTimer()

	// Run the code b.N times
	for i := 0; i < benchmark.N; i++ {
		api.HandleEndpoint("/api/user", nil, nil)
	}
}

func BenchmarkAPIPath_User_Empty(benchmark *testing.B) {
	// Reset the timer before running the code
	benchmark.ResetTimer()

	// Run the code b.N times
	for i := 0; i < benchmark.N; i++ {
		api.HandleEndpoint("/api/user/", nil, nil)
	}
}

func BenchmarkAPIPath_User_995(benchmark *testing.B) {
	// Reset the timer before running the code
	benchmark.ResetTimer()

	// Run the code b.N times
	for i := 0; i < benchmark.N; i++ {
		api.HandleEndpoint("/api/user/995", nil, nil)
	}
}

func BenchmarkAPIPath_User_NonExisting(benchmark *testing.B) {
	// Reset the timer before running the code
	benchmark.ResetTimer()

	// Run the code b.N times
	for i := 0; i < benchmark.N; i++ {
		api.HandleEndpoint("/api/user/996", nil, nil)
	}
}

func BenchmarkAPIPath_Users(benchmark *testing.B) {
	// Reset the timer before running the code
	benchmark.ResetTimer()

	// Run the code b.N times
	for i := 0; i < benchmark.N; i++ {
		api.HandleEndpoint("/api/users", nil, nil)
	}
}

func BenchmarkAPIPath_Places(benchmark *testing.B) {
	// Reset the timer before running the code
	benchmark.ResetTimer()

	// Run the code b.N times
	for i := 0; i < benchmark.N; i++ {
		api.HandleEndpoint("/api/places", nil, nil)
	}
}
