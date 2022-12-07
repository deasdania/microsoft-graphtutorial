package graphhelper

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	auth "github.com/microsoft/kiota-authentication-azure-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/me"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

type GraphHelper struct {
	deviceCodeCredential   *azidentity.DeviceCodeCredential
	userClient             *msgraphsdk.GraphServiceClient
	graphUserScopes        []string
	clientSecretCredential *azidentity.ClientSecretCredential
	appClient              *msgraphsdk.GraphServiceClient
}

func NewGraphHelper() *GraphHelper {
	g := &GraphHelper{}
	return g
}
func (g *GraphHelper) EnsureGraphForAppOnlyAuth() error {
    if g.clientSecretCredential == nil {
        clientId := os.Getenv("CLIENT_ID")
        tenantId := os.Getenv("TENANT_ID")
        clientSecret := os.Getenv("CLIENT_SECRET")
        credential, err := azidentity.NewClientSecretCredential(tenantId, clientId, clientSecret, nil)
        if err != nil {
            return err
        }

        g.clientSecretCredential = credential
    }

    if g.appClient == nil {
        // Create an auth provider using the credential
        authProvider, err := auth.NewAzureIdentityAuthenticationProviderWithScopes(g.clientSecretCredential, []string{
            "https://graph.microsoft.com/.default",
        })
        if err != nil {
            return err
        }

        // Create a request adapter using the auth provider
        adapter, err := msgraphsdk.NewGraphRequestAdapter(authProvider)
        if err != nil {
            return err
        }

        // Create a Graph client using request adapter
        client := msgraphsdk.NewGraphServiceClient(adapter)
        g.appClient = client
    }

    return nil
}

func (g *GraphHelper) InitializeGraphForUserAuth() error {
	clientId := os.Getenv("CLIENT_ID")
	authTenant := os.Getenv("AUTH_TENANT")
	scopes := os.Getenv("GRAPH_USER_SCOPES")
	g.graphUserScopes = strings.Split(scopes, ",")

	// Create the device code credential
	credential, err := azidentity.NewDeviceCodeCredential(&azidentity.DeviceCodeCredentialOptions{
		ClientID: clientId,
		TenantID: authTenant,
		UserPrompt: func(ctx context.Context, message azidentity.DeviceCodeMessage) error {
			fmt.Println(message.Message)
			return nil
		},
	})
	if err != nil {
		return err
	}

	g.deviceCodeCredential = credential

	// Create an auth provider using the credential
	authProvider, err := auth.NewAzureIdentityAuthenticationProviderWithScopes(credential, g.graphUserScopes)
	if err != nil {
		return err
	}

	// Create a request adapter using the auth provider
	adapter, err := msgraphsdk.NewGraphRequestAdapter(authProvider)
	if err != nil {
		return err
	}

	// Create a Graph client using request adapter
	client := msgraphsdk.NewGraphServiceClient(adapter)
	g.userClient = client

	return nil
}
func (g *GraphHelper) GetUserToken() (*string, error) {
	token, err := g.deviceCodeCredential.GetToken(context.Background(), policy.TokenRequestOptions{
		Scopes: g.graphUserScopes,
	})
	if err != nil {
		return nil, err
	}

	return &token.Token, nil
}

func (g *GraphHelper) GetUser() (models.Userable, error) {
	query := me.MeRequestBuilderGetQueryParameters{
		// Only request specific properties
		Select: []string{"displayName", "mail", "userPrincipalName"},
	}

	return g.userClient.Me().Get(context.Background(),
		&me.MeRequestBuilderGetRequestConfiguration{
			QueryParameters: &query,
		})
}
func (g *GraphHelper) GetInbox() (models.MessageCollectionResponseable, error) {
	var topValue int32 = 25
	q := me.MeMailFoldersItemMessagesRequestBuilderGetQueryParameters{
		Select:  []string{"from", "isRead", "receivedDateTime", "subject"},
		Top:     &topValue,
		Orderby: []string{"receivedDateTime DESC"},
	}
	query := me.MeMailFoldersItemChildFoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilderGetQueryParameters{
		Select: []string{"from", "isRead", "receivedDateTime", "subject"},
	}
	fmt.Println(query)
	return g.userClient.Me().
		MailFoldersById("inbox").
		Messages().
		Get(context.Background(),
			&me.MeMailFoldersItemMessagesRequestBuilderGetRequestConfiguration{
				QueryParameters: &q,
			})
}

func (g *GraphHelper) GetAttachment(ctx context.Context, msgId string) (models.AttachmentCollectionResponseable, error) {
	config := &me.MeMessagesItemAttachmentsRequestBuilderGetRequestConfiguration{
		QueryParameters: &me.MeMessagesItemAttachmentsRequestBuilderGetQueryParameters{},
	}
	result, err := g.userClient.Me().MessagesById(msgId).Attachments().Get(ctx, config)
	if err != nil {
		return nil, err
	}
	fmt.Println(result.GetOdataNextLink())
	return result, nil
}

// func (g *GraphHelper) GetAttachmentRaw(ctx context.Context, msgId string, attId string) (models.AttachmentCollectionResponseable, error) {
// 	result, err := g.userClient.Me().MessagesById(msgId).AttachmentsById(attId).Get(ctx, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	result.GetFieldDeserializers()

// }

// func (g *GraphHelper) GetDownloadUrl(ctx context.Context, attId string, msgId string) (models.Attachmentable, error) {
// 	result, err := g.userClient.Me().MessagesById(msgId).AttachmentsById(attId).Get(ctx, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// result.GetOdataType()
// 	// result.Get
// 	// result.GetAdditionalData()
// 	return result, nil
// }

// func (g *GraphHelper) GetInboxByUserId(userId string) (models.MessageCollectionResponseable, error) {
// 	return g.userClient.Me().
// 		MailFoldersById("inbox").
// 		Messages().
// 		Get(context.Background(),
// 			&me.MeMailFoldersItemMessagesRequestBuilderGetRequestConfiguration{
// 				QueryParameters: &q,
// 			})
// }
