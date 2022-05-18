/*
SPDX-FileCopyrightText: 2021 SAP SE or an SAP affiliate company and Gardener contributors

SPDX-License-Identifier: Apache-2.0
*/

// this file contains code snippets from https://github.com/kubernetes/kubernetes/blob/master/pkg/serviceaccount/jwt.go to not have a dependency to k/k only for the sake of the tests

package test

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/x509"
	"encoding/base64"
	"fmt"

	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	corev1 "k8s.io/api/core/v1"
)

type JwtTokenGenerator struct {
	iss    string
	signer jose.Signer
}

func (j *JwtTokenGenerator) GenerateToken(claims *jwt.Claims, privateClaims interface{}) (string, error) {
	// claims are applied in reverse precedence
	return jwt.Signed(j.signer).
		Claims(privateClaims).
		Claims(claims).
		Claims(&jwt.Claims{
			Issuer: j.iss,
		}).
		CompactSerialize()
}

// NewJWTTokenGenerator returns a TokenGenerator that generates signed JWT tokens, using the given privateKey.
// privateKey is a PEM-encoded byte array of a private RSA key.
func NewJWTTokenGenerator(iss string, privateKey interface{}) (*JwtTokenGenerator, error) {
	var (
		signer jose.Signer
		err    error
	)

	switch pk := privateKey.(type) {
	case *ecdsa.PrivateKey:
		signer, err = signerFromECDSAPrivateKey(pk)
		if err != nil {
			return nil, fmt.Errorf("could not generate signer for ECDSA keypair: %v", err)
		}
	default:
		return nil, fmt.Errorf("unknown private key type %T, must be *rsa.PrivateKey, *ecdsa.PrivateKey, or jose.OpaqueSigner", privateKey)
	}

	return &JwtTokenGenerator{
		iss:    iss,
		signer: signer,
	}, nil
}

// keyIDFromPublicKey derives a key ID non-reversibly from a public key.
//
// The Key ID is field on a given on JWTs and JWKs that help relying parties
// pick the correct key for verification when the identity party advertises
// multiple keys.
//
// Making the derivation non-reversible makes it impossible for someone to
// accidentally obtain the real key from the key ID and use it for token
// validation.
func keyIDFromPublicKey(publicKey interface{}) (string, error) {
	publicKeyDERBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", fmt.Errorf("failed to serialize public key to DER format: %v", err)
	}

	hasher := crypto.SHA256.New()
	hasher.Write(publicKeyDERBytes)
	publicKeyDERHash := hasher.Sum(nil)

	keyID := base64.RawURLEncoding.EncodeToString(publicKeyDERHash)

	return keyID, nil
}

func signerFromECDSAPrivateKey(keyPair *ecdsa.PrivateKey) (jose.Signer, error) {
	var alg jose.SignatureAlgorithm

	switch keyPair.Curve {
	case elliptic.P256():
		alg = jose.ES256
	case elliptic.P384():
		alg = jose.ES384
	case elliptic.P521():
		alg = jose.ES512
	default:
		return nil, fmt.Errorf("unknown private key curve, must be 256, 384, or 521")
	}

	keyID, err := keyIDFromPublicKey(&keyPair.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to derive keyID: %v", err)
	}

	// Wrap the ECDSA keypair in a JOSE JWK with the designated key ID.
	privateJWK := &jose.JSONWebKey{
		Algorithm: string(alg),
		Key:       keyPair,
		KeyID:     keyID,
		Use:       "sig",
	}

	signer, err := jose.NewSigner(
		jose.SigningKey{
			Algorithm: alg,
			Key:       privateJWK,
		},
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create signer: %v", err)
	}

	return signer, nil
}

func LegacyClaims(serviceAccount corev1.ServiceAccount, secret corev1.Secret) (*jwt.Claims, interface{}) {
	return &jwt.Claims{
			Subject: MakeUsername(serviceAccount.Namespace, serviceAccount.Name),
		}, &legacyPrivateClaims{
			Namespace:          serviceAccount.Namespace,
			ServiceAccountName: serviceAccount.Name,
			ServiceAccountUID:  string(serviceAccount.UID),
			SecretName:         secret.Name,
		}
}

func MakeUsername(namespace, name string) string {
	return "system:serviceaccount:" + namespace + ":" + name
}

const LegacyIssuer = "kubernetes/serviceaccount"

type legacyPrivateClaims struct {
	ServiceAccountName string `json:"kubernetes.io/serviceaccount/service-account.name"`
	ServiceAccountUID  string `json:"kubernetes.io/serviceaccount/service-account.uid"`
	SecretName         string `json:"kubernetes.io/serviceaccount/secret.name"`
	Namespace          string `json:"kubernetes.io/serviceaccount/namespace"`
}