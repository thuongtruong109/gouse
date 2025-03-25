package gouse

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"io"
	"os"

	// "crypto/ecdsa"
	// "crypto/ed25519"
	// "crypto/elliptic"
	// "crypto/rsa"
	// "crypto/x509"
	// "crypto/x509/pkix"
	// "encoding/pem"
	// "flag"
	// "log"
	// "math/big"
	// "net"
	// "strings"
	// "time"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/pbkdf2"
)

func EncodeData(input []byte) ([]byte, error) {
	encodedData := base64.StdEncoding.EncodeToString(input)
	return []byte(encodedData), nil
}

func DecodeData(input []byte) ([]byte, error) {
	decodedData, err := base64.StdEncoding.DecodeString(string(input))
	if err != nil {
		return nil, err
	}
	return decodedData, nil
}

func EncryptFile(source string, password []byte) {
	if _, err := os.Stat(source); os.IsNotExist(err) {
		panic(err.Error())
	}

	plaintext, err := os.ReadFile(source)
	if err != nil {
		panic(err.Error())
	}

	key := password
	nonce := make([]byte, 12)

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	dk := pbkdf2.Key(key, nonce, 4096, 32, sha1.New)

	block, err := aes.NewCipher(dk)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	ciphertext = append(ciphertext, nonce...)

	f, err := os.Create(source)
	if err != nil {
		panic(err.Error())
	}
	_, err = io.Copy(f, bytes.NewReader(ciphertext))
	if err != nil {
		panic(err.Error())
	}
}

func DecryptFile(source string, password []byte) {
	if _, err := os.Stat(source); os.IsNotExist(err) {
		panic(err.Error())
	}

	ciphertext, err := os.ReadFile(source)
	if err != nil {
		panic(err.Error())
	}

	key := password
	salt := ciphertext[len(ciphertext)-12:]
	str := hex.EncodeToString(salt)

	nonce, err := hex.DecodeString(str)
	if err != nil {
		panic(err.Error())
	}

	dk := pbkdf2.Key(key, nonce, 4096, 32, sha1.New)

	block, err := aes.NewCipher(dk)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext[:len(ciphertext)-12], nil)
	if err != nil {
		panic(err.Error())
	}

	f, err := os.Create(source)
	if err != nil {
		panic(err.Error())
	}
	_, err = io.Copy(f, bytes.NewReader(plaintext))
	if err != nil {
		panic(err.Error())
	}
}

func EncryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", ErrorHashPassword
	}

	return string(hashedPassword), nil
}

func ComparePassword(password, receivedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(receivedPassword)); err != nil {
		return ErrorComparePassword
	}

	return nil
}

// var (
// 	host       = flag.String("host", "443", "Comma-separated hostnames and IPs to generate a certificate for")
// 	validFrom  = flag.String("start-date", "", "Creation date formatted as Jan 1 15:04:05 2011")
// 	validFor   = flag.Duration("duration", 365*24*time.Hour, "Duration that certificate is valid for")
// 	isCA       = flag.Bool("ca", false, "whether this cert should be its own Certificate Authority")
// 	rsaBits    = flag.Int("rsa-bits", 2048, "Size of RSA key to generate. Ignored if --ecdsa-curve is set")
// 	ecdsaCurve = flag.String("ecdsa-curve", "", "ECDSA curve to use to generate a key. Valid values are P224, P256 (recommended), P384, P521")
// 	ed25519Key = flag.Bool("ed25519", false, "Generate an Ed25519 key")
// )

// type ITls struct {
// 	Host       string
// 	ValidFrom  string
// 	ValidFor   time.Duration
// 	IsCA       bool
// 	RsaBits    int
// 	EcdsaCurve string
// }

// func (tls *ITls) TLSConfig(certPath, keyPath string) {
// 	flag.Parse()

// 	if len(*tls.Host) == 0 {
// 		log.Fatalf("Missing required --host parameter")
// 	}

// 	var priv any
// 	var err error
// 	switch *tls.EcdsaCurve {
// 	case "":
// 		if *ed25519Key {
// 			_, priv, err = ed25519.GenerateKey(rand.Reader)
// 		} else {
// 			priv, err = rsa.GenerateKey(rand.Reader, *rsaBits)
// 		}
// 	case "P224":
// 		priv, err = ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
// 	case "P256":
// 		priv, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
// 	case "P384":
// 		priv, err = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
// 	case "P521":
// 		priv, err = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
// 	default:
// 		log.Fatalf("Unrecognized elliptic curve: %q", *ecdsaCurve)
// 	}
// 	if err != nil {
// 		log.Fatalf("Failed to generate private key: %v", err)
// 	}

// 	// ECDSA, ED25519 and RSA subject keys should have the DigitalSignature
// 	// KeyUsage bits set in the x509.Certificate template
// 	keyUsage := x509.KeyUsageDigitalSignature
// 	// Only RSA subject keys should have the KeyEncipherment KeyUsage bits set. In
// 	// the context of TLS this KeyUsage is particular to RSA key exchange and
// 	// authentication.
// 	if _, isRSA := priv.(*rsa.PrivateKey); isRSA {
// 		keyUsage |= x509.KeyUsageKeyEncipherment
// 	}

// 	var notBefore time.Time
// 	if len(*validFrom) == 0 {
// 		notBefore = time.Now()
// 	} else {
// 		notBefore, err = time.Parse("Jan 2 15:04:05 2006", *validFrom)
// 		if err != nil {
// 			log.Fatalf("Failed to parse creation date: %v", err)
// 		}
// 	}

// 	notAfter := notBefore.Add(*validFor)

// 	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
// 	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
// 	if err != nil {
// 		log.Fatalf("Failed to generate serial number: %v", err)
// 	}

// 	template := x509.Certificate{
// 		SerialNumber: serialNumber,
// 		Subject: pkix.Name{
// 			Organization: []string{"zoomer"},
// 		},
// 		NotBefore: notBefore,
// 		NotAfter:  notAfter,

// 		KeyUsage:              keyUsage,
// 		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
// 		BasicConstraintsValid: true,
// 	}

// 	hosts := strings.Split(*host, ",")
// 	for _, h := range hosts {
// 		if ip := net.ParseIP(h); ip != nil {
// 			template.IPAddresses = append(template.IPAddresses, ip)
// 		} else {
// 			template.DNSNames = append(template.DNSNames, h)
// 		}
// 	}

// 	if *isCA {
// 		template.IsCA = true
// 		template.KeyUsage |= x509.KeyUsageCertSign
// 	}

// 	publicKey := func(priv any) any {
// 		switch k := priv.(type) {
// 		case *rsa.PrivateKey:
// 			return &k.PublicKey
// 		case *ecdsa.PrivateKey:
// 			return &k.PublicKey
// 		case ed25519.PrivateKey:
// 			return k.Public().(ed25519.PublicKey)
// 		default:
// 			return nil
// 		}
// 	}

// 	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, publicKey(priv), priv)
// 	if err != nil {
// 		log.Fatalf("Failed to create certificate: %v", err)
// 	}

// 	certOut, err := os.Create(certPath)
// 	if err != nil {
// 		log.Fatalf("Failed to open cert.pem for writing: %v", err)
// 	}
// 	if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
// 		log.Fatalf("Failed to write data to cert.pem: %v", err)
// 	}

// 	if err := certOut.Close(); err != nil {
// 		log.Fatalf("Error closing cert.pem: %v", err)
// 	}
// 	log.Print("wrote cert.pem\n")

// 	keyOut, err := os.OpenFile(keyPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
// 	if err != nil {
// 		log.Fatalf("Failed to open key.pem for writing: %v", err)
// 	}
// 	privBytes, err := x509.MarshalPKCS8PrivateKey(priv)
// 	if err != nil {
// 		log.Fatalf("Unable to marshal private key: %v", err)
// 	}
// 	if err := pem.Encode(keyOut, &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes}); err != nil {
// 		log.Fatalf("Failed to write data to key.pem: %v", err)
// 	}
// 	if err := keyOut.Close(); err != nil {
// 		log.Fatalf("Error closing key.pem: %v", err)
// 	}
// 	log.Print("wrote key.pem\n")
// }
