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
 * @file task.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Wed May 10 15:54:39 2017
 *
 **/

package server

//===================================================================
// Public APIs
//===================================================================

const (
	TASK_NOT_STARTED = iota
	TASK_STARTED
	TASK_RUNNING
	TASK_FAILED
	TASK_DONE
)

type Task struct {
	Name   string
	Status int
}

//===================================================================
// Private
//===================================================================
