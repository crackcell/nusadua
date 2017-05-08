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
 * @file rpc_test.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Mon May  8 15:01:39 2017
 *
 **/

package router

import (
	"fmt"
	"testing"
	"time"
)

//===================================================================
// Public APIs
//===================================================================

func TestRpc(t *testing.T) {
	r := NewRpc()
	r.Start("127.0.0.1", 9099)
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Println("waiting", i)
			time.Sleep(time.Second)
		}
		r.Stop()
	}()
	r.Wait()
}

//===================================================================
// Private
//===================================================================
