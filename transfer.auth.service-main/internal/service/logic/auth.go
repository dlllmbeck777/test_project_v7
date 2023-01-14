package logic

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/transferMVP/transfer.webapp/internal/models"
	"github.com/transferMVP/transfer.webapp/internal/pkg/errorsApp"
	"github.com/transferMVP/transfer.webapp/internal/pkg/response"
	"github.com/transferMVP/transfer.webapp/internal/repository/pg"
	"github.com/valyala/fasthttp"
)

//	@version	2.0
//	@host		localhost:8889
//	@basePath	/api/v1

// RegistrUser godoc
//	@Tags		Users
//	@Summary	Регистрация пользователя
//	@Accept		json
//	@Produce	json
//	@Param		Body		body	models.User	true	"Тело"
//	@Success	200				{object}	models.User
//	@Router		/registr [post]
func RegistrUser(resp *response.Response, ctx *fasthttp.RequestCtx) {

	var user models.User
	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
		resp.SetError(errorsApp.GetErr(2, err.Error()))
		return
	}
	hp := sha256.Sum256([]byte(user.Password))
	user.Password = hex.EncodeToString(hp[:])

	id, err := pg.AddUser(&user)
	if err != nil {
		resp.SetError(errorsApp.GetErr(3, err.Error()))
		return
	}
	//redis2.AddToken(user.AccessToken)
	user.Id = id
	user.Password = ""
	resp.SetValue(user)
}

//	@version	2.0
//	@host		localhost:8889
//	@basePath	/api/v1

// GetUser godoc
//	@Tags		Users
//	@Summary	Получение данных о пользователе
//	@Accept		json
//	@Produce	json
//	@Param		id		body	string	true	"Тело"
//	@Success	200				{object}	models.User
//	@Router		/getuser [post]
func GetUser(resp *response.Response, ctx *fasthttp.RequestCtx) {
	var mpData map[string]string
	if err := getData(ctx, &mpData); err != nil {
		resp.SetError(errorsApp.GetErr(2, err.Error()))
		return
	}

	// get token
	id, ok := mpData["id"]
	if !ok || len(id) == 0 {
		resp.SetError(errorsApp.GetErr(4, errors.New("undefined id").Error()))
		return
	}

	user, err := pg.GetUser(id)
	if err != nil {
		resp.SetError(errorsApp.GetErr(5, err.Error()))
		return
	}
	resp.SetValue(user)
}

func getData(ctx *fasthttp.RequestCtx, in interface{}) error {
	return json.Unmarshal(ctx.PostBody(), in)
}

//func OAuthSignIn(resp *response.Response, ctx *fasthttp.RequestCtx) {
//	fmt.Println("OAuthSignIn")
//	rootUrl := `https://accounts.google.com/o/oauth2/v2/auth`
//
//	URL, err := url.Parse(rootUrl)
//	if err != nil {
//		fmt.Println(1)
//		fmt.Println(err)
//
//	}
//	fmt.Println(URL.String())
//	parameters := url.Values{}
//	parameters.Add("client_id", "924872859032-k6vqkop29oeu6am6tf88j3borrsdk99c.apps.googleusercontent.com")
//	parameters.Add("scope", strings.Join([]string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"}, " "))
//	parameters.Add("redirect_uri", "http://localhost:9091/api/sessions/oauth/google")
//	parameters.Add("response_type", "code")
//	parameters.Add("access_type", "offline")
//	parameters.Add("prompt", "consent")
//	parameters.Add("state", "")
//	URL.RawQuery = parameters.Encode()
//	url := URL.String()
//
//	resp1, err1 := http.Get(url)
//	if err1 != nil {
//		fmt.Println(2)
//		fmt.Println(err)
//	}
//	dat, err2 := ioutil.ReadAll(resp1.Body)
//	if err2 != nil {
//		fmt.Println(3)
//		fmt.Println(err)
//	}
//	resp.SetValue(dat)
//}

//func b2s(b []byte) string {
//	/* #nosec G103 */
//	return *(*string)(unsafe.Pointer(&b))
//}

//func GoogleRedirect(resp *response.Response, ctx *fasthttp.RequestCtx) {
//	fmt.Println("GoogleRedirect")
//
//	code := ctx.QueryArgs().Peek("code")
//	codeStr := string(code)
//	var pathUrl string = "/"
//
//	state := ctx.QueryArgs().Peek("state")
//	stateStr := string(state)
//	if stateStr != "" {
//		pathUrl = stateStr
//	}
//
//	if codeStr == "" {
//		fmt.Println("\"message\": \"Authorization code not provided!")
//		return
//	}
//
//	// Use the code to get the id and access tokens
//	tokenRes, err := GetGoogleOauthToken(codeStr)
//
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//
//	user, err := GetGoogleUser(tokenRes.Access_token, tokenRes.Id_token)
//	fmt.Println(pathUrl)
//	resp.SetValue(user)
//}
//
//type GoogleOauthToken struct {
//	Access_token string
//	Id_token     string
//}
//
//func GetGoogleOauthToken(code string) (*GoogleOauthToken, error) {
//	const rootURl = "https://oauth2.googleapis.com/token"
//
//	values := url.Values{}
//	values.Add("grant_type", "authorization_code")
//	values.Add("code", code)
//	values.Add("client_id", "124989408133-kf59mge02rsof2ltpmaobu0v9klebp36.apps.googleusercontent.com")
//	values.Add("client_secret", "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")
//	values.Add("redirect_uri", "http://127.0.0.1:9091/api/sessions/oauth/google")
//
//	query := values.Encode()
//
//	req, err := http.NewRequest("POST", rootURl, bytes.NewBufferString(query))
//	if err != nil {
//		return nil, err
//	}
//
//	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
//	client := http.Client{
//		Timeout: time.Second * 30,
//	}
//
//	res, err := client.Do(req)
//	if err != nil {
//		return nil, err
//	}
//
//	if res.StatusCode != http.StatusOK {
//		return nil, errors.New("could not retrieve token")
//	}
//
//	resBody, err := ioutil.ReadAll(res.Body)
//	if err != nil {
//		return nil, err
//	}
//
//	var GoogleOauthTokenRes map[string]interface{}
//
//	if err := json.Unmarshal(resBody, &GoogleOauthTokenRes); err != nil {
//		return nil, err
//	}
//
//	tokenBody := &GoogleOauthToken{
//		Access_token: GoogleOauthTokenRes["access_token"].(string),
//		Id_token:     GoogleOauthTokenRes["id_token"].(string),
//	}
//
//	return tokenBody, nil
//}
//
//type GoogleUserResult struct {
//	Id             string
//	Email          string
//	Verified_email bool
//	Name           string
//	Given_name     string
//	Family_name    string
//	Picture        string
//	Locale         string
//}
