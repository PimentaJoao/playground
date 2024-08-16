package main

import "testing"

// balanced workload
var workloads []int = []int{1, 2, 3, 4, 5, 6, 7}

// bestcase workload
// var workloads []int = []int{1}

// worstcase workload
// var workloads []int = []int{7}

var worksize int = len(workloads)

func BenchmarkIf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		daySelectorIf(workloads[i%worksize])
	}
}

func BenchmarkSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		daySelectorSwitch(workloads[i%worksize])
	}
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		daySelectorMap(workloads[i%worksize])
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		daySelectorSlice(workloads[i%worksize])
	}
}
