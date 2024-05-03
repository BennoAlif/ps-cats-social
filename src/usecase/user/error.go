package userusecase

import "errors"

var ErrUserNotFound = errors.New("user tidak ditemukan")
var ErrEmailAlreadyUsed = errors.New("email sudah digunakan")
var ErrInvalidPassword = errors.New("password salah")
var ErrInvalidToken = errors.New("token tidak valid")
var ErrExpiredToken = errors.New("token kadaluwarsa")
var ErrTokenNotFound = errors.New("token tidak ditemukan")
var ErrInvalidUser = errors.New("email atau password salah")
