// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package types

import (
	"fmt"
	"io"
	"math"
	"sort"

	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

var lengthBufMetalist = []byte{129}

func (t *Metalist) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufMetalist); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.MetaList ([]cid.Cid) (slice)
	if len(t.MetaList) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.MetaList was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.MetaList))); err != nil {
		return err
	}
	for _, v := range t.MetaList {
		if err := cbg.WriteCidBuf(scratch, w, v); err != nil {
			return xerrors.Errorf("failed writing cid field t.MetaList: %w", err)
		}
	}
	return nil
}

func (t *Metalist) UnmarshalCBOR(r io.Reader) error {
	*t = Metalist{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.MetaList ([]cid.Cid) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.MetaList: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.MetaList = make([]cid.Cid, extra)
	}

	for i := 0; i < int(extra); i++ {

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("reading cid field t.MetaList failed: %w", err)
		}
		t.MetaList[i] = c
	}

	return nil
}

var lengthBufSnapShot = []byte{132}

func (t *SnapShot) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufSnapShot); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Update (map[string]*types.Metalist) (map)
	{
		if len(t.Update) > 4096 {
			return xerrors.Errorf("cannot marshal t.Update map too large")
		}

		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajMap, uint64(len(t.Update))); err != nil {
			return err
		}

		keys := make([]string, 0, len(t.Update))
		for k := range t.Update {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			v := t.Update[k]

			if len(k) > cbg.MaxLength {
				return xerrors.Errorf("Value in field k was too long")
			}

			if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(k))); err != nil {
				return err
			}
			if _, err := io.WriteString(w, string(k)); err != nil {
				return err
			}

			if err := v.MarshalCBOR(w); err != nil {
				return err
			}

		}
	}

	// t.Height (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Height)); err != nil {
		return err
	}

	// t.CreateTime (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.CreateTime)); err != nil {
		return err
	}

	// t.PrevSnapShot (string) (string)
	if len(t.PrevSnapShot) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.PrevSnapShot was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.PrevSnapShot))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.PrevSnapShot)); err != nil {
		return err
	}
	return nil
}

func (t *SnapShot) UnmarshalCBOR(r io.Reader) error {
	*t = SnapShot{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 4 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Update (map[string]*types.Metalist) (map)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("expected a map (major type 5)")
	}
	if extra > 4096 {
		return fmt.Errorf("t.Update: map too large")
	}

	t.Update = make(map[string]*Metalist, extra)

	for i, l := 0, int(extra); i < l; i++ {

		var k string

		{
			sval, err := cbg.ReadStringBuf(br, scratch)
			if err != nil {
				return err
			}

			k = string(sval)
		}

		var v *Metalist

		{

			b, err := br.ReadByte()
			if err != nil {
				return err
			}
			if b != cbg.CborNull[0] {
				if err := br.UnreadByte(); err != nil {
					return err
				}
				v = new(Metalist)
				if err := v.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("unmarshaling v pointer: %w", err)
				}
			}

		}

		t.Update[k] = v

	}
	// t.Height (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Height = uint64(extra)

	}
	// t.CreateTime (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.CreateTime = uint64(extra)

	}
	// t.PrevSnapShot (string) (string)

	{
		sval, err := cbg.ReadStringBuf(br, scratch)
		if err != nil {
			return err
		}

		t.PrevSnapShot = string(sval)
	}
	return nil
}
