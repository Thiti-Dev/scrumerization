package entities

import (
	jetModel "github.com/Thiti-Dev/scrumerization-core-service/.gen/scrumerization/public/model"
)

type PopulatedRoom struct {
	jetModel.Rooms

	Creator jetModel.Users // No matter the struct property name is, under the hood these will be iterated and evaluate by its reflected type
}

type PaginatedRoomResult struct {
	Data []PopulatedRoom
	PaginationResult
}
