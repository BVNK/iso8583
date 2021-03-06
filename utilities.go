package iso8583

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	yaml "gopkg.in/yaml.v2"
)

func SetConfig(specFiles []string) (err error) {
	specs = make(map[string]Spec)
	for _, specfile := range specFiles {
		if _, err = os.Stat(specfile); os.IsNotExist(err) {
			fmt.Printf("%s does not exist: %v", specfile, err)
		}

		var spec Spec
		spec, err = SpecFromFile(specfile)
		if err != nil {
			return
		}

		specs[spec.Version()] = spec
	}
	return
}

func (iso Message) Spec() Spec {
	return iso.spec
}
func (iso Message) Mti() string {
	return iso.mti
}
func (iso Message) Bitmap() []byte {
	return iso.bitmap
}

func (iso Message) Data() map[int]Field {
	return iso.data
}

func (message Message) Error(code string) (response string, err error) {

	if response, err = message.ToString(); err != nil {
		return
	}

	date := time.Now()
	dateTime := fmt.Sprintf("%04d%02d%02d%02d%02d%02d", date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), date.Second())

	if err = message.AddMTI(fmt.Sprintf("%s644", string(message.Mti()[0]))); err != nil {
		return
	}

	if err = message.AddField(39, code); err != nil {
		return
	}

	if err = message.AddField(12, dateTime); err != nil {
		return
	}

	if response, err = message.ToString(); err != nil {
		return
	}

	return response, nil
}

//returns a field description for a field
//field id must be in the format spec.index.subindex.subindex...
//Example: 2003.30.1.2
func DescribeField(fieldID string) (desc FieldDescription, indices []int, level int, err error) {
	specName, indices, level, err := DecodeFieldID(fieldID)
	if err != nil {
		return
	}

	//Get spec
	spec, ok := specs[specName]
	if !ok {
		err = fmt.Errorf("Spect %s is not supported", specName)
		return
	}

	//Get top field
	desc, ok = spec.fields[indices[0]]
	if !ok {
		err = fmt.Errorf("Spec %s does not contain field %d", spec.Version(), indices[0])
		return
	}

	if len(indices) == 1 {
		return
	}

	//Get Subfield
	desc, ok = desc.Subfields[indices[1]]
	if !ok {
		err = fmt.Errorf("Spec %s does not contain field %s", spec.Version(), fieldID)
		return
	}

	if len(indices) == 2 {
		return
	}

	//Get sub subfield
	desc, ok = desc.Subfields[indices[2]]
	if !ok {
		err = fmt.Errorf("Spec %s does not contain field %s", spec.Version(), fieldID)
		return
	}

	return
}

func GetFieldSpec(fieldID string) (spec Spec, err error) {
	specName, _, _, err := DecodeFieldID(fieldID)
	if err != nil {
		return
	}

	spec, ok := specs[specName]
	if !ok {
		err = fmt.Errorf("Spect %s is not supported", specName)
		return
	}

	return
}

func DecodeFieldID(fieldID string) (spec string, indices []int, level int, err error) {
	data := strings.Split(fieldID, ".")
	if len(data) < 2 || len(data) > 4 { //specs only have 4 levels
		err = errors.New("Invalid field ID")
		return
	}

	if data[0] != SPEC1987 && data[0] != SPEC1993 && data[0] != SPEC2003 {
		err = errors.New("Invalid field SPEC. Must be 1987, 1997 or 2003")
		return
	}
	spec = data[0]

	level = len(data) - 2
	for _, index := range data[1:] {
		val, err2 := strconv.Atoi(index)
		if err2 != nil {
			return
		}
		indices = append(indices, val)
	}

	if level >= len(indices) || level < 0 {
		err = errors.New("Invlaid field level.")
	}

	return
}

func (iso *Message) MessageCode() (code string) {
	if err := ValidateMti(iso.mti); err != nil {
		return ""
	}

	return iso.mti[1:4]
}

func (iso *Message) GetFlow() (flow, response MessageFlow, err error) {
	if iso.spec.messageFlows == nil {
		err = fmt.Errorf("Message flows are not supported for spec %s", iso.spec.Version())
		return
	}

	if err = ValidateMti(iso.mti); err != nil {
		return
	}

	message := iso.mti[1:4]

	flow, ok := iso.spec.messageFlows[message]
	if !ok {
		err = fmt.Errorf("No flow available for mti %s", iso.mti)
		return
	}

	if flow.ResponseMTI != "" {
		if response, ok = iso.spec.messageFlows[flow.ResponseMTI]; !ok {
			err = fmt.Errorf("Message %s requires a response but there is no flow for response %s", message, flow.ResponseMTI)
			return
		}
	}

	return
}

func (iso *Message) GetField(index int) (field Field, err error) {
	data := iso.Data()
	if data == nil {
		err = fmt.Errorf("Message has no data")
		return
	}

	field, ok := data[index]
	if !ok {
		err = fmt.Errorf("Field %d not found in message data", index)
		return
	}

	return field, nil
}

func (iso *Message) GetTransactionAmount(index int) (amount float64, currencyCode string, err error) {
	if index != BIT87_TRANSACTION_AMOUNT && index != BIT87_ORIGINAL_AMOUNT {
		err = errors.New("Can only get transaction amounts from fields 4 and 30")
		return
	}

	field, err := iso.GetField(index)
	if err != nil {
		return
	}

	var tempAmount decimal.Decimal

	switch iso.spec.version {
	case "1987", "1993":
		currency, err2 := iso.GetField(BIT87_TXN_CURRENCY_CODE)
		if err = err2; err != nil {
			return
		}
		currencyCode = currency.Value
		tempAmount, err = decimal.NewFromString(field.Value)
		if err != nil {
			return
		}
	case "2003":
		if len(field.Value) != 16 {
			err = errors.New("Invalid amount")
			return
		}
		currencyCode = field.Value[:3]
		tempAmount, err = decimal.NewFromString(field.Value[4:16])
		if err != nil {
			return
		}
	}

	decimals, ok := Currencies[currencyCode]
	if !ok {
		err = errors.New("Invalid or unsupported currency")
		return
	}

	decimalUnits := decimal.NewFromFloat(float64(decimals.Decimals))
	if decimalUnits.LessThanOrEqual(decimal.NewFromFloat(0.0)) {
		err = errors.New("Invalid decimal units")
		return
	}

	currencyCode = strings.ToLower(decimals.Code)
	amount, _ = tempAmount.Div(decimalUnits).Float64()
	return
}

//This gets very complex, so won't validate now, but framework is prepared
func (iso *Message) ValidateFlowFields() (err error) {
	if iso.data == nil {
		return fmt.Errorf("ISO message has no data")
	}

	flow, _, err := iso.GetFlow()
	if err != nil {
		return
	}

	if flow.MandatoryFields == nil {
		return fmt.Errorf("No mandatory fields found for messat %s. Check if fields have been configured", iso.mti)
	}

	for field, requirement := range flow.MandatoryFields {
		value, ok := iso.data[field]
		if requirement == "M" && (!ok || value.Value == "") {
			return fmt.Errorf("Message %s is missing mandatory field %d with requirement status %s", iso.mti, field, requirement)
		}
	}

	return
}

func leftPad(s string, length int, pad string) string {
	if len(s) >= length {
		return s
	}
	padding := strings.Repeat(pad, length-len(s))
	return padding + s
}

func rightPad(s string, length int, pad string) string {
	if len(s) >= length {
		return s
	}
	padding := strings.Repeat(pad, length-len(s))
	return s + padding
}

// MtiValidator validates and iso8583 mti
func ValidateMti(mti string) error {
	len := len(mti)
	if len != 4 {
		return errors.New("MTI must be length (4)")
	}

	_, err := strconv.ParseInt(mti, 10, 64)
	if err != nil {
		return errors.New("MTI can only contain integers")
	}

	return nil
}

func (f Field) Validate() (err error) {
	desc, _, _, err := DescribeField(f.ID)
	if err != nil {
		return
	}

	switch desc.ContentType {
	case "ans", "anb", "ansb", "b", "bmps":
	case "n":
		err = validation.Validate(f.Value, is.UTFLetterNumeric)
	case "a":
		err = validation.Validate(f.Value, is.Alpha)
	case "an":
		err = validation.Validate(f.Value, is.Alphanumeric)
	case "anp": //alphanumeric with right space padding
		err = validation.Validate(strings.TrimRight(f.Value, " "), is.UTFLetterNumeric)
	case "xn":
		if string(f.Value[0]) != "c" && string(f.Value[0]) != "d" {
			err = errors.New("XN value must be prefixed with a c or d")
			break
		}
		err = validation.Validate(f.Value[1:], is.UTFLetterNumeric)
	case "z":
		//Tracks 2 and 3 code set as specified in ISO 4909, ISO 7811-2 and ISO 7813.
		//Most of track 2 data is already in track 1. Track 3 is almost never used,
		//so we will simply store and not validate for now because it is a bit more complex
	default:
		err = errors.New("Invalid content type")
	}

	return
}

//We only retrieve date time values, so if a field has separate date and time, join them before passing
func parseTimeFormat(value string) (result time.Time, err error) {
	format := "CCYYMMDDhhmmss"
	if len(value) != len(format) {
		err = fmt.Errorf("%s is not formatted as %s", value, format)
		return
	}

	layout := "2006-01-02T15:04:05.000Z"
	datestring := fmt.Sprintf("%s-%s-%sT%s:%s:%s.000Z", value[0:4], value[4:6], value[6:8], value[8:10], value[10:12], value[12:14])
	return time.Parse(layout, datestring)
}

func validateSubField(data string, subIndex int, desc FieldDescription) (err error) {
	return
}

func (s Spec) Version() string {
	return s.version
}

// readFromFile reads a yaml specfile and loads
// and iso8583 spec from it
func (s *Spec) readFromFile(filename string) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	type TempSpec struct {
		Version string                   `yaml:"Version"`
		Fields  map[int]FieldDescription `yaml:"Fields"`
	}

	var temp TempSpec
	err = yaml.Unmarshal(content, &temp) // expecting content to be valid yaml
	if err != nil {
		return err
	}

	if !strings.Contains(strings.Join([]string{SPEC1987, SPEC1993, SPEC2003}, ","), temp.Version) {
		return fmt.Errorf("Invalid spec version %v.", temp.Version)
	}

	if temp.Fields == nil || len(temp.Fields) == 0 {
		return fmt.Errorf("Invalid spec file %s", filename)
	}

	s.version, s.fields = temp.Version, temp.Fields

	s.messageFlows, err = loadFlows(filename)
	return nil
}

func loadFlows(filename string) (flows map[string]MessageFlow, err error) {
	filename = strings.TrimSuffix(strings.TrimSuffix(filename, ".yaml"), ".yml")
	flowsFilename := filename + ".flow.json"
	mandatoryFieldsFilename := filename + ".mandatory.yaml"

	content, err := ioutil.ReadFile(flowsFilename)
	if err != nil {
		return
	}

	if err = json.Unmarshal(content, &flows); err != nil {
		return nil, err
	}

	var mandatoryFields map[int]map[string]string

	//If an error occurs, we simply do not set the mandatory fields to the specific message
	content, _ = ioutil.ReadFile(mandatoryFieldsFilename)
	if err = yaml.Unmarshal(content, &mandatoryFields); err != nil {
		return flows, nil
	}

	for index, flow := range flows {
		for fieldIndex, field := range mandatoryFields {
			if fieldRequirement, ok := field[index]; ok {
				flowFields := flow.MandatoryFields
				if flowFields == nil {
					flowFields = make(map[int]string)
				}
				flowFields[fieldIndex] = fieldRequirement
				flow.MandatoryFields = flowFields
			}

		}
		flows[index] = flow
	}

	return
}

// SpecFromFile returns a brand new empty spec
func SpecFromFile(filename string) (Spec, error) {
	s := Spec{}
	err := s.readFromFile(filename)
	if err != nil {
		return s, err
	}
	return s, nil
}

// Get returns the value of bit i from map m
// i = 0 gets the left most bit or most significant bit
func GetBit(m []byte, i int) bool {
	return m[i/8]&tA[i%8] != 0
}

// Set sets bit i of map m to value v.
// It doesn't check the bounds of the slice.
func SetBit(m []byte, i int) {
	index := i / 8
	bit := i % 8
	m[index] = m[index] | tA[bit]
}

func ClearBit(m []byte, i int) {
	index := i / 8
	bit := i % 8
	m[index] = m[index] & tB[bit]
}

// Len returns the length (in bits) of the provided byteslice.
// It will always be a multipile of 8 bits.
func BitMapLen(m []byte) int {
	return len(m) * 8
	//return len(m) * 8 * 2 //Because the map was converted from hex, we have 2 characters per byte
}
