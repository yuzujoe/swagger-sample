/*
 * inazma-sample
 *
 * this api practice inazma
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"codegen/go/db"
	"log"

	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
	openapi "codegen/go"
)

func main() {
	log.Printf("Server started")
	db.Init()
	openapi.NewRouter()
}
