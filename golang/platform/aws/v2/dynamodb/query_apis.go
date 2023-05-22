package dynamodb_lib

import (
	"context"
	"encoding/json"
	"fmt"
	"golang/platform/aws/v2/dynamodb/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (ddbClient *DynamoDBClient) Put(item models.TransactionLog) error {

	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		return err
	}

	params := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(item.GetTableName()),
	}

	output, err := ddbClient.sess.PutItem(context.Background(), params)
	if err != nil {
		return err
	}

	_, err = json.MarshalIndent(output, "", "  ")

	return err
}

func (ddbClient *DynamoDBClient) Query() (interface{}, error) {

	var expr expression.Expression
	var err error
	keyExpr := expression.Key("txId").Equal(expression.Value("a120f758-88cb-4ae2-9dc4-006159f05624")).
		And(
			expression.Key("timestamp").GreaterThan(expression.Value(0)),
		)

	// filterExpr := expression.Name("ca_addr").Equal(expression.Value(contract))
	// filterExpr = filterExpr.And(expression.Name("method").Equal(expression.Value(method)))
	// filterExpr = filterExpr.And(expression.Name("event_type").Equal(expression.Value(actionList[0])))
	// if len(actionList) > 1 {
	// 	filterExpr = filterExpr.Or(expression.Name("event_type").Equal(expression.Value(actionList[1])))
	// }

	// expr, err = expression.NewBuilder().WithKeyCondition(keyExpr).WithFilter(filterExpr).Build()
	expr, err = expression.NewBuilder().WithKeyCondition(keyExpr).Build()
	// expr, err = expression.NewBuilder().Build()

	// limitNum, err := strconv.Atoi(limit)
	// if limitNum == 0 {
	// 	return nil, fmt.Errorf("limit number is 0")
	// }

	fmt.Println(expr.KeyCondition()) // nil
	fmt.Println(expr.Filter())       // nil
	fmt.Println(len(expr.Names()))   // 0
	fmt.Println(len(expr.Values()))  // 0

	params := &dynamodb.QueryInput{
		TableName: aws.String("TRANSACTION_LOG_DEV_BY_OWEN"),
		// FilterExpression:          expr.Filter(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		// Limit:                     aws.Int32(int32(1)),
		ScanIndexForward: aws.Bool(false),
	}

	output, err := ddbClient.sess.Query(context.Background(), params)
	if err != nil {
		return nil, err
	}

	var items []models.TransactionLog
	// var items interface{}
	attributevalue.UnmarshalListOfMaps(output.Items, &items)
	// items := make([]models.LogTxnEvent, 0, len(output.Items))

	// for _, item := range output.Items {
	// 	data := models.LogTxnEvent{}
	// 	err := attributevalue.UnmarshalMap(item, &data)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	items = append(items, data)
	// }

	return items, nil
}

// func (ddbClient *DynamoDBClient) Query2(item DynamoDBModel) ([]DynamoDBModel, error) {

// 	var expr expression.Expression
// 	var err error

// 	keyExpr := expression.Key("chain_addr").Equal(expression.Value(chainName + "#" + account)).
// 		And(
// 			expression.Key("ts_idx").GreaterThan(expression.Value("0")),
// 		)

// 	filterExpr := expression.Name("ca_addr").Equal(expression.Value(contract))
// 	filterExpr = filterExpr.And(expression.Name("method").Equal(expression.Value(method)))
// 	filterExpr = filterExpr.And(expression.Name("event_type").Equal(expression.Value(actionList[0])))
// 	if len(actionList) > 1 {
// 		filterExpr = filterExpr.Or(expression.Name("event_type").Equal(expression.Value(actionList[1])))
// 	}

// 	expr, err = expression.NewBuilder().WithKeyCondition(keyExpr).WithFilter(filterExpr).Build()
// 	limitNum, err := strconv.Atoi(limit)
// 	if limitNum == 0 {
// 		return nil, fmt.Errorf("limit number is 0")
// 	}
// 	params := &dynamodb.QueryInput{
// 		TableName:                 aws.String(models.TableNameOfLogTxnEvent()),
// 		FilterExpression:          expr.Filter(),
// 		ExpressionAttributeNames:  expr.Names(),
// 		ExpressionAttributeValues: expr.Values(),
// 		KeyConditionExpression:    expr.KeyCondition(),
// 		Limit:                     aws.Int32(int32(limitNum)),
// 		ScanIndexForward:          aws.Bool(false),
// 	}
// 	output, err := client.Query(context.TODO(), params)

// 	if err != nil {
// 		return nil, err
// 	}

// 	items := make([]models.LogTxnEvent, 0, len(output.Items))

// 	for _, item := range output.Items {
// 		data := models.LogTxnEvent{}
// 		err := attributevalue.UnmarshalMap(item, &data)
// 		if err != nil {
// 			return nil, err
// 		}
// 		items = append(items, data)
// 	}

// 	return items, nil
// }

// func UpdateContractTokenCursor(contractAddress string, chainName string, tokenType string, newCursor int) error {
// 	var err error
// 	var expr expression.Expression
// 	var update expression.UpdateBuilder

// 	update = expression.Set(expression.Name("cursor"), expression.Value(newCursor))
// 	expr, err = expression.NewBuilder().WithUpdate(update).Build()

// 	client, err := GetClient()
// 	if err != nil {
// 		return err
// 	}
// 	params := &dynamodb.UpdateItemInput{
// 		TableName: aws.String(models.TableNameContractToken()),
// 		Key: map[string]types.AttributeValue{
// 			"chain_token": &types.AttributeValueMemberS{
// 				Value: chainName + "#" + tokenType,
// 			},
// 			"ca_addr": &types.AttributeValueMemberS{
// 				Value: contractAddress,
// 			},
// 		},
// 		UpdateExpression:          expr.Update(),
// 		ExpressionAttributeNames:  expr.Names(),
// 		ExpressionAttributeValues: expr.Values(),
// 	}

// 	_, err = client.UpdateItem(context.TODO(), params)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func ScanContractToken() ([]models.ContractToken, error) {
// 	var err error

// 	client, err := GetClient()
// 	if err != nil {
// 		return nil, err
// 	}
// 	params := &dynamodb.ScanInput{
// 		TableName: aws.String(models.TableNameContractToken()),
// 	}

// 	output, err := client.Scan(context.TODO(), params)

// 	if err != nil {
// 		return nil, err
// 	}

// 	items := make([]models.ContractToken, 0, len(output.Items))

// 	for _, item := range output.Items {
// 		data := models.ContractToken{}
// 		err := attributevalue.UnmarshalMap(item, &data)
// 		if err != nil {
// 			return nil, err
// 		}
// 		items = append(items, data)
// 	}

// 	return items, nil
// }
