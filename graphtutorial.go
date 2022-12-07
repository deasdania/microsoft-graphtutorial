package main

import (
	"context"
	"fmt"
	"graphtutorial/graphhelper"
	"log"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Go Graph Tutorial")
	fmt.Println()

	// Load .env files
	// .env.local takes precedence (if present)
	// godotenv.Load(".env")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	graphHelper := graphhelper.NewGraphHelper()

	initializeGraph(graphHelper)

	greetUser(graphHelper)

	var choice int64 = -1

	for {
		fmt.Println("Please choose one of the following options:")
		fmt.Println("0. Exit")
		fmt.Println("1. Display access token")
		fmt.Println("2. List my inbox")
		fmt.Println("3. Send mail")
		fmt.Println("4. List users (requires app-only)")
		fmt.Println("5. Make a Graph call")

		_, err = fmt.Scanf("%d", &choice)
		if err != nil {
			choice = -1
		}

		switch choice {
		case 0:
			// Exit the program
			fmt.Println("Goodbye...")
		case 1:
			// Display access token
			displayAccessToken(graphHelper)
		case 2:
			// List emails from user's inbox
			listInbox(graphHelper)
		case 3:
			// Send an email message
			sendMail(graphHelper)
		case 4:
			// List users
			listUsers(graphHelper)
		case 5:
			// Run any Graph code
			makeGraphCall(graphHelper)
		default:
			fmt.Println("Invalid choice! Please try again.")
		}

		if choice == 0 {
			break
		}
	}
}

func initializeGraph(g *graphhelper.GraphHelper) {
	err := g.InitializeGraphForUserAuth()
	if err != nil {
		log.Panicf("Error initializing Graph for user auth: %v\n", err)
	}
}

func displayAccessToken(g *graphhelper.GraphHelper) {
	token, err := g.GetUserToken()
	if err != nil {
		log.Panicf("Error getting user token: %v\n", err)
	}

	fmt.Printf("User token: %s", *token)
	fmt.Println()
}
func listInbox(g *graphhelper.GraphHelper) {
	messages, err := g.GetInbox()
	if err != nil {
		log.Panicf("Error getting user's inbox: %v", err)
	}

	// Load local time zone
	// Dates returned by Graph are in UTC, use this
	// to convert to local
	location, err := time.LoadLocation("Local")
	if err != nil {
		log.Panicf("Error getting local timezone: %v", err)
	}

	// Output each message's details
	for x, message := range messages.GetValue() {
		fmt.Printf("Message: %s\n", *message.GetSubject())
		msgId := *message.GetId()
		// fmt.Printf("Id: %s\n", *message.GetId())
		fmt.Printf("  From: %s\n", *message.GetFrom().GetEmailAddress().GetName())

		status := "Unknown"
		if *message.GetIsRead() {
			status = "Read"
		} else {
			status = "Unread"
		}
		// att := *message.GetHasAttachments()
		// if att {
		// 	fmt.Println("has attachment for this one")
		// }
		att, err := g.GetAttachment(context.Background(), msgId)
		if err != nil {
			fmt.Printf("  ErrorGetAttachment: %s\n", err.Error())
		}
		val := att.GetValue()
		if len(val) > 0 {
			for _, v := range val {
				fmt.Printf("  type: %s\n", *v.GetContentType())
				fmt.Printf("  size: %d\n", *v.GetSize())
				fmt.Printf("  id: %s\n", *v.GetId())
				fmt.Printf("v.GetAdditionalData(): %v\n", v.GetAdditionalData())
				fmt.Printf("v.GetFieldDeserializers(): %v\n", v.GetFieldDeserializers())
				fmt.Printf("v.GetOdataType(): %v\n", *v.GetOdataType())
				fmt.Printf("v.GetName(): %v\n", *v.GetName())
				mymap := v.GetFieldDeserializers()
				keys := make([]string, len(mymap))
				// internal.NewCallRecord()
				// person := internal.NewPerson()
				// person.GetCallRecord().GetFieldDeserializers()
				i := 0
				for k := range mymap {
					keys[i] = k
					fmt.Printf("key %s\n", k)
					_, ok := mymap[k]
					if ok {
						fmt.Println("need to parse first")
						// var parseFa abs.ParseNodeFactoryRegistry
						// d, err := parseFa.GetRootParseNode(k)
						// err = z(d)
						// if err != nil {
						// 	continue
						// }
						// fmt.Println("raw value")
						// fmt.Println(d.GetRawValue())
					}
					// x, err := abs.ParseNode.GetRawValue()

					// err = z()
					// 	if k ==
					// 	parser := z(serialization.ParseNode.GetRawValue())
					// 	// b, _ := serialization.SerializationWriter.GetSerializedContent()
					// 	// raw := z(parser)
					// 	// fmt.Printf(raw)
					i++
				}

				fmt.Printf("  id: %s\n", *v.GetId())
			}
		}
		fmt.Printf("  Status: %s\n", status)
		fmt.Printf("  Received: %s\n", (*message.GetReceivedDateTime()).In(location))
		if x == 3 {
			break
		}
	}

	// If GetOdataNextLink does not return nil,
	// there are more messages available on the server
	nextLink := messages.GetOdataNextLink()

	fmt.Println()
	fmt.Printf("More messages available? %t\n", nextLink != nil)
	fmt.Println()
}

// graphClient := msgraphsdk.NewGraphServiceClient(requestAdapter)

// 	result, err := graphClient.Me().MessagesById("message-id").AttachmentsById("attachment-id").Get(context.Background(), nil)

func sendMail(g *graphhelper.GraphHelper) {
	// TODO
}

func listUsers(g *graphhelper.GraphHelper) {
	// TODO
}

func makeGraphCall(g *graphhelper.GraphHelper) {
	// TODO
}

func greetUser(g *graphhelper.GraphHelper) {
	user, err := g.GetUser()
	if err != nil {
		log.Panicf("Error getting user: %v\n", err)
	}

	fmt.Printf("Hello, %s!\n", *user.GetDisplayName())

	// For Work/school accounts, email is in Mail property
	// Personal accounts, email is in UserPrincipalName
	email := user.GetMail()
	if email == nil {
		email = user.GetUserPrincipalName()
	}

	fmt.Printf("Email: %s\n", *email)
	fmt.Println()
}

// GET /me/messages/{id}/attachments/{id}
