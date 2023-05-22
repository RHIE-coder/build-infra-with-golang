package dynamodb_lib

import "github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"

type Inputs struct {
	keyMap        map[string]interface{}
	expressionMap map[string]interface{}
	keyExpr       expression.KeyBuilder
	nameExpr      *expression.NameBuilder
	valueExpr     *expression.ValueBuilder
}

func Builder() Inputs {
	return Inputs{}
}

func (inp Inputs) Key(key string) {
	inp.keyExpr = expression.Key(key)
	return
}

func (inp Inputs) Name(name string) {
	// inp.valueExpr = inp.keyExpr.Equal()
}

func (inp Inputs) Value(value interface{}) {

}
