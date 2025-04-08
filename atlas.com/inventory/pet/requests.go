package pet

import (
	"atlas-inventory/rest"
	"fmt"

	"github.com/Chronicle20/atlas-rest/requests"
)

const (
	Resource        = "pets"
	ById            = Resource + "/%d"
	ByOwnerResource = "characters/%d/pets"
)

func getBaseRequest() string {
	return requests.RootUrl("PETS")
}

func requestById(petId uint32) requests.Request[RestModel] {
	return rest.MakeGetRequest[RestModel](fmt.Sprintf(getBaseRequest()+ById, petId))
}
