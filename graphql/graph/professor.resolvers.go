package graph

import (
	"context"
	"graphql/graph/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *queryResolver) Professor(ctx context.Context, id string) (*model.Professor, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	dbProfessor := &model.DBProfessor{}

	err = r.ProfessorCollection.FindOne(
		timeoutCtx,	bson.M{"_id": objectId}).Decode(dbProfessor)
	if err != nil {
		return nil, err
	}

	return model.Tra

}
