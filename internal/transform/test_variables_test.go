package transform

import (
	"time"

	"github.com/stellar/stellar-etl/internal/utils"

	"github.com/stellar/go/ingest"
	"github.com/stellar/go/xdr"
)

var genericSourceAccount, _ = xdr.NewMuxedAccount(xdr.CryptoKeyTypeKeyTypeEd25519, xdr.Uint256([32]byte{}))
var genericAccountID, _ = xdr.NewAccountId(xdr.PublicKeyTypePublicKeyTypeEd25519, xdr.Uint256([32]byte{}))
var genericAccountAddress, _ = genericAccountID.GetAddress()
var genericManageBuyOfferOperation = xdr.Operation{
	SourceAccount: &genericSourceAccount,
	Body: xdr.OperationBody{
		Type:             xdr.OperationTypeManageBuyOffer,
		ManageBuyOfferOp: &xdr.ManageBuyOfferOp{},
	},
}
var genericBumpOperation = xdr.Operation{
	SourceAccount: &genericSourceAccount,
	Body: xdr.OperationBody{
		Type:           xdr.OperationTypeBumpSequence,
		BumpSequenceOp: &xdr.BumpSequenceOp{},
	},
}
var genericBumpOperationEnvelope = xdr.TransactionV1Envelope{
	Tx: xdr.Transaction{
		SourceAccount: genericSourceAccount,
		Memo:          xdr.Memo{},
		Operations: []xdr.Operation{
			genericBumpOperation,
		},
	},
}
var genericManageBuyOfferEnvelope = xdr.TransactionV1Envelope{
	Tx: xdr.Transaction{
		SourceAccount: genericSourceAccount,
		Memo:          xdr.Memo{},
		Operations: []xdr.Operation{
			genericManageBuyOfferOperation,
		},
	},
}
var genericLedgerTransaction = ingest.LedgerTransaction{
	Index: 1,
	Envelope: xdr.TransactionEnvelope{
		Type: xdr.EnvelopeTypeEnvelopeTypeTx,
		V1:   &genericBumpOperationEnvelope,
	},
	Result: utils.CreateSampleResultMeta(true, 10).Result,
}
var genericLedgerHeaderHistoryEntry = xdr.LedgerHeaderHistoryEntry{}
var genericCloseTime = time.Unix(0, 0)

// a selection of hardcoded accounts with their IDs and addresses
var testAccount1Address = "GCEODJVUUVYVFD5KT4TOEDTMXQ76OPFOQC2EMYYMLPXQCUVPOB6XRWPQ"
var testAccount1ID, _ = xdr.AddressToAccountId(testAccount1Address)
var testAccount1 = testAccount1ID.ToMuxedAccount()

var testAccount2Address = "GAOEOQMXDDXPVJC3HDFX6LZFKANJ4OOLQOD2MNXJ7PGAY5FEO4BRRAQU"
var testAccount2ID, _ = xdr.AddressToAccountId(testAccount2Address)
var testAccount2 = testAccount2ID.ToMuxedAccount()

var testAccount3Address = "GBT4YAEGJQ5YSFUMNKX6BPBUOCPNAIOFAVZOF6MIME2CECBMEIUXFZZN"
var testAccount3ID, _ = xdr.AddressToAccountId(testAccount3Address)
var testAccount3 = testAccount3ID.ToMuxedAccount()

var testAccount4Address = "GBVVRXLMNCJQW3IDDXC3X6XCH35B5Q7QXNMMFPENSOGUPQO7WO7HGZPA"
var testAccount4ID, _ = xdr.AddressToAccountId(testAccount4Address)
var testAccount4 = testAccount4ID.ToMuxedAccount()

// a selection of hardcoded assets and their AssetOutput representations

var usdtAsset = xdr.Asset{
	Type: xdr.AssetTypeAssetTypeCreditAlphanum4,
	AlphaNum4: &xdr.AlphaNum4{
		AssetCode: xdr.AssetCode4([4]byte{0x55, 0x53, 0x44, 0x54}),
		Issuer:    testAccount4ID,
	},
}

var usdtTrustLineAsset = xdr.TrustLineAsset{
	Type: xdr.AssetTypeAssetTypeCreditAlphanum4,
	AlphaNum4: &xdr.AlphaNum4{
		AssetCode: xdr.AssetCode4([4]byte{0x55, 0x53, 0x54, 0x54}),
		Issuer:    testAccount3ID,
	},
}

var usdtChangeTrustAsset = xdr.ChangeTrustAsset{
	Type: xdr.AssetTypeAssetTypeCreditAlphanum4,
	AlphaNum4: &xdr.AlphaNum4{
		AssetCode: xdr.AssetCode4([4]byte{0x55, 0x53, 0x53, 0x44}),
		Issuer:    testAccount4ID,
	},
}

var usdtAssetPath = Path{
	AssetType:   "credit_alphanum4",
	AssetCode:   "USDT",
	AssetIssuer: testAccount4Address,
}

var ethAsset = xdr.Asset{
	Type: xdr.AssetTypeAssetTypeCreditAlphanum4,
	AlphaNum4: &xdr.AlphaNum4{
		AssetCode: xdr.AssetCode4([4]byte{0x45, 0x54, 0x48}),
		Issuer:    testAccount3ID,
	},
}

var ethTrustLineAsset = xdr.TrustLineAsset{
	Type: xdr.AssetTypeAssetTypeCreditAlphanum4,
	AlphaNum4: &xdr.AlphaNum4{
		AssetCode: xdr.AssetCode4([4]byte{0x45, 0x54, 0x48}),
		Issuer:    testAccount3ID,
	},
}

var ethAssetPath = Path{
	AssetType:   "credit_alphanum4",
	AssetCode:   "ETH",
	AssetIssuer: testAccount1Address,
}

var nativeAsset = xdr.MustNewNativeAsset()
var nativeAssetPath = Path{
	AssetType: "native",
}

var nativeTrustLineAsset = xdr.MustNewNativeAsset().ToTrustLineAsset()

var genericClaimableBalance = xdr.ClaimableBalanceId{
	Type: xdr.ClaimableBalanceIdTypeClaimableBalanceIdTypeV0,
	V0:   &xdr.Hash{1, 2, 3, 4, 5, 6, 7, 8, 9},
}

var testClaimant = xdr.Claimant{
	Type: xdr.ClaimantTypeClaimantTypeV0,
	V0: &xdr.ClaimantV0{
		Destination: testAccount1ID,
		Predicate: xdr.ClaimPredicate{
			Type: xdr.ClaimPredicateTypeClaimPredicateUnconditional,
		},
	},
}

var testClaimantDetails = Claimant{
	Destination: testAccount1Address,
	Predicate:   xdr.ClaimPredicate{},
}
