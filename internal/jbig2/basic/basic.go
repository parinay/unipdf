//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package basic ;import _c "github.com/unidoc/unipdf/v3/internal/jbig2/errors";func Ceil (numerator ,denominator int )int {if numerator %denominator ==0{return numerator /denominator ;};return (numerator /denominator )+1;};func (_fbd *Stack )Pop ()(_bg interface{},_bag bool ){_bg ,_bag =_fbd .peek ();if !_bag {return nil ,_bag ;};_fbd .Data =_fbd .Data [:_fbd .top ()];return _bg ,true ;};func (_be *Stack )Len ()int {return len (_be .Data )};func NewNumSlice (i int )*NumSlice {_efc :=NumSlice (make ([]float32 ,i ));return &_efc };func (_bc *IntSlice )Add (v int )error {if _bc ==nil {return _c .Error ("\u0049\u006e\u0074S\u006c\u0069\u0063\u0065\u002e\u0041\u0064\u0064","\u0073\u006c\u0069\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064");};*_bc =append (*_bc ,v );return nil ;};func (_fb IntsMap )Delete (key uint64 ){delete (_fb ,key )};type IntSlice []int ;func (_b IntsMap )GetSlice (key uint64 )([]int ,bool ){_ac ,_g :=_b [key ];if !_g {return nil ,false ;};return _ac ,true ;};func Max (x ,y int )int {if x > y {return x ;};return y ;};func Sign (v float32 )float32 {if v >=0.0{return 1.0;};return -1.0;};func (_d IntSlice )Get (index int )(int ,error ){if index > len (_d )-1{return 0,_c .Errorf ("\u0049\u006e\u0074S\u006c\u0069\u0063\u0065\u002e\u0047\u0065\u0074","\u0069\u006e\u0064\u0065x:\u0020\u0025\u0064\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006eg\u0065",index );};return _d [index ],nil ;};type Stack struct{Data []interface{};Aux *Stack ;};func (_bb *NumSlice )AddInt (v int ){*_bb =append (*_bb ,float32 (v ))};func Abs (v int )int {if v > 0{return v ;};return -v ;};func (_a IntsMap )Get (key uint64 )(int ,bool ){_ad ,_f :=_a [key ];if !_f {return 0,false ;};if len (_ad )==0{return 0,false ;};return _ad [0],true ;};type NumSlice []float32 ;func (_adb *Stack )top ()int {return len (_adb .Data )-1};func (_bcc NumSlice )GetInt (i int )(int ,error ){const _ag ="\u0047\u0065\u0074\u0049\u006e\u0074";if i < 0||i > len (_bcc )-1{return 0,_c .Errorf (_ag ,"\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065",i );};_cc :=_bcc [i ];return int (_cc +Sign (_cc )*0.5),nil ;};func Min (x ,y int )int {if x < y {return x ;};return y ;};func (_fg *Stack )Peek ()(_bd interface{},_da bool ){return _fg .peek ()};func (_ga NumSlice )Get (i int )(float32 ,error ){if i < 0||i > len (_ga )-1{return 0,_c .Errorf ("\u004e\u0075\u006dS\u006c\u0069\u0063\u0065\u002e\u0047\u0065\u0074","\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065",i );};return _ga [i ],nil ;};func (_fa *Stack )Push (v interface{}){_fa .Data =append (_fa .Data ,v )};func NewIntSlice (i int )*IntSlice {_efb :=IntSlice (make ([]int ,i ));return &_efb };func (_gf *IntSlice )Copy ()*IntSlice {_eb :=IntSlice (make ([]int ,len (*_gf )));copy (_eb ,*_gf );return &_eb ;};func (_faf *Stack )peek ()(interface{},bool ){_fgc :=_faf .top ();if _fgc ==-1{return nil ,false ;};return _faf .Data [_fgc ],true ;};func (_agf NumSlice )GetIntSlice ()[]int {_ca :=make ([]int ,len (_agf ));for _dc ,_ba :=range _agf {_ca [_dc ]=int (_ba );};return _ca ;};type IntsMap map[uint64 ][]int ;func (_gc IntSlice )Size ()int {return len (_gc )};func (_ef IntsMap )Add (key uint64 ,value int ){_ef [key ]=append (_ef [key ],value )};func (_aa *NumSlice )Add (v float32 ){*_aa =append (*_aa ,v )};