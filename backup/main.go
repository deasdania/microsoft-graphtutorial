package main

// "context"
// "fmt"
// "log"
// "net/url"
// "os"

// azidentity "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
// "github.com/joho/godotenv"
// a "github.com/microsoft/kiota-authentication-azure-go"
// msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
// core "github.com/microsoftgraph/msgraph-sdk-go-core"

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env")
	// }
	// cred, err := azidentity.NewDeviceCodeCredential(&azidentity.DeviceCodeCredentialOptions{
	// 	TenantID: os.Getenv("TENANT_ID"),
	// 	ClientID: os.Getenv("CLIENT_ID"),
	// 	UserPrompt: func(ctx context.Context, message azidentity.DeviceCodeMessage) error {
	// 		fmt.Println(message.Message)
	// 		return nil
	// 	},
	// })

	// if err != nil {
	// 	fmt.Printf("Error creating credentials: %v\n", err)
	// }
	// client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, []string{"Files.Read", "Mail.Read"})
	// if err != nil {
	// 	fmt.Printf("Error creating client: %v\n", err)
	// 	return
	// }
	// auth, err := a.NewAzureIdentityAuthenticationProvider(cred)
	// adapter, err := core.NewGraphRequestAdapterBaseWithParseNodeFactory(auth)
	// if err != nil {
	// 	fmt.Printf("Error creating adapter: %v\n", err)
	// 	return
	// }
	// requestInf := abs.NewRequestInformation()
	// targetUrl, err := url.Parse("https://graph.microsoft.com/v1.0/me")
	// if err != nil {
	// 	fmt.Printf("Error parsing URL: %v\n", err)
	// }
	// requestInf.SetUri(*targetUrl)
	// 	auth, err := a.NewAzureIdentityAuthenticationProviderWithScopes(cred, []string{"Mail.Read", "Mail.Send"})
	// 	if err != nil {
	// 		fmt.Printf("Error authentication provider: %v\n", err)
	// 		return
	// 	}
	// 	auth2, err := a.NewAzureIdentityAuthenticationProvider(cred)
	// 	if err != nil {
	// 		fmt.Printf("Error authentication provider: %v\n", err)
	// 		return
	// 	}

	// 	adapter, err := core.NewGraphRequestAdapterBase(auth2, core.GraphClientOptions{})
	// 	if err != nil {
	// 		fmt.Printf("Error creating adapter: %v\n", err)
	// 		return
	// 	}
	// 	requestInf := abs.NewRequestInformation()
	// 	targetUrl, err := url.Parse("https://graph.microsoft.com/v1.0/me")
	// 	if err != nil {
	// 		fmt.Printf("Error parsing URL: %v\n", err)
	// 	}
	// 	requestInf.SetUri(*targetUrl)

	// 	x := adapter.GetBaseUrl()
	// 	fmt.Println(x)
	// 	// User is your own type that implements Parsable or comes from the service library
	// 	// user, err := adapter.SendAsync(*requestInf, func() { return &User }, nil)

	// // if err != nil {
	// // 	fmt.Printf("Error getting the user: %v\n", err)
	// // }
}

// type User interface{}
