package usermongo

import (
	"delivery-go/entities_layer/user/userentity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserSystem struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty"`
	IDCl        primitive.ObjectID   `bson:"_id_client"`
	IDsEst      []primitive.ObjectID `bson:"_ids_estb"`
	Login       string
	Password    string
	Email       string
	PhoneNumber string
	Name        string
	Surname     string
	Patronymic  string
	Roles       []string
}

func ToEntityUserSystem(us *UserSystem) *userentity.UserSystem {
	IdsEst := []string{}
	for _, id := range us.IDsEst {
		IdsEst = append(IdsEst, id.Hex())
	}
	return &userentity.UserSystem{
		ID:          us.ID.Hex(),
		IDCl:        us.IDCl.Hex(),
		IDsEst:      IdsEst,
		Login:       us.Login,
		Password:    us.Password,
		Email:       us.Email,
		PhoneNumber: us.PhoneNumber,
		Name:        us.Name,
		Surname:     us.Surname,
		Patronymic:  us.Patronymic,
		Roles:       us.Roles,
	}
}

func toUserSystem(usm *userentity.UserSystem) *UserSystem {
	id, _ := primitive.ObjectIDFromHex(usm.ID)
	idCl, _ := primitive.ObjectIDFromHex(usm.IDCl)

	IdsEst := []primitive.ObjectID{}
	for _, idStr := range usm.IDsEst {
		idObj, _ := primitive.ObjectIDFromHex(idStr)
		IdsEst = append(IdsEst, idObj)
	}

	return &UserSystem{
		ID:          id,
		IDCl:        idCl,
		IDsEst:      IdsEst,
		Login:       usm.Login,
		Password:    usm.Password,
		Email:       usm.Email,
		PhoneNumber: usm.PhoneNumber,
		Name:        usm.Name,
		Surname:     usm.Surname,
		Patronymic:  usm.Patronymic,
		Roles:       usm.Roles,
	}
}
