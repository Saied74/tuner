package main

import (
	"fmt"
	"os"
)

//csv first line for the base case with no errors
var impedance = []string{"swr", "theta", "r0", "x0", "r1", "x1", "region",
	"parallel", "series"}

//writes the header for the base case with no errors
func writeImpedanceHeader(f *os.File) error {
	for _, item := range impedance {
		_, err := f.WriteString(item + ",")
		if err != nil {
			return err
		}
	}
	return nil
}

//writes the header for the tolerance study.  It generates the text from
//the tolerance values.  G stands for Gamma, T stands for Theta
func writeToleranceHeader(f *os.File) error {
	for _, t := range tolerance {
		tt := int(t * 100)
		item := fmt.Sprintf("Region G %d,Region T %d,", tt, tt)
		_, err := f.WriteString(item)
		if err != nil {
			return err
		}
	}
	return nil
}

//header for when the actual value of L and C are calculated
//past use, may not have any future use
func writeFreqHeader(f *os.File) error {
	for _, item := range rcValues {
		_, err := f.WriteString(item + ",")
		if err != nil {
			return err
		}
	}
	return nil
}

//writes the base case (no errors) data
func (s *smith) writeImpedance(f *os.File) error {
	line := fmt.Sprintf("%.1f,%0.0f,%0.2f,%0.2f,%0.2f,%0.2f,%d,%0.2f,%0.2f,",
		s.s, s.theta, s.point0.r, s.point0.x, s.point1.r, s.point1.x, s.region, s.parallelReact, s.seriesReact)
	_, err := f.WriteString(line)
	if err != nil {
		return err
	}
	return nil
}

//write the results of the tolerance study (region number)
func (s *smith) writeTolerance(f *os.File) error {
	for _, item := range s.tolerance {
		_, err := f.WriteString(fmt.Sprintf("%d,", item.region))
		if err != nil {
			return err
		}
	}
	return nil
}

//writes the actual Ls and Cs based on freequency of bands
//past use, may not have any future use
func (s *smith) writeFreqs(f *os.File) {
	var m float64
	for i, freq := range s.freqs {
		m = 1e9
		if i%2 == 0 {
			m = 1e12
		}
		f.WriteString(fmt.Sprintf("%f,", freq*m))
	}
}