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
 * @file rpc.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Fri May  5 17:40:17 2017
 *
 **/

package shepherd

//===================================================================
// Public APIs
//===================================================================

type Rpc struct {
	addr string
	port int
}

func NewRpc(addr string, port int) *Rpc {
	return &Rpc{
		addr: addr,
		port: port,
	}
}

//===================================================================
// Private
//===================================================================
