package pjsval_test

import (
	"log"
	"os"

	"github.com/moqada/pjsval"
)

func ExampleGenerate() {
	in, err := os.Open("./example.json")
	if err != nil {
		log.Fatal(err)
		return
	}
	if err := pjsval.Generate(in, os.Stdout, "validator"); err != nil {
		log.Fatal(err)
		return
	}
	// Output:
	// package validator
	//
	// import "github.com/lestrrat/go-jsval"
	//
	// var UserGetSelf *jsval.JSVal
	// var UserPostCreate *jsval.JSVal
	// var M *jsval.ConstraintMap
	// var R0 jsval.Constraint
	// var R1 jsval.Constraint
	// var R2 jsval.Constraint
	//
	// func init() {
	// 	M = &jsval.ConstraintMap{}
	// 	R0 = jsval.String().RegexpString("^[0-9]{4}-[0-9]{2}-[0-9]{2}$")
	// 	R1 = jsval.String()
	// 	R2 = jsval.String()
	// 	M.SetReference("#/definitions/user/definitions/birthday", R0)
	// 	M.SetReference("#/definitions/user/definitions/firstName", R1)
	// 	M.SetReference("#/definitions/user/definitions/lastName", R2)
	// 	UserGetSelf = jsval.New().
	// 		SetConstraintMap(M).
	// 		SetRoot(
	// 			jsval.Object().
	// 				AdditionalProperties(
	// 					jsval.EmptyConstraint,
	// 				).
	// 				AddProp(
	// 					"fields",
	// 					jsval.String(),
	// 				),
	// 		)
	//
	// 	UserPostCreate = jsval.New().
	// 		SetConstraintMap(M).
	// 		SetRoot(
	// 			jsval.Object().
	// 				Required("birthday", "email", "firstName", "lastName").
	// 				AdditionalProperties(
	// 					jsval.EmptyConstraint,
	// 				).
	// 				AddProp(
	// 					"birthday",
	// 					jsval.Reference(M).RefersTo("#/definitions/user/definitions/birthday"),
	// 				).
	// 				AddProp(
	// 					"firstName",
	// 					jsval.Reference(M).RefersTo("#/definitions/user/definitions/firstName"),
	// 				).
	// 				AddProp(
	// 					"lastName",
	// 					jsval.Reference(M).RefersTo("#/definitions/user/definitions/lastName"),
	// 				),
	// 		)
	//
	// }
}
