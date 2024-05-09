package constant

import "errors"

//General Error Code
var ErrBadRequest error = errors.New("Bad Request")

//User Errors Code
var UserNotFound error = errors.New("User not found")
var ErrLoginEmptyInput error = errors.New("Email and password cannot be empty")
var ErrLoginNotFound error = errors.New("User not found")
var ErrLoginIncorrectPassword error = errors.New("Incorrect Password")
var ErrLoginJWT error = errors.New("Failed to generate JWT token, please try again")
var ErrRegisterUserExists error = errors.New("User already exists")
var ErrHashPassword error = errors.New("Failed to hash password, please use another password")
var ErrUpdateUserEmailExists error = errors.New("Email already exists, please use another email")
var ErrUpdateUser error = errors.New("Failed to update user")

//Collector Errors Code
var ErrorCollectorRegister error = errors.New("Failed to register collector")
var ErrCollectorUserEmailExists error = errors.New("Email already exists, please use another email")
var ErrCollectorUserNotFound error = errors.New("Collector not found")
var ErrCollectorIncorrectPassword error = errors.New("Incorrect Password")
var ErrCollectorLoginJWT error = errors.New("Failed to generate JWT token, please try again")
var ErrUpdateCollectorEmailExists error = errors.New("Email already exists, please use another email")
var ErrorUpdateCollector error = errors.New("Failed to update collector")

//Pickup Request Errors Code
var ErrPickupRequestNotFound error = errors.New("Pickup request not found")
var ErrPickupRequestCreate error = errors.New("Failed to create pickup request")
var ErrPickupRequestGet error = errors.New("Failed to get pickup request")
var ErrPickupRequestDelete error = errors.New("Failed to delete pickup request")

//Pickup Transaction Errors Code
var ErrPickupTransactionCreate error = errors.New("Failed to create pickup transaction")
var ErrPickupTransactionCreateChat error = errors.New("Failed to create chat")
var ErrPickupTransactionUpdateStatus error = errors.New("Failed to update status")
var ErrPickupTransactionGet error = errors.New("Failed to get pickup transaction")
var ErrPickupTransactionGetAll error = errors.New("Failed to get all pickup transaction")

//Payment Transaction Success
var ErrPaymentTransactionCreate error = errors.New("Failed to create payment transaction")
var ErrAlreadyVerified error = errors.New("Payment already verified")
var ErrPaymentTransactionUpdate error = errors.New("Failed to update payment transaction")
var ErrGetPaymentTransaction error = errors.New("Failed to get payment transaction")

