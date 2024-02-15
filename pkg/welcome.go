package pkg

import (
	"github.com/matcornic/hermes/v2"
)

type Welcome struct {
}

func (w *Welcome) Name() string {
	return "welcome"
}

func (w *Welcome) Email() hermes.Email {
	return hermes.Email{
		Body: hermes.Body{
			Name: "John Doe",
			Intros: []string{
				"Welcome to Tribal! on board!",
			},
			Dictionary: []hermes.Entry{
				{Key: "Firstname", Value: "John"},
				{Key: "Lastname", Value: "Doe"},
				{Key: "Birthday", Value: "01/01/1990"},
			},
			Actions: []hermes.Action{
				{
					Instructions: "To get started with Tribal, please click here:",
					Button: hermes.Button{
						Text: "Confirm your account",
						Link: "https://hermes-example.com/confirm?token=d9729feb74992cc3482b350163a1a010",
					},
				},
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}
}
