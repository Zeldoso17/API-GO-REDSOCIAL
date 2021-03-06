package bd

import (
	"context"
	"time"
	"github.com/Zeldoso17/API-GO-REDSOCIAL/models"
)

func DeleteRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t) // Here I'm deleting the relation
	if err != nil {
		return false, err
	}
	return true, nil
}