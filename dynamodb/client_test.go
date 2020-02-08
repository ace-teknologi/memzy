package dynamodb

import (
	"testing"

	"github.com/ace-teknologi/memzy"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	dynamock "github.com/gusaul/go-dynamock"
)

var (
	mock *dynamock.DynaMock
	svc  dynamodbiface.DynamoDBAPI
)

func init() {
	svc, mock = dynamock.New()
}

func testClient(tableName string) *Client {
	c := New(tableName)

	c.Service = svc
	return c
}

type MegaTest struct {
	MegaString string
	MegaBool   bool
	MegaInt    int
}

func TestNew(t *testing.T) {
	c := testClient("test-table")
	if *c.TableName != "test-table" {
		t.Errorf("Expected test-table, got %s", *c.TableName)
	}
}

func mockGetItem() {
	res := dynamodb.GetItemOutput{
		Item: map[string]*dynamodb.AttributeValue{
			"MegaString": &dynamodb.AttributeValue{
				S: aws.String("ABC"),
			},
			"MegaInt": &dynamodb.AttributeValue{
				N: aws.String("123"),
			},
			"MegaBool": &dynamodb.AttributeValue{
				BOOL: aws.Bool(true),
			},
		},
	}
	key := map[string]*dynamodb.AttributeValue{
		"MegaString": {
			S: aws.String("ABC"),
		},
	}
	mock.ExpectGetItem().ToTable("test-table").WithKeys(key).WillReturns(res)
}

func mockMissingItem() {
	res := dynamodb.GetItemOutput{}
	key := map[string]*dynamodb.AttributeValue{
		"MegaString": {
			S: aws.String("missing"),
		},
	}
	mock.ExpectGetItem().ToTable("test-table").WithKeys(key).WillReturns(res)
}

func TestGetItem(t *testing.T) {
	mockGetItem()
	c := testClient("test-table")
	obj := &MegaTest{}

	err := c.GetItem(obj, map[string]interface{}{"MegaString": "ABC"})
	if err != nil {
		t.Error(err)
	}

	if obj.MegaInt != 123 {
		t.Errorf("Expected 123, got %d", obj.MegaInt)
	}

	if obj.MegaString != "ABC" {
		t.Errorf("Expected ABC, got %v", obj.MegaString)
	}

	if !obj.MegaBool {
		t.Errorf("Expected true, got %t", obj.MegaBool)
	}
}

func TestGetMissingItem(t *testing.T) {
	mockMissingItem()
	c := testClient("test-table")

	obj := &MegaTest{}

	err := c.GetItem(obj, map[string]interface{}{"MegaString": "missing"})
	if err == nil || err != ErrNotFound {
		t.Errorf("Expected ErrNotFound, got %v", err)
	}
}

func TestPutItem(t *testing.T) {
	c := testClient("test-table")

	obj := MegaTest{"test", true, 3}

	res := dynamodb.PutItemOutput{}
	mock.ExpectPutItem().ToTable("test-table").WillReturns(res)

	err := c.PutItem(obj)
	if err != nil {
		t.Error(err)
	}
}

func TestStoreImplementsMemzy(t *testing.T) {
	c := testClient("ImplementsMemzy")
	if _, ok := interface{}(c).(memzy.Memzy); !ok {
		t.Errorf("Client does not implement memzy.Memzy")
	}
}
