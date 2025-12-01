//Barrier.go Template Code
//Copyright (C) 2024 Dr. Joseph Kehoe

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by:
// Description:
// A simple barrier implemented using mutex and unbuffered channel
// Issues:
// None I hope
//1. Change mutex to atomic variable
//2. Make it a reusable barrier
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, arrived *atomic.Int64, totalRoutines int, wg *sync.WaitGroup, theChan chan bool) bool {
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)
	//we wait here until everyone has completed part A
	arrived.Add(1)
	if arrived.Load() == int64(totalRoutines) { //last to arrive -signal others to go
		theChan <- true
		<-theChan
		arrived.Store(0)
	} else { //not all here yet we wait until signal
		<-theChan
		theChan <- true //once we get through send signal to next routine to continue
	} //end of if-else
	fmt.Println("PartB", goNum)
	wg.Done()
	return true
} //end-doStuff

func main() {
	totalRoutines := 10
	var wg sync.WaitGroup
	var arrived atomic.Int64
	wg.Add(totalRoutines)
	//we will need some of these
	theChan := make(chan bool, 1)        //use unbuffered channel in place of semaphore
	for i := 0; i < totalRoutines; i++ { //create the go Routines here
		go doStuff(i, &arrived, totalRoutines, &wg, theChan)
	}
	wg.Wait() //wait for everyone to finish before exiting
	fmt.Println("Firstly")

	wg.Add(totalRoutines)
	for i := range totalRoutines {
		go doStuff(i, &arrived, totalRoutines, &wg, theChan)
	}
	wg.Wait()
	fmt.Println("Secondly")
} //end-main
