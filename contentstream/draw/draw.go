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

// Package draw has handy features for defining paths which can be used to draw content on a PDF page.  Handles
// defining paths as points, vector calculations and conversion to PDF content stream data which can be used in
// page content streams and XObject forms and thus also in annotation appearance streams.
//
// Also defines utility functions for drawing common shapes such as rectangles, lines and circles (ovals).
package draw ;import (_g "fmt";_c "github.com/unidoc/unipdf/v3/contentstream";_ba "github.com/unidoc/unipdf/v3/core";_e "github.com/unidoc/unipdf/v3/internal/transform";_d "github.com/unidoc/unipdf/v3/model";_bd "math";);

// Line defines a line shape between point 1 (X1,Y1) and point 2 (X2,Y2).  The line ending styles can be none (regular line),
// or arrows at either end.  The line also has a specified width, color and opacity.
type Line struct{X1 float64 ;Y1 float64 ;X2 float64 ;Y2 float64 ;LineColor *_d .PdfColorDeviceRGB ;Opacity float64 ;LineWidth float64 ;LineEndingStyle1 LineEndingStyle ;LineEndingStyle2 LineEndingStyle ;LineStyle LineStyle ;};

// NewPoint returns a new point with the coordinates x, y.
func NewPoint (x ,y float64 )Point {return Point {X :x ,Y :y }};

// NewPath returns a new empty path.
func NewPath ()Path {return Path {}};

// FlipY flips the sign of the Dy component of the vector.
func (_gcac Vector )FlipY ()Vector {_gcac .Dy =-_gcac .Dy ;return _gcac };

// Draw draws the polygon. A graphics state name can be specified for
// setting the polygon properties (e.g. setting the opacity). Otherwise leave
// empty (""). Returns the content stream as a byte array and the polygon
// bounding box.
func (_fga Polygon )Draw (gsName string )([]byte ,*_d .PdfRectangle ,error ){_ge :=_c .NewContentCreator ();_ge .Add_q ();_fga .FillEnabled =_fga .FillEnabled &&_fga .FillColor !=nil ;if _fga .FillEnabled {_ge .Add_rg (_fga .FillColor .R (),_fga .FillColor .G (),_fga .FillColor .B ());};_fga .BorderEnabled =_fga .BorderEnabled &&_fga .BorderColor !=nil ;if _fga .BorderEnabled {_ge .Add_RG (_fga .BorderColor .R (),_fga .BorderColor .G (),_fga .BorderColor .B ());_ge .Add_w (_fga .BorderWidth );};if len (gsName )> 1{_ge .Add_gs (_ba .PdfObjectName (gsName ));};_eec :=NewPath ();for _ ,_fge :=range _fga .Points {for _cegf ,_cf :=range _fge {_eec =_eec .AppendPoint (_cf );if _cegf ==0{_ge .Add_m (_cf .X ,_cf .Y );}else {_ge .Add_l (_cf .X ,_cf .Y );};};_ge .Add_h ();};if _fga .FillEnabled &&_fga .BorderEnabled {_ge .Add_B ();}else if _fga .FillEnabled {_ge .Add_f ();}else if _fga .BorderEnabled {_ge .Add_S ();};_ge .Add_Q ();return _ge .Bytes (),_eec .GetBoundingBox ().ToPdfRectangle (),nil ;};

// Polygon is a multi-point shape that can be drawn to a PDF content stream.
type Polygon struct{Points [][]Point ;FillEnabled bool ;FillColor *_d .PdfColorDeviceRGB ;BorderEnabled bool ;BorderColor *_d .PdfColorDeviceRGB ;BorderWidth float64 ;};

// LineStyle refers to how the line will be created.
type LineStyle int ;

// NewVector returns a new vector with the direction specified by dx and dy.
func NewVector (dx ,dy float64 )Vector {_dab :=Vector {};_dab .Dx =dx ;_dab .Dy =dy ;return _dab };

// RemovePoint removes the point at the index specified by number from the
// path. The index is 1-based.
func (_eb Path )RemovePoint (number int )Path {if number < 1||number > len (_eb .Points ){return _eb ;};_eeb :=number -1;_eb .Points =append (_eb .Points [:_eeb ],_eb .Points [_eeb +1:]...);return _eb ;};

// PolyBezierCurve represents a composite curve that is the result of
// joining multiple cubic Bezier curves.
type PolyBezierCurve struct{Curves []CubicBezierCurve ;BorderWidth float64 ;BorderColor *_d .PdfColorDeviceRGB ;FillEnabled bool ;FillColor *_d .PdfColorDeviceRGB ;};

// AppendPoint adds the specified point to the path.
func (_ed Path )AppendPoint (point Point )Path {_ed .Points =append (_ed .Points ,point );return _ed };

// Offset shifts the Bezier path with the specified offsets.
func (_ecf CubicBezierPath )Offset (offX ,offY float64 )CubicBezierPath {for _dg ,_ca :=range _ecf .Curves {_ecf .Curves [_dg ]=_ca .AddOffsetXY (offX ,offY );};return _ecf ;};func (_gab Point )String ()string {return _g .Sprintf ("(\u0025\u002e\u0031\u0066\u002c\u0025\u002e\u0031\u0066\u0029",_gab .X ,_gab .Y );};

// NewCubicBezierPath returns a new empty cubic Bezier path.
func NewCubicBezierPath ()CubicBezierPath {_bg :=CubicBezierPath {};_bg .Curves =[]CubicBezierCurve {};return _bg ;};

// Draw draws the line to PDF contentstream. Generates the content stream which can be used in page contents or
// appearance stream of annotation. Returns the stream content, XForm bounding box (local), bounding box and an error
// if one occurred.
func (_cgc Line )Draw (gsName string )([]byte ,*_d .PdfRectangle ,error ){_abe ,_aae :=_cgc .X1 ,_cgc .X2 ;_ecc ,_ad :=_cgc .Y1 ,_cgc .Y2 ;_fd :=_ad -_ecc ;_cce :=_aae -_abe ;_cee :=_bd .Atan2 (_fd ,_cce );L :=_bd .Sqrt (_bd .Pow (_cce ,2.0)+_bd .Pow (_fd ,2.0));_abec :=_cgc .LineWidth ;_fc :=_bd .Pi ;_cfe :=1.0;if _cce < 0{_cfe *=-1.0;};if _fd < 0{_cfe *=-1.0;};VsX :=_cfe *(-_abec /2*_bd .Cos (_cee +_fc /2));VsY :=_cfe *(-_abec /2*_bd .Sin (_cee +_fc /2)+_abec *_bd .Sin (_cee +_fc /2));V1X :=VsX +_abec /2*_bd .Cos (_cee +_fc /2);V1Y :=VsY +_abec /2*_bd .Sin (_cee +_fc /2);V2X :=VsX +_abec /2*_bd .Cos (_cee +_fc /2)+L *_bd .Cos (_cee );V2Y :=VsY +_abec /2*_bd .Sin (_cee +_fc /2)+L *_bd .Sin (_cee );V3X :=VsX +_abec /2*_bd .Cos (_cee +_fc /2)+L *_bd .Cos (_cee )+_abec *_bd .Cos (_cee -_fc /2);V3Y :=VsY +_abec /2*_bd .Sin (_cee +_fc /2)+L *_bd .Sin (_cee )+_abec *_bd .Sin (_cee -_fc /2);V4X :=VsX +_abec /2*_bd .Cos (_cee -_fc /2);V4Y :=VsY +_abec /2*_bd .Sin (_cee -_fc /2);_eee :=NewPath ();_eee =_eee .AppendPoint (NewPoint (V1X ,V1Y ));_eee =_eee .AppendPoint (NewPoint (V2X ,V2Y ));_eee =_eee .AppendPoint (NewPoint (V3X ,V3Y ));_eee =_eee .AppendPoint (NewPoint (V4X ,V4Y ));_cdg :=_cgc .LineEndingStyle1 ;_dba :=_cgc .LineEndingStyle2 ;_df :=3*_abec ;_edb :=3*_abec ;_cfc :=(_edb -_abec )/2;if _dba ==LineEndingStyleArrow {_ebe :=_eee .GetPointNumber (2);_cca :=NewVectorPolar (_df ,_cee +_fc );_fcc :=_ebe .AddVector (_cca );_ace :=NewVectorPolar (_edb /2,_cee +_fc /2);_dea :=NewVectorPolar (_df ,_cee );_gga :=NewVectorPolar (_cfc ,_cee +_fc /2);_ag :=_fcc .AddVector (_gga );_daf :=_dea .Add (_ace .Flip ());_dbe :=_ag .AddVector (_daf );_fdg :=_ace .Scale (2).Flip ().Add (_daf .Flip ());_dbd :=_dbe .AddVector (_fdg );_ccae :=_fcc .AddVector (NewVectorPolar (_abec ,_cee -_fc /2));_ebc :=NewPath ();_ebc =_ebc .AppendPoint (_eee .GetPointNumber (1));_ebc =_ebc .AppendPoint (_fcc );_ebc =_ebc .AppendPoint (_ag );_ebc =_ebc .AppendPoint (_dbe );_ebc =_ebc .AppendPoint (_dbd );_ebc =_ebc .AppendPoint (_ccae );_ebc =_ebc .AppendPoint (_eee .GetPointNumber (4));_eee =_ebc ;};if _cdg ==LineEndingStyleArrow {_dbg :=_eee .GetPointNumber (1);_aaea :=_eee .GetPointNumber (_eee .Length ());_gaa :=NewVectorPolar (_abec /2,_cee +_fc +_fc /2);_bac :=_dbg .AddVector (_gaa );_feaa :=NewVectorPolar (_df ,_cee ).Add (NewVectorPolar (_edb /2,_cee +_fc /2));_aga :=_bac .AddVector (_feaa );_cfed :=NewVectorPolar (_cfc ,_cee -_fc /2);_ceb :=_aga .AddVector (_cfed );_fdb :=NewVectorPolar (_df ,_cee );_adc :=_aaea .AddVector (_fdb );_gafd :=NewVectorPolar (_cfc ,_cee +_fc +_fc /2);_be :=_adc .AddVector (_gafd );_ccg :=_bac ;_ggg :=NewPath ();_ggg =_ggg .AppendPoint (_bac );_ggg =_ggg .AppendPoint (_aga );_ggg =_ggg .AppendPoint (_ceb );for _ ,_cfd :=range _eee .Points [1:len (_eee .Points )-1]{_ggg =_ggg .AppendPoint (_cfd );};_ggg =_ggg .AppendPoint (_adc );_ggg =_ggg .AppendPoint (_be );_ggg =_ggg .AppendPoint (_ccg );_eee =_ggg ;};_eff :=_c .NewContentCreator ();_eff .Add_q ().Add_rg (_cgc .LineColor .R (),_cgc .LineColor .G (),_cgc .LineColor .B ());if len (gsName )> 1{_eff .Add_gs (_ba .PdfObjectName (gsName ));};_eee =_eee .Offset (_cgc .X1 ,_cgc .Y1 );_ade :=_eee .GetBoundingBox ();DrawPathWithCreator (_eee ,_eff );if _cgc .LineStyle ==LineStyleDashed {_eff .Add_d ([]int64 {1,1},0).Add_S ().Add_f ().Add_Q ();}else {_eff .Add_f ().Add_Q ();};return _eff .Bytes (),_ade .ToPdfRectangle (),nil ;};

// AddVector adds vector to a point.
func (_aff Point )AddVector (v Vector )Point {_aff .X +=v .Dx ;_aff .Y +=v .Dy ;return _aff };const (LineStyleSolid LineStyle =0;LineStyleDashed LineStyle =1;);

// ToPdfRectangle returns the bounding box as a PDF rectangle.
func (_da BoundingBox )ToPdfRectangle ()*_d .PdfRectangle {return &_d .PdfRectangle {Llx :_da .X ,Lly :_da .Y ,Urx :_da .X +_da .Width ,Ury :_da .Y +_da .Height };};

// NewVectorBetween returns a new vector with the direction specified by
// the subtraction of point a from point b (b-a).
func NewVectorBetween (a Point ,b Point )Vector {_efd :=Vector {};_efd .Dx =b .X -a .X ;_efd .Dy =b .Y -a .Y ;return _efd ;};

// LineEndingStyle defines the line ending style for lines.
// The currently supported line ending styles are None, Arrow (ClosedArrow) and Butt.
type LineEndingStyle int ;

// Scale scales the vector by the specified factor.
func (_gce Vector )Scale (factor float64 )Vector {_gfa :=_gce .Magnitude ();_bae :=_gce .GetPolarAngle ();_gce .Dx =factor *_gfa *_bd .Cos (_bae );_gce .Dy =factor *_gfa *_bd .Sin (_bae );return _gce ;};

// AddOffsetXY adds X,Y offset to all points on a curve.
func (_ec CubicBezierCurve )AddOffsetXY (offX ,offY float64 )CubicBezierCurve {_ec .P0 .X +=offX ;_ec .P1 .X +=offX ;_ec .P2 .X +=offX ;_ec .P3 .X +=offX ;_ec .P0 .Y +=offY ;_ec .P1 .Y +=offY ;_ec .P2 .Y +=offY ;_ec .P3 .Y +=offY ;return _ec ;};

// BoundingBox represents the smallest rectangular area that encapsulates an object.
type BoundingBox struct{X float64 ;Y float64 ;Width float64 ;Height float64 ;};

// FlipX flips the sign of the Dx component of the vector.
func (_ea Vector )FlipX ()Vector {_ea .Dx =-_ea .Dx ;return _ea };

// Draw draws the basic line to PDF. Generates the content stream which can be used in page contents or appearance
// stream of annotation. Returns the stream content, XForm bounding box (local), bounding box and an error if
// one occurred.
func (_gffe BasicLine )Draw (gsName string )([]byte ,*_d .PdfRectangle ,error ){_dfd :=_gffe .LineWidth ;_bfa :=NewPath ();_bfa =_bfa .AppendPoint (NewPoint (_gffe .X1 ,_gffe .Y1 ));_bfa =_bfa .AppendPoint (NewPoint (_gffe .X2 ,_gffe .Y2 ));_beg :=_c .NewContentCreator ();_bdag :=_bfa .GetBoundingBox ();DrawPathWithCreator (_bfa ,_beg );if _gffe .LineStyle ==LineStyleDashed {_beg .Add_d ([]int64 {1,1},0);};_beg .Add_RG (_gffe .LineColor .R (),_gffe .LineColor .G (),_gffe .LineColor .B ()).Add_w (_dfd ).Add_S ().Add_Q ();return _beg .Bytes (),_bdag .ToPdfRectangle (),nil ;};

// ToPdfRectangle returns the rectangle as a PDF rectangle.
func (_cbe Rectangle )ToPdfRectangle ()*_d .PdfRectangle {return &_d .PdfRectangle {Llx :_cbe .X ,Lly :_cbe .Y ,Urx :_cbe .X +_cbe .Width ,Ury :_cbe .Y +_cbe .Height };};

// AppendCurve appends the specified Bezier curve to the path.
func (_ef CubicBezierPath )AppendCurve (curve CubicBezierCurve )CubicBezierPath {_ef .Curves =append (_ef .Curves ,curve );return _ef ;};

// Copy returns a clone of the path.
func (_eg Path )Copy ()Path {_fg :=Path {};_fg .Points =[]Point {};for _ ,_ecd :=range _eg .Points {_fg .Points =append (_fg .Points ,_ecd );};return _fg ;};

// Draw draws the rectangle. Can specify a graphics state (gsName) for setting opacity etc.
// Otherwise leave empty (""). Returns the content stream as a byte array, bounding box and an error on failure.
func (_gfe Rectangle )Draw (gsName string )([]byte ,*_d .PdfRectangle ,error ){_gaf :=NewPath ();_gaf =_gaf .AppendPoint (NewPoint (0,0));_gaf =_gaf .AppendPoint (NewPoint (0,_gfe .Height ));_gaf =_gaf .AppendPoint (NewPoint (_gfe .Width ,_gfe .Height ));_gaf =_gaf .AppendPoint (NewPoint (_gfe .Width ,0));_gaf =_gaf .AppendPoint (NewPoint (0,0));if _gfe .X !=0||_gfe .Y !=0{_gaf =_gaf .Offset (_gfe .X ,_gfe .Y );};_fea :=_c .NewContentCreator ();_fea .Add_q ();if _gfe .FillEnabled {_fea .Add_rg (_gfe .FillColor .R (),_gfe .FillColor .G (),_gfe .FillColor .B ());};if _gfe .BorderEnabled {_fea .Add_RG (_gfe .BorderColor .R (),_gfe .BorderColor .G (),_gfe .BorderColor .B ());_fea .Add_w (_gfe .BorderWidth );};if len (gsName )> 1{_fea .Add_gs (_ba .PdfObjectName (gsName ));};DrawPathWithCreator (_gaf ,_fea );_fea .Add_h ();if _gfe .FillEnabled &&_gfe .BorderEnabled {_fea .Add_B ();}else if _gfe .FillEnabled {_fea .Add_f ();}else if _gfe .BorderEnabled {_fea .Add_S ();};_fea .Add_Q ();return _fea .Bytes (),_gaf .GetBoundingBox ().ToPdfRectangle (),nil ;};

// GetPointNumber returns the path point at the index specified by number.
// The index is 1-based.
func (_bb Path )GetPointNumber (number int )Point {if number < 1||number > len (_bb .Points ){return Point {};};return _bb .Points [number -1];};

// Length returns the number of points in the path.
func (_afe Path )Length ()int {return len (_afe .Points )};

// Rectangle is a shape with a specified Width and Height and a lower left corner at (X,Y) that can be
// drawn to a PDF content stream.  The rectangle can optionally have a border and a filling color.
// The Width/Height includes the border (if any specified), i.e. is positioned inside.
type Rectangle struct{X float64 ;Y float64 ;Width float64 ;Height float64 ;FillEnabled bool ;FillColor *_d .PdfColorDeviceRGB ;BorderEnabled bool ;BorderWidth float64 ;BorderColor *_d .PdfColorDeviceRGB ;Opacity float64 ;};

// GetBoundingBox returns the bounding box of the Bezier path.
func (_ab CubicBezierPath )GetBoundingBox ()Rectangle {_gad :=Rectangle {};_dd :=0.0;_efe :=0.0;_fed :=0.0;_fb :=0.0;for _cac ,_cb :=range _ab .Curves {_af :=_cb .GetBounds ();if _cac ==0{_dd =_af .Llx ;_efe =_af .Urx ;_fed =_af .Lly ;_fb =_af .Ury ;continue ;};if _af .Llx < _dd {_dd =_af .Llx ;};if _af .Urx > _efe {_efe =_af .Urx ;};if _af .Lly < _fed {_fed =_af .Lly ;};if _af .Ury > _fb {_fb =_af .Ury ;};};_gad .X =_dd ;_gad .Y =_fed ;_gad .Width =_efe -_dd ;_gad .Height =_fb -_fed ;return _gad ;};

// Rotate rotates the vector by the specified angle.
func (_gcc Vector )Rotate (phi float64 )Vector {_afea :=_gcc .Magnitude ();_eebd :=_gcc .GetPolarAngle ();return NewVectorPolar (_afea ,_eebd +phi );};

// Polyline defines a slice of points that are connected as straight lines.
type Polyline struct{Points []Point ;LineColor *_d .PdfColorDeviceRGB ;LineWidth float64 ;};

// Point represents a two-dimensional point.
type Point struct{X float64 ;Y float64 ;};

// GetBoundingBox returns the bounding box of the path.
func (_bc Path )GetBoundingBox ()BoundingBox {_cacb :=BoundingBox {};_ege :=0.0;_gc :=0.0;_gf :=0.0;_gca :=0.0;for _dc ,_gdg :=range _bc .Points {if _dc ==0{_ege =_gdg .X ;_gc =_gdg .X ;_gf =_gdg .Y ;_gca =_gdg .Y ;continue ;};if _gdg .X < _ege {_ege =_gdg .X ;};if _gdg .X > _gc {_gc =_gdg .X ;};if _gdg .Y < _gf {_gf =_gdg .Y ;};if _gdg .Y > _gca {_gca =_gdg .Y ;};};_cacb .X =_ege ;_cacb .Y =_gf ;_cacb .Width =_gc -_ege ;_cacb .Height =_gca -_gf ;return _cacb ;};

// Add adds the specified vector to the current one and returns the result.
func (_bdagf Vector )Add (other Vector )Vector {_bdagf .Dx +=other .Dx ;_bdagf .Dy +=other .Dy ;return _bdagf ;};

// DrawBezierPathWithCreator makes the bezier path with the content creator.
// Adds the PDF commands to draw the path to the creator instance.
func DrawBezierPathWithCreator (bpath CubicBezierPath ,creator *_c .ContentCreator ){for _ffd ,_gcd :=range bpath .Curves {if _ffd ==0{creator .Add_m (_gcd .P0 .X ,_gcd .P0 .Y );};creator .Add_c (_gcd .P1 .X ,_gcd .P1 .Y ,_gcd .P2 .X ,_gcd .P2 .Y ,_gcd .P3 .X ,_gcd .P3 .Y );};};

// Copy returns a clone of the Bezier path.
func (_ga CubicBezierPath )Copy ()CubicBezierPath {_bf :=CubicBezierPath {};_bf .Curves =[]CubicBezierCurve {};for _ ,_cge :=range _ga .Curves {_bf .Curves =append (_bf .Curves ,_cge );};return _bf ;};

// NewVectorPolar returns a new vector calculated from the specified
// magnitude and angle.
func NewVectorPolar (length float64 ,theta float64 )Vector {_ccaa :=Vector {};_ccaa .Dx =length *_bd .Cos (theta );_ccaa .Dy =length *_bd .Sin (theta );return _ccaa ;};const (LineEndingStyleNone LineEndingStyle =0;LineEndingStyleArrow LineEndingStyle =1;LineEndingStyleButt LineEndingStyle =2;);

// GetBounds returns the bounding box of the Bezier curve.
func (_a CubicBezierCurve )GetBounds ()_d .PdfRectangle {_de :=_a .P0 .X ;_ee :=_a .P0 .X ;_cg :=_a .P0 .Y ;_bda :=_a .P0 .Y ;for _cc :=0.0;_cc <=1.0;_cc +=0.001{Rx :=_a .P0 .X *_bd .Pow (1-_cc ,3)+_a .P1 .X *3*_cc *_bd .Pow (1-_cc ,2)+_a .P2 .X *3*_bd .Pow (_cc ,2)*(1-_cc )+_a .P3 .X *_bd .Pow (_cc ,3);Ry :=_a .P0 .Y *_bd .Pow (1-_cc ,3)+_a .P1 .Y *3*_cc *_bd .Pow (1-_cc ,2)+_a .P2 .Y *3*_bd .Pow (_cc ,2)*(1-_cc )+_a .P3 .Y *_bd .Pow (_cc ,3);if Rx < _de {_de =Rx ;};if Rx > _ee {_ee =Rx ;};if Ry < _cg {_cg =Ry ;};if Ry > _bda {_bda =Ry ;};};_fe :=_d .PdfRectangle {};_fe .Llx =_de ;_fe .Lly =_cg ;_fe .Urx =_ee ;_fe .Ury =_bda ;return _fe ;};

// Offset shifts the path with the specified offsets.
func (_ac Path )Offset (offX ,offY float64 )Path {for _gd ,_bgg :=range _ac .Points {_ac .Points [_gd ]=_bgg .Add (offX ,offY );};return _ac ;};

// Add shifts the coordinates of the point with dx, dy and returns the result.
func (_ceg Point )Add (dx ,dy float64 )Point {_ceg .X +=dx ;_ceg .Y +=dy ;return _ceg };

// DrawPathWithCreator makes the path with the content creator.
// Adds the PDF commands to draw the path to the creator instance.
func DrawPathWithCreator (path Path ,creator *_c .ContentCreator ){for _fa ,_baa :=range path .Points {if _fa ==0{creator .Add_m (_baa .X ,_baa .Y );}else {creator .Add_l (_baa .X ,_baa .Y );};};};

// CubicBezierCurve is defined by:
// R(t) = P0*(1-t)^3 + P1*3*t*(1-t)^2 + P2*3*t^2*(1-t) + P3*t^3
// where P0 is the current point, P1, P2 control points and P3 the final point.
type CubicBezierCurve struct{P0 Point ;P1 Point ;P2 Point ;P3 Point ;};

// Magnitude returns the magnitude of the vector.
func (_fee Vector )Magnitude ()float64 {return _bd .Sqrt (_bd .Pow (_fee .Dx ,2.0)+_bd .Pow (_fee .Dy ,2.0));};

// Draw draws the polyline. A graphics state name can be specified for
// setting the polyline properties (e.g. setting the opacity). Otherwise leave
// empty (""). Returns the content stream as a byte array and the polyline
// bounding box.
func (_ecfg Polyline )Draw (gsName string )([]byte ,*_d .PdfRectangle ,error ){if _ecfg .LineColor ==nil {_ecfg .LineColor =_d .NewPdfColorDeviceRGB (0,0,0);};_dggc :=NewPath ();for _ ,_fgd :=range _ecfg .Points {_dggc =_dggc .AppendPoint (_fgd );};_add :=_c .NewContentCreator ();_add .Add_q ();_add .Add_RG (_ecfg .LineColor .R (),_ecfg .LineColor .G (),_ecfg .LineColor .B ());_add .Add_w (_ecfg .LineWidth );if len (gsName )> 1{_add .Add_gs (_ba .PdfObjectName (gsName ));};DrawPathWithCreator (_dggc ,_add );_add .Add_S ();_add .Add_Q ();return _add .Bytes (),_dggc .GetBoundingBox ().ToPdfRectangle (),nil ;};

// Vector represents a two-dimensional vector.
type Vector struct{Dx float64 ;Dy float64 ;};

// CubicBezierPath represents a collection of cubic Bezier curves.
type CubicBezierPath struct{Curves []CubicBezierCurve ;};

// GetPolarAngle returns the angle the magnitude of the vector forms with the
// positive X-axis going counterclockwise.
func (_dgeb Vector )GetPolarAngle ()float64 {return _bd .Atan2 (_dgeb .Dy ,_dgeb .Dx )};

// Draw draws the composite Bezier curve. A graphics state name can be
// specified for setting the curve properties (e.g. setting the opacity).
// Otherwise leave empty (""). Returns the content stream as a byte array and
// the curve bounding box.
func (_gff PolyBezierCurve )Draw (gsName string )([]byte ,*_d .PdfRectangle ,error ){if _gff .BorderColor ==nil {_gff .BorderColor =_d .NewPdfColorDeviceRGB (0,0,0);};_ede :=NewCubicBezierPath ();for _ ,_cgd :=range _gff .Curves {_ede =_ede .AppendCurve (_cgd );};_fbd :=_c .NewContentCreator ();_fbd .Add_q ();_gff .FillEnabled =_gff .FillEnabled &&_gff .FillColor !=nil ;if _gff .FillEnabled {_fbd .Add_rg (_gff .FillColor .R (),_gff .FillColor .G (),_gff .FillColor .B ());};_fbd .Add_RG (_gff .BorderColor .R (),_gff .BorderColor .G (),_gff .BorderColor .B ());_fbd .Add_w (_gff .BorderWidth );if len (gsName )> 1{_fbd .Add_gs (_ba .PdfObjectName (gsName ));};for _ ,_ffb :=range _ede .Curves {_fbd .Add_m (_ffb .P0 .X ,_ffb .P0 .Y );_fbd .Add_c (_ffb .P1 .X ,_ffb .P1 .Y ,_ffb .P2 .X ,_ffb .P2 .Y ,_ffb .P3 .X ,_ffb .P3 .Y );};if _gff .FillEnabled {_fbd .Add_h ();_fbd .Add_B ();}else {_fbd .Add_S ();};_fbd .Add_Q ();return _fbd .Bytes (),_ede .GetBoundingBox ().ToPdfRectangle (),nil ;};

// Draw draws the circle. Can specify a graphics state (gsName) for setting opacity etc.  Otherwise leave empty ("").
// Returns the content stream as a byte array, the bounding box and an error on failure.
func (_fba Circle )Draw (gsName string )([]byte ,*_d .PdfRectangle ,error ){_cbd :=_fba .Width /2;_cd :=_fba .Height /2;if _fba .BorderEnabled {_cbd -=_fba .BorderWidth /2;_cd -=_fba .BorderWidth /2;};_bde :=0.551784;_bag :=_cbd *_bde ;_ddc :=_cd *_bde ;_aca :=NewCubicBezierPath ();_aca =_aca .AppendCurve (NewCubicBezierCurve (-_cbd ,0,-_cbd ,_ddc ,-_bag ,_cd ,0,_cd ));_aca =_aca .AppendCurve (NewCubicBezierCurve (0,_cd ,_bag ,_cd ,_cbd ,_ddc ,_cbd ,0));_aca =_aca .AppendCurve (NewCubicBezierCurve (_cbd ,0,_cbd ,-_ddc ,_bag ,-_cd ,0,-_cd ));_aca =_aca .AppendCurve (NewCubicBezierCurve (0,-_cd ,-_bag ,-_cd ,-_cbd ,-_ddc ,-_cbd ,0));_aca =_aca .Offset (_cbd ,_cd );if _fba .BorderEnabled {_aca =_aca .Offset (_fba .BorderWidth /2,_fba .BorderWidth /2);};if _fba .X !=0||_fba .Y !=0{_aca =_aca .Offset (_fba .X ,_fba .Y );};_gg :=_c .NewContentCreator ();_gg .Add_q ();if _fba .FillEnabled {_gg .Add_rg (_fba .FillColor .R (),_fba .FillColor .G (),_fba .FillColor .B ());};if _fba .BorderEnabled {_gg .Add_RG (_fba .BorderColor .R (),_fba .BorderColor .G (),_fba .BorderColor .B ());_gg .Add_w (_fba .BorderWidth );};if len (gsName )> 1{_gg .Add_gs (_ba .PdfObjectName (gsName ));};DrawBezierPathWithCreator (_aca ,_gg );_gg .Add_h ();if _fba .FillEnabled &&_fba .BorderEnabled {_gg .Add_B ();}else if _fba .FillEnabled {_gg .Add_f ();}else if _fba .BorderEnabled {_gg .Add_S ();};_gg .Add_Q ();_dge :=_aca .GetBoundingBox ();if _fba .BorderEnabled {_dge .Height +=_fba .BorderWidth ;_dge .Width +=_fba .BorderWidth ;_dge .X -=_fba .BorderWidth /2;_dge .Y -=_fba .BorderWidth /2;};return _gg .Bytes (),_dge .ToPdfRectangle (),nil ;};

// Path consists of straight line connections between each point defined in an array of points.
type Path struct{Points []Point ;};

// NewCubicBezierCurve returns a new cubic Bezier curve.
func NewCubicBezierCurve (x0 ,y0 ,x1 ,y1 ,x2 ,y2 ,x3 ,y3 float64 )CubicBezierCurve {_f :=CubicBezierCurve {};_f .P0 =NewPoint (x0 ,y0 );_f .P1 =NewPoint (x1 ,y1 );_f .P2 =NewPoint (x2 ,y2 );_f .P3 =NewPoint (x3 ,y3 );return _f ;};

// BasicLine defines a line between point 1 (X1,Y1) and point 2 (X2,Y2). The line has a specified width, color and opacity.
type BasicLine struct{X1 float64 ;Y1 float64 ;X2 float64 ;Y2 float64 ;LineColor *_d .PdfColorDeviceRGB ;Opacity float64 ;LineWidth float64 ;LineStyle LineStyle ;};

// Flip changes the sign of the vector: -vector.
func (_bdeg Vector )Flip ()Vector {_gcae :=_bdeg .Magnitude ();_gb :=_bdeg .GetPolarAngle ();_bdeg .Dx =_gcae *_bd .Cos (_gb +_bd .Pi );_bdeg .Dy =_gcae *_bd .Sin (_gb +_bd .Pi );return _bdeg ;};

// Circle represents a circle shape with fill and border properties that can be drawn to a PDF content stream.
type Circle struct{X float64 ;Y float64 ;Width float64 ;Height float64 ;FillEnabled bool ;FillColor *_d .PdfColorDeviceRGB ;BorderEnabled bool ;BorderWidth float64 ;BorderColor *_d .PdfColorDeviceRGB ;Opacity float64 ;};

// Rotate returns a new Point at `p` rotated by `theta` degrees.
func (_dgg Point )Rotate (theta float64 )Point {_bcb :=_e .NewPoint (_dgg .X ,_dgg .Y ).Rotate (theta );return NewPoint (_bcb .X ,_bcb .Y );};