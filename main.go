package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"os"
	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
	"github.com/magiclabs/magic-admin-go"
	"github.com/magiclabs/magic-admin-go/client"
	"github.com/magiclabs/magic-admin-go/token"
)

type httpHandlerFunc func(http.ResponseWriter, *http.Request)
type key string

const authBearer = "Bearer"

// Load .env file from given path
var err = godotenv.Load(".env")

// Get env variables
var magicSecretKey = os.Getenv("MAGIC_SECRET_KEY")

// Instantiate Magic ✨
var magicSDK = client.New(magicSecretKey, magic.NewDefaultClient())

func checkBearerToken(next httpHandlerFunc) httpHandlerFunc {
	
	return func(res http.ResponseWriter, req *http.Request) {
		
		// Check whether or not DIDT exists in HTTP Header Request
		if !strings.HasPrefix(req.Header.Get("Authorization"), authBearer) {
			fmt.Fprintf(res, "Bearer token is required to access this route.")
			return
		}

		// Retrieve DIDT token from HTTP Header Request
		didToken := req.Header.Get("Authorization")[len(authBearer)+1:]

		// Create a Token instance to interact with the DID token
		tk, err := token.NewToken(didToken)
		if err != nil {
				fmt.Fprintf(res, "Malformed DID token error: %s", err.Error())
				res.Write([]byte(err.Error()))
				return
		}

		// Validate the Token instance before using it
		if err := tk.Validate(); err != nil {
			fmt.Fprintf(res, "DID token failed validation: %s", err.Error())
			return
	  }

		userInfo, err := magicSDK.User.GetMetadataByToken(didToken)
    if err != nil {
        fmt.Fprintf(res, "Error: %s", err.Error())
        return
    }

		fmt.Fprintf(res, "Email: %s \n", userInfo.Email)
		fmt.Fprintf(res, "Issuer: %s \n", userInfo.Issuer)
		fmt.Fprintf(res, "PublicAddress: %s \n", userInfo.PublicAddress)
		
		next(res, req)
	}
}

func handleRequests() {

	// Create a new instance of a mux router ✨
	myRouter := mux.NewRouter().StrictSlash(true)

	// Unprotected Home page ✨
	myRouter.HandleFunc("/", homePage)

	// Protected page ✨
	myRouter.HandleFunc("/protected", checkBearerToken(protectedPage))

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
		log.Printf("Defaulting to port %s", port)
	}

	if err := http.ListenAndServe(":"+port, myRouter); err != nil {
		log.Fatal(err)
	}
}

// Unprotected Page
func homePage(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Welcome to Home page!")
}

// Protected Page
func protectedPage(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "You have access to see Protected page!!!")
}

func main() {
  
	handleRequests()
}