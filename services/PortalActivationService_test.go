package services

import (
	"rose-be-go/models/dto"
	"testing"
)

func TestSenmail_PackMsgTemp(t *testing.T) {

	templateEmail := `
		<p>Hi, Kami telah menerima permintaanmu untuk melakukan reset password akun Business Portal.</p>
		<p>Silahkan login ulang menggunakan username dan password dibawah: </p>
		<p>Username: ` + "Test" + `</p>
		<p>Password: ` + "test" + `</p>
		<p>Klik link ` + "test" + `untuk menuju halaman login.</p>
		<p>Abaikan email ini jika kamu tidak pernah meminta untuk melakukan reset password.</p>`

	req := dto.ReqPortalActivation{Email:"fikri.alimudin@ottodigital.id"}
	err := sendMail("fikri.alimudin@ottodigital.id", "Test123", templateEmail, req)

	if err != nil {
		t.Errorf("Error %v", err)
	}
}