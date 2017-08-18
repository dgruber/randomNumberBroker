package main

import (
	"fmt"
	"github.com/pivotal-cf/brokerapi"
)

func createServiceDescription() []brokerapi.Service {
	planList := []brokerapi.ServicePlan{}

	planList = append(planList, brokerapi.ServicePlan{
		ID:          "0768E956-6650-4010-8E5F-2BBED9D03031",
		Name:        "default",
		Description: "Creates a non-negative pseudo-random 63-bit integer.",
		Metadata: &brokerapi.ServicePlanMetadata{
			Bullets: []string{
				"Random numbers from an external source which serves",
				"multiple services is extremely useful!!!",
			},
			DisplayName: "RandomNumber",
		},
	})

	return []brokerapi.Service{
		brokerapi.Service{
			ID:          "FD288BB5-D6D3-479B-9F1C-2A6B21D868FA",
			Name:        "RandomNumberBroker",
			Description: "Example service broker which creates random numbers.",
			Bindable:    true,
			Plans:       planList,
			Metadata: &brokerapi.ServiceMetadata{
				DisplayName:         "Random Number Broker",
				LongDescription:     "Creates a random number when bound to an application.",
				DocumentationUrl:    "http://github.com/dgruber/randomNumberBroker/README.md",
				SupportUrl:          "http://github.com/dgruber/randomNumberBroker/README.md",
				ImageUrl:            fmt.Sprintf("data:image/png;base64,%s", "AAABAAEAEBAQAAEABAAoAQAAFgAAACgAAAAQAAAAIAAAAAEABAAAAAAAgAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAA6OjoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABEREREREQAAERERERERAAAREREREREAABEREREREQAAERERERERAAAREREREREAABEREREREQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADwDwAA8A8AAP5/AAD+fwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAD//wAA"),
				ProviderDisplayName: "Example",
			},
			Tags: []string{
				"random",
				"example",
			},
		},
	}
}

func planExists(planID string) bool {
	return planID == "0768E956-6650-4010-8E5F-2BBED9D03031"
}
