package main

import (
	"errors"
	"fmt"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin"
	"strconv"
	"strings"
)

var sdk wpwithin.WPWithin
var menuItems []MenuItem

type MenuItem struct {
	Label  string
	Action func() (int, error)
}

func doUI() {

	menuItems = make([]MenuItem, 0)

	menuItems = append(menuItems, MenuItem{"-------------------- GENERAL  --------------------", mInvalidSelection})
	menuItems = append(menuItems, MenuItem{"Init default device", mInitDefaultDevice})
	menuItems = append(menuItems, MenuItem{"Start RPC Service", mStartRPCService})
	menuItems = append(menuItems, MenuItem{"Init new device", mInitNewDevice})
	menuItems = append(menuItems, MenuItem{"Get device info", mGetDeviceInfo})
	menuItems = append(menuItems, MenuItem{"Sample demo, car wash (Producer)", mCarWashDemoProducer})
	menuItems = append(menuItems, MenuItem{"Sample demo, car wash (Consumer)", mCarWashDemoConsumer})
	menuItems = append(menuItems, MenuItem{"Reset session", mResetSessionState})
	menuItems = append(menuItems, MenuItem{"Load configuration", mLoadConfig})
	menuItems = append(menuItems, MenuItem{"Read loaded configuration", mReadConfig})
	menuItems = append(menuItems, MenuItem{"-------------------- PRODUCER --------------------", mInvalidSelection})
	menuItems = append(menuItems, MenuItem{"Create default producer", mDefaultProducer})
	menuItems = append(menuItems, MenuItem{"Create new producer", mNewProducer})
	menuItems = append(menuItems, MenuItem{"Add default HTE credentials", mDefaultHTECredentials})
	menuItems = append(menuItems, MenuItem{"Add new HTE credentials", mNewHTECredentials})
	menuItems = append(menuItems, MenuItem{"Add RoboWash service", mAddRoboWashService})
	menuItems = append(menuItems, MenuItem{"Add RoboAir service", mAddRoboAirService})
	menuItems = append(menuItems, MenuItem{"Initialise producer", mBroadcast})
	menuItems = append(menuItems, MenuItem{"Start service broadcast", mStartBroadcast})
	menuItems = append(menuItems, MenuItem{"Stop broadcast", mStopBroadcast})
	menuItems = append(menuItems, MenuItem{"Producer status", mProducerStatus})
	menuItems = append(menuItems, MenuItem{"-------------------- CONSUMER --------------------", mInvalidSelection})
	menuItems = append(menuItems, MenuItem{"Create default consumer", mDefaultConsumer})
	menuItems = append(menuItems, MenuItem{"Create new consumer", mNewConsumer})
	menuItems = append(menuItems, MenuItem{"Scan services", mScanService})
	menuItems = append(menuItems, MenuItem{"Create default HCE credential", mDefaultHCECredential})
	menuItems = append(menuItems, MenuItem{"Create new HCE credential", mNewHCECredential})
	menuItems = append(menuItems, MenuItem{"Discover services", mDiscoverSvcs})
	menuItems = append(menuItems, MenuItem{"Get service prices", mGetSvcPrices})
	menuItems = append(menuItems, MenuItem{"Select service", mSelectService})
	menuItems = append(menuItems, MenuItem{"Make payment", mMakePayment})
	menuItems = append(menuItems, MenuItem{"Consumer status", mConsumerStatus})
	menuItems = append(menuItems, MenuItem{"--------------------------------------------------", mInvalidSelection})
	menuItems = append(menuItems, MenuItem{"Exit", mQuit})

	renderMenu()
}
func mInvalidSelection() (int, error) {

	return 0, errors.New("*** Invalid menu selection - please choose another item ***")
}

func promptContinue() bool {

	fmt.Print("Continue (y/n): ")
	var cont string
	_, err := fmt.Scanf("%s", &cont)

	if err != nil {

		fmt.Println("Please enter \"y\" or \"n\"")
		return promptContinue()
	} else if strings.EqualFold(cont, "y") {

		return true
	} else {

		return false
	}
}

func renderMenu() {

	fmt.Println("----------------------------- Worldpay Within SDK Client ----------------------------")

	for i, item := range menuItems {

		fmt.Printf("%d - %s\n", i, item.Label)
	}

	fmt.Println("-------------------------------------------------------------------------------------")

	fmt.Print("Please select choice: ")
	var input string

	if _, err := fmt.Scanln(&input); err != nil {

		fmt.Printf("Selection error: %q\n", err.Error())
		renderMenu()
		return
	}

	inputInt, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Please type an integer choice!")
		renderMenu()
		return
	}

	if inputInt >= len(menuItems) {
		fmt.Println("Index out of bounds!")
		renderMenu()
		return
	}

	var exitCode int

	if exitCode, err = menuItems[inputInt].Action(); err != nil {

		fmt.Println(err.Error())
	}

	if exitCode != 1 {
		renderMenu()
	}
}

func setUp() {

	fmt.Println("Setup")
}

func mQuit() (int, error) {

	fmt.Println("")
	fmt.Println("Goodbye...")
	return 1, nil
}