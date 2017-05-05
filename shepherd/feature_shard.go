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
 * @file feature_shard.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Thu May  4 19:26:19 2017
 *
 **/

package shepherd

import (
	"fmt"

	"github.com/crackcell/kihaadhoo/collections/hashring"
)

//===================================================================
// Public APIs
//===================================================================

type FeatureShard struct {
	nodes      []string
	ring       *hashring.HashRing
	replicaNum int
}

func NewFeatureShard(nodes []string, replicaNum int) *FeatureShard {
	return &FeatureShard{
		nodes:      nodes,
		ring:       hashring.New(nodes),
		replicaNum: replicaNum,
	}
}

func (m *FeatureShard) GetNodesByFeature(key []int64) (nodes []string, err error) {
	if nodes, err = m.ring.GetNodes(fmt.Sprintf("%v", key), m.replicaNum); err != nil {
		return nodes, nil
	} else {
		return nodes, err
	}
}

//===================================================================
// Private
//===================================================================
