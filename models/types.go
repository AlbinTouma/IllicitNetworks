package db_models
/*
  A name used for person or company. This is assumed to be as complete as available. 
  Things having multiple names must be considered perfectly ordinary.
  */
type Name []string
/*
Properties to define countries and territories. zz global iso
I think I'll store this as string rather than as normalised iso */
type Country []string

type Gender []string
// Is an array of number where number is a string ie 1.70cm or 1000
type Height []Number
/*A numeric value, like the size of a piece of land, or the value of a contract. Since all property values in FtM are strings, this is also a string and there is no specified format (e.g. 1,000.00 vs. 1.000,00).
In the future we might want to enable annotations for format, units, or even to introduce a separate property type for monetary values.
*/
type Number string
/*A geographic address used to describe a location of residence or post box. 
There's no specific order for the sub-parts of address (street, city, postal code)
*/
type Address []string

/* There's also AddressEntity, the address associated with an entity. I don't think
I need this yet. It treats address/place as an entity and useful in graphs. */

type Alias []string

// A name used for person/company. Assumed to be as complete a name as available
type WeakAlias []string

type  PreviousName []string


/*A date or time stamp based on ISO 8601, meant to allow for different degrees of precision by specifying prefix.
  This means that 2021, 2021-02, 2021-02-16, 2021-02-16T21, 2021-02-16T21:48 and 2021-02-16T21:48:52 are all valid values, with an implied precision.
  Time zone is always expected to be UTC and cannot be specified otherwise.
*/
type CreatedAt []string

 /*A date or time stamp based on ISO 8601, meant to allow for different degrees of precision by specifying prefix.
  This means that 2021, 2021-02, 2021-02-16, 2021-02-16T21, 2021-02-16T21:48 and 2021-02-16T21:48:52 are all valid values, with an implied precision.
  Time zone is always expected to be UTC and cannot be specified otherwise.
*/

type ModifiedAt []string


/* Longer text such as document text or descriptions. Full text search material */
type Text string

type Description Text


type IndexText string

type Keywords string 

//? 
type Notes string

//?
type Program []string
//?
type Publisher string

// Technically the type here is url
type PublisherUrl string

type Date []string
type RetrievedAt string
  /*A date or time stamp based on ISO 8601, meant to allow for different degrees of precision by specifying prefix.
  This means that 2021, 2021-02, 2021-02-16, 2021-02-16T21, 2021-02-16T21:48 and 2021-02-16T21:48:52 are all valid values, with an implied precision.
  Time zone is always expected to be UTC and cannot be specified otherwise.
*/

type SourceUrl string

// Technically the type here is Topic
type Topics []string

type WikiDataId string
type WikipediaUrl string

/* This is Document type (lots to add here)
type Proof struct {
  Program Document  `json:"Proof"` 
}
*/
type LegalForm string
type Status string



