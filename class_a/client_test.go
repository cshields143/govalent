package class_a

import (
	"github.com/cshields143/govalent/client"
	"github.com/google/go-cmp/cmp"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestClassA_Blocks(t *testing.T) {
    tests := []struct {
        desc        string
        response    string
        want        Blocks
    }{
        {
            desc: "WhenNonEmpty",
            response: `{"data":{"updated_at":"2022-02-18T02:51:16.364252395Z","items":[{"signed_at":"2021-01-01T00:20:01Z","height":11565118}],"pagination":null},"error":false,"error_message":"","error_code":0}`,
            want: Blocks{
                UpdatedAt: time.Date(2022, 2, 18, 2, 51, 16, 364252395, time.UTC),
                Items: []struct {
                    SignedAt time.Time `json:"signed_at"`
                    Height int `json:"height"`
                }{{
                    time.Date(2021, 1, 1, 0, 20, 1, 0, time.UTC),
                    11565118,
                }},
                Pagination: Pagination{},
            },
        },
    }
    for _, tc := range tests {
        tc := tc
        t.Run(tc.desc, func (t *testing.T) {
		    s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			    w.WriteHeader(http.StatusOK)
			    w.Header().Set("Content-Type", "application/json")
			    _, err := io.WriteString(w, tc.response)
			    if err != nil {
				    t.Fatalf("while write response got err: %v", err)
			    }
		    }))
		    defer s.Close()
		    api := client.New(s.URL, "ckey_test", s.Client())
		    classA := Client{API: *api}
		    got, err := classA.Blocks("chain", "start", "end", PaginateParams{})
		    if diff := cmp.Diff(got, tc.want); diff != "" || err != nil {
			    t.Errorf("%v.Blocks(chain, startDate, endDate) has diff (-got/+want)\n: %v", classA, diff)
			    t.Errorf("%v.Blocks(chain, startDate, endDate) got err: %v want nil", classA, err)
		    }
        })
    }
}

func TestClassA_GetTokenBalances_WhenValid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		desc     string
		response string
		want     Portfolios
	}{
		{
			desc:     "WhenEmptyResponse",
			response: `{"data": null}`,
			want:     Portfolios{},
		},
		{
			desc:     "WhenNoNFT",
			response: `{"data":{"address":"0x01","updated_at":"2021-05-05T18:00:00Z","quote_currency":"USD","chain_id":56,"items":[{"contract_name":"Binance Coin","contract_ticker_symbol":"BNB","contract_address":"0xb1","type":"cryptocurrency","balance":"10","quote_rate":645,"quote":0.02}]}}`,
			want: Portfolios{
				Address:       "0x01",
				UpdatedAt:     time.Date(2021, 5, 5, 18, 0, 0, 0, time.UTC),
				QuoteCurrency: "USD",
				ChainID:       56,
				Items: []Portfolio{
					{
						ContractDecimals:     0,
						ContractName:         "Binance Coin",
						ContractTickerSymbol: "BNB",
						ContractAddress:      "0xb1",
						SupportsErc:          nil,
						LogoUrl:              "",
						Type:                 "cryptocurrency",
						Balance:              "10",
						QuoteRate:            645,
						Quote:                0.02,
						NftData:              nil,
					},
				},
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				_, err := io.WriteString(w, tc.response)
				if err != nil {
					t.Fatalf("while write response got err: %v", err)
				}
			}))
			defer s.Close()
			api := client.New(s.URL, "ckey_test", s.Client())
			classA := Client{API: *api}
			got, err := classA.TokenBalances("chain", "address", BalanceParams{Nft: true})

			if diff := cmp.Diff(got, tc.want); diff != "" || err != nil {
				t.Errorf("%v.TokenBalances(chain, address) has diff (-got/+want)\n: %v", classA, diff)
				t.Errorf("%v.TokenBalances(chain, address) got err: %v want nil", classA, err)
			}
		})
	}
}
