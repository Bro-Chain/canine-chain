package cli

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	eciesgo "github.com/ecies/go/v2"
	filetypes "github.com/jackal-dao/canine/x/filetree/types"
	"github.com/spf13/cobra"
)

func MakePrivateKey(clientCtx client.Context) (*eciesgo.PrivateKey, error) {
	signed, _, err := clientCtx.Keyring.Sign(clientCtx.GetFromName(), []byte("jackal_init"))
	if err != nil {
		return nil, err
	}

	k := secp256k1.GenPrivKeyFromSecret(signed)

	newKey := eciesgo.NewPrivateKeyFromBytes(k.Bytes())

	return newKey, nil
}

func merkleHelper(argHashpath string) (string, string) {

	//Cut out the / at the end for compatibility with types/merkle-paths.go
	trimPath := strings.TrimSuffix(argHashpath, "/")
	chunks := strings.Split(trimPath, "/")

	parentString := strings.Join(chunks[0:len(chunks)-1], "/")
	childString := string(chunks[len(chunks)-1])
	parentHash := filetypes.MerklePath(parentString)

	h := sha256.New()
	h.Write([]byte(childString))
	childHash := fmt.Sprintf("%x", h.Sum(nil))

	return parentHash, childHash

}

// Owner address is whoever owns this file/folder
func MakeOwnerAddress(merklePath string, user string) string {

	h := sha256.New()
	h.Write([]byte(user))
	hash := h.Sum(nil)
	accountHash := fmt.Sprintf("%x", hash)

	//h1 is so named as to differentiate it from h above--else compiler complains
	h1 := sha256.New()
	h1.Write([]byte(fmt.Sprintf("o%s%s", merklePath, accountHash)))
	hash1 := h1.Sum(nil)
	ownerAddress := fmt.Sprintf("%x", hash1)

	return ownerAddress
}

func encryptFileAESKey(cmd *cobra.Command, key string, argKeys string) ([]byte, error) {

	clientCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return nil, err
	}

	queryClient := filetypes.NewQueryClient(clientCtx)

	res, err := queryClient.Pubkey(cmd.Context(), &filetypes.QueryGetPubkeyRequest{Address: key})
	if err != nil {
		return nil, filetypes.ErrPubKeyNotFound
	}

	pkey, err := eciesgo.NewPublicKeyFromHex(res.Pubkey.Key)
	if err != nil {
		return nil, err
	}

	encrypted, err := clientCtx.Keyring.Encrypt(pkey.Bytes(false), []byte(argKeys))
	if err != nil {
		return nil, err
	}
	return encrypted, nil

}

func getCallerAddress(ctx client.Context, cmd *cobra.Command) (*string, error) {

	fromAddress := ctx.GetFromAddress().String()
	return &fromAddress, nil

}

func JSONMarshalViewersAndEditors(viewers map[string]string, editors map[string]string, viewersToNotify []string, editorsToNotify []string) ([]byte, []byte, []byte, []byte, error) {

	jsonViewers, err := json.Marshal(viewers)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	jsonEditors, err := json.Marshal(editors)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	jsonViewersToNotify, err := json.Marshal(viewersToNotify)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	jsonEditorsToNotify, err := json.Marshal(editorsToNotify)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return jsonViewers, jsonEditors, jsonViewersToNotify, jsonEditorsToNotify, nil

}
