package mailsend

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

type GraphHelper struct {
	clientSecretCredential *azidentity.ClientSecretCredential
}

func NewGraphHelper() *GraphHelper {
	g := &GraphHelper{}
	return g
}

func (g *GraphHelper) InitializeGraphForAppAuth(
	tenantId string,
	clientId string,
	clientSecret string,
) error {
	credential, err := azidentity.NewClientSecretCredential(tenantId, clientId, clientSecret, nil)
	if err != nil {
		return err
	}

	g.clientSecretCredential = credential
	return nil
}

func (g *GraphHelper) GetAppToken() (*string, error) {
	token, err := g.clientSecretCredential.GetToken(
		context.Background(),
		policy.TokenRequestOptions{
			Scopes: []string{
				"https://graph.microsoft.com/.default",
			},
		},
	)
	if err != nil {
		return nil, err
	}

	return &token.Token, nil
}
