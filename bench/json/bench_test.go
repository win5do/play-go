package json

import (
	"encoding/json"
	"fmt"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/require"
)

func BenchmarkUnmarshal(b *testing.B) {
	for _, v := range []struct {
		name     string
		data     []byte
		getModel func() any
	}{
		{
			name: "S",
			data: smallFixture,
			getModel: func() any {
				return SmallPayload{}
			},
		},
		{
			name: "M",
			data: mediumFixture,
			getModel: func() any {
				return MediumPayload{}
			},
		},
		{
			name: "L",
			data: largeFixture,
			getModel: func() any {
				return LargePayload{}
			},
		},
	} {
		b.Run(fmt.Sprintf("%s=std", v.name), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				model := v.getModel()
				err := json.Unmarshal(v.data, &model)
				require.NoError(b, err)
			}
		})

		b.Run(fmt.Sprintf("%s=jtr", v.name), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				model := v.getModel()
				err := jsoniter.Unmarshal(v.data, &model)
				require.NoError(b, err)
			}
		})
	}
}

func BenchmarkMarshal(b *testing.B) {
	var err error
	var modelS SmallPayload
	var modelM MediumPayload
	var modelL LargePayload
	err = json.Unmarshal(smallFixture, &modelS)
	require.NoError(b, err)
	err = json.Unmarshal(mediumFixture, &modelM)
	require.NoError(b, err)
	err = json.Unmarshal(largeFixture, &modelL)
	require.NoError(b, err)

	for _, v := range []struct {
		name string
		data any
	}{
		{
			name: "S",
			data: modelS,
		},
		{
			name: "M",
			data: modelM,
		},
		{
			name: "L",
			data: modelL,
		},
	} {
		b.Run(fmt.Sprintf("%s=std", v.name), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_, err := json.Marshal(v.data)
				require.NoError(b, err)
			}
		})

		b.Run(fmt.Sprintf("%s=jtr", v.name), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_, err := jsoniter.Marshal(v.data)
				require.NoError(b, err)
			}
		})
	}
}
