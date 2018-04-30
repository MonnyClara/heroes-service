package main

import (
	"fmt"
	"github.com/chainHero/heroes-service/blockchain"
	"github.com/chainHero/heroes-service/web"
	"github.com/chainHero/heroes-service/web/controllers"
	"os"
)

func main() {
	// Definition of the Fabric SDK properties
	fSetup := blockchain.FabricSetup{
		// Channel parameters
		ChannelID:     "chainhero",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/chainHero/heroes-service/fixtures/artifacts/chainhero.channel.tx",

		// Chaincode parameters
		ChainCodeID:     "heroes-service",
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "github.com/chainHero/heroes-service/chaincode/",
		OrgAdmin:        "Admin",
		OrgName:         "Org1",
		ConfigFile:      "config.yaml",

		// User parameters
		UserName: "User1",
	}

        // Initialization of the Fabric SDK from the previously set properties
        err := fSetup.Initialize()
        if err != nil {
                fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
        }


        // Install and instantiate the chaincode
        err = fSetup.InstallAndInstantiateCC()
        if err != nil {
                fmt.Printf("Unable to install and instantiate the chaincode: %v\n", err)
        }

        response, err := fSetup.QueryState("LV1")
        if err != nil {
                fmt.Printf("Unable to query hello on the chaincode: %v\n", err)
        } else {
                fmt.Printf("Response from the query hello: %s\n", response)
        }


        // Invoke the chaincode
        txId, err2 := fSetup.InvokeHello("chainHero")
        if err2 != nil {
                fmt.Printf("Unable to invoke hello on the chaincode: %v\n", err2)
        } else {
                fmt.Printf("Successfully invoke hello, transaction ID: %s\n", txId)
        }


        response, err = fSetup.QueryState("LV1")
        if err != nil {
                fmt.Printf("Unable to query hello on the chaincode: %v\n", err)
        } else {
                fmt.Printf("Response from the query hello: %s\n", response)
        }



        // Launch the web application listening
        app := &controllers.Application{
                Fabric: &fSetup,
        }
        web.Serve(app)



}
