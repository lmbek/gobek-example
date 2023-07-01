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
		api.HandleEndpoint("/api/links")
	}
}

func BenchmarkAPIPath_User_NoIdentifier(benchmark *testing.B) {
	// Reset the timer before running the code
	benchmark.ResetTimer()

	// Run the code b.N times
	for i := 0; i < benchmark.N; i++ {
		api.HandleEndpoint("/api/user")
	}
}

func BenchmarkAPIPath_User_Empty(benchmark *testing.B) {
	// Reset the timer before running the code
	benchmark.ResetTimer()

	// Run the code b.N times
	for i := 0; i < benchmark.N; i++ {
		api.HandleEndpoint("/api/user/")
	}
}

func BenchmarkAPIPath_User_995(benchmark *testing.B) {
	// Reset the timer before running the code
	benchmark.ResetTimer()

	// Run the code b.N times
	for i := 0; i < benchmark.N; i++ {
		api.HandleEndpoint("/api/user/995")
	}
}

func BenchmarkAPIPath_User_NonExisting(benchmark *testing.B) {
	// Reset the timer before running the code
	benchmark.ResetTimer()

	// Run the code b.N times
	for i := 0; i < benchmark.N; i++ {
		api.HandleEndpoint("/api/user/996")
	}
}

func BenchmarkAPIPath_Users(benchmark *testing.B) {
	// Reset the timer before running the code
	benchmark.ResetTimer()

	// Run the code b.N times
	for i := 0; i < benchmark.N; i++ {
		api.HandleEndpoint("/api/users")
	}
}

func BenchmarkAPIPath_Places(benchmark *testing.B) {
	// Reset the timer before running the code
	benchmark.ResetTimer()

	// Run the code b.N times
	for i := 0; i < benchmark.N; i++ {
		api.HandleEndpoint("/api/places")
	}
}
