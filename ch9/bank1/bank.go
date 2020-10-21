// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

import "fmt"

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func Deposit(amount int) { deposits <- amount } //往通道写数据
func Balance() int       { return <-balances }  //从通道取数据

//使用 channel 对控制共享数据  balance变量被限制在了monitor goroutine中，名为 teller：
func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits: //从通道取数据
			balance += amount
		case balances <- balance: //往通道写数据
			fmt.Println("账户余额写入balances channel")
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
