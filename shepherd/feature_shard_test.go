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
 * @file feature_shard_test.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Thu May  4 19:48:51 2017
 *
 **/

package shepherd

import (
	"fmt"
	"testing"
)

//===================================================================
// Public APIs
//===================================================================

func TestMetaGetNodesByFeature(t *testing.T) {
	nodes := []string{
		"127.0.0.1:1988",
		"127.0.0.1:1989",
		"127.0.0.1:1990",
	}
	key := []int64{1, 2, 3}
	meta := NewMeta(nodes, 3)
	fmt.Println(meta.GetNodesByFeature(key))
}

//===================================================================
// Private
//===================================================================
