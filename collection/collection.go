package collection

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"guest-ledger-book-api/db"
	"guest-ledger-book-api/models"
	"log"
	"net/http"
)

var GuestCollection = db.Db().Database("GuestLedgerBook").Collection("Guests")

func GetAllGuests(c echo.Context) error {
	filter := bson.D{{}}

	cursor, err := GuestCollection.Find(db.Ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	var guests []models.Guest

	if err = cursor.All(db.Ctx, &guests); err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, guests)
}

func CreateNewGuest(c echo.Context) error {
	var newGuest = models.Guest{
		ID: primitive.NewObjectID(),
	}

	if err := c.Bind(&newGuest); err != nil {
		return err
	}

	result, err := GuestCollection.InsertOne(db.Ctx, newGuest)

	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, result)
}
