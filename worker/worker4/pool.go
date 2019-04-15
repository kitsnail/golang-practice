package main

type Pool struct {
	Queue chan func() error
}
