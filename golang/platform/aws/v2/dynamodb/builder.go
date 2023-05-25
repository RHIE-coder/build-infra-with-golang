package dynamodb_lib

import "github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"

func KeyEqual(key string, value interface{}) expression.KeyConditionBuilder {
	return expression.KeyEqual(expression.Key(key), expression.Value(value))
}

/*
Equal             -,NameBuilder,ValueBuilder,SizeBuilder  ==> ConditionBuilder
NotEqual          -,NameBuilder,ValueBuilder,SizeBuilder  ==> ConditionBuilder
LessThan          -,NameBuilder,ValueBuilder,SizeBuilder  ==> ConditionBuilder
LessThanEqual     -,NameBuilder,ValueBuilder,SizeBuilder  ==> ConditionBuilder
GreaterThan       -,NameBuilder,ValueBuilder,SizeBuilder  ==> ConditionBuilder
GreaterThanEqual  -,NameBuilder,ValueBuilder,SizeBuilder  ==> ConditionBuilder
And               -,ConditionBuilder                      ==> ConditionBuilder
Or                -,ConditionBuilder                      ==> ConditionBuilder
Not               -,ConditionBuilder                      ==> ConditionBuilder
Between           -,NameBuilder,ValueBuilder,SizeBuilder  ==> ConditionBuilder
In                -,NameBuilder,ValueBuilder,SizeBuilder  ==> ConditionBuilder
BeginsWith        -,NameBuilder                           ==> ConditionBuilder
Contains          -,NameBuilder                           ==> ConditionBuilder
*/
