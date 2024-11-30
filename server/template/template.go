package template

import "fmt"

func GenerateEmailRegister(baseClientURL string, token string) string {
	return fmt.Sprintf(`
	<div>
<p>Hi there,</p>
</br>
<p>Thank you for signing up for Simple Dashboard App. Click on the link below to verify your email :</p>
</br>
<a target="_blank" href="%s/verify/%s">Verify Email</a>
</br>
<p>This link will expire in 24 hours. If you did not sign up for a Simple Dashboard App,
you can safely ignore this email.</p>
</br>
</br>
<p>Best,</p>
</br>
<p>Afista Pratama</p>
<div>
	`, baseClientURL, token)
}

func GenerateEmailForgotPassword(baseClientURL string, token string) string {
	return fmt.Sprintf(`
	<div>
<p>Hi there,</p>

<p>Click on the link below to reset your password :</p>
</br>
<a target="_blank" href="%s/reset-password/%s">Reset Password</a>
</br>
<p>This link will expire in 24 hours. If you did not request a password reset,
you can safely ignore this email.</p>
</br>
</br>
<p>Best,</p>
</br>
<p>Afista Pratama</p>
<div>
	`, baseClientURL, token)
}
