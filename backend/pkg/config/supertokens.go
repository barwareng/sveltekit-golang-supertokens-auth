package config

import (
	"os"
	"strings"

	"github.com/barwareng/sveltekit-golang-supertokens-auth/app/models"
	"github.com/barwareng/sveltekit-golang-supertokens-auth/pkg/database"
	"github.com/gofiber/fiber/v2/log"
	"github.com/rs/xid"
	"github.com/supertokens/supertokens-golang/ingredients/emaildelivery"
	"github.com/supertokens/supertokens-golang/recipe/emailverification"
	"github.com/supertokens/supertokens-golang/recipe/emailverification/evmodels"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/session/sessmodels"
	"github.com/supertokens/supertokens-golang/recipe/thirdparty/tpmodels"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartyemailpassword"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartyemailpassword/tpepmodels"
	"github.com/supertokens/supertokens-golang/recipe/userroles"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func SupertokensInit() {
	cookieSameSite := "lax"
	cookieDomain := os.Getenv("SUPERTOKENS_COOKIE_DOMAIN")
	cookieSecure := true
	apiBasePath := "/auth"
	websiteBasePath := "/auth"
	err := supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			ConnectionURI: os.Getenv("SUPERTOKENS_CONNECTION_URI"),
		},
		AppInfo: supertokens.AppInfo{
			AppName:         os.Getenv("APP_NAME"),
			APIDomain:       os.Getenv("SUPERTOKENS_API_DOMAIN"),
			WebsiteDomain:   os.Getenv("SUPERTOKENS_FRONTEND_DOMAIN"),
			APIBasePath:     &apiBasePath,
			WebsiteBasePath: &websiteBasePath,
		},
		RecipeList: []supertokens.Recipe{
			emailverification.Init(evmodels.TypeInput{
				Mode: evmodels.ModeOptional,
				Override: &evmodels.OverrideStruct{
					APIs: func(originalImplementation evmodels.APIInterface) evmodels.APIInterface {
						ogVerifyEmailPOST := *originalImplementation.VerifyEmailPOST
						(*originalImplementation.VerifyEmailPOST) = func(token string, sessionContainer sessmodels.SessionContainer, tenantId string, options evmodels.APIOptions, userContext supertokens.UserContext) (evmodels.VerifyEmailPOSTResponse, error) {
							resp, err := ogVerifyEmailPOST(token, sessionContainer, tenantId, options, userContext)
							if err != nil {
								return evmodels.VerifyEmailPOSTResponse{}, err
							}

							if resp.OK != nil {
								id := resp.OK.User.ID
								email := resp.OK.User.Email
								user := models.User{ID: id, Email: email}
								database.DB.Create(&user)
							}

							return resp, nil
						}
						return originalImplementation
					},
				},
				EmailDelivery: &emaildelivery.TypeInput{
					Override: func(originalImplementation emaildelivery.EmailDeliveryInterface) emaildelivery.EmailDeliveryInterface {
						ogSendEmail := *originalImplementation.SendEmail

						(*originalImplementation.SendEmail) = func(input emaildelivery.EmailType, userContext supertokens.UserContext) error {
							// This is: `${websiteDomain}${websiteBasePath}/verify-email`
							log.Info(input.EmailVerification.EmailVerifyLink)
							// log.Info()
							input.EmailVerification.EmailVerifyLink = strings.Replace(
								input.EmailVerification.EmailVerifyLink,
								"auth/",
								"", 1,
							)
							return ogSendEmail(input, userContext)
						}
						return originalImplementation
					},
				},
			}),
			thirdpartyemailpassword.Init(&tpepmodels.TypeInput{
				Override: &tpepmodels.OverrideStruct{
					Functions: func(originalImplementation tpepmodels.RecipeInterface) tpepmodels.RecipeInterface {
						// create a copy of the originalImplementation
						originalEmailPasswordSignUp := *originalImplementation.EmailPasswordSignUp
						originalThirdPartySignInUp := *originalImplementation.ThirdPartySignInUp

						// override the email password sign up function
						(*originalImplementation.EmailPasswordSignUp) = func(email, password string, tenantId string, userContext supertokens.UserContext) (tpepmodels.SignUpResponse, error) {

							resp, err := originalEmailPasswordSignUp(email, password, tenantId, userContext)
							if err != nil {
								return tpepmodels.SignUpResponse{}, err
							}

							if resp.OK != nil {
								externalUserId := xid.New().String()
								_, err := supertokens.CreateUserIdMapping(resp.OK.User.ID, externalUserId, nil, nil)
								if err != nil {
									return tpepmodels.SignUpResponse{}, err
								}
								resp.OK.User.ID = externalUserId
							}

							return resp, err
						}

						// override the thirdparty sign in / up function
						(*originalImplementation.ThirdPartySignInUp) = func(thirdPartyID, thirdPartyUserID, email string, oAuthTokens tpmodels.TypeOAuthTokens, rawUserInfoFromProvider tpmodels.TypeRawUserInfoFromProvider, tenantId string, userContext supertokens.UserContext) (tpepmodels.SignInUpResponse, error) {

							resp, err := originalThirdPartySignInUp(thirdPartyID, thirdPartyUserID, email, oAuthTokens, rawUserInfoFromProvider, tenantId, userContext)
							if err != nil {
								return tpepmodels.SignInUpResponse{}, err
							}

							if resp.OK != nil {
								if resp.OK.CreatedNewUser {
									externalUserId := xid.New().String()
									_, err := supertokens.CreateUserIdMapping(resp.OK.User.ID, externalUserId, nil, nil)
									if err != nil {
										return tpepmodels.SignInUpResponse{}, err
									}
									resp.OK.User.ID = externalUserId
								}
							}

							return resp, err
						}

						return originalImplementation
					},
				},
				Providers: []tpmodels.ProviderInput{
					{
						Config: tpmodels.ProviderConfig{
							ThirdPartyId: "google",
							Clients: []tpmodels.ProviderClientConfig{
								{
									ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
									ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
								},
							},
						},
					},
				},
			}),
			userroles.Init(nil),
			session.Init(&sessmodels.TypeInput{
				CookieSameSite: &cookieSameSite,
				CookieSecure:   &cookieSecure,
				CookieDomain:   &cookieDomain,
			}), // initializes session features
		},
	})

	if err != nil {
		panic(err.Error()) //TODO resolve this. Use Fiber recover
	}
}
