package driver

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/fy23-gw-gackathon/reportify-backend/config"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/net/context"
	"log"
)

type CognitoClient struct {
	config.Config
	*cognitoidentityprovider.Client
}

func NewCognitoClient(conf config.Config) *CognitoClient {
	cfg, err := awsConfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}
	return &CognitoClient{
		Config: conf,
		Client: cognitoidentityprovider.NewFromConfig(cfg),
	}
}

type CognitoUser struct {
	ID     string `json:"cognito:username"`
	UserID string `json:"custom:id"`
	Email  string `json:"email"`
}

func (c *CognitoClient) GetUserFromToken(ctx context.Context, idToken string) (CognitoUser, error) {
	CustomClaims := jwt.MapClaims{}

	// tokenからjwt形式へ変換する
	token, _ := jwt.ParseWithClaims(idToken, CustomClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte("test"), nil
	})

	if token == nil {
		return CognitoUser{}, errors.New("token is nil")
	}

	// 一度JSONへ変換
	jsonString, err := json.Marshal(token.Claims)
	if err != nil {
		return CognitoUser{}, err
	}

	var user CognitoUser

	if err := json.Unmarshal(jsonString, &user); err != nil {
		return CognitoUser{}, err
	}
	return user, nil
}

func (c *CognitoClient) CreateUser(ctx context.Context, userID, email string) (CognitoUser, error) {
	output, err := c.Client.AdminCreateUser(ctx, &cognitoidentityprovider.AdminCreateUserInput{
		UserPoolId:             aws.String(c.Config.Cognito.UserPoolID),
		Username:               aws.String(email),
		DesiredDeliveryMediums: []types.DeliveryMediumType{types.DeliveryMediumTypeEmail},
		UserAttributes: []types.AttributeType{
			{
				Name:  aws.String("custom:id"),
				Value: aws.String(userID),
			},
		},
	})

	if err != nil {
		return CognitoUser{}, err
	}

	return CognitoUser{
		ID:     *output.User.Username,
		Email:  email,
		UserID: userID,
	}, nil
}
