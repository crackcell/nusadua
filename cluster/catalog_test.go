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
 * @file catalog_test.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Tue May  9 17:01:25 2017
 *
 **/

package cluster

import (
	"fmt"
	"testing"
)

//===================================================================
// Public APIs
//===================================================================

func TestCatalog_GetInstances(t *testing.T) {
	catalog, err := NewServiceDiscover("127.0.0.1:8500")
	if err != nil {
		t.Error(err)
	}
	instances, err := catalog.GetInstances("nusadua-test")
	if err != nil {
		t.Error(err)
	}
	for _, i := range instances {
		fmt.Println(i.ID, i.Host, i.Port)
	}
}

//===================================================================
// Private
//===================================================================
