package matchusecase

import "errors"

var ErrCatNotBelongToUser = errors.New("cat is not belong to the user")
var ErrSameCatGender = errors.New("cat genders cannot be the same")
var ErrAlreadyMatch = errors.New("cat's is already match")
var ErrSameOwner = errors.New("cat is from the same owner")
var ErrNotFound = errors.New("match data not found")
var ErrForbidden = errors.New("forbidden")
