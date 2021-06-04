## Constant time statistical testing
The `go` folder implements a simple test to estimate whether modular exponentiation is (not) constant time. It compares exponentiating by 0 and by a random number.

To run the test:

1. clone the tool `dudect` with support for Golang integration (https://github.com/ansemjo/dudect).
2. remove all files in `dudect/dut/go` and copy the content of the `dudect_interface/go` directory into `dudect/dut/go`
3. replace `dudect/Makefile` with `dudect_interface/Makefile`
4. from the `dudect` directory, run `make go`
5. run the generated executable: `./dudect_go_-O2` (in MacOS, you might have to run it from inside the `dut/go` directory, with `../../dudect_go_-O2`)

To see how the native (non-constant time) `big.Int` implementation compares, edit the `do_one_computation` function in `go/dut_go.go` and call `unsafeExp` instead of `safeExp`.

Sample test outputs with `chunksize = 16, measurements = 1e4`:
<details>
  <summary><code>big.Int</code></summary>
  
```
meas:    0.01 M, max t:   +1.45, max tau: 1.45e-02, (5/tau)^2: 1.19e+05. For the moment, maybe constant time.
meas:    0.02 M, max t:  +78.32, max tau: 5.54e-01, (5/tau)^2: 8.15e+01. Probably not constant time.
meas:    0.03 M, max t: +149.81, max tau: 8.65e-01, (5/tau)^2: 3.34e+01. Probably not constant time.
meas:    0.04 M, max t: +218.69, max tau: 1.09e+00, (5/tau)^2: 2.09e+01. Probably not constant time.
meas:    0.05 M, max t: +286.15, max tau: 1.28e+00, (5/tau)^2: 1.52e+01. Probably not constant time.
meas:    0.06 M, max t: +351.30, max tau: 1.44e+00, (5/tau)^2: 1.21e+01. Probably not constant time.
meas:    0.07 M, max t: +414.97, max tau: 1.57e+00, (5/tau)^2: 1.01e+01. Probably not constant time.
meas:    0.08 M, max t: +478.99, max tau: 1.70e+00, (5/tau)^2: 8.67e+00. Probably not constant time.
meas:    0.09 M, max t: +545.65, max tau: 1.83e+00, (5/tau)^2: 7.48e+00. Definitely not constant time.
```
</details>

<details>
  <summary><code>safenum</code></summary>
  
```
meas:    0.01 M, max t:   +0.11, max tau: 1.10e-03, (5/tau)^2: 2.08e+07. For the moment, maybe constant time.
meas:    0.02 M, max t:   +0.26, max tau: 1.84e-03, (5/tau)^2: 7.37e+06. For the moment, maybe constant time.
meas:    0.03 M, max t:   +2.32, max tau: 1.43e-02, (5/tau)^2: 1.22e+05. For the moment, maybe constant time.
meas:    0.03 M, max t:   +2.20, max tau: 1.35e-02, (5/tau)^2: 1.37e+05. For the moment, maybe constant time.
meas:    0.04 M, max t:   +2.13, max tau: 1.03e-02, (5/tau)^2: 2.34e+05. For the moment, maybe constant time.
meas:    0.05 M, max t:   +2.18, max tau: 9.53e-03, (5/tau)^2: 2.75e+05. For the moment, maybe constant time.
meas:    0.06 M, max t:   +1.99, max tau: 8.05e-03, (5/tau)^2: 3.85e+05. For the moment, maybe constant time.
meas:    0.05 M, max t:   +1.71, max tau: 7.81e-03, (5/tau)^2: 4.09e+05. For the moment, maybe constant time.
meas:    0.08 M, max t:   +2.29, max tau: 8.16e-03, (5/tau)^2: 3.75e+05. For the moment, maybe constant time.
meas:    0.09 M, max t:   +2.45, max tau: 8.39e-03, (5/tau)^2: 3.55e+05. For the moment, maybe constant time.
meas:    0.09 M, max t:   +2.58, max tau: 8.41e-03, (5/tau)^2: 3.54e+05. For the moment, maybe constant time.
meas:    0.10 M, max t:   +2.24, max tau: 7.00e-03, (5/tau)^2: 5.11e+05. For the moment, maybe constant time.
meas:    0.10 M, max t:   +1.68, max tau: 5.22e-03, (5/tau)^2: 9.18e+05. For the moment, maybe constant time.
meas:    0.06 M, max t:   +2.04, max tau: 8.43e-03, (5/tau)^2: 3.52e+05. For the moment, maybe constant time.
meas:    0.12 M, max t:   +2.19, max tau: 6.30e-03, (5/tau)^2: 6.30e+05. For the moment, maybe constant time.
meas:    0.13 M, max t:   +2.28, max tau: 6.33e-03, (5/tau)^2: 6.25e+05. For the moment, maybe constant time.
meas:    0.14 M, max t:   +2.43, max tau: 6.55e-03, (5/tau)^2: 5.83e+05. For the moment, maybe constant time.
meas:    0.14 M, max t:   +2.32, max tau: 6.11e-03, (5/tau)^2: 6.70e+05. For the moment, maybe constant time.
meas:    0.15 M, max t:   +2.26, max tau: 5.79e-03, (5/tau)^2: 7.45e+05. For the moment, maybe constant time.
meas:    0.16 M, max t:   +2.13, max tau: 5.32e-03, (5/tau)^2: 8.82e+05. For the moment, maybe constant time.
meas:    0.17 M, max t:   +2.21, max tau: 5.43e-03, (5/tau)^2: 8.48e+05. For the moment, maybe constant time.
meas:    0.17 M, max t:   +2.56, max tau: 6.15e-03, (5/tau)^2: 6.60e+05. For the moment, maybe constant time.
meas:    0.18 M, max t:   +2.69, max tau: 6.31e-03, (5/tau)^2: 6.27e+05. For the moment, maybe constant time.
meas:    0.19 M, max t:   +2.19, max tau: 5.04e-03, (5/tau)^2: 9.84e+05. For the moment, maybe constant time.
meas:    0.20 M, max t:   +1.98, max tau: 4.46e-03, (5/tau)^2: 1.26e+06. For the moment, maybe constant time.
meas:    0.20 M, max t:   +1.73, max tau: 3.87e-03, (5/tau)^2: 1.67e+06. For the moment, maybe constant time.
meas:    0.21 M, max t:   +2.15, max tau: 4.72e-03, (5/tau)^2: 1.12e+06. For the moment, maybe constant time.
meas:    0.21 M, max t:   +2.39, max tau: 5.16e-03, (5/tau)^2: 9.39e+05. For the moment, maybe constant time.
meas:    0.22 M, max t:   +2.37, max tau: 5.04e-03, (5/tau)^2: 9.84e+05. For the moment, maybe constant time.
meas:    0.23 M, max t:   +2.27, max tau: 4.74e-03, (5/tau)^2: 1.12e+06. For the moment, maybe constant time.
meas:    0.24 M, max t:   +2.25, max tau: 4.63e-03, (5/tau)^2: 1.16e+06. For the moment, maybe constant time.
meas:    0.24 M, max t:   +2.43, max tau: 4.91e-03, (5/tau)^2: 1.04e+06. For the moment, maybe constant time.
meas:    0.25 M, max t:   +2.50, max tau: 4.97e-03, (5/tau)^2: 1.01e+06. For the moment, maybe constant time.
meas:    0.26 M, max t:   +2.63, max tau: 5.17e-03, (5/tau)^2: 9.36e+05. For the moment, maybe constant time.
meas:    0.27 M, max t:   +2.69, max tau: 5.22e-03, (5/tau)^2: 9.18e+05. For the moment, maybe constant time.
meas:    0.27 M, max t:   +2.57, max tau: 4.94e-03, (5/tau)^2: 1.02e+06. For the moment, maybe constant time.
meas:    0.28 M, max t:   +2.35, max tau: 4.45e-03, (5/tau)^2: 1.26e+06. For the moment, maybe constant time.
meas:    0.03 M, max t:   +2.23, max tau: 1.31e-02, (5/tau)^2: 1.45e+05. For the moment, maybe constant time.
meas:    0.02 M, max t:   +2.93, max tau: 2.07e-02, (5/tau)^2: 5.84e+04. For the moment, maybe constant time.
meas:    0.02 M, max t:   +2.80, max tau: 1.94e-02, (5/tau)^2: 6.64e+04. For the moment, maybe constant time.
meas:    0.02 M, max t:   +2.75, max tau: 1.89e-02, (5/tau)^2: 7.01e+04. For the moment, maybe constant time.
meas:    0.02 M, max t:   +2.46, max tau: 1.67e-02, (5/tau)^2: 9.00e+04. For the moment, maybe constant time.
meas:    0.02 M, max t:   +2.57, max tau: 1.72e-02, (5/tau)^2: 8.48e+04. For the moment, maybe constant time.
meas:    0.02 M, max t:   +2.49, max tau: 1.64e-02, (5/tau)^2: 9.24e+04. For the moment, maybe constant time.
meas:    0.02 M, max t:   +2.63, max tau: 1.72e-02, (5/tau)^2: 8.44e+04. For the moment, maybe constant time.
meas:    0.02 M, max t:   +2.67, max tau: 1.73e-02, (5/tau)^2: 8.40e+04. For the moment, maybe constant time.
meas:    0.02 M, max t:   +2.44, max tau: 1.55e-02, (5/tau)^2: 1.04e+05. For the moment, maybe constant time.
meas:    0.03 M, max t:   +2.44, max tau: 1.54e-02, (5/tau)^2: 1.06e+05. For the moment, maybe constant time.
meas:    0.03 M, max t:   +2.44, max tau: 1.52e-02, (5/tau)^2: 1.08e+05. For the moment, maybe constant time.
meas:    0.03 M, max t:   +2.13, max tau: 1.31e-02, (5/tau)^2: 1.46e+05. For the moment, maybe constant time.
meas:    0.03 M, max t:   +2.18, max tau: 1.33e-02, (5/tau)^2: 1.41e+05. For the moment, maybe constant time.
meas:    0.03 M, max t:   +2.06, max tau: 1.25e-02, (5/tau)^2: 1.61e+05. For the moment, maybe constant time.
meas:    0.03 M, max t:   +1.95, max tau: 1.18e-02, (5/tau)^2: 1.79e+05. For the moment, maybe constant time.
meas:    0.03 M, max t:   +1.84, max tau: 1.10e-02, (5/tau)^2: 2.07e+05. For the moment, maybe constant time.
meas:    0.03 M, max t:   +1.98, max tau: 1.17e-02, (5/tau)^2: 1.82e+05. For the moment, maybe constant time.
meas:    0.03 M, max t:   +2.15, max tau: 1.26e-02, (5/tau)^2: 1.59e+05. For the moment, maybe constant time.
meas:    0.03 M, max t:   +2.01, max tau: 1.16e-02, (5/tau)^2: 1.87e+05. For the moment, maybe constant time.
meas:    0.03 M, max t:   +1.97, max tau: 1.12e-02, (5/tau)^2: 1.99e+05. For the moment, maybe constant time.
meas:    0.03 M, max t:   +1.92, max tau: 1.08e-02, (5/tau)^2: 2.12e+05. For the moment, maybe constant time.
meas:    0.03 M, max t:   +1.68, max tau: 9.37e-03, (5/tau)^2: 2.85e+05. For the moment, maybe constant time.
meas:    0.03 M, max t:   +1.82, max tau: 1.00e-02, (5/tau)^2: 2.50e+05. For the moment, maybe constant time.
meas:    0.03 M, max t:   +1.80, max tau: 9.80e-03, (5/tau)^2: 2.60e+05. For the moment, maybe constant time.
meas:    0.03 M, max t:   +1.93, max tau: 1.04e-02, (5/tau)^2: 2.31e+05. For the moment, maybe constant time.
meas:    0.04 M, max t:   +1.82, max tau: 9.68e-03, (5/tau)^2: 2.67e+05. For the moment, maybe constant time.
meas:    0.04 M, max t:   +1.90, max tau: 1.01e-02, (5/tau)^2: 2.46e+05. For the moment, maybe constant time.
meas:    0.04 M, max t:   +1.88, max tau: 9.88e-03, (5/tau)^2: 2.56e+05. For the moment, maybe constant time.
meas:    0.04 M, max t:   +1.88, max tau: 9.88e-03, (5/tau)^2: 2.56e+05. For the moment, maybe constant time.
meas:    0.04 M, max t:   +1.82, max tau: 9.50e-03, (5/tau)^2: 2.77e+05. For the moment, maybe constant time.
meas:    0.52 M, max t:   +1.79, max tau: 2.47e-03, (5/tau)^2: 4.09e+06. For the moment, maybe constant time.
meas:    0.53 M, max t:   +1.86, max tau: 2.56e-03, (5/tau)^2: 3.80e+06. For the moment, maybe constant time.
meas:    0.04 M, max t:   +1.72, max tau: 8.87e-03, (5/tau)^2: 3.18e+05. For the moment, maybe constant time.
meas:    0.54 M, max t:   +1.78, max tau: 2.43e-03, (5/tau)^2: 4.24e+06. For the moment, maybe constant time.
meas:    0.70 M, max t:   +1.74, max tau: 2.08e-03, (5/tau)^2: 5.78e+06. For the moment, maybe constant time.
meas:    0.55 M, max t:   +1.81, max tau: 2.44e-03, (5/tau)^2: 4.21e+06. For the moment, maybe constant time.
meas:    0.56 M, max t:   +1.77, max tau: 2.36e-03, (5/tau)^2: 4.48e+06. For the moment, maybe constant time.
meas:    0.73 M, max t:   +1.76, max tau: 2.06e-03, (5/tau)^2: 5.91e+06. For the moment, maybe constant time.
meas:    0.58 M, max t:   +1.90, max tau: 2.51e-03, (5/tau)^2: 3.98e+06. For the moment, maybe constant time.
meas:    0.58 M, max t:   +1.70, max tau: 2.22e-03, (5/tau)^2: 5.06e+06. For the moment, maybe constant time.
meas:    0.59 M, max t:   +1.63, max tau: 2.12e-03, (5/tau)^2: 5.57e+06. For the moment, maybe constant time.
meas:    0.77 M, max t:   +1.78, max tau: 2.03e-03, (5/tau)^2: 6.04e+06. For the moment, maybe constant time.
meas:    0.78 M, max t:   +1.82, max tau: 2.07e-03, (5/tau)^2: 5.86e+06. For the moment, maybe constant time.
meas:    0.78 M, max t:   +1.81, max tau: 2.04e-03, (5/tau)^2: 6.00e+06. For the moment, maybe constant time.
meas:    0.79 M, max t:   +1.83, max tau: 2.05e-03, (5/tau)^2: 5.95e+06. For the moment, maybe constant time.
meas:    0.80 M, max t:   +1.84, max tau: 2.05e-03, (5/tau)^2: 5.95e+06. For the moment, maybe constant time.
meas:    0.81 M, max t:   +1.81, max tau: 2.01e-03, (5/tau)^2: 6.17e+06. For the moment, maybe constant time.
meas:    0.82 M, max t:   +1.81, max tau: 1.99e-03, (5/tau)^2: 6.30e+06. For the moment, maybe constant time.
meas:    0.83 M, max t:   +2.06, max tau: 2.26e-03, (5/tau)^2: 4.90e+06. For the moment, maybe constant time.
meas:    0.84 M, max t:   +1.85, max tau: 2.01e-03, (5/tau)^2: 6.17e+06. For the moment, maybe constant time.
meas:    0.85 M, max t:   +1.78, max tau: 1.93e-03, (5/tau)^2: 6.73e+06. For the moment, maybe constant time.
meas:    0.86 M, max t:   +1.84, max tau: 1.99e-03, (5/tau)^2: 6.33e+06. For the moment, maybe constant time.
meas:    0.87 M, max t:   +1.87, max tau: 2.01e-03, (5/tau)^2: 6.20e+06. For the moment, maybe constant time.
meas:    0.88 M, max t:   +1.89, max tau: 2.01e-03, (5/tau)^2: 6.18e+06. For the moment, maybe constant time.
meas:    0.89 M, max t:   +1.96, max tau: 2.07e-03, (5/tau)^2: 5.81e+06. For the moment, maybe constant time.
meas:    0.90 M, max t:   +2.05, max tau: 2.16e-03, (5/tau)^2: 5.36e+06. For the moment, maybe constant time.
meas:    0.91 M, max t:   +1.96, max tau: 2.06e-03, (5/tau)^2: 5.89e+06. For the moment, maybe constant time.
meas:    0.92 M, max t:   +1.93, max tau: 2.02e-03, (5/tau)^2: 6.14e+06. For the moment, maybe constant time.
meas:    0.93 M, max t:   +1.91, max tau: 1.99e-03, (5/tau)^2: 6.34e+06. For the moment, maybe constant time.
meas:    0.94 M, max t:   +2.08, max tau: 2.15e-03, (5/tau)^2: 5.41e+06. For the moment, maybe constant time.
meas:    0.94 M, max t:   +2.11, max tau: 2.17e-03, (5/tau)^2: 5.29e+06. For the moment, maybe constant time.
meas:    0.95 M, max t:   +2.17, max tau: 2.22e-03, (5/tau)^2: 5.07e+06. For the moment, maybe constant time.
meas:    0.96 M, max t:   +2.03, max tau: 2.07e-03, (5/tau)^2: 5.82e+06. For the moment, maybe constant time.
meas:    0.97 M, max t:   +1.95, max tau: 1.98e-03, (5/tau)^2: 6.38e+06. For the moment, maybe constant time.
meas:    0.98 M, max t:   +2.02, max tau: 2.04e-03, (5/tau)^2: 5.98e+06. For the moment, maybe constant time.
meas:    0.99 M, max t:   +1.94, max tau: 1.96e-03, (5/tau)^2: 6.54e+06. For the moment, maybe constant time.
meas:    1.00 M, max t:   +1.87, max tau: 1.87e-03, (5/tau)^2: 7.16e+06. For the moment, maybe constant time.
meas:    1.01 M, max t:   +1.95, max tau: 1.94e-03, (5/tau)^2: 6.61e+06. For the moment, maybe constant time.
meas:    1.02 M, max t:   +1.87, max tau: 1.86e-03, (5/tau)^2: 7.23e+06. For the moment, maybe constant time.
meas:    1.02 M, max t:   +2.02, max tau: 1.99e-03, (5/tau)^2: 6.30e+06. For the moment, maybe constant time.
meas:    1.03 M, max t:   +1.94, max tau: 1.91e-03, (5/tau)^2: 6.85e+06. For the moment, maybe constant time.
meas:    1.04 M, max t:   +1.95, max tau: 1.92e-03, (5/tau)^2: 6.82e+06. For the moment, maybe constant time.
meas:    1.05 M, max t:   +1.96, max tau: 1.92e-03, (5/tau)^2: 6.78e+06. For the moment, maybe constant time.
meas:    1.06 M, max t:   +2.15, max tau: 2.09e-03, (5/tau)^2: 5.70e+06. For the moment, maybe constant time.
meas:    1.06 M, max t:   +2.51, max tau: 2.44e-03, (5/tau)^2: 4.20e+06. For the moment, maybe constant time.
meas:    1.07 M, max t:   +2.64, max tau: 2.55e-03, (5/tau)^2: 3.85e+06. For the moment, maybe constant time.
meas:    1.08 M, max t:   +2.67, max tau: 2.56e-03, (5/tau)^2: 3.80e+06. For the moment, maybe constant time.
meas:    1.09 M, max t:   +2.68, max tau: 2.57e-03, (5/tau)^2: 3.80e+06. For the moment, maybe constant time.
meas:    1.10 M, max t:   +2.62, max tau: 2.49e-03, (5/tau)^2: 4.02e+06. For the moment, maybe constant time.
meas:    1.11 M, max t:   +2.47, max tau: 2.35e-03, (5/tau)^2: 4.55e+06. For the moment, maybe constant time.
meas:    1.12 M, max t:   +2.42, max tau: 2.29e-03, (5/tau)^2: 4.78e+06. For the moment, maybe constant time.
meas:    1.13 M, max t:   +2.50, max tau: 2.35e-03, (5/tau)^2: 4.53e+06. For the moment, maybe constant time.
meas:    1.14 M, max t:   +2.67, max tau: 2.50e-03, (5/tau)^2: 4.01e+06. For the moment, maybe constant time.
meas:    1.15 M, max t:   +2.78, max tau: 2.60e-03, (5/tau)^2: 3.71e+06. For the moment, maybe constant time.
meas:    1.16 M, max t:   +2.75, max tau: 2.56e-03, (5/tau)^2: 3.83e+06. For the moment, maybe constant time.
meas:    1.17 M, max t:   +2.68, max tau: 2.48e-03, (5/tau)^2: 4.06e+06. For the moment, maybe constant time.
meas:    1.18 M, max t:   +2.63, max tau: 2.43e-03, (5/tau)^2: 4.25e+06. For the moment, maybe constant time.
meas:    1.19 M, max t:   +2.66, max tau: 2.44e-03, (5/tau)^2: 4.19e+06. For the moment, maybe constant time.
meas:    1.20 M, max t:   +2.81, max tau: 2.57e-03, (5/tau)^2: 3.77e+06. For the moment, maybe constant time.
meas:    1.20 M, max t:   +2.65, max tau: 2.41e-03, (5/tau)^2: 4.30e+06. For the moment, maybe constant time.
meas:    1.21 M, max t:   +2.51, max tau: 2.28e-03, (5/tau)^2: 4.81e+06. For the moment, maybe constant time.
meas:    1.22 M, max t:   +2.50, max tau: 2.26e-03, (5/tau)^2: 4.90e+06. For the moment, maybe constant time.
meas:    1.23 M, max t:   +2.50, max tau: 2.25e-03, (5/tau)^2: 4.94e+06. For the moment, maybe constant time.
meas:    1.24 M, max t:   +2.31, max tau: 2.08e-03, (5/tau)^2: 5.79e+06. For the moment, maybe constant time.
meas:    1.25 M, max t:   +2.30, max tau: 2.06e-03, (5/tau)^2: 5.90e+06. For the moment, maybe constant time.
meas:    1.26 M, max t:   +2.21, max tau: 1.97e-03, (5/tau)^2: 6.46e+06. For the moment, maybe constant time.
meas:    1.27 M, max t:   +2.13, max tau: 1.89e-03, (5/tau)^2: 7.03e+06. For the moment, maybe constant time.
meas:    1.28 M, max t:   +2.17, max tau: 1.92e-03, (5/tau)^2: 6.77e+06. For the moment, maybe constant time.
meas:    1.29 M, max t:   +2.15, max tau: 1.90e-03, (5/tau)^2: 6.96e+06. For the moment, maybe constant time.
meas:    1.30 M, max t:   +2.18, max tau: 1.91e-03, (5/tau)^2: 6.82e+06. For the moment, maybe constant time.
meas:    1.31 M, max t:   +2.11, max tau: 1.84e-03, (5/tau)^2: 7.35e+06. For the moment, maybe constant time.
meas:    1.32 M, max t:   +2.19, max tau: 1.91e-03, (5/tau)^2: 6.85e+06. For the moment, maybe constant time.
meas:    1.33 M, max t:   +2.19, max tau: 1.90e-03, (5/tau)^2: 6.91e+06. For the moment, maybe constant time.
meas:    1.34 M, max t:   +2.12, max tau: 1.83e-03, (5/tau)^2: 7.45e+06. For the moment, maybe constant time.
meas:    1.34 M, max t:   +2.06, max tau: 1.78e-03, (5/tau)^2: 7.92e+06. For the moment, maybe constant time.
meas:    1.35 M, max t:   +2.05, max tau: 1.77e-03, (5/tau)^2: 8.02e+06. For the moment, maybe constant time.
meas:    1.36 M, max t:   +2.11, max tau: 1.81e-03, (5/tau)^2: 7.66e+06. For the moment, maybe constant time.
meas:    1.37 M, max t:   +2.12, max tau: 1.81e-03, (5/tau)^2: 7.63e+06. For the moment, maybe constant time.
meas:    1.38 M, max t:   +2.24, max tau: 1.91e-03, (5/tau)^2: 6.88e+06. For the moment, maybe constant time.
meas:    1.39 M, max t:   +2.36, max tau: 2.00e-03, (5/tau)^2: 6.25e+06. For the moment, maybe constant time.
meas:    1.40 M, max t:   +2.33, max tau: 1.97e-03, (5/tau)^2: 6.47e+06. For the moment, maybe constant time.
meas:    1.41 M, max t:   +2.27, max tau: 1.91e-03, (5/tau)^2: 6.84e+06. For the moment, maybe constant time.
meas:    1.42 M, max t:   +2.25, max tau: 1.89e-03, (5/tau)^2: 6.97e+06. For the moment, maybe constant time.
meas:    1.43 M, max t:   +2.23, max tau: 1.87e-03, (5/tau)^2: 7.17e+06. For the moment, maybe constant time.
meas:    1.44 M, max t:   +2.18, max tau: 1.82e-03, (5/tau)^2: 7.56e+06. For the moment, maybe constant time.
meas:    1.45 M, max t:   +2.09, max tau: 1.74e-03, (5/tau)^2: 8.30e+06. For the moment, maybe constant time.
meas:    1.46 M, max t:   +2.01, max tau: 1.67e-03, (5/tau)^2: 9.00e+06. For the moment, maybe constant time.
meas:    1.46 M, max t:   +2.15, max tau: 1.78e-03, (5/tau)^2: 7.93e+06. For the moment, maybe constant time.
meas:    1.47 M, max t:   +2.10, max tau: 1.73e-03, (5/tau)^2: 8.35e+06. For the moment, maybe constant time.
meas:    1.48 M, max t:   +2.21, max tau: 1.81e-03, (5/tau)^2: 7.60e+06. For the moment, maybe constant time.
meas:    1.49 M, max t:   +2.19, max tau: 1.79e-03, (5/tau)^2: 7.81e+06. For the moment, maybe constant time.
meas:    1.50 M, max t:   +2.17, max tau: 1.77e-03, (5/tau)^2: 7.94e+06. For the moment, maybe constant time.
meas:    1.51 M, max t:   +2.05, max tau: 1.67e-03, (5/tau)^2: 8.98e+06. For the moment, maybe constant time.
meas:    1.52 M, max t:   +2.06, max tau: 1.67e-03, (5/tau)^2: 8.98e+06. For the moment, maybe constant time.
meas:    1.53 M, max t:   +2.08, max tau: 1.68e-03, (5/tau)^2: 8.84e+06. For the moment, maybe constant time.
meas:    1.54 M, max t:   +1.96, max tau: 1.58e-03, (5/tau)^2: 9.97e+06. For the moment, maybe constant time.
meas:    1.55 M, max t:   +1.89, max tau: 1.52e-03, (5/tau)^2: 1.08e+07. For the moment, maybe constant time.
meas:    1.56 M, max t:   +1.89, max tau: 1.52e-03, (5/tau)^2: 1.09e+07. For the moment, maybe constant time.
meas:    1.57 M, max t:   +1.72, max tau: 1.38e-03, (5/tau)^2: 1.32e+07. For the moment, maybe constant time.
meas:    1.58 M, max t:   +1.67, max tau: 1.33e-03, (5/tau)^2: 1.42e+07. For the moment, maybe constant time.
meas:    1.59 M, max t:   +1.65, max tau: 1.31e-03, (5/tau)^2: 1.46e+07. For the moment, maybe constant time.
meas:    1.60 M, max t:   +1.61, max tau: 1.28e-03, (5/tau)^2: 1.54e+07. For the moment, maybe constant time.
meas:    1.61 M, max t:   +1.58, max tau: 1.24e-03, (5/tau)^2: 1.62e+07. For the moment, maybe constant time.
meas:    1.62 M, max t:   +1.60, max tau: 1.25e-03, (5/tau)^2: 1.59e+07. For the moment, maybe constant time.
meas:    1.63 M, max t:   +1.67, max tau: 1.31e-03, (5/tau)^2: 1.46e+07. For the moment, maybe constant time.
meas:    1.64 M, max t:   +1.64, max tau: 1.28e-03, (5/tau)^2: 1.52e+07. For the moment, maybe constant time.
meas:    1.65 M, max t:   +1.62, max tau: 1.27e-03, (5/tau)^2: 1.56e+07. For the moment, maybe constant time.
meas:    1.65 M, max t:   +1.50, max tau: 1.17e-03, (5/tau)^2: 1.83e+07. For the moment, maybe constant time.
meas:    1.04 M, max t:   +1.52, max tau: 1.49e-03, (5/tau)^2: 1.12e+07. For the moment, maybe constant time.
meas:    1.04 M, max t:   +1.48, max tau: 1.45e-03, (5/tau)^2: 1.19e+07. For the moment, maybe constant time.
meas:    0.86 M, max t:   +1.51, max tau: 1.62e-03, (5/tau)^2: 9.48e+06. For the moment, maybe constant time.
meas:    0.87 M, max t:   +1.41, max tau: 1.52e-03, (5/tau)^2: 1.08e+07. For the moment, maybe constant time.
meas:    0.18 M, max t:   +1.41, max tau: 3.37e-03, (5/tau)^2: 2.21e+06. For the moment, maybe constant time.
meas:    0.18 M, max t:   +1.41, max tau: 3.37e-03, (5/tau)^2: 2.21e+06. For the moment, maybe constant time.
meas:    0.18 M, max t:   +1.41, max tau: 3.37e-03, (5/tau)^2: 2.21e+06. For the moment, maybe constant time.
meas:    0.18 M, max t:   +1.41, max tau: 3.37e-03, (5/tau)^2: 2.21e+06. For the moment, maybe constant time.
meas:    0.18 M, max t:   +1.41, max tau: 3.37e-03, (5/tau)^2: 2.21e+06. For the moment, maybe constant time.
meas:    0.18 M, max t:   +1.43, max tau: 3.41e-03, (5/tau)^2: 2.15e+06. For the moment, maybe constant time.
meas:    0.18 M, max t:   +1.43, max tau: 3.41e-03, (5/tau)^2: 2.15e+06. For the moment, maybe constant time.
meas:    0.18 M, max t:   +1.43, max tau: 3.41e-03, (5/tau)^2: 2.15e+06. For the moment, maybe constant time.
meas:    0.18 M, max t:   +1.43, max tau: 3.41e-03, (5/tau)^2: 2.15e+06. For the moment, maybe constant time.
meas:    1.11 M, max t:   +1.44, max tau: 1.37e-03, (5/tau)^2: 1.34e+07. For the moment, maybe constant time.
meas:    0.18 M, max t:   +1.43, max tau: 3.41e-03, (5/tau)^2: 2.15e+06. For the moment, maybe constant time.
meas:    0.18 M, max t:   +1.43, max tau: 3.41e-03, (5/tau)^2: 2.15e+06. For the moment, maybe constant time.
meas:    1.13 M, max t:   +1.58, max tau: 1.48e-03, (5/tau)^2: 1.13e+07. For the moment, maybe constant time.
meas:    1.13 M, max t:   +1.54, max tau: 1.45e-03, (5/tau)^2: 1.19e+07. For the moment, maybe constant time.
meas:    1.14 M, max t:   +1.48, max tau: 1.39e-03, (5/tau)^2: 1.29e+07. For the moment, maybe constant time.
meas:    0.18 M, max t:   +1.43, max tau: 3.41e-03, (5/tau)^2: 2.15e+06. For the moment, maybe constant time.
meas:    0.18 M, max t:   +1.44, max tau: 3.43e-03, (5/tau)^2: 2.12e+06. For the moment, maybe constant time.
meas:    0.18 M, max t:   +1.44, max tau: 3.43e-03, (5/tau)^2: 2.12e+06. For the moment, maybe constant time.
meas:    0.18 M, max t:   +1.44, max tau: 3.43e-03, (5/tau)^2: 2.12e+06. For the moment, maybe constant time.
meas:    0.18 M, max t:   +1.44, max tau: 3.43e-03, (5/tau)^2: 2.12e+06. For the moment, maybe constant time.
meas:    0.18 M, max t:   +1.45, max tau: 3.46e-03, (5/tau)^2: 2.08e+06. For the moment, maybe constant time.
meas:    0.18 M, max t:   +1.45, max tau: 3.46e-03, (5/tau)^2: 2.08e+06. For the moment, maybe constant time.
meas:    0.18 M, max t:   +1.45, max tau: 3.46e-03, (5/tau)^2: 2.08e+06. For the moment, maybe constant time.
meas:    0.18 M, max t:   +1.45, max tau: 3.46e-03, (5/tau)^2: 2.08e+06. For the moment, maybe constant time.
meas:    0.18 M, max t:   +1.44, max tau: 3.44e-03, (5/tau)^2: 2.11e+06. For the moment, maybe constant time.
meas:    0.18 M, max t:   +1.44, max tau: 3.44e-03, (5/tau)^2: 2.11e+06. For the moment, maybe constant time.
meas:    0.18 M, max t:   +1.44, max tau: 3.44e-03, (5/tau)^2: 2.11e+06. For the moment, maybe constant time.
meas:    0.18 M, max t:   +1.44, max tau: 3.44e-03, (5/tau)^2: 2.11e+06. For the moment, maybe constant time.
meas:    0.18 M, max t:   +1.44, max tau: 3.44e-03, (5/tau)^2: 2.11e+06. For the moment, maybe constant time.
meas:    0.18 M, max t:   +1.44, max tau: 3.44e-03, (5/tau)^2: 2.11e+06. For the moment, maybe constant time.
meas:    0.18 M, max t:   +1.44, max tau: 3.44e-03, (5/tau)^2: 2.11e+06. For the moment, maybe constant time.
meas:    1.23 M, max t:   +1.45, max tau: 1.31e-03, (5/tau)^2: 1.46e+07. For the moment, maybe constant time.
meas:    0.18 M, max t:   +1.92, max tau: 4.58e-03, (5/tau)^2: 1.19e+06. For the moment, maybe constant time.
meas:    0.18 M, max t:   +1.92, max tau: 4.58e-03, (5/tau)^2: 1.19e+06. For the moment, maybe constant time.
meas:    0.18 M, max t:   +1.92, max tau: 4.58e-03, (5/tau)^2: 1.19e+06. For the moment, maybe constant time.
meas:    0.18 M, max t:   +1.92, max tau: 4.58e-03, (5/tau)^2: 1.19e+06. For the moment, maybe constant time.
```
</details>
