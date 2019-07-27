package codec

import (
	"fmt"
	"testing"
)

func TestXMLPrivateKey(t *testing.T) {
	priKeyXML := "<RSAKeyValue><Modulus>q54XcahjRw7B5cFTw2S9nxlf2ApNJs/vwq8Ka3Xz1egVhNA1NTT6NtmtuETMYbdalTPE3b4qYHuBBIa6Vmp580RLKmQLFtm0haF9XJwbEvXnhpdU1AaZHJQYMOv1pKwvX6fP7jWhP8oFn4S480v5pU1QqVpZ6+Pjk9hSF5P1TNs=</Modulus><Exponent>AQAB</Exponent><P>wB9eFPA7KxnIhnb1/rAiBBi7ABUjFuTZV0tQA+Qz3B/OSH2FvjOuMFR3UtjjIlYvXs8UsOOoy8xK+zBNUaFm0Q==</P><Q>5K1t3p/thPECHc8B1D0iS6xSdn2oRQTlctUsa7x6hz0G8yunt6jHKGicAyCaNRs7X4wCGaCab3myw1NlPm376w==</Q><DP>lni61gllPheej3oipsxKiAzagXEMn4SzmQQ6ciHMYE5k4S+2jUaq2d1961xFYQJBba8g1H4qGRwzadLl69ZO0Q==</DP><DQ>afoFxPqQ9N9Mf8wDqPx7/F3sCYXKidHvE/y/DV7X2fzqT5+XpoHNamiwhLKDRMFODsWv43iOmcLomaT4cqo96Q==</DQ><InverseQ>CKGetXVwzs3rp9Iwrko4tqKtoypiChPByA9i3CtSQLBMrdnnmJciQIfqyoQzhXpbuWFrf706Fmn12oIHjpsMKg==</InverseQ><D>TfOZIn8h2PzTBdEyguYXW1TjO/Yx5Rc+CGgJi5YgE2E/pPwNrwVJQfDN/40AaMIn2u2Q1keyZ/CxQaUQrn/es+mGNM9ShytVdSHfL2BR9jNwJSTD/bvJvAgpA0oWLWgWedl2IMb6VBnwofv561IPLsoQsWmz2s/8BFnkbe2/WQE=</D></RSAKeyValue>"
	priKey := XMLToPrivateKey(priKeyXML)
	fmt.Println(priKey)
}

func TestXMLPubKeyExport(t *testing.T) {
	priKeyXML := "<RSAKeyValue><Modulus>q54XcahjRw7B5cFTw2S9nxlf2ApNJs/vwq8Ka3Xz1egVhNA1NTT6NtmtuETMYbdalTPE3b4qYHuBBIa6Vmp580RLKmQLFtm0haF9XJwbEvXnhpdU1AaZHJQYMOv1pKwvX6fP7jWhP8oFn4S480v5pU1QqVpZ6+Pjk9hSF5P1TNs=</Modulus><Exponent>AQAB</Exponent><P>wB9eFPA7KxnIhnb1/rAiBBi7ABUjFuTZV0tQA+Qz3B/OSH2FvjOuMFR3UtjjIlYvXs8UsOOoy8xK+zBNUaFm0Q==</P><Q>5K1t3p/thPECHc8B1D0iS6xSdn2oRQTlctUsa7x6hz0G8yunt6jHKGicAyCaNRs7X4wCGaCab3myw1NlPm376w==</Q><DP>lni61gllPheej3oipsxKiAzagXEMn4SzmQQ6ciHMYE5k4S+2jUaq2d1961xFYQJBba8g1H4qGRwzadLl69ZO0Q==</DP><DQ>afoFxPqQ9N9Mf8wDqPx7/F3sCYXKidHvE/y/DV7X2fzqT5+XpoHNamiwhLKDRMFODsWv43iOmcLomaT4cqo96Q==</DQ><InverseQ>CKGetXVwzs3rp9Iwrko4tqKtoypiChPByA9i3CtSQLBMrdnnmJciQIfqyoQzhXpbuWFrf706Fmn12oIHjpsMKg==</InverseQ><D>TfOZIn8h2PzTBdEyguYXW1TjO/Yx5Rc+CGgJi5YgE2E/pPwNrwVJQfDN/40AaMIn2u2Q1keyZ/CxQaUQrn/es+mGNM9ShytVdSHfL2BR9jNwJSTD/bvJvAgpA0oWLWgWedl2IMb6VBnwofv561IPLsoQsWmz2s/8BFnkbe2/WQE=</D></RSAKeyValue>"
	pubKeyXML := XMLPublicKey(priKeyXML)
	fmt.Println(pubKeyXML)
}
