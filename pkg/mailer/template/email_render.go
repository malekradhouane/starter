package template

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
)

//go:embed activateAccount.html
var activateAccountTpl []byte

//go:embed resetPassword.html
var resetPasswordTpl []byte

// RenderInviteUser returns the activateAccount HTML with the link injected
func RenderActivateAccount(link string) (string, error) {
	if len(activateAccountTpl) == 0 {
		return "", fmt.Errorf("" +
			" template not embedded")
	}
	tpl, err := template.New("activateAccount").Parse(string(activateAccountTpl))
	if err != nil {
		return "", fmt.Errorf("parse invite template: %w", err)
	}

	var buf bytes.Buffer
	data := struct {
		Link string
	}{Link: link}

	if err := tpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("execute activateAccount template: %w", err)
	}
	return buf.String(), nil
}

// RenderResetPassword returns the reset password HTML with the link injected
func RenderResetPassword(link string) (string, error) {
	if len(resetPasswordTpl) == 0 {
		return "", fmt.Errorf("reset password template not embedded")
	}
	tpl, err := template.New("resetPassword").Parse(string(resetPasswordTpl))
	if err != nil {
		return "", fmt.Errorf("parse reset password template: %w", err)
	}

	var buf bytes.Buffer
	data := struct {
		Link string
	}{Link: link}

	if err := tpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("execute reset password template: %w", err)
	}
	return buf.String(), nil
}
