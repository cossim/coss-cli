package pgp

import (
	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"github.com/ProtonMail/gopenpgp/v2/helper"
	"os"
	"strings"
)

func GenerateKeyPair() error {
	rsaKey, err := helper.GenerateKey("coss-im", "max.mustermann@example.com", []byte("LongSecret"), "rsa", 2048)
	if err != nil {
		return err
	}

	cacheDir := "./config/pgp"
	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		err := os.Mkdir(cacheDir, 0755) // 创建文件夹并设置权限
		if err != nil {
			return err
		}
	}
	// 保存私钥到文件
	privateKeyFile, err := os.Create(cacheDir + "/private_key")
	if err != nil {
		return err
	}

	_, err = privateKeyFile.WriteString(rsaKey)
	if err != nil {
		privateKeyFile.Close()
		return err
	}
	privateKeyFile.Close()

	// 保存公钥到文件
	publicKeyFile, err := os.Create(cacheDir + "/public_key")
	if err != nil {
		return err
	}
	keyRing, err := crypto.NewKeyFromArmoredReader(strings.NewReader(rsaKey))
	if err != nil {
		return err
	}

	publicKey, err := keyRing.GetArmoredPublicKey()
	if err != nil {
		return err
	}
	_, err = publicKeyFile.WriteString(publicKey)
	if err != nil {
		publicKeyFile.Close()
		return err
	}
	publicKeyFile.Close()
	return nil
}
