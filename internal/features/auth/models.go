package auth

/*
package auth dùng để làm mẫu, còn theo hay không thì tùy @@
[Client]
↓ SignUpRequest --- ↑ signUpResponse
[Handler]
↓ SignUpInput --- ↑ signUpResult
[Usecase]
↓ createUserParams --- ↑ createUserResult
[Repository]
*/

type signUpRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signUpInput struct {
	Username string
	Email    string
	Password string
}

type createUserParams struct {
	Username string
	Email    string
	Password string
}
type createUserResult struct {
	Username string
	Email    string
}

type signUpResult struct {
	Username string
	Email    string
}

type signUpResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
