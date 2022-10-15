// Michael Valdron, Copyright 2022
//
// Package db provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package db

// MessageResponse defines model for messageResponse.
type MessageResponse struct {
	// An error message block
	Error *string `json:"error,omitempty"`

	// A descriptive message
	Message string `json:"message"`
}

// Student defines model for student.
type Student struct {
	// Student email
	Email StudentEmailAddress `json:"email"`

	// Student Unique identifier
	Id StudentId `json:"id"`

	// Name of student
	Name string `json:"name"`
}

// Student email
type StudentEmailAddress = string

// Student Unique identifier
type StudentId = int64
