# Go lookup methods benchmark
Benchmarking various ways of translating an integer into a day of the week.

### How?

Currently, there are four different types of lookup methods implemented, `if/else`, `switch/case`, `Map` and `Slice`. Each method has it's own function which receives a integer and returns the corresponding day of the week (e.g. 1 returns Sunday, 4 returns Wednesday, and so on). Using the built-in Golang benchmarking solution, it was possible to see how long our application takes to resolve a day of the week.

There are three different types of workloads. The balanced workload offers the same lookup possibility for all week days. The bestcase workload only has Sunday as a lookup possibility, meaning that the first lookup is always the one desired. The worstcase workload only has Saturday as a lookup possibility, following the same principle as the bestcase workload.

### Results

For the *balanced workload*, we get the following results:

| Lookup Method | Throughput (ns/op) |
|---|:-:|
| `If` | 7.49 |
| `Switch/case` | 7.38 |
| `Map` | 14.17 |
| `Slice` | 7.72 |

Now, for the *bestcase workload* results, we were presented with:

| Lookup Method | Throughput (ns/op) |
|---|:-:|
| `If` | 7.47 |
| `Switch/case` | 7.49 |
| `Map` | 12.64 |
| `Slice` | 7.44 |

At last, considering the *worstcase workload* results:

| Lookup Method | Throughput (ns/op) |
|---|:-:|
| `If` | 7.45 |
| `Switch/case` | 7.77 |
| `Map` | 15.15 |
| `Slice` | 7.65 |

### Why?

This was done as a practical learning exercise to understand more about both the built-in benchmarking tool and about the different speeds of different lookup methods. Take my results with a grain of salt as this is just a learning focused project, not a scientific one.

This project was inspired by RaiTamarindo's [Atoi x Regex for integer validation](https://github.com/RaiTamarindo/atoi-x-regex-benchmark).
