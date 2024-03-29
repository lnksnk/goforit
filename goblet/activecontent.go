package goblet

import (
	"fmt"
	"io"
	"regexp"
	"strings"

	goja "github.com/dop251/goja"

	"github.com/efjoubert/goforit/goio"
)

func serverActiveContent(r io.Reader, active bool, a ...interface{}) (atvr io.Reader, err error) {
	if active {
		var pr, pw = io.Pipe()
		var pwio, _ = goio.NewIORW(pw)
		var unpio, _ = goio.NewIORW()
		var atvElems = map[string]interface{}{}
		var orgrio, _ = goio.NewIORW(r)
		for len(a) > 0 && len(a)%2 == 0 {
			if s, sok := a[0].(string); sok && a[1] != nil {
				atvElems[s] = a[1]
				a = a[2:]
			} else {
				break
			}
		}
		var atvrf = &ActiveReader{orgr: r, orgrio: orgrio, pr: pr, pw: pw, pwio: pwio, isatv: active, prvb: []byte{0, 0}, lblbytes: [][]byte{([]byte)("<@"), ([]byte)("@>"), ([]byte)("<"), ([]byte)(">")}, lbli: []int{0, 0, 0, 0}, unpio: unpio, cntntstarti: -1, atvElems: atvElems, atvinterupt: make(chan bool, 1), atvpsvmde: mdeatv}
		go func() {
			err = readActiveContent(atvrf)
		}()
		pr = nil
		pw = nil
		pwio = nil
		atvr = atvrf
		unpio = nil
	} else {
		atvr = r
	}
	return
}

type modeactivepassive int

const (
	mdeatv modeactivepassive = iota
	mdepsv
	mdepsvsngl
	mdepsvcmplx
)

//ActiveReader ActiveReader
type ActiveReader struct {
	pr          *io.PipeReader
	pw          *io.PipeWriter
	pwio        *goio.IORW
	orgr        io.Reader
	orgrio      *goio.IORW
	isatv       bool
	prvb        []byte
	lblbytes    [][]byte
	lbli        []int
	unpio       *goio.IORW
	unvlio      *goio.IORW
	cdeio       *goio.IORW
	cntntstarti int64
	cntntio     *goio.IORW
	foundcode   bool
	hasCode     bool
	hasContent  bool
	atvElems    map[string]interface{}
	atvinterupt chan bool
	atvpsvmde   modeactivepassive
}

func (atvr *ActiveReader) interuptReader() {
	if atvr.atvinterupt != nil {
		atvr.atvinterupt <- true
	}
}

//Read refer to io.Reader Read
func (atvr *ActiveReader) Read(p []byte) (n int, err error) {
	if atvr.isatv {
		return atvr.pr.Read(p)
	}
	return atvr.orgr.Read(p)
}

//Close close ActiveReader
func (atvr *ActiveReader) Close() (err error) {
	if atvr.orgr != nil {
		if orgrclose, orgrcloseok := atvr.orgr.(io.ReadCloser); orgrcloseok {
			orgrclose.Close()
			orgrclose = nil
		}
		atvr.orgr = nil
	}
	if atvr.isatv {
		if atvr.orgrio != nil {
			atvr.orgrio.Close()
			atvr.orgrio = nil
		}
		if atvr.pr != nil {
			atvr.pr.Close()
		}
		if atvr.pw != nil {
			atvr.pw.Close()
			atvr.pw = nil
		}
		if atvr.pwio != nil {
			atvr.pwio.Close()
			atvr.pwio = nil
		}
		if atvr.lblbytes != nil {
			atvr.lblbytes = nil
		}
		if atvr.lbli != nil {
			atvr.lbli = nil
		}
		if atvr.unpio != nil {
			atvr.unpio.Close()
			atvr.unpio = nil
		}
		if atvr.unvlio != nil {
			atvr.unvlio.Close()
			atvr.unvlio = nil
		}
		if atvr.cdeio != nil {
			atvr.cdeio.Close()
			atvr.cdeio = nil
		}
		if atvr.cntntio != nil {
			atvr.cntntio.Close()
			atvr.cntntio = nil
		}
		if atvr.atvElems != nil {
			if len(atvr.atvElems) > 0 {
				var keys = make([]string, len(atvr.atvElems), len(atvr.atvElems))
				var keysi = 0
				for k := range atvr.atvElems {
					keys[keysi] = k
					keysi++
				}

				for _, k := range keys {
					atvr.atvElems[k] = nil
					delete(atvr.atvElems, k)
				}
				keys = nil
			}
			atvr.atvElems = nil
		}
		if atvr.atvinterupt != nil {
			close(atvr.atvinterupt)
			atvr.atvinterupt = nil
		}
	}
	return
}

func readNextByte(atvr *ActiveReader, p []byte) (n int, err error) {
	switch atvr.atvpsvmde {
	case mdeatv:
		n, err = atvr.orgrio.Read(p)
	}
	return
}

func readActiveContent(atvr *ActiveReader) (err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}
		}
	}()
	var p = make([]byte, 1)
	defer atvr.pw.Close()
	var interupted = false
	for {
		select {
		case intrptd := <-atvr.atvinterupt:
			if intrptd {
				interupted = intrptd
			}
		default:
		}
		if interupted {
			break
		}
		var n, err = atvr.orgr.Read(p)
		if n > 0 {
			for bn := range p[0:n] {
				parseAtiveReaderByte(atvr, atvr.pwio, p[0:n][bn:bn+1], atvr.lblbytes, atvr.lbli)
			}
		}
		if (err == io.EOF || err != nil) || n == 0 {
			break
		}
	}
	if !interupted {
		if atvr.foundcode {
			flushContentFound(atvr)

			if err = evalCode(atvr); err != nil {
				println(err.Error())
			}

		} else {
			if atvr.unpio.Size() > 0 {
				atvr.pwio.Print(atvr.unpio)
				atvr.unpio.Close()
			}
		}
	}
	return
}

func evalCode(atvr *ActiveReader) (err error) {
	var pgrm, pgrmerr = goja.Compile("", atvr.cdeio.String(), false)
	if pgrmerr == nil {
		var vm = goja.New()
		vm.Set("_atvr", atvr)
		if len(atvr.atvElems) > 0 {
			for elemnm := range atvr.atvElems {
				vm.Set(elemnm, atvr.atvElems[elemnm])
			}
		}
		_, err = vm.RunProgram(pgrm)
		vm = nil
	} else {
		err = pgrmerr
	}
	return
}

func flushContentFound(atvr *ActiveReader) {
	if atvr.unpio.Size() > 0 {
		if atvr.cntntio == nil {
			atvr.cntntio, _ = goio.NewIORW()
		}
		atvr.cntntio.Print(atvr.unpio)
		atvr.unpio.Close()
		if atvr.cdeio == nil {
			atvr.cdeio, _ = goio.NewIORW()
		}
		if atvr.cntntstarti > -1 {
			atvr.cdeio.Print("_atvr.FlushContent(", atvr.cntntstarti, ",", atvr.cntntio.Size(), ");")
			atvr.cntntstarti = -1
		}
	}
}

//FlushContent FlushContent
func (atvr *ActiveReader) FlushContent(cntntstarti int64, cntntendi int64) {
	if atvr.cntntio != nil {
		if n, nerr := atvr.cntntio.Seek(cntntstarti, 0); n >= 0 && nerr == nil {
			if n < cntntendi {
				if _, writtenerr := io.CopyN(atvr.pwio, atvr.cntntio, cntntendi-n); writtenerr != nil {
					panic(writtenerr)
				}
			}
		}
	}
}

func parseAtiveReaderByte(atvr *ActiveReader, wo io.Writer, b []byte, lblbytes [][]byte, lbli []int) {
	if lbli[1] == 0 && lbli[0] < len(lblbytes[0]) {
		if lbli[0] > 0 && lblbytes[0][lbli[0]-1] == atvr.prvb[0] && lblbytes[0][lbli[0]] != b[0] {
			for n := range lblbytes[0][0:lbli[0]] {
				interpratePassiveReaderByte(atvr, wo, lblbytes[0][n:n+1], lblbytes, lbli)
			}
			lbli[0] = 0
			atvr.prvb[0] = 0
		}
		if lblbytes[0][lbli[0]] == b[0] {
			lbli[0]++
			if lbli[0] == len(lblbytes[0]) {
				if !atvr.foundcode {
					if atvr.unpio.Size() > 0 {
						atvr.pwio.Print(atvr.unpio)
						atvr.unpio.Close()
					}
				}
				atvr.atvpsvmde = mdeatv
				atvr.prvb[0] = 0
			} else {
				atvr.prvb[0] = b[0]
			}
		} else {
			if lbli[0] > 0 {
				for n := range lblbytes[0][0:lbli[0]] {
					interpratePassiveReaderByte(atvr, wo, lblbytes[0][n:n+1], lblbytes, lbli)
				}
				lbli[0] = 0
			}
			atvr.prvb[0] = b[0]
			interpratePassiveReaderByte(atvr, wo, b, lblbytes, lbli)
		}
	} else if lbli[0] == len(lblbytes[0]) && lbli[1] < len(lblbytes[1]) {
		if lblbytes[1][lbli[1]] == b[0] {
			lbli[1]++
			if lbli[1] == len(lblbytes[1]) {
				lbli[0] = 0
				lbli[1] = 0
				if atvr.hasCode && !atvr.foundcode {
					atvr.foundcode = true
				}
				atvr.hasCode = false
				atvr.atvpsvmde = mdepsv
			}
		} else {
			if lbli[1] > 0 {
				if atvr.hasCode {
					if atvr.cdeio == nil {
						atvr.cdeio, _ = goio.NewIORW()
					}
					atvr.cdeio.Print(lblbytes[1][0:lbli[1]])
				} else {
					for cbn := range lblbytes[1][0:lbli[1]] {
						if !atvr.hasCode && strings.TrimSpace(string(lblbytes[1][0:lbli[1]][cbn:cbn+1])) != "" {
							atvr.hasCode = true
							flushContentFound(atvr)
							if atvr.cdeio == nil {
								atvr.cdeio, _ = goio.NewIORW()
							}
							atvr.cdeio.Print(lblbytes[1][0:lbli[1]][cbn:])
							break
						}
					}
				}
			}
			if atvr.hasCode {
				if atvr.cdeio == nil {
					atvr.cdeio, _ = goio.NewIORW()
				}
				atvr.cdeio.Print(b)
			} else {
				if !atvr.hasCode && strings.TrimSpace(string(b)) != "" {
					atvr.hasCode = true
					flushContentFound(atvr)
					if atvr.cdeio == nil {
						atvr.cdeio, _ = goio.NewIORW()
					}
					atvr.cdeio.Print(b)
				}
			}
		}
	}
}

func interpratePassiveReaderByte(atvr *ActiveReader, wo io.Writer, b []byte, lblbytes [][]byte, lbli []int) {
	if lbli[3] == 0 && lbli[2] < len(lblbytes[2]) {
		if lbli[2] > 0 && lblbytes[2][lbli[2]-1] == atvr.prvb[1] && lblbytes[2][lbli[2]] != b[0] {
			atvr.unpio.Print(lblbytes[2][0:lbli[2]])
			if !atvr.foundcode {
				if atvr.unpio.Size() >= 4096 {
					atvr.pwio.Print(atvr.unpio)
					atvr.unpio.Close()
				}
			} else {
				if atvr.cntntstarti == -1 {
					if atvr.cntntio == nil {
						atvr.cntntstarti = 0
					} else {
						atvr.cntntstarti = atvr.cntntio.Size()
					}
				}
			}
			lbli[2] = 0
			atvr.prvb[1] = 0
		}
		if lblbytes[2][lbli[2]] == b[0] {
			lbli[2]++
			if lbli[2] == len(lblbytes[2]) {
				if atvr.unpio.Size() > 0 {
					if !atvr.foundcode {
						if atvr.unpio.Size() >= 4096 {
							atvr.pwio.Print(atvr.unpio)
							atvr.unpio.Close()
						}
					} else {
						if atvr.cntntstarti == -1 {
							if atvr.cntntio == nil {
								atvr.cntntstarti = 0
							} else {
								atvr.cntntstarti = atvr.cntntio.Size()
							}
						}
					}
				}
				atvr.prvb[1] = 0
			} else {
				atvr.prvb[1] = b[0]
			}
		} else {
			if lbli[2] > 0 {
				atvr.unpio.Print(lblbytes[0][0:lbli[0]])
				if !atvr.foundcode {
					if atvr.unpio.Size() >= 4096 {
						atvr.pwio.Print(atvr.unpio)
						atvr.unpio.Close()
					}
				} else {
					if atvr.cntntstarti == -1 {
						if atvr.cntntio == nil {
							atvr.cntntstarti = 0
						} else {
							atvr.cntntstarti = atvr.cntntio.Size()
						}
					}
				}
				lbli[2] = 0
			}
			atvr.prvb[1] = b[0]
			atvr.unpio.Print(b)
			if !atvr.foundcode {
				if atvr.unpio.Size() >= 4096 {
					atvr.pwio.Print(atvr.unpio)
					atvr.unpio.Close()
				}
			} else {
				if atvr.cntntstarti == -1 {
					if atvr.cntntio == nil {
						atvr.cntntstarti = 0
					} else {
						atvr.cntntstarti = atvr.cntntio.Size()
					}
				}
			}
		}
	} else if lbli[2] == len(lblbytes[2]) && lbli[3] < len(lblbytes[3]) {
		if lblbytes[3][lbli[3]] == b[0] {
			lbli[3]++
			if lbli[3] == len(lblbytes[3]) {
				if atvr.hasContent {
					atvr.hasContent = false
					validPassiveConnent(atvr, lblbytes, lbli, true)
				} else {
					validPassiveConnent(atvr, lblbytes, lbli, false)
				}
				atvr.prvb[1] = 0
			}
		} else {
			if lbli[3] > 0 {
				var vlbn = 0
				for vlbn < lbli[3] {
					var fb = lblbytes[3][0:lbli[3]][vlbn : vlbn+1]
					switch {
					case !atvr.hasContent && strings.TrimSpace(string(fb[:])) != "":
						atvr.hasContent = true
						fallthrough
					case atvr.hasContent && strings.TrimSpace(string(fb[:])) != "":
						if atvr.unvlio == nil {
							atvr.unvlio, _ = goio.NewIORW()
						}
						atvr.unvlio.Print(fb[:])
					case atvr.hasContent:
						atvr.hasContent = false
						fb = lblbytes[3][vlbn:lbli[3]]
						vlbn = lbli[3]
						validPassiveConnent(atvr, lblbytes, lbli, true)
						atvr.unpio.Print(fb)
						if !atvr.foundcode {
							if atvr.unpio.Size() >= 4096 {
								atvr.pwio.Print(atvr.unpio)
								atvr.unpio.Close()
							}
						} else {
							if atvr.cntntstarti == -1 {
								if atvr.cntntio == nil {
									atvr.cntntstarti = 0
								} else {
									atvr.cntntstarti = atvr.cntntio.Size()
								}
							}
						}
						continue
					default:
						fb = lblbytes[3][vlbn:lbli[3]]
						vlbn = lbli[3]
						atvr.unpio.Print(fb)
						if !atvr.foundcode {
							if atvr.unpio.Size() >= 4096 {
								atvr.pwio.Print(atvr.unpio)
								atvr.unpio.Close()
							}
						} else {
							if atvr.cntntstarti == -1 {
								if atvr.cntntio == nil {
									atvr.cntntstarti = 0
								} else {
									atvr.cntntstarti = atvr.cntntio.Size()
								}
							}
						}
						lbli[2] = 0
						continue
					}
					vlbn++
				}
				lbli[3] = 0
			}
			switch {
			case !atvr.hasContent && strings.TrimSpace(string(b[0:1])) != "":
				atvr.hasContent = true
				fallthrough
			case atvr.hasContent && strings.TrimSpace(string(b[0:1])) != "":
				if atvr.unvlio == nil {
					atvr.unvlio, _ = goio.NewIORW()
				}
				atvr.unvlio.Print(b)
			case atvr.hasContent:
				atvr.hasContent = false
				validPassiveConnent(atvr, lblbytes, lbli, true)
				atvr.unpio.Print(b)
			default:
				validPassiveConnent(atvr, lblbytes, lbli, false)
				atvr.unpio.Print(b)
			}
			if !atvr.foundcode {
				if atvr.unpio.Size() >= 4096 {
					atvr.pwio.Print(atvr.unpio)
					atvr.unpio.Close()
				}
			} else {
				if atvr.cntntstarti == -1 {
					if atvr.cntntio == nil {
						atvr.cntntstarti = 0
					} else {
						atvr.cntntstarti = atvr.cntntio.Size()
					}
				}
			}
		}
	}
}

func validPassiveConnent(atvr *ActiveReader, lblbytes [][]byte, lbli []int, performTest bool) {
	var valid = false
	if performTest && atvr.unvlio != nil && !atvr.unvlio.Empty() {
		if valid = atvr.unvlio.HasPrefixExp(regexptagstart); valid {
			fmt.Println(atvr.unvlio)
		}
	}
	if valid {
		if atvr.unvlio != nil && !atvr.unvlio.Empty() {
			atvr.unvlio.Close()
		}
		lbli[2] = 0
		lbli[3] = 0
	} else {
		if lbli[2] > 0 {
			atvr.unpio.Print(lblbytes[2][0:lbli[2]])
			lbli[2] = 0
		}
		if atvr.unvlio != nil && !atvr.unvlio.Empty() {
			atvr.unpio.Print(atvr.unvlio)
			atvr.unvlio.Close()
		}
		if lbli[3] > 0 {
			atvr.unpio.Print(lblbytes[3][0:lbli[3]])
			lbli[3] = 0
		}
		if atvr.unpio.Size() > 0 {
			if !atvr.foundcode {
				if atvr.unpio.Size() >= 4096 {
					atvr.pwio.Print(atvr.unpio)
					atvr.unpio.Close()
				}
			} else {
				if atvr.cntntstarti == -1 {
					if atvr.cntntio == nil {
						atvr.cntntstarti = 0
					} else {
						atvr.cntntstarti = atvr.cntntio.Size()
					}
				}
			}
		}
	}
	return
}

const tagstartregexp string = `^((:(([a-z]|[A-Z])\w*)+)|(.:(([a-z]|[A-Z])\w*)+)|(([a-z]|[A-Z])+(:(([a-z]|[A-Z])\w*)+)+))+(:(([a-z]|[A-Z])\w*)+)*(-(([a-z]|[A-Z])\w*)+)?(.([a-z]|[A-Z])+)?$`

var regexptagstart *regexp.Regexp

//const fulltagstartregexp string = `^(([<]|[</])+((:(([a-z]|[A-Z])\w*)+)|(.:(([a-z]|[A-Z])\w*)+)|(([a-z]|[A-Z])+(:(([a-z]|[A-Z])\w*)+)+))+(:(([a-z]|[A-Z])\w*)+)+)*(-(([a-z]|[A-Z])\w*)+)?(.([a-z]|[A-Z])+)?([>]|[/>])+$`

//var regexpfulltagstart *regexp.Regexp

const propregexp string = `^-?-?(([a-z]+[0-9]*)[a-z]*)+(-([a-z]+[0-9]*)[a-z]*)?$`

var regexprop *regexp.Regexp

const propvalnumberexp string = `^[-+]?\d+([.]\d+)?$`

var regexpropvalnumberexp *regexp.Regexp

func init() {

	if regexptagstart == nil {
		regexptagstart = regexp.MustCompile(tagstartregexp)
	}

	//if regexpfulltagstart == nil {
	//	regexpfulltagstart = regexp.MustCompile(fulltagstartregexp)
	//}

	if regexprop == nil {
		regexprop = regexp.MustCompile(propregexp)
	}

	if regexpropvalnumberexp == nil {
		regexpropvalnumberexp = regexp.MustCompile(propvalnumberexp)
	}
}
