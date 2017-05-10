/***************************************************************
 *
 * Copyright (c) 2015, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the BSD licence
 *
 **************************************************************/

/**
 *
 *
 * @file flag.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Thu Jul 23 23:05:22 2015
 *
 **/

package flag

import (
	"flag"
)

//===================================================================
// Public APIs
//===================================================================

func BoolVar(p *bool, short string, long string, value bool, usage string) {
	flag.BoolVar(p, short, value, usage)
	flag.BoolVar(p, long, value, usage)
}

func Float64Var(p *float64, short string, long string, value float64, usage string) {
	flag.Float64Var(p, short, value, usage)
	flag.Float64Var(p, long, value, usage)
}

func IntVar(p *int, short string, long string, value int, usage string) {
	flag.IntVar(p, short, value, usage)
	flag.IntVar(p, long, value, usage)
}

func Int64Var(p *int64, short string, long string, value int64, usage string) {
	flag.Int64Var(p, short, value, usage)
	flag.Int64Var(p, long, value, usage)
}

func StringVar(p *string, short string, long string, value string, usage string) {
	flag.StringVar(p, short, value, usage)
	flag.StringVar(p, long, value, usage)
}

func UintVar(p *uint, short string, long string, value uint, usage string) {
	flag.UintVar(p, short, value, usage)
	flag.UintVar(p, long, value, usage)
}

func Uint64Var(p *uint64, short string, long string, value uint64, usage string) {
	flag.Uint64Var(p, short, value, usage)
	flag.Uint64Var(p, long, value, usage)
}

func Parse() {
	flag.Parse()
}

//===================================================================
// Private
//===================================================================
