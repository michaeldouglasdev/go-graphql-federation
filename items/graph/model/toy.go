package model

/*import (
	"fmt"
	"io"
	"strconv"
)

type Toy struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Price       *ItemPrice     `json:"price"`
	Category    ToyCategory    `json:"category"`
	Material    ToyMaterial    `json:"material"`
	SuitableFor ToySuitableFor `json:"suitableFor"`
}

func (Toy) IsItem()                   {}
func (this Toy) GetID() string        { return this.ID }
func (this Toy) GetName() string      { return this.Name }
func (this Toy) GetPrice() *ItemPrice { return this.Price }

func (Toy) IsEntity() {}

type ToyCategory string

const (
	ToyCategoryInteractive ToyCategory = "INTERACTIVE"
)

var AllToyCategory = []ToyCategory{
	ToyCategoryInteractive,
}

func (e ToyCategory) IsValid() bool {
	switch e {
	case ToyCategoryInteractive:
		return true
	}
	return false
}

func (e ToyCategory) String() string {
	return string(e)
}

func (e *ToyCategory) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ToyCategory(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ToyCategory", str)
	}
	return nil
}

func (e ToyCategory) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ToyMaterial string

const (
	ToyMaterialPlastic ToyMaterial = "PLASTIC"
)

var AllToyMaterial = []ToyMaterial{
	ToyMaterialPlastic,
}

func (e ToyMaterial) IsValid() bool {
	switch e {
	case ToyMaterialPlastic:
		return true
	}
	return false
}

func (e ToyMaterial) String() string {
	return string(e)
}

func (e *ToyMaterial) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ToyMaterial(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ToyMaterial", str)
	}
	return nil
}

func (e ToyMaterial) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ToySuitableFor string

const (
	ToySuitableForCats ToySuitableFor = "CATS"
	ToySuitableForDogs ToySuitableFor = "DOGS"
)

var AllToySuitableFor = []ToySuitableFor{
	ToySuitableForCats,
	ToySuitableForDogs,
}

func (e ToySuitableFor) IsValid() bool {
	switch e {
	case ToySuitableForCats, ToySuitableForDogs:
		return true
	}
	return false
}

func (e ToySuitableFor) String() string {
	return string(e)
}

func (e *ToySuitableFor) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ToySuitableFor(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ToySuitableFor", str)
	}
	return nil
}

func (e ToySuitableFor) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
*/
