package main

import (
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws/sts"
)

func getSessionToken(tokenCode string) (*sts.Credentials) {
    session = sts.New(getSession())
    credentials, err := session.GetSessionToken()
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
    return credentials
}

func getSession() (sess *session.Session) {
    session, err := session.NewSession()
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
    return session
}
