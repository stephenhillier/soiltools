# soiltools
Open source geotechnical analysis tools

## package: lab

The `lab` package contains tools for laboratory and index testing.

### Grain size analysis

A grain size analysis, or sieve test, helps to characterize the distribution of particles in a sample of soil. See [Sieve analysis](https://en.wikipedia.org/wiki/Sieve_analysis).

To calculate the results of a sieve test, create a variable of type SieveTest:

```go
t := lab.SieveTest{}
```

the `SieveTest` type holds some important information about the sample being tested:
```go
t.InitialMass = 2100.0 // grams
t.DryMass = 2000.0
t.WashedMass = 1900.0 
```

Add weighed sieves (collected as user input after performing the test):

```go
t.AddSieve(20, 0) // 20 mm, 0 g retained
t.AddSieve(16, 255.5) // 16 mm, 255.5 g retained
t.AddSieve(12, 172.9)
t.AddSieve(5, 194.1)
```

Adding one sieve large enough to not capture any soil is important for plotting. In our example above, the 20 mm sieve has a retained weight of 0.

Finally, calculate the results:

```go
t.Passing()
```

## Testing

To run unit tests, enter the `lab` directory and run `go test`.