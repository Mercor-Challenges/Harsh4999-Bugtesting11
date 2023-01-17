/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// +build use_codec

package benchmark

import (
	"testing"

	"github.com/ugorji/go/codec"
)

func BenchmarkCodec_Unmarshal_M(b *testing.B) {
	var h codec.Handle = new(codec.JsonHandle)
	dec := codec.NewDecoderBytes(nil, h)

	b.SetBytes(int64(len(largeStructText)))
	for i := 0; i < b.N; i++ {
		var s LargeStruct
		dec.ResetBytes(largeStructText)
		if err := dec.Decode(&s); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkCodec_Unmarshal_S(b *testing.B) {
	var h codec.Handle = new(codec.JsonHandle)
	dec := codec.NewDecoderBytes(nil, h)

	b.SetBytes(int64(len(smallStructText)))
	for i := 0; i < b.N; i++ {
		var s LargeStruct
		dec.ResetBytes(smallStructText)
		if err := dec.Decode(&s); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkCodec_Marshal_S(b *testing.B) {
	var h codec.Handle = new(codec.JsonHandle)

	var out []byte
	enc := codec.NewEncoderBytes(&out, h)

	var l int64
	for i := 0; i < b.N; i++ {
		enc.ResetBytes(&out)
		if err := enc.Encode(&smallStructData); err != nil {
			b.Error(err)
		}
		l = int64(len(out))
		out = nil
	}

	b.SetBytes(l)
}

func BenchmarkCodec_Marshal_M(b *testing.B) {
	var h codec.Handle = new(codec.JsonHandle)

	var out []byte
	enc := codec.NewEncoderBytes(&out, h)

	var l int64
	for i := 0; i < b.N; i++ {
		enc.ResetBytes(&out)
		if err := enc.Encode(&largeStructData); err != nil {
			b.Error(err)
		}
		l = int64(len(out))
		out = nil
	}

	b.SetBytes(l)
}

func BenchmarkCodec_Marshal_L(b *testing.B) {
	var h codec.Handle = new(codec.JsonHandle)

	var out []byte
	enc := codec.NewEncoderBytes(&out, h)

	var l int64
	for i := 0; i < b.N; i++ {
		enc.ResetBytes(&out)
		if err := enc.Encode(&xlStructData); err != nil {
			b.Error(err)
		}
		l = int64(len(out))
		out = nil
	}

	b.SetBytes(l)
}

func BenchmarkCodec_Marshal_S_Reuse(b *testing.B) {
	var h codec.Handle = new(codec.JsonHandle)

	var out []byte
	enc := codec.NewEncoderBytes(&out, h)

	var l int64
	for i := 0; i < b.N; i++ {
		enc.ResetBytes(&out)
		if err := enc.Encode(&smallStructData); err != nil {
			b.Error(err)
		}
		l = int64(len(out))
		out = out[:0]
	}

	b.SetBytes(l)
}

func BenchmarkCodec_Marshal_M_Reuse(b *testing.B) {
	var h codec.Handle = new(codec.JsonHandle)

	var out []byte
	enc := codec.NewEncoderBytes(&out, h)

	var l int64
	for i := 0; i < b.N; i++ {
		enc.ResetBytes(&out)
		if err := enc.Encode(&largeStructData); err != nil {
			b.Error(err)
		}
		l = int64(len(out))
		out = out[:0]
	}

	b.SetBytes(l)
}

func BenchmarkCodec_Marshal_L_Reuse(b *testing.B) {
	var h codec.Handle = new(codec.JsonHandle)

	var out []byte
	enc := codec.NewEncoderBytes(&out, h)

	var l int64
	for i := 0; i < b.N; i++ {
		enc.ResetBytes(&out)
		if err := enc.Encode(&xlStructData); err != nil {
			b.Error(err)
		}
		l = int64(len(out))
		out = out[:0]
	}

	b.SetBytes(l)
}

func BenchmarkCodec_Marshal_S_Parallel(b *testing.B) {
	var l int64

	b.RunParallel(func(pb *testing.PB) {
		var out []byte

		var h codec.Handle = new(codec.JsonHandle)
		enc := codec.NewEncoderBytes(&out, h)

		for pb.Next() {
			enc.ResetBytes(&out)
			if err := enc.Encode(&smallStructData); err != nil {
				b.Error(err)
			}
			l = int64(len(out))
			out = nil
		}
	})

	b.SetBytes(l)
}

func BenchmarkCodec_Marshal_M_Parallel(b *testing.B) {
	var l int64

	b.RunParallel(func(pb *testing.PB) {
		var h codec.Handle = new(codec.JsonHandle)

		var out []byte
		enc := codec.NewEncoderBytes(&out, h)

		for pb.Next() {
			enc.ResetBytes(&out)
			if err := enc.Encode(&largeStructData); err != nil {
				b.Error(err)
			}
			l = int64(len(out))
			out = nil
		}
	})
	b.SetBytes(l)
}

func BenchmarkCodec_Marshal_L_Parallel(b *testing.B) {
	var l int64

	b.RunParallel(func(pb *testing.PB) {
		var h codec.Handle = new(codec.JsonHandle)

		var out []byte
		enc := codec.NewEncoderBytes(&out, h)

		for pb.Next() {
			enc.ResetBytes(&out)
			if err := enc.Encode(&xlStructData); err != nil {
				b.Error(err)
			}
			l = int64(len(out))
			out = nil
		}
	})
	b.SetBytes(l)
}

func BenchmarkCodec_Marshal_S_Parallel_Reuse(b *testing.B) {
	var l int64

	b.RunParallel(func(pb *testing.PB) {
		var out []byte

		var h codec.Handle = new(codec.JsonHandle)
		enc := codec.NewEncoderBytes(&out, h)

		for pb.Next() {
			enc.ResetBytes(&out)
			if err := enc.Encode(&smallStructData); err != nil {
				b.Error(err)
			}
			l = int64(len(out))
			out = out[:0]
		}
	})

	b.SetBytes(l)
}

func BenchmarkCodec_Marshal_M_Parallel_Reuse(b *testing.B) {
	var l int64

	b.RunParallel(func(pb *testing.PB) {
		var h codec.Handle = new(codec.JsonHandle)

		var out []byte
		enc := codec.NewEncoderBytes(&out, h)

		for pb.Next() {
			enc.ResetBytes(&out)
			if err := enc.Encode(&largeStructData); err != nil {
				b.Error(err)
			}
			l = int64(len(out))
			out = out[:0]
		}
	})
	b.SetBytes(l)
}

func BenchmarkCodec_Marshal_L_Parallel_Reuse(b *testing.B) {
	var l int64

	b.RunParallel(func(pb *testing.PB) {
		var h codec.Handle = new(codec.JsonHandle)

		var out []byte
		enc := codec.NewEncoderBytes(&out, h)

		for pb.Next() {
			enc.ResetBytes(&out)
			if err := enc.Encode(&xlStructData); err != nil {
				b.Error(err)
			}
			l = int64(len(out))
			out = out[:0]
		}
	})
	b.SetBytes(l)
}
