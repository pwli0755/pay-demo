package service

import (
	"context"
	"fmt"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/alipay/cert"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math/rand"

	pb "pay/api/pay/v1"
	"pay/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

func NewPayService(product *biz.ProductUsecase, logger log.Logger) *PayService {
	return &PayService{
		product: product,
		log:     log.NewHelper(logger),
	}
}

func (s *PayService) ListProduct(ctx context.Context, req *pb.ListProductRequest) (*pb.ListProductReply, error) {
	ps, err := s.product.List(ctx)
	reply := &pb.ListProductReply{}
	for _, p := range ps {
		reply.Results = append(reply.Results, &pb.Product{
			Id:        p.ID,
			Title:     p.Title,
			Price:     p.Price,
			CreatedAt: timestamppb.New(p.CreatedAt),
			UpdatedAt: timestamppb.New(p.UpdatedAt),
		})
	}
	return reply, err
}

const (
	kServerPort   = "8000"
	kServerDomain = "http://127.0.0.1" + ":" + kServerPort
)

var aliClient *alipay.Client

func init() {
	cert.Appid = "2021000121693429"
	cert.PrivateKey = "MIIEpgIBAAKCAQEAs4jckwbJcRRNFB2GB8XCkEsCCUSRA/Qp+pYOm8b2tp+SlKLGtJP6oCl8UnPO07FLc7Iw8zFOSO8TVypzbdCRB02UloHi8/Ii9X251jKXoiw8KvYh5AQmNr1M7dQhoZeRpBVjbZAkRkxTq6NoElilgAV20BdSAgjwSIBkmYxMpVeGLl0X78Tu3tTPHqdEICoAFeY04gAizqjDOw0e8i6P5sO+cXGtcw4A3CIxiDrHsJ3/grjGcMnJ+WAr62QmWQoEBKbZvxC9jJE38viKOiC54LwcStVXjHrSwo8h+alcrfyKDIVVWxxXSOoWdAF3ryDtRiyNVOI/RETHrtzUZRRuIwIDAQABAoIBAQCwOmWZhI3zi6PlXN2Vf49uJ9KF2mImaWNTcDxCuNivho9Riz5VPvRChrZcEQUyUtPna0AVV46qlNJ9O1Q2tQXHD2YNHs3x+vpJ0vG5ycuCCr28xgGaWmBQVxzOTu38OlVhpQUGJPkWcBGpZyre51j9A2AO/vUmvjNuV2loN4l8uXa3Zo9Icqu1Q5DbIlWwKivpGjvg1IIQSB9tF7t0cbpsspxFyYRhWsfS+CuLrfvfVpIyVPwUdxwzEh3pvQDV0C9Y/c8LOOLghfuoT1VAIlZpaeXb43PfUJLDrLhjwTAK1PUyhay/IoG4wP1PouzMzyr7I9n/TPBr5zAIlf8aonpBAoGBAPQNWJgQJyDOa0u2ZA7STTYXnPhA5l6SmSVpLf4Cda1+MQ/ehzfR3XHwsUgNYkIzzPtvv44AObfVaFh+RiwjXdB6GN92VSdl+sSrUnioTtRF1sAtBCRKYCEce21KG38y+SAG8iQU3oXCyPTlRKNg3yWnv9k5qZzD/X5+792O0DfRAoGBALxS7lzn4MVhkpTBSWesp3zRcdj2g8Kr6FUL0vspmFY1rePHKMkTbZEmMbMtLRhbYxattwQkv7n6dfi04yqUuv4Pbs1iLYbQ8BluFB5zB0Bo655T8uZAkQDrRCrsFXBc8sehEt12mmTImY7Q0eclRsFEMMixC4IxO0q/gVQU57ezAoGBAOh1reVpvVtqQpkjabsUZacYZtOwPb3nNSiFPuGrxhszD9hlxbZNl/hnovVWijk0zhLRJkxDurZ396QS4xQ3u6xQIFD5jbKxWGLsLOnwpuVagGscdc17aoUfdBFtfTNzgggXlZz4o9wU5QUfPHnCU8qyNfLbEcvYgyRyFFedKIeBAoGBAJSTtT0VL+9pqJS13ezueYFvWKu86X0X1YfreTvwuCAj35oaUaI6MrJWeNWM6cwSpZ2J1h9twtm+sX5Tb/nzN4gjst5U++gmRZc6kqLnS6xUWrgiMTvZas1X0AMxGUT6AAzhlpmk7fBfl07mjwQXE7h8zSQ5EgRYRRgW+LjWeW4jAoGBAMZ5VU+CjKuIjvx/YLchd0/5+ID2vwmwGl247QOCeQnshneA71v5rP4TDp5IbI/iqP9q+Odi9XX5LV9N6D6Yuirub7GtIaL4dqYN/+1k6RERBwnXB1wi2uZJgzvzh8pWSwsyl2d67uacQvNCHXmn9YyC8EaKcGH+xsb69ZlqKlG7"
	cert.AppPublicContent = []byte(`-----BEGIN CERTIFICATE-----
MIIDmTCCAoGgAwIBAgIQICISBupnbWo6FSml0XR6WTANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UE
BhMCQ04xGzAZBgNVBAoMEkFudCBGaW5hbmNpYWwgdGVzdDElMCMGA1UECwwcQ2VydGlmaWNhdGlv
biBBdXRob3JpdHkgdGVzdDE+MDwGA1UEAww1QW50IEZpbmFuY2lhbCBDZXJ0aWZpY2F0aW9uIEF1
dGhvcml0eSBDbGFzcyAyIFIxIHRlc3QwHhcNMjIxMjA2MDMwOTE5WhcNMjMxMjExMDMwOTE5WjBr
MQswCQYDVQQGEwJDTjEfMB0GA1UECgwWa2NkdWppMTM3N0BzYW5kYm94LmNvbTEPMA0GA1UECwwG
QWxpcGF5MSowKAYDVQQDDCEyMDg4NjIxOTkzODg0OTY1LTIwMjEwMDAxMjE2OTM0MjkwggEiMA0G
CSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCziNyTBslxFE0UHYYHxcKQSwIJRJED9Cn6lg6bxva2
n5KUosa0k/qgKXxSc87TsUtzsjDzMU5I7xNXKnNt0JEHTZSWgeLz8iL1fbnWMpeiLDwq9iHkBCY2
vUzt1CGhl5GkFWNtkCRGTFOro2gSWKWABXbQF1ICCPBIgGSZjEylV4YuXRfvxO7e1M8ep0QgKgAV
5jTiACLOqMM7DR7yLo/mw75xca1zDgDcIjGIOsewnf+CuMZwycn5YCvrZCZZCgQEptm/EL2MkTfy
+Io6ILngvBxK1VeMetLCjyH5qVyt/IoMhVVbHFdI6hZ0AXevIO1GLI1U4j9ERMeu3NRlFG4jAgMB
AAGjEjAQMA4GA1UdDwEB/wQEAwIE8DANBgkqhkiG9w0BAQsFAAOCAQEAolq0m07A2/JQaUGzHp5y
oY6yq7lnmpIpC7csHUW/Fi+DdWv2e0F2/hy0Wsc5dTUYtG6bLsoRKb23oVVPl8DCw+24GT6U3eGv
OlVQyjJDdKKyhvciVpU/V2TqEOP5yrsyMr0n09trM1eogS8XWCV4c4uTVgoUgG2rCGunpNQy2uaT
5TKxQS1B8P92EBWxJUcQEN1LjEn7l1gG3t5tUM3IghXdyNxbgYzcqptGc8SuapE0xvSgW5c/SobW
u5s8w4UKQd3gfz6tq5BAPBGvTVn3IucJPqRnWAvy+Ue8/Meqab/jSe838cX6NgHfYyt862S9cRSA
pZbe1kXb1xUm8arfWQ==
-----END CERTIFICATE-----`)
	cert.AlipayRootContent = []byte(`-----BEGIN CERTIFICATE-----
MIIBszCCAVegAwIBAgIIaeL+wBcKxnswDAYIKoEcz1UBg3UFADAuMQswCQYDVQQG
EwJDTjEOMAwGA1UECgwFTlJDQUMxDzANBgNVBAMMBlJPT1RDQTAeFw0xMjA3MTQw
MzExNTlaFw00MjA3MDcwMzExNTlaMC4xCzAJBgNVBAYTAkNOMQ4wDAYDVQQKDAVO
UkNBQzEPMA0GA1UEAwwGUk9PVENBMFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAE
MPCca6pmgcchsTf2UnBeL9rtp4nw+itk1Kzrmbnqo05lUwkwlWK+4OIrtFdAqnRT
V7Q9v1htkv42TsIutzd126NdMFswHwYDVR0jBBgwFoAUTDKxl9kzG8SmBcHG5Yti
W/CXdlgwDAYDVR0TBAUwAwEB/zALBgNVHQ8EBAMCAQYwHQYDVR0OBBYEFEwysZfZ
MxvEpgXBxuWLYlvwl3ZYMAwGCCqBHM9VAYN1BQADSAAwRQIgG1bSLeOXp3oB8H7b
53W+CKOPl2PknmWEq/lMhtn25HkCIQDaHDgWxWFtnCrBjH16/W3Ezn7/U/Vjo5xI
pDoiVhsLwg==
-----END CERTIFICATE-----

-----BEGIN CERTIFICATE-----
MIIF0zCCA7ugAwIBAgIIH8+hjWpIDREwDQYJKoZIhvcNAQELBQAwejELMAkGA1UE
BhMCQ04xFjAUBgNVBAoMDUFudCBGaW5hbmNpYWwxIDAeBgNVBAsMF0NlcnRpZmlj
YXRpb24gQXV0aG9yaXR5MTEwLwYDVQQDDChBbnQgRmluYW5jaWFsIENlcnRpZmlj
YXRpb24gQXV0aG9yaXR5IFIxMB4XDTE4MDMyMTEzNDg0MFoXDTM4MDIyODEzNDg0
MFowejELMAkGA1UEBhMCQ04xFjAUBgNVBAoMDUFudCBGaW5hbmNpYWwxIDAeBgNV
BAsMF0NlcnRpZmljYXRpb24gQXV0aG9yaXR5MTEwLwYDVQQDDChBbnQgRmluYW5j
aWFsIENlcnRpZmljYXRpb24gQXV0aG9yaXR5IFIxMIICIjANBgkqhkiG9w0BAQEF
AAOCAg8AMIICCgKCAgEAtytTRcBNuur5h8xuxnlKJetT65cHGemGi8oD+beHFPTk
rUTlFt9Xn7fAVGo6QSsPb9uGLpUFGEdGmbsQ2q9cV4P89qkH04VzIPwT7AywJdt2
xAvMs+MgHFJzOYfL1QkdOOVO7NwKxH8IvlQgFabWomWk2Ei9WfUyxFjVO1LVh0Bp
dRBeWLMkdudx0tl3+21t1apnReFNQ5nfX29xeSxIhesaMHDZFViO/DXDNW2BcTs6
vSWKyJ4YIIIzStumD8K1xMsoaZBMDxg4itjWFaKRgNuPiIn4kjDY3kC66Sl/6yTl
YUz8AybbEsICZzssdZh7jcNb1VRfk79lgAprm/Ktl+mgrU1gaMGP1OE25JCbqli1
Pbw/BpPynyP9+XulE+2mxFwTYhKAwpDIDKuYsFUXuo8t261pCovI1CXFzAQM2w7H
DtA2nOXSW6q0jGDJ5+WauH+K8ZSvA6x4sFo4u0KNCx0ROTBpLif6GTngqo3sj+98
SZiMNLFMQoQkjkdN5Q5g9N6CFZPVZ6QpO0JcIc7S1le/g9z5iBKnifrKxy0TQjtG
PsDwc8ubPnRm/F82RReCoyNyx63indpgFfhN7+KxUIQ9cOwwTvemmor0A+ZQamRe
9LMuiEfEaWUDK+6O0Gl8lO571uI5onYdN1VIgOmwFbe+D8TcuzVjIZ/zvHrAGUcC
AwEAAaNdMFswCwYDVR0PBAQDAgEGMAwGA1UdEwQFMAMBAf8wHQYDVR0OBBYEFF90
tATATwda6uWx2yKjh0GynOEBMB8GA1UdIwQYMBaAFF90tATATwda6uWx2yKjh0Gy
nOEBMA0GCSqGSIb3DQEBCwUAA4ICAQCVYaOtqOLIpsrEikE5lb+UARNSFJg6tpkf
tJ2U8QF/DejemEHx5IClQu6ajxjtu0Aie4/3UnIXop8nH/Q57l+Wyt9T7N2WPiNq
JSlYKYbJpPF8LXbuKYG3BTFTdOVFIeRe2NUyYh/xs6bXGr4WKTXb3qBmzR02FSy3
IODQw5Q6zpXj8prYqFHYsOvGCEc1CwJaSaYwRhTkFedJUxiyhyB5GQwoFfExCVHW
05ZFCAVYFldCJvUzfzrWubN6wX0DD2dwultgmldOn/W/n8at52mpPNvIdbZb2F41
T0YZeoWnCJrYXjq/32oc1cmifIHqySnyMnavi75DxPCdZsCOpSAT4j4lAQRGsfgI
kkLPGQieMfNNkMCKh7qjwdXAVtdqhf0RVtFILH3OyEodlk1HYXqX5iE5wlaKzDop
PKwf2Q3BErq1xChYGGVS+dEvyXc/2nIBlt7uLWKp4XFjqekKbaGaLJdjYP5b2s7N
1dM0MXQ/f8XoXKBkJNzEiM3hfsU6DOREgMc1DIsFKxfuMwX3EkVQM1If8ghb6x5Y
jXayv+NLbidOSzk4vl5QwngO/JYFMkoc6i9LNwEaEtR9PhnrdubxmrtM+RjfBm02
77q3dSWFESFQ4QxYWew4pHE0DpWbWy/iMIKQ6UZ5RLvB8GEcgt8ON7BBJeMc+Dyi
kT9qhqn+lw==
-----END CERTIFICATE-----

-----BEGIN CERTIFICATE-----
MIICiDCCAgygAwIBAgIIQX76UsB/30owDAYIKoZIzj0EAwMFADB6MQswCQYDVQQG
EwJDTjEWMBQGA1UECgwNQW50IEZpbmFuY2lhbDEgMB4GA1UECwwXQ2VydGlmaWNh
dGlvbiBBdXRob3JpdHkxMTAvBgNVBAMMKEFudCBGaW5hbmNpYWwgQ2VydGlmaWNh
dGlvbiBBdXRob3JpdHkgRTEwHhcNMTkwNDI4MTYyMDQ0WhcNNDkwNDIwMTYyMDQ0
WjB6MQswCQYDVQQGEwJDTjEWMBQGA1UECgwNQW50IEZpbmFuY2lhbDEgMB4GA1UE
CwwXQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkxMTAvBgNVBAMMKEFudCBGaW5hbmNp
YWwgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkgRTEwdjAQBgcqhkjOPQIBBgUrgQQA
IgNiAASCCRa94QI0vR5Up9Yr9HEupz6hSoyjySYqo7v837KnmjveUIUNiuC9pWAU
WP3jwLX3HkzeiNdeg22a0IZPoSUCpasufiLAnfXh6NInLiWBrjLJXDSGaY7vaokt
rpZvAdmjXTBbMAsGA1UdDwQEAwIBBjAMBgNVHRMEBTADAQH/MB0GA1UdDgQWBBRZ
4ZTgDpksHL2qcpkFkxD2zVd16TAfBgNVHSMEGDAWgBRZ4ZTgDpksHL2qcpkFkxD2
zVd16TAMBggqhkjOPQQDAwUAA2gAMGUCMQD4IoqT2hTUn0jt7oXLdMJ8q4vLp6sg
wHfPiOr9gxreb+e6Oidwd2LDnC4OUqCWiF8CMAzwKs4SnDJYcMLf2vpkbuVE4dTH
Rglz+HGcTLWsFs4KxLsq7MuU+vJTBUeDJeDjdA==
-----END CERTIFICATE-----

-----BEGIN CERTIFICATE-----
MIIDxTCCAq2gAwIBAgIUEMdk6dVgOEIS2cCP0Q43P90Ps5YwDQYJKoZIhvcNAQEF
BQAwajELMAkGA1UEBhMCQ04xEzARBgNVBAoMCmlUcnVzQ2hpbmExHDAaBgNVBAsM
E0NoaW5hIFRydXN0IE5ldHdvcmsxKDAmBgNVBAMMH2lUcnVzQ2hpbmEgQ2xhc3Mg
MiBSb290IENBIC0gRzMwHhcNMTMwNDE4MDkzNjU2WhcNMzMwNDE4MDkzNjU2WjBq
MQswCQYDVQQGEwJDTjETMBEGA1UECgwKaVRydXNDaGluYTEcMBoGA1UECwwTQ2hp
bmEgVHJ1c3QgTmV0d29yazEoMCYGA1UEAwwfaVRydXNDaGluYSBDbGFzcyAyIFJv
b3QgQ0EgLSBHMzCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAOPPShpV
nJbMqqCw6Bz1kehnoPst9pkr0V9idOwU2oyS47/HjJXk9Rd5a9xfwkPO88trUpz5
4GmmwspDXjVFu9L0eFaRuH3KMha1Ak01citbF7cQLJlS7XI+tpkTGHEY5pt3EsQg
wykfZl/A1jrnSkspMS997r2Gim54cwz+mTMgDRhZsKK/lbOeBPpWtcFizjXYCqhw
WktvQfZBYi6o4sHCshnOswi4yV1p+LuFcQ2ciYdWvULh1eZhLxHbGXyznYHi0dGN
z+I9H8aXxqAQfHVhbdHNzi77hCxFjOy+hHrGsyzjrd2swVQ2iUWP8BfEQqGLqM1g
KgWKYfcTGdbPB1MCAwEAAaNjMGEwHQYDVR0OBBYEFG/oAMxTVe7y0+408CTAK8hA
uTyRMB8GA1UdIwQYMBaAFG/oAMxTVe7y0+408CTAK8hAuTyRMA8GA1UdEwEB/wQF
MAMBAf8wDgYDVR0PAQH/BAQDAgEGMA0GCSqGSIb3DQEBBQUAA4IBAQBLnUTfW7hp
emMbuUGCk7RBswzOT83bDM6824EkUnf+X0iKS95SUNGeeSWK2o/3ALJo5hi7GZr3
U8eLaWAcYizfO99UXMRBPw5PRR+gXGEronGUugLpxsjuynoLQu8GQAeysSXKbN1I
UugDo9u8igJORYA+5ms0s5sCUySqbQ2R5z/GoceyI9LdxIVa1RjVX8pYOj8JFwtn
DJN3ftSFvNMYwRuILKuqUYSHc2GPYiHVflDh5nDymCMOQFcFG3WsEuB+EYQPFgIU
1DHmdZcz7Llx8UOZXX2JupWCYzK1XhJb+r4hK5ncf/w8qGtYlmyJpxk3hr1TfUJX
Yf4Zr0fJsGuv
-----END CERTIFICATE-----`)
	cert.AlipayPublicContentRSA2 = []byte(`-----BEGIN CERTIFICATE-----
MIIDszCCApugAwIBAgIQICISBid77Ii2blp8Za+onzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UE
BhMCQ04xGzAZBgNVBAoMEkFudCBGaW5hbmNpYWwgdGVzdDElMCMGA1UECwwcQ2VydGlmaWNhdGlv
biBBdXRob3JpdHkgdGVzdDE+MDwGA1UEAww1QW50IEZpbmFuY2lhbCBDZXJ0aWZpY2F0aW9uIEF1
dGhvcml0eSBDbGFzcyAyIFIxIHRlc3QwHhcNMjIxMjA2MDMwODM5WhcNMjMxMjA2MDMwODM5WjCB
hDELMAkGA1UEBhMCQ04xHzAdBgNVBAoMFmtjZHVqaTEzNzdAc2FuZGJveC5jb20xDzANBgNVBAsM
BkFsaXBheTFDMEEGA1UEAww65pSv5LuY5a6dKOS4reWbvSnnvZHnu5zmioDmnK/mnInpmZDlhazl
j7gtMjA4ODYyMTk5Mzg4NDk2NTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBANsvbtHI
lf0/f1V5Gvby7CtxVa01zwFIjr9LJd0170oL1cModwZSaZagGfuc9r96AYkFK7dwKJ6DyMjU1m6Y
5CS9w7vKRA+ksfbhbywOjQBRxSbpIb7f1M2kAh+/vT7CecY/vuBEU0xtSLPBUrINBf6elamA9ZnW
ALlnoZWnLqWUWsvxZVKXKdiLdSOc+ZYCocsMKryMGexB8gh0oTMW4pXdGgQi5WGwyRDTmyQave/H
sX0dvDaoohqPeY4R3PnoYiIE5/ybZLYm3c7Q//b+LHl8O5W++zIJtsFxlZup4K4ul2UYRAeoujh3
hNtVx+I2tmjGxGXrudvs4WZPtBQ22R8CAwEAAaMSMBAwDgYDVR0PAQH/BAQDAgTwMA0GCSqGSIb3
DQEBCwUAA4IBAQAF7Pyh5yfn13rjiVwFi8K5D5igY3SJtvCmte1//OuCscgNinmiEg/kquJi1wm0
pTQ9SogsR95p+I2Vo8MnV5uQEyVsX45CmChkITck29Lm3L1sP6jaILDNqrI31mSGo9xbSHWp+COU
dvBfIgUSG+J6b9cuSVW7a6ixLmFqfrGaEW79Q256u1+yaSmhWn3mn9Im3C7Is4kXyeF8KR0oLZPR
DV2kf1fO81TQAotkR59bfbenr9uux9Zgnl4gWj3WE5HNQuCK4BYWv5itS6qWSX/CM7U4a7Fwuj2f
AmNnGj5x0znKBd6PLrOCh0XWgy+Aca/fIBoRcmdErzobZZc+CPKB
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIDszCCApugAwIBAgIQIBkIGbgVxq210KxLJ+YA/TANBgkqhkiG9w0BAQsFADCBhDELMAkGA1UE
BhMCQ04xFjAUBgNVBAoMDUFudCBGaW5hbmNpYWwxJTAjBgNVBAsMHENlcnRpZmljYXRpb24gQXV0
aG9yaXR5IHRlc3QxNjA0BgNVBAMMLUFudCBGaW5hbmNpYWwgQ2VydGlmaWNhdGlvbiBBdXRob3Jp
dHkgUjEgdGVzdDAeFw0xOTA4MTkxMTE2MDBaFw0yNDA4MDExMTE2MDBaMIGRMQswCQYDVQQGEwJD
TjEbMBkGA1UECgwSQW50IEZpbmFuY2lhbCB0ZXN0MSUwIwYDVQQLDBxDZXJ0aWZpY2F0aW9uIEF1
dGhvcml0eSB0ZXN0MT4wPAYDVQQDDDVBbnQgRmluYW5jaWFsIENlcnRpZmljYXRpb24gQXV0aG9y
aXR5IENsYXNzIDIgUjEgdGVzdDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAMh4FKYO
ZyRQHD6eFbPKZeSAnrfjfU7xmS9Yoozuu+iuqZlb6Z0SPLUqqTZAFZejOcmr07ln/pwZxluqplxC
5+B48End4nclDMlT5HPrDr3W0frs6Xsa2ZNcyil/iKNB5MbGll8LRAxntsKvZZj6vUTMb705gYgm
VUMILwi/ZxKTQqBtkT/kQQ5y6nOZsj7XI5rYdz6qqOROrpvS/d7iypdHOMIM9Iz9DlL1mrCykbBi
t25y+gTeXmuisHUwqaRpwtCGK4BayCqxRGbNipe6W73EK9lBrrzNtTr9NaysesT/v+l25JHCL9tG
wpNr1oWFzk4IHVOg0ORiQ6SUgxZUTYcCAwEAAaMSMBAwDgYDVR0PAQH/BAQDAgTwMA0GCSqGSIb3
DQEBCwUAA4IBAQBWThEoIaQoBX2YeRY/I8gu6TYnFXtyuCljANnXnM38ft+ikhE5mMNgKmJYLHvT
yWWWgwHoSAWEuml7EGbE/2AK2h3k0MdfiWLzdmpPCRG/RJHk6UB1pMHPilI+c0MVu16OPpKbg5Vf
LTv7dsAB40AzKsvyYw88/Ezi1osTXo6QQwda7uefvudirtb8FcQM9R66cJxl3kt1FXbpYwheIm/p
j1mq64swCoIYu4NrsUYtn6CV542DTQMI5QdXkn+PzUUly8F6kDp+KpMNd0avfWNL5+O++z+F5Szy
1CPta1D7EQ/eYmMP+mOQ35oifWIoFCpN6qQVBS/Hob1J/UUyg7BW
-----END CERTIFICATE-----
`)
	var err error

	aliClient, err = alipay.NewClient(cert.Appid, cert.PrivateKey, false)
	if err != nil {
		log.Fatal("初始化支付宝失败", err)
	}

	// Debug开关，输出/关闭日志
	aliClient.DebugSwitch = gopay.DebugOn

	// 配置公共参数
	aliClient.SetCharset("utf-8").
		SetSignType("RSA2").
		// SetAppAuthToken("")
		SetReturnUrl("http://10.7.14.40:8080/#/").
		SetNotifyUrl("https://ea12-113-200-76-118.jp.ngrok.io/v1/notify")

	// 自动同步验签（只支持证书模式）
	// 传入 alipayCertPublicKey_RSA2.crt 内容
	aliClient.AutoVerifySign(cert.AlipayPublicContentRSA2)
}

func (s *PayService) CreateTrade(ctx context.Context, req *pb.CreateTradeRequest) (*pb.CreateTradeReply, error) {

	// 传入证书内容
	err := aliClient.SetCertSnByContent(cert.AppPublicContent, cert.AlipayRootContent, cert.AlipayPublicContentRSA2)
	// 传入证书文件路径
	//err := client.SetCertSnByPath("cert/appCertPublicKey_2021000117673683.crt", "cert/alipayRootCert.crt", "cert/alipayCertPublicKey_RSA2.crt")
	if err != nil {
		s.log.Error("SetCertSn:", err)
		return nil, err
	}
	bm := make(gopay.BodyMap)
	bm.Set("subject", "订阅测试支付").
		Set("out_trade_no", "GZ201909081743434320102"+fmt.Sprintf("%s", rand.Int31())).
		Set("total_amount", "88.88").
		Set("return_url", "http://10.7.14.40:8080/#/").
		Set("notify_url", "http://82.157.153.93:8000/v1/notify").
		Set("qr_pay_mode", "4").
		Set("qrcode_width", "200").
		//Set("time_expire", "2022-12-07 15:27:59").
		Set("product_code", "FAST_INSTANT_TRADE_PAY")

	aliRes, err := aliClient.TradePagePay(ctx, bm)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	reply := &pb.CreateTradeReply{Data: &pb.CreateTradeReplyContent{}}
	reply.Data.FormStr = aliRes
	return reply, err
}

func (s *PayService) Notify(ctx context.Context, req *pb.NotifyRequest) (*pb.NotifyReply, error) {
	s.log.Debug("req", req)
	reply := &pb.NotifyReply{}
	// 验签
	ok, err := alipay.VerifySignWithCert(cert.AlipayPublicContentRSA2, req)
	if err != nil {
		s.log.Error("Notify VerifySignWithCert err", err)
		return nil, err
	}
	log.Debug("支付宝验签是否通过:", ok)
	if ok {
		reply.Value = "success"
	}
	return reply, err
}
