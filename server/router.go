/***************************************************************
 *
 * Copyright (c) 2017, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the Apache licence
 *
 **************************************************************/

/**
 *
 *
 * @file router.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Wed May 10 17:24:58 2017
 *
 **/

package server

import (
	"github.com/crackcell/kihaadhoo/collections/hashring"
)

//===================================================================
// Public APIs
//===================================================================

type Router struct {
	nodes   []string
	ring    *hashring.HashRing
	replica int
}

func NewRouter() *Router {
	return &Router{}
}

//===================================================================
// Private
//===================================================================
