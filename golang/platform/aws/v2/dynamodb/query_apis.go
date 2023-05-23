package dynamodb_lib

import (
	"context"
	"encoding/json"
	"golang/platform/aws/v2/dynamodb/models"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func defineType(model interface{}) interface{} {

	switch model.(type) {
	case models.TransactionLog:
		return &[]models.TransactionLog{}
	}

	return nil
}

func (ddbClient *DynamoDBClient) Put(item DynamoDBModel) error {

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

func (ddbClient *DynamoDBClient) Query(model DynamoDBModel, keyBuilder expression.KeyConditionBuilder, filterBuilder expression.ConditionBuilder, limit int, isAscending bool) (interface{}, error) {

	var expr expression.Expression
	var err error
	items := defineType(model)

	if model == nil {
		log.Fatal("DynamoDBModel must be set")
	}

	params := &dynamodb.QueryInput{
		TableName: aws.String(model.GetTableName()),
	}

	if !keyBuilder.IsSet() {
		log.Fatal("key must be set")
	}

	if filterBuilder.IsSet() {
		expr, err = expression.NewBuilder().WithKeyCondition(keyBuilder).WithFilter(filterBuilder).Build()
		params.FilterExpression = expr.Filter()
	} else {
		expr, err = expression.NewBuilder().WithKeyCondition(keyBuilder).Build()
	}

	// filterExpr := expression.Name("ca_addr").Equal(expression.Value(contract))
	// filterExpr = filterExpr.And(expression.Name("method").Equal(expression.Value(method)))
	// filterExpr = filterExpr.And(expression.Name("event_type").Equal(expression.Value(actionList[0])))
	// if len(actionList) > 1 {
	// 	filterExpr = filterExpr.Or(expression.Name("event_type").Equal(expression.Value(actionList[1])))
	// }

	if limit > 0 {
		params.Limit = aws.Int32(int32(limit))
	}

	params.KeyConditionExpression = expr.KeyCondition()
	params.ExpressionAttributeNames = expr.Names()
	params.ExpressionAttributeValues = expr.Values()
	params.ScanIndexForward = aws.Bool(isAscending)

	output, err := ddbClient.sess.Query(context.Background(), params)
	if err != nil {
		return nil, err
	}

	err = attributevalue.UnmarshalListOfMaps(output.Items, &items)
	if err != nil {
		log.Fatal(err.Error())
	}

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
